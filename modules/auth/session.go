package auth

import (
	"encoding/json"
	"github.com/chenhg5/go-admin/modules/connections"
	"time"
	"github.com/chenhg5/go-admin/context"
	"net/http"
	"github.com/chenhg5/go-admin/plugins/admin/modules"
	"github.com/chenhg5/go-admin/modules/config"
)

var (
	driver MysqlDriver
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
		Domain:   config.Get().DOMAIN,
		Expires:  time.Now().Add(ses.Expires),
		HttpOnly: false,
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
		ses.Sid = modules.Uuid(15)
	}
	ses.Context = ctx
	return ses
}

func InitSession(ctx *context.Context) *Session {

	sessions := new(Session)
	sessions.UpdateConfig(Config{
		Expires: time.Hour * 10,
		Cookie:  "go_admin_session",
	})

	sessions.UseDatabase(&driver)
	sessions.Values = make(map[string]interface{}, 0)

	return sessions.StartCtx(ctx)
}

type MysqlDriver struct{}

func (driver *MysqlDriver) Load(sid string) map[string]interface{} {
	sesModel, _ := connections.GetConnection().Query("select * from goadmin_session where sid = ?", sid)
	if len(sesModel) < 1 {
		return map[string]interface{}{}
	} else {
		var values map[string]interface{}
		json.Unmarshal([]byte(sesModel[0]["values"].(string)), &values)
		return values
	}
}

func (driver *MysqlDriver) Update(sid string, values map[string]interface{}) {
	if sid != "" {
		if len(values) == 0 {
			connections.GetConnection().Exec("delete from goadmin_session where sid = ?", sid)
			return
		}
		valuesByte, _ := json.Marshal(values)
		sesModel, _ := connections.GetConnection().Query("select * from goadmin_session where sid = ?", sid)
		if len(sesModel) < 1 {
			connections.GetConnection().Exec("insert into goadmin_session (`values`, sid) values (?, ?)", string(valuesByte), sid)
		} else {
			connections.GetConnection().Exec("update goadmin_session set `values` = ? where sid = ?", string(valuesByte), sid)
		}
	}
}
