package auth

import (
	"github.com/valyala/fasthttp"
	"goAdmin/connections/mysql"
)

type User struct {
	ID        string
	Level     string
	Name      string
	LevelName string
	CreateAt  string
}

func Filter(ctx *fasthttp.RequestCtx) (User, bool) {
	cookieSec := string(ctx.Request.Header.Cookie("go_admin_session")[:])
	id := InitSessionHelper(ctx).GetUserIdFromSession(cookieSec)
	user, ok := GetCurUserById(id)

	if !ok {
		return user, false
	}

	return user, CheckLevel(user) && CheckLimit(user)
}

func GetCurUserById(id string) (user User, ok bool) {
	admin, _ := mysql.Query("select * from goadmin_users where id = ?", id)

	if len(admin) < 1 {
		ok = false
		return
	}

	user.ID = id
	user.Level = "super"
	user.LevelName = "超级管理员"
	user.Name = admin[0]["name"].(string)
	user.CreateAt = admin[0]["created_at"].(string)
	ok = true

	return
}

func CheckLevel(user User) bool {
	if user.Level == "super" {
		return true
	} else {
		return false
	}
}

func CheckLimit(user User) bool {
	return true
}
