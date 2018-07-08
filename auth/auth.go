package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/valyala/fasthttp"
	"goAdmin/connections/mysql"
	"strconv"
	"time"
)

func Check(password []byte, username string) (user User, ok bool) {

	admin, _ := mysql.Query("select * from goadmin_users where username = ?", username)

	if len(admin) < 1 {
		ok = false
	} else {
		hashpwd := EncodePassword(password)
		if hashpwd == admin[0]["password"].(string) {
			ok = true
			user.ID = strconv.FormatInt(admin[0]["id"].(int64), 10)
			user.Level = "super"
			user.LevelName = "超级管理员"
			user.Name = admin[0]["name"].(string)
			user.CreateAt = admin[0]["created_at"].(string)
		} else {
			ok = false
		}
	}
	return
}

func EncodePassword(pwd []byte) string {
	hash := sha256.New()
	hash.Write(pwd)
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func SetCookie(ctx *fasthttp.RequestCtx, user User) bool {

	sessionKey := InitSessionHelper(ctx).PutIntoSession(user.ID)

	var c fasthttp.Cookie
	c.SetKey("go_admin_session")
	c.SetValue(sessionKey)
	c.SetDomain("localhost")
	c.SetExpire(time.Now().Add(time.Hour * 48))
	ctx.Response.Header.SetCookie(&c)

	return true
}

func DelCookie(ctx *fasthttp.RequestCtx) bool {
	var c fasthttp.Cookie
	c.SetKey("go_admin_session")
	c.SetValue("")
	c.SetDomain("localhost")
	c.SetExpire(time.Now())
	ctx.Response.Header.SetCookie(&c)

	return true
}
