package controller

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"fmt"
	"github.com/chenhg5/go-admin/context"
	"bytes"
	"net/http"
	"github.com/chenhg5/go-admin/template"
)

func Auth(ctx *context.Context) {

	password := ctx.Request.FormValue("password")
	username := ctx.Request.FormValue("username")

	if user, ok := auth.Check(password, username); ok {

		auth.SetCookie(ctx, user)

		ctx.Write(http.StatusOK, map[string]string{
			"Content-Type": "application/json",
		}, `{"code":200, "msg":"登录成功", "url":"` + Config.ADMIN_PREFIX + `"}`)
		return
	}
	ctx.Write(http.StatusBadRequest, map[string]string{
		"Content-Type": "application/json",
	}, `{"code":400, "msg":"登录失败"}`)
	return
}

func Logout(ctx *context.Context) {
	auth.DelCookie(ctx)
	ctx.Response.Header.Add("Location", "/login")
	ctx.Response.StatusCode = 302
}

func ShowLogin(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	tmpl, name := template.GetComp("login").GetTemplate()
	buf := new(bytes.Buffer)
	fmt.Println(tmpl.ExecuteTemplate(buf, name, struct {
		AssertRootUrl string
	}{Config.ADMIN_PREFIX}))
	ctx.Write(http.StatusOK, map[string]string{}, buf.String())

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}
