package controller

import (
	"bytes"
	"errors"
	template2 "html/template"
	"net/http"
	"net/url"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/captcha"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

var (
	ErrWrongCaptcha = errors.New("wrong captcha")
)

// Auth check the input password and username for authentication.
func (h *Handler) Auth(ctx *context.Context) {

	var (
		user              models.UserModel
		ok                bool
		errMsg            = "fail"
		customAuth, exist = h.services.GetOrNot(auth.ServiceKey)
		err               error
	)

	if capDriver, has := h.captchaConfig["driver"]; has {
		if capt, has := captcha.Get(capDriver); has && !capt.Validate(ctx.FormValue("token")) {
			response.BadRequest(ctx, ErrWrongCaptcha.Error())
			return
		}
	}

	if !exist {
		method := ctx.FormValue("method")
		var authenticator auth.Authenticator
		switch method {
		case auth.GeneralMethod:
			authenticator = auth.NewGeneralAuth(h.conn)
		case auth.LdapMethod:
			ldapCfg := h.config.Ldap
			if !ldapCfg.Enable {
				response.BadRequest(ctx, "ldap login disable")
				return
			}
			authenticator = auth.NewLdapAuth(h.conn, auth.NewLdapConfig(ldapCfg.ServerUrls, ldapCfg.BindDN, ldapCfg.BindPwd, ldapCfg.BaseDN))
		default:
			authenticator = auth.NewGeneralAuth(h.conn)
		}
		if user, err = authenticator.Authenticate(ctx.Request); err != nil {
			response.BadRequest(ctx, err.Error())
			return
		}
		user = user.SetConn(h.conn).WithRoles().WithPermissions().WithMenus()
		ok = true
	} else {
		user, ok, errMsg = auth.GetService(customAuth).P(ctx)
	}

	if !ok {
		response.BadRequest(ctx, errMsg)
		return
	}

	if err := auth.NewCookieManger(h.conn).SetCookie(ctx, user); err != nil {
		response.Error(ctx, err.Error())
		return
	}

	if ref := ctx.Referer(); ref != "" {
		if u, err := url.Parse(ref); err == nil {
			v := u.Query()
			if r := v.Get("ref"); r != "" {
				rr, _ := url.QueryUnescape(r)
				response.OkWithData(ctx, map[string]interface{}{
					"url": rr,
				})
				return
			}
		}
	}

	response.OkWithData(ctx, map[string]interface{}{
		"url": h.config.GetIndexURL(),
	})
}

// Logout delete the cookie.
func (h *Handler) Logout(ctx *context.Context) {
	err := auth.NewCookieManger(db.GetConnection(h.services)).DelCookie(ctx)
	if err != nil {
		logger.Error("logout error", err)
	}
	ctx.AddHeader("Location", h.config.Url(config.GetLoginUrl()))
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
