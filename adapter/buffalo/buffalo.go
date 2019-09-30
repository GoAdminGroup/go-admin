// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package buffalo

import (
	"bytes"
	"errors"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/logger"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/modules/system"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/plugins/admin/modules/constant"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
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

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.Request().URL.Path, ctx.Request().Method) {
		alert := template.Get(globalConfig.Theme).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template2.HTML("no permission")).GetContent()

		panel = types.Panel{
			Content:     alert,
			Description: "Error",
			Title:       "Error",
		}
	} else {
		panel = c()
	}

	tmpl, tmplName := template.Get(globalConfig.Theme).GetTemplate(ctx.Request().Header.Get(constant.PjaxHeader) == "true")

	ctx.Response().Header().Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *(menu.GetGlobalMenu(user).SetActiveClass(globalConfig.UrlRemovePrefix(ctx.Request().URL.String()))),
		System: types.SystemInfo{
			Version: system.Version,
		},
		Panel:       panel,
		UrlPrefix:   globalConfig.Prefix(),
		Title:       globalConfig.Title,
		Logo:        globalConfig.Logo,
		MiniLogo:    globalConfig.MiniLogo,
		ColorScheme: globalConfig.ColorScheme,
	})
	if err != nil {
		logger.Error("Buffalo Content", err)
	}
	_ = ctx.Render(http.StatusOK, render.String(buf.String()))
}
