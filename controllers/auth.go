package controller

import (
	"github.com/valyala/fasthttp"
	"bytes"
	"goAdmin/auth"
	"goAdmin/template"
)

func Auth(ctx *fasthttp.RequestCtx) {

	password := ctx.FormValue("password")
	username := string(ctx.FormValue("username")[:])

	if user, ok := auth.Check(password, username); ok {

		auth.SetCookie(ctx, user)

		ctx.WriteString(`{"code":200, "msg":"登录成功", "url":"/user/info"}`)
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

	defer handle(ctx)

	buffer := new(bytes.Buffer)
	template.GetLoginPage(buffer)

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}