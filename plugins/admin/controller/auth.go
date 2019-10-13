package controller

import (
	"bytes"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
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
		Title     string
		Logo      template2.HTML
		CdnUrl    string
		System    types.SystemInfo
	}{
		UrlPrefix: config.Prefix(),
		Title:     config.LoginTitle,
		Logo:      config.LoginLogo,
		System: types.SystemInfo{
			Version: system.Version,
		},
		CdnUrl: config.AssetUrl,
	}); err == nil {
		ctx.Html(http.StatusOK, buf.String())
	} else {
		logger.Error(err)
		ctx.Html(http.StatusOK, "parse template error (；′⌒`)")
	}
}
