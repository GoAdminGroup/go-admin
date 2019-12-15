package controller

import (
	"bytes"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/captcha"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
)

// Auth check the input password and username for authentication.
func Auth(ctx *context.Context) {

	password := ctx.FormValue("password")
	username := ctx.FormValue("username")

	if password == "" || username == "" {
		response.BadRequest(ctx, "wrong password or username")
		return
	}

	conn := db.GetConnection(services)

	if user, ok := auth.Check(password, username, conn); ok {

		cd, ok := captcha.Get(captchaConfig["driver"])

		if ok {
			if !cd.Validate(ctx.FormValue("token")) {
				response.BadRequest(ctx, "wrong captcha")
			}
		}

		auth.SetCookie(ctx, user, conn)

		response.OkWithData(ctx, map[string]interface{}{
			"url": config.GetIndexURL(),
		})
		return
	}
	response.BadRequest(ctx, "fail")
}

// Logout delete the cookie.
func Logout(ctx *context.Context) {
	auth.DelCookie(ctx, db.GetConnection(services))
	ctx.AddHeader("Location", config.Url("/login"))
	ctx.SetStatusCode(302)
}

// ShowLogin show the login page.
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
		UrlPrefix: config.AssertPrefix(),
		Title:     config.LoginTitle,
		Logo:      config.LoginLogo,
		System: types.SystemInfo{
			Version: system.Version(),
		},
		CdnUrl: config.AssetUrl,
	}); err == nil {
		ctx.HTML(http.StatusOK, buf.String())
	} else {
		logger.Error(err)
		ctx.HTML(http.StatusOK, "parse template error (；′⌒`)")
	}
}
