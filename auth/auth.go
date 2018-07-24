package auth

import (
	"github.com/golang/crypto/bcrypt"
	"github.com/valyala/fasthttp"
	"goAdmin/connections/mysql"
	"strconv"
	"time"
	"goAdmin/config"
)

func Check(password []byte, username string) (user User, ok bool) {

	admin, _ := mysql.Query("select * from goadmin_users where username = ?", username)

	if len(admin) < 1 {
		ok = false
	} else {
		if ComparePassword(password, admin[0]["password"].(string)) {
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

func ComparePassword(comPwd []byte, pwdHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), comPwd)
	if err != nil {
		return false
	} else {
		return true
	}
}

func EncodePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash[:])
}

func SetCookie(ctx *fasthttp.RequestCtx, user User) bool {

	sessionKey := InitSessionHelper(ctx).PutIntoSession(user.ID)

	var c fasthttp.Cookie
	c.SetKey("go_admin_session")
	c.SetValue(sessionKey)
	c.SetDomain(config.EnvConfig["AUTH_DOMAIN"].(string))
	c.SetExpire(time.Now().Add(time.Hour * 48))
	ctx.Response.Header.SetCookie(&c)

	return true
}

func DelCookie(ctx *fasthttp.RequestCtx) bool {
	var c fasthttp.Cookie
	c.SetKey("go_admin_session")
	c.SetValue("")
	c.SetDomain(config.EnvConfig["AUTH_DOMAIN"].(string))
	c.SetExpire(time.Now())
	ctx.Response.Header.SetCookie(&c)

	return true
}
