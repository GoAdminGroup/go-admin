package auth

import (
	"github.com/fasthttp-contrib/sessions"
	"github.com/valyala/fasthttp"
	"goAdmin/modules"
)

type SessionHelper struct {
	Sess sessions.Session
}

func InitSessionHelper(ctx *fasthttp.RequestCtx) *SessionHelper {
	return &SessionHelper{
		sessions.StartFasthttp(ctx),
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
