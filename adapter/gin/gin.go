// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gin

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
		var plugCopy = plug
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

				ctx.SetHandlers(plugCopy.GetHandler(c.Request.URL.Path, strings.ToLower(c.Request.Method))).Next()
				for key, head := range ctx.Response.Header {
					c.Header(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
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
		ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		ctx.Abort()
		return
	}

	userId, ok := auth.Driver.Load(sesKey)["user_id"]

	if !ok {
		ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		ctx.Abort()
		return
	}

	user, ok := auth.GetCurUserById(int64(userId.(float64)))

	if !ok {
		ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		ctx.Abort()
		return
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.Request.URL.Path, ctx.Request.Method) {
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

	tmpl, tmplName := template.Get(globalConfig.Theme).GetTemplate(ctx.Request.Header.Get(constant.PjaxHeader) == "true")

	ctx.Header("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user, *(menu.GetGlobalMenu(user).SetActiveClass(globalConfig.UrlRemovePrefix(ctx.Request.URL.String()))), panel, globalConfig))
	if err != nil {
		logger.Error("Gin Content", err)
	}
	ctx.String(http.StatusOK, buf.String())
}
