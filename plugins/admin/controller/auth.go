package controller

import (
	"bytes"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/modules/response"
	"github.com/chenhg5/go-admin/template"
	"net/http"
)

func Auth(ctx *context.Context) {

	password := ctx.FormValue("password")
	username := ctx.FormValue("username")

	if password == "" || username == "" {
		response.BadRequest(ctx, "fail")
		return
	}

	if user, ok := auth.Check(password, username); ok {

		auth.SetCookie(ctx, user)

		response.OkWithData(ctx, map[string]interface{}{
			"url": config.GetIndexUrl(),
		})
		return
	}
	response.BadRequest(ctx, "fail")
}

func Logout(ctx *context.Context) {
	auth.DelCookie(ctx)
	ctx.AddHeader("Location", config.Url("/login"))
	ctx.SetStatusCode(302)
}

func ShowLogin(ctx *context.Context) {

	tmpl, name := template.GetComp("login").GetTemplate()
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, name, struct {
		UrlPrefix string
	}{config.Prefix()}); err == nil {
		ctx.Html(http.StatusOK, buf.String())
	} else {
		ctx.Html(http.StatusOK, "resolve template error (；′⌒`)")
	}
}
