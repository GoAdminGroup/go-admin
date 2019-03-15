// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

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
	"github.com/gin-gonic/gin"
	template2 "html/template"
	"net/http"
	"strings"
)

type Gin struct {
}

func init() {
	engine.Register(new(Gin))
}

func (gins *Gin) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *gin.Engine
		ok  bool
	)
	if eng, ok = router.(*gin.Engine); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy plugins.Plugin
		plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Handle(strings.ToUpper(req.Method), req.URL, func(c *gin.Context) {
				ctx := context.NewContext(c.Request)

				for _, param := range c.Params {
					if c.Request.URL.RawQuery == "" {
						c.Request.URL.RawQuery += strings.Replace(param.Key, ":", "", -1) + "=" + param.Value
					} else {
						c.Request.URL.RawQuery += "&" + strings.Replace(param.Key, ":", "", -1) + "=" + param.Value
					}
				}

				plugCopy.GetHandler(c.Request.URL.Path, strings.ToLower(c.Request.Method))(ctx)
				for key, head := range ctx.Response.Header {
					c.Header(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					buf.ReadFrom(ctx.Response.Body)
					c.String(ctx.Response.StatusCode, buf.String())
				} else {
					c.Status(ctx.Response.StatusCode)
				}
			})
		}
	}

	return nil
}

func (gins *Gin) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx *gin.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(*gin.Context); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey, err := ctx.Cookie("go_admin_session")

	if err != nil || sesKey == "" {
		ctx.Redirect(http.StatusFound, "/"+globalConfig.PREFIX+"/login")
		ctx.Abort()
	}

	userId, ok := auth.Driver.Load(sesKey)["user_id"]

	if !ok {
		ctx.Redirect(http.StatusFound, "/"+globalConfig.PREFIX+"/login")
		ctx.Abort()
	}

	user, ok := auth.GetCurUserById(userId.(string))

	if !ok {
		ctx.Redirect(http.StatusFound, "/"+globalConfig.PREFIX+"/login")
		ctx.Abort()
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.Request.URL.Path, ctx.Request.Method) {
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

	tmpl, tmplName := template.Get(globalConfig.THEME).GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

	ctx.Header("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *(menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Request.URL.String(), "/"+globalConfig.PREFIX, "", 1))),
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: "/" + globalConfig.PREFIX,
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
	})
	ctx.String(http.StatusOK, buf.String())
}
