package auth

import (
	"github.com/golang/crypto/bcrypt"
	"github.com/valyala/fasthttp"
	"goAdmin/config"
	"goAdmin/connections/mysql"
	"strconv"
	"time"
	"goAdmin/modules"
)

func Check(password []byte, username string) (user User, ok bool) {

	admin, _ := mysql.Query("select * from goadmin_users where username = ?", username)

	if len(admin) < 1 {
		ok = false
	} else {
		if ComparePassword(password, admin[0]["password"].(string)) {
			ok = true

			roleModel, _ := mysql.Query("select r.id, r.name, r.slug from goadmin_role_users as u left join goadmin_roles as r on u.role_id = r.id where user_id = ?", admin[0]["id"])

			user.ID = strconv.FormatInt(admin[0]["id"].(int64), 10)
			user.Level = roleModel[0]["slug"].(string)
			user.LevelName = roleModel[0]["name"].(string)
			user.Name = admin[0]["name"].(string)
			user.CreateAt = admin[0]["created_at"].(string)
			user.Avatar = admin[0]["avatar"].(string)
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

type CSRFToken []string

var TokenHelper = new(CSRFToken)

func (token *CSRFToken) AddToken() string {
	tokenStr := modules.Uuid(35)
	if len(*token) == 1 && (*token)[0] == "" {
		(*token)[0] = tokenStr
	} else {
		*token = append(*token, tokenStr)
	}
	return tokenStr
}

func (token *CSRFToken) CheckToken(tocheck string) bool {
	for i := 0; i < len(*token); i++ {
		if (*token)[i] == tocheck {
			*token = append((*token)[0:i], (*token)[i:len((*token))]...)
			return true
		}
	}
	return false
}