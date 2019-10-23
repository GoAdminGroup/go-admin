// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package echo

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/labstack/echo"
	template2 "html/template"
	"net/http"
	"strings"
)

type Echo struct {
}

func init() {
	engine.Register(new(Echo))
}

func (e *Echo) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *echo.Echo
		ok  bool
	)
	if eng, ok = router.(*echo.Echo); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Add(strings.ToUpper(req.Method), req.URL, func(c echo.Context) error {
				ctx := context.NewContext(c.Request())

				for _, key := range c.ParamNames() {
					if c.Request().URL.RawQuery == "" {
						c.Request().URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + c.Param(key)
					} else {
						c.Request().URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + c.Param(key)
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))).Next()
				for key, head := range ctx.Response.Header {
					c.Response().Header().Set(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
					_ = c.String(ctx.Response.StatusCode, buf.String())
				}
				return nil
			})
		}
	}

	return nil
}

func (e *Echo) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx echo.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(echo.Context); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey, err := ctx.Cookie("go_admin_session")

	if err != nil || sesKey == nil {
		_ = ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		return
	}

	userId, ok := auth.Driver.Load(sesKey.Value)["user_id"]

	if !ok {
		_ = ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		return
	}

	user, ok := auth.GetCurUserById(int64(userId.(float64)))

	if !ok {
		_ = ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		return
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.Path(), ctx.Request().Method) {
		alert := template.Get(globalConfig.Theme).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").SetContent(template2.HTML("no permission")).GetContent()

		panel = types.Panel{
			Content:     alert,
			Description: language.Get("error"),
			Title:       language.Get("error"),
		}
	} else {
		panel, err = c(ctx)
		if err != nil {
			alert := template.Get(globalConfig.Theme).
				Alert().
				SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
				SetTheme("warning").SetContent(template2.HTML(err.Error())).GetContent()
			panel = types.Panel{
				Content:     alert,
				Description: language.Get("error"),
				Title:       language.Get("error"),
			}
		}
	}

	tmpl, tmplName := template.Get(globalConfig.Theme).GetTemplate(ctx.Request().Header.Get(constant.PjaxHeader) == "true")

	ctx.Response().Header().Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user, *(menu.GetGlobalMenu(user).SetActiveClass(globalConfig.UrlRemovePrefix(ctx.Request().URL.String()))), panel, globalConfig))
	if err != nil {
		logger.Error("Echo Content", err)
	}
	_ = ctx.String(http.StatusOK, buf.String())
}
