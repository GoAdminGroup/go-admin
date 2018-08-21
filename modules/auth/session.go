package auth

import (
	"encoding/json"
	"github.com/fasthttp-contrib/sessions"
	"github.com/valyala/fasthttp"
	"goAdmin/modules/connections/mysql"
	"time"
)

var (
	Session sessions.Session
	driver  MysqlDriver
)

func InitSession(ctx *fasthttp.RequestCtx) sessions.Session {

	sessions.UpdateConfig(sessions.Config{
		Expires: time.Hour * 10,
		Cookie:  "go_admin_session",
	})

	sessions.UseDatabase(&driver)

	return sessions.StartFasthttp(ctx)
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
