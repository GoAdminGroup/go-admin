package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/modules/auth"
	tmp "html/template"
	"fmt"
	"goAdmin/template"
)

func Auth(ctx *fasthttp.RequestCtx) {

	password := ctx.FormValue("password")
	username := string(ctx.FormValue("username"))

	if user, ok := auth.Check(password, username); ok {

		auth.SetCookie(ctx, user)

		ctx.WriteString(`{"code":200, "msg":"登录成功", "url":"/"}`)
		return
	}
	ctx.WriteString(`{"code":400, "msg":"登录失败"`)
	return
}

func Logout(ctx *fasthttp.RequestCtx) {
	auth.DelCookie(ctx)
	ctx.Response.Header.Add("Location", "/login")
	ctx.Response.SetStatusCode(302)
}

func ShowLogin(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	tmpl, err := tmp.New("login").Parse(template.Adminlte["login/theme1"])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tmpl.ExecuteTemplate(ctx.Response.BodyWriter(), "login_theme1", nil))

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}
