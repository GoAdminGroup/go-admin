package controller

import (
	"bytes"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/template"
	"net/http"
)

func Auth(ctx *context.Context) {

	password := ctx.FormValue("password")
	username := ctx.FormValue("username")

	if user, ok := auth.Check(password, username); ok {

		auth.SetCookie(ctx, user)

		menu.Unlock()

		ctx.Json(http.StatusOK, map[string]interface{}{
			"code": 200,
			"msg":  "登录成功",
			"url":  Config.PREFIX + Config.INDEX,
		})
		return
	}
	ctx.Json(http.StatusBadRequest, map[string]interface{}{
		"code": 400,
		"msg":  "登录失败",
	})
	return
}

func Logout(ctx *context.Context) {
	auth.DelCookie(ctx)
	ctx.AddHeader("Location", Config.PREFIX+"/login")
	ctx.SetStatusCode(302)
}

func ShowLogin(ctx *context.Context) {

	tmpl, name := template.GetComp("login").GetTemplate()
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, name, struct {
		AssertRootUrl string
	}{Config.PREFIX}); err == nil {
		ctx.Html(http.StatusOK, buf.String())
	} else {
		ctx.Html(http.StatusOK, "resolve template error (；′⌒`)")
	}
}
