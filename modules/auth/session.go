// Copyright 2019 GoAdmin Core Team.  All rights reserved.
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

var (
	Driver MysqlDriver
)

type PersistenceDriver interface {
	Load(string) map[string]interface{}
	Update(sid string, values map[string]interface{})
}

type SessionInterface interface {
	Get(string) interface{}
	Set(string, interface{})
	UseDatabase(PersistenceDriver)
	StartCtx(*context.Context) Session
}

type Session struct {
	Expires time.Duration
	Cookie  string
	Values  map[string]interface{}
	Driver  PersistenceDriver
	Sid     string
	Context *context.Context
}

type Config struct {
	Expires time.Duration
	Cookie  string
}

func (ses *Session) UpdateConfig(config Config) {
	ses.Expires = config.Expires
	ses.Cookie = config.Cookie
}

func (ses *Session) Get(key string) interface{} {
	return ses.Values[key]
}

func (ses *Session) Set(key string, value interface{}) {
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

func (ses *Session) Clear() {
	ses.Values = map[string]interface{}{}
	ses.Driver.Update(ses.Sid, ses.Values)
}

func (ses *Session) UseDatabase(driver PersistenceDriver) {
	ses.Driver = driver
}

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

func InitSession(ctx *context.Context) *Session {

	sessions := new(Session)
	sessions.UpdateConfig(Config{
		Expires: time.Second * time.Duration(config.Get().SessionLifeTime),
		Cookie:  "go_admin_session",
	})

	sessions.UseDatabase(&Driver)
	sessions.Values = make(map[string]interface{})

	return sessions.StartCtx(ctx)
}

type MysqlDriver struct{}

func (driver *MysqlDriver) Load(sid string) map[string]interface{} {
	sesModel, _ := db.Table("goadmin_session").Where("sid", "=", sid).First()

	if sesModel == nil {
		return map[string]interface{}{}
	} else {
		var values map[string]interface{}
		_ = json.Unmarshal([]byte(sesModel["values"].(string)), &values)
		return values
	}
}

func deleteOverdueSession() {

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

	if db.DriverPostgresql == driver {
		cmd = `delete from goadmin_session where extract(epoch from now()) - ` + duration + ` > extract(epoch from created_at)`
	} else if db.DriverMysql == driver {
		cmd = `delete from goadmin_session where unix_timestamp(created_at) < unix_timestamp() - ` + duration
	} else if db.DriverSqlite == driver {
		cmd = `delete from goadmin_session where strftime('%s', created_at) < strftime('%s', 'now') - ` + duration
	}

	logger.LogSql(cmd, nil)

	_, _ = db.Query(cmd)
}

func (driver *MysqlDriver) Update(sid string, values map[string]interface{}) {

	go deleteOverdueSession()

	if sid != "" {
		if len(values) == 0 {
			_ = db.Table("goadmin_session").Where("sid", "=", sid).Delete()
			return
		}
		valuesByte, _ := json.Marshal(values)
		sesModel, _ := db.Table("goadmin_session").Where("sid", "=", sid).First()
		if sesModel == nil {
			_, _ = db.Table("goadmin_session").Insert(dialect.H{
				"values": string(valuesByte),
				"sid":    sid,
			})
		} else {
			_, _ = db.Table("goadmin_session").
				Where("sid", "=", sid).
				Update(dialect.H{
					"values": string(valuesByte),
				})
		}
	}
}
