// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package buffalo

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
	"github.com/gobuffalo/buffalo"
	template2 "html/template"
	"net/http"
	neturl "net/url"
	"regexp"
	"strings"
)

type Buffalo struct {
}

func init() {
	engine.Register(new(Buffalo))
}

func (bu *Buffalo) Use(router interface{}, plugin []plugins.Plugin) error {

	var (
		eng *buffalo.App
		ok  bool
	)
	if eng, ok = router.(*buffalo.App); !ok {
		return errors.New("wrong parameter")
	}

	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {

			url := req.URL
			url = reg1.ReplaceAllString(url, "{$1}/")
			url = reg2.ReplaceAllString(url, "{$1}")

			getHandleFunc(eng, strings.ToUpper(req.Method))(url, func(c buffalo.Context) error {

				if c.Request().URL.Path[len(c.Request().URL.Path)-1] == '/' {
					c.Request().URL.Path = c.Request().URL.Path[:len(c.Request().URL.Path)-1]
				}

				ctx := context.NewContext(c.Request())

				params := c.Params().(neturl.Values)

				for key, param := range params {
					if c.Request().URL.RawQuery == "" {
						c.Request().URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + param[0]
					} else {
						c.Request().URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + param[0]
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))).Next()
				for key, head := range ctx.Response.Header {
					c.Response().Header().Set(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
					c.Response().WriteHeader(ctx.Response.StatusCode)
					_, _ = c.Response().Write(buf.Bytes())
				} else {
					c.Response().WriteHeader(ctx.Response.StatusCode)
				}
				return nil
			})
		}
	}

	return nil
}

type HandleFun func(p string, h buffalo.Handler) *buffalo.RouteInfo

func getHandleFunc(eng *buffalo.App, method string) HandleFun {
	switch method {
	case "GET":
		return eng.GET
	case "POST":
		return eng.POST
	case "PUT":
		return eng.PUT
	case "DELETE":
		return eng.DELETE
	case "HEAD":
		return eng.HEAD
	case "OPTIONS":
		return eng.OPTIONS
	case "PATCH":
		return eng.PATCH
	default:
		panic("wrong method")
	}
}

func (bu *Buffalo) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx buffalo.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(buffalo.Context); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey, err := ctx.Cookies().Get("go_admin_session")

	if err != nil || sesKey == "" {
		_ = ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		return
	}

	userId, ok := auth.Driver.Load(sesKey)["user_id"]

	if !ok {
		_ = ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		return
	}

	user, ok := auth.GetCurUserById(int64(userId.(float64)))

	if !ok {
		_ = ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		return
	}

	var (
		panel types.Panel
	)

	if !auth.CheckPermissions(user, ctx.Request().URL.Path, ctx.Request().Method) {
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
		logger.Error("Buffalo Content", err)
	}
	_, _ = ctx.Response().Write(buf.Bytes())
}
