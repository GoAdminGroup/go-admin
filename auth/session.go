package auth

import (
	"encoding/json"
	"github.com/fasthttp-contrib/sessions"
	"github.com/valyala/fasthttp"
	"goAdmin/connections/mysql"
	"goAdmin/modules"
	"time"
)

type SessionHelper struct {
	Sess sessions.Session
}

func InitSessionHelper(ctx *fasthttp.RequestCtx) *SessionHelper {

	sessions.UpdateConfig(sessions.Config{
		Expires: time.Hour * 10,
	})

	var driver MysqlDriver
	sessions.UseDatabase(&driver)

	session := sessions.StartFasthttp(ctx)

	return &SessionHelper{
		session,
	}
}

func (helper *SessionHelper) GetUserIdFromSession(cookieSec string) (id string) {
	var ok bool
	if id, ok = helper.Sess.Get(cookieSec).(string); ok {
		return
	} else {
		return ""
	}
}

func GenerateSessionId() string {
	return modules.Uuid(60)
}

func (helper *SessionHelper) PutIntoSession(value string) string {
	sessionKey := GenerateSessionId()
	helper.Sess.Set(sessionKey, value)
	return sessionKey
}

type MysqlDriver struct{}

func (driver *MysqlDriver) Load(sid string) map[string]interface{} {
	sesModel, _ := mysql.Query("select * from goadmin_session where sid = ?", sid)
	if len(sesModel) < 1 {
		return map[string]interface{}{}
	} else {
		var values map[string]interface{}
		json.Unmarshal([]byte(sesModel[0]["values"].(string)), &values)
		return values
	}
}

func (driver *MysqlDriver) Update(sid string, values map[string]interface{}) {
	if sid != "" && len(values) != 0 {
		valuesByte, _ := json.Marshal(values)
		sesModel, _ := mysql.Query("select * from goadmin_session where sid = ?", sid)
		if len(sesModel) < 1 {
			mysql.Exec("insert into goadmin_session (`values`, sid) values (?, ?)", string(valuesByte), sid)
		} else {
			mysql.Exec("update goadmin_session set `values` = ? where sid = ?", string(valuesByte), sid)
		}
	}
}
