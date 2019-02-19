// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package iris

import (
	"bytes"
	"errors"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
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
		var plugCopy plugins.Plugin
		plugCopy = plug
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

				plugCopy.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))(ctx)
				for key, head := range ctx.Response.Header {
					c.Header(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					buf.ReadFrom(ctx.Response.Body)
					c.WriteString(buf.String())
				}
				c.StatusCode(ctx.Response.StatusCode)
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
		ctx.Redirect("/"+globalConfig.PREFIX+"/login", http.StatusFound)
		return
	}

	userId, ok := auth.Driver.Load(sesKey)["user_id"]

	if !ok {
		ctx.Redirect("/"+globalConfig.PREFIX+"/login", http.StatusFound)
		return
	}

	user, ok := auth.GetCurUserById(userId.(string))

	if !ok {
		ctx.Redirect("/"+globalConfig.PREFIX+"/login", http.StatusFound)
		return
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.Path(), ctx.Method()) {
		alert := template.Get(globalConfig.THEME).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template2.HTML("没有权限")).GetContent()

		panel = types.Panel{
			Content:     alert,
			Description: "Error",
			Title:       "Error",
		}
	} else {
		panel = c()
	}

	tmpl, tmplName := template.Get(globalConfig.THEME).GetTemplate(ctx.GetHeader("X-PJAX") == "true")

	ctx.Header("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *(menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Request().URL.String(), "/"+globalConfig.PREFIX, "", 1))),
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: "/" + globalConfig.PREFIX,
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
	})
	ctx.WriteString(buf.String())
}
