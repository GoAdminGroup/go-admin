// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package iris

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
	"github.com/kataras/iris"
	template2 "html/template"
	"net/http"
	"strings"
)

type Iris struct {
}

func init() {
	engine.Register(new(Iris))
}

func (is *Iris) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *iris.Application
		ok  bool
	)
	if eng, ok = router.(*iris.Application); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Handle(strings.ToUpper(req.Method), req.URL, func(c iris.Context) {
				ctx := context.NewContext(c.Request())

				var params = map[string]string{}
				c.Params().Visit(func(key string, value string) {
					params[key] = value
				})

				for key, value := range params {
					if c.Request().URL.RawQuery == "" {
						c.Request().URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
					} else {
						c.Request().URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))).Next()
				for key, head := range ctx.Response.Header {
					c.Header(key, head[0])
				}
				c.StatusCode(ctx.Response.StatusCode)
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
					_, _ = c.WriteString(buf.String())
				}
			})
		}
	}

	return nil
}

func (is *Iris) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx iris.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(iris.Context); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey := ctx.GetCookie("go_admin_session")

	if sesKey == "" {
		ctx.Redirect(globalConfig.Url("/login"), http.StatusFound)
		return
	}

	userId, ok := auth.Driver.Load(sesKey)["user_id"]

	if !ok {
		ctx.Redirect(globalConfig.Url("/login"), http.StatusFound)
		return
	}

	user, ok := auth.GetCurUserById(int64(userId.(float64)))

	if !ok {
		ctx.Redirect(globalConfig.Url("/login"), http.StatusFound)
		return
	}

	var (
		panel types.Panel
		err   error
	)

	if !auth.CheckPermissions(user, ctx.Path(), ctx.Method()) {
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

	tmpl, tmplName := template.Get(globalConfig.Theme).GetTemplate(ctx.GetHeader(constant.PjaxHeader) == "true")

	ctx.Header("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName,
		types.NewPage(user, *(menu.GetGlobalMenu(user).SetActiveClass(globalConfig.UrlRemovePrefix(ctx.Request().URL.String()))),
			panel, globalConfig))
	if err != nil {
		logger.Error("Iris Content", err)
	}
	_, _ = ctx.WriteString(buf.String())
}
