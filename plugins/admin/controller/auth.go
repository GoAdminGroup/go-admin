package controller

import (
	"bytes"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/captcha"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
)

// Auth check the input password and username for authentication.
func (h *Handler) Auth(ctx *context.Context) {

	var (
		user     models.UserModel
		ok       bool
		errMsg   = "fail"
		s, exist = h.services.GetOrNot(auth.ServiceKey)
	)

	if capDriver, ok := h.captchaConfig["driver"]; ok {
		capt, ok := captcha.Get(capDriver)

		if ok {
			if !capt.Validate(ctx.FormValue("token")) {
				response.BadRequest(ctx, "wrong captcha")
				return
			}
		}
	}

	if !exist {
		password := ctx.FormValue("password")
		username := ctx.FormValue("username")

		if password == "" || username == "" {
			response.BadRequest(ctx, "wrong password or username")
			return
		}
		user, ok = auth.Check(password, username, h.conn)
	} else {
		user, ok, errMsg = auth.GetService(s).P(ctx)
	}

	if !ok {
		response.BadRequest(ctx, errMsg)
		return
	}

	auth.SetCookie(ctx, user, h.conn)

	response.OkWithData(ctx, map[string]interface{}{
		"url": h.config.GetIndexURL(),
	})
	return

}

// Logout delete the cookie.
func (h *Handler) Logout(ctx *context.Context) {
	auth.DelCookie(ctx, db.GetConnection(h.services))
	ctx.AddHeader("Location", h.config.Url("/login"))
	ctx.SetStatusCode(302)
}

// ShowLogin show the login page.
func (h *Handler) ShowLogin(ctx *context.Context) {

	tmpl, name := template.GetComp("login").GetTemplate()
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, name, struct {
		UrlPrefix string
		Title     string
		Logo      template2.HTML
		CdnUrl    string
		System    types.SystemInfo
	}{
		UrlPrefix: h.config.AssertPrefix(),
		Title:     h.config.LoginTitle,
		Logo:      h.config.LoginLogo,
		System: types.SystemInfo{
			Version: system.Version(),
		},
		CdnUrl: h.config.AssetUrl,
	}); err == nil {
		ctx.HTML(http.StatusOK, buf.String())
	} else {
		logger.Error(err)
		ctx.HTML(http.StatusOK, "parse template error (；′⌒`)")
	}
}
