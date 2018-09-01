package controller

import (
	"github.com/chenhg5/go-admin/modules/auth"
	tmp "html/template"
	"fmt"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/template/adminlte"
	"bytes"
	"net/http"
)

func Auth(ctx *context.Context) {

	password := ctx.Request.FormValue("password")
	username := ctx.Request.FormValue("username")

	if user, ok := auth.Check(password, username); ok {

		auth.SetCookie(ctx, user)

		ctx.Write(http.StatusOK, map[string]string{
			"Content-Type": "application/json",
		}, `{"code":200, "msg":"登录成功", "url":"` + AssertRootUrl + `"}`)
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

	tmpl, err := tmp.New("login").Parse(adminlte.Adminlte["login/theme1"])
	if err != nil {
		fmt.Println(err)
	}
	buf := new(bytes.Buffer)
	fmt.Println(tmpl.ExecuteTemplate(buf, "login_theme1", struct {
		AssertRootUrl string
	}{AssertRootUrl}))
	ctx.Write(http.StatusOK, map[string]string{}, buf.String())

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}
