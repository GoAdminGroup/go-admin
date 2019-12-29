// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"net/http"
	"strconv"
	"time"
)

const DefaultCookieKey = "go_admin_session"

// NewDBDriver return the default PersistenceDriver.
func newDBDriver(conn db.Connection) *DBDriver {
	return &DBDriver{
		conn: conn,
	}
}

// PersistenceDriver is a driver of storing and getting the session info.
type PersistenceDriver interface {
	Load(string) map[string]interface{}
	Update(sid string, values map[string]interface{})
}

// GetSessionByKey get the session value by key.
func GetSessionByKey(sesKey, key string, conn db.Connection) interface{} {
	return newDBDriver(conn).Load(sesKey)[key]
}

// Session contains info of session.
type Session struct {
	Expires time.Duration
	Cookie  string
	Values  map[string]interface{}
	Driver  PersistenceDriver
	Sid     string
	Context *context.Context
}

// Config wraps the Session info.
type Config struct {
	Expires time.Duration
	Cookie  string
}

// UpdateConfig update the Expires and Cookie of Session.
func (ses *Session) UpdateConfig(config Config) {
	ses.Expires = config.Expires
	ses.Cookie = config.Cookie
}

// Get get the session value.
func (ses *Session) Get(key string) interface{} {
	return ses.Values[key]
}

// Add add the session value of key.
func (ses *Session) Add(key string, value interface{}) {
	ses.Values[key] = value
	ses.Driver.Update(ses.Sid, ses.Values)
	cookie := http.Cookie{
		Name:     ses.Cookie,
		Value:    ses.Sid,
		MaxAge:   7200,
		Expires:  time.Now().Add(ses.Expires),
		HttpOnly: true,
		Path:     "/",
	}
	if config.Get().Domain != "" {
		cookie.Domain = config.Get().Domain
	}
	ses.Context.SetCookie(&cookie)
}

// Clear clear a Session.
func (ses *Session) Clear() {
	ses.Values = map[string]interface{}{}
	ses.Driver.Update(ses.Sid, ses.Values)
}

// UseDriver set the driver of the Session.
func (ses *Session) UseDriver(driver PersistenceDriver) {
	ses.Driver = driver
}

// StartCtx return a Session from the given Context.
func (ses *Session) StartCtx(ctx *context.Context) *Session {
	if cookie, err := ctx.Request.Cookie(ses.Cookie); err == nil && cookie.Value != "" {
		ses.Sid = cookie.Value
		valueFromDriver := ses.Driver.Load(cookie.Value)
		if len(valueFromDriver) > 0 {
			ses.Values = valueFromDriver
		}
	} else {
		ses.Sid = modules.Uuid()
	}
	ses.Context = ctx
	return ses
}

// InitSession return the default Session.
func InitSession(ctx *context.Context, conn db.Connection) *Session {

	sessions := new(Session)
	sessions.UpdateConfig(Config{
		Expires: time.Second * time.Duration(config.Get().SessionLifeTime),
		Cookie:  DefaultCookieKey,
	})

	sessions.UseDriver(newDBDriver(conn))
	sessions.Values = make(map[string]interface{})

	return sessions.StartCtx(ctx)
}

// DBDriver is a driver which uses database as a persistence tool.
type DBDriver struct {
	conn db.Connection
}

// Load implements the PersistenceDriver.Load.
func (driver *DBDriver) Load(sid string) map[string]interface{} {
	sesModel, _ := driver.table("goadmin_session").Where("sid", "=", sid).First()

	if sesModel == nil {
		return map[string]interface{}{}
	}

	var values map[string]interface{}
	_ = json.Unmarshal([]byte(sesModel["values"].(string)), &values)
	return values
}

func deleteOverdueSession(conn db.Connection) {

	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
			panic(err)
		}
	}()

	var (
		duration = strconv.Itoa(config.Get().SessionLifeTime + 1000)
		driver   = config.Get().Databases.GetDefault().Driver
		cmd      = ``
	)

	// TODO: use dialect module
	if db.DriverPostgresql == driver {
		cmd = `delete from goadmin_session where extract(epoch from now()) - ` + duration + ` > extract(epoch from created_at)`
	} else if db.DriverMysql == driver {
		cmd = `delete from goadmin_session where unix_timestamp(created_at) < unix_timestamp() - ` + duration
	} else if db.DriverSqlite == driver {
		cmd = `delete from goadmin_session where strftime('%s', created_at) < strftime('%s', 'now') - ` + duration
	}

	logger.LogSQL(cmd, nil)

	_, _ = conn.Query(cmd)
}

// Update implements the PersistenceDriver.Update.
func (driver *DBDriver) Update(sid string, values map[string]interface{}) {

	go deleteOverdueSession(driver.conn)

	if sid != "" {
		if len(values) == 0 {
			_ = driver.table("goadmin_session").Where("sid", "=", sid).Delete()
			return
		}
		valuesByte, _ := json.Marshal(values)
		sesModel, _ := driver.table("goadmin_session").Where("sid", "=", sid).First()
		if sesModel == nil {
			_, _ = driver.table("goadmin_session").Insert(dialect.H{
				"values": string(valuesByte),
				"sid":    sid,
			})
		} else {
			_, _ = driver.table("goadmin_session").
				Where("sid", "=", sid).
				Update(dialect.H{
					"values": string(valuesByte),
				})
		}
	}
}

func (driver *DBDriver) table(table string) *db.SQL {
	return db.Table(table).WithDriver(driver.conn)
}
