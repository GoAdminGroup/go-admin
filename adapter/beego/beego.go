// Copyright 2019 GoAdmin.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package beego

import (
	"bytes"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	gctx "github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/modules/logger"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/plugins/admin/modules/constant"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	template2 "html/template"
	"net/http"
	"strings"
)

type Beego struct {
}

func init() {
	engine.Register(new(Beego))
}

func (bee *Beego) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *beego.App
		ok  bool
	)
	if eng, ok = router.(*beego.App); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Handlers.AddMethod(req.Method, req.URL, func(c *context.Context) {
				for key, value := range c.Input.Params() {
					if c.Request.URL.RawQuery == "" {
						c.Request.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
					} else {
						c.Request.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
					}
				}
				ctx := gctx.NewContext(c.Request)
				ctx.SetHandlers(plugCopy.GetHandler(c.Request.URL.Path, strings.ToLower(c.Request.Method))).Next()
				for key, head := range ctx.Response.Header {
					c.ResponseWriter.Header().Add(key, head[0])
				}
				c.ResponseWriter.WriteHeader(ctx.Response.StatusCode)
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
					c.WriteString(buf.String())
				}
			})
		}
	}

	return nil
}

func (bee *Beego) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx *context.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(*context.Context); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey := ctx.GetCookie("go_admin_session")

	if sesKey == "" {
		ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		return
	}

	userId, ok := auth.Driver.Load(sesKey)["user_id"]

	if !ok {
		ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		return
	}

	user, ok := auth.GetCurUserById(int64(userId.(float64)))

	if !ok {
		ctx.Redirect(http.StatusFound, globalConfig.Url("/login"))
		return
	}

	var (
		panel types.Panel
		err   error
	)

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

	ctx.ResponseWriter.Header().Add("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user,
		*(menu.GetGlobalMenu(user).SetActiveClass(globalConfig.UrlRemovePrefix(ctx.Request.URL.String()))),
		panel, globalConfig))
	if err != nil {
		logger.Error("Beego Content", err)
	}
	ctx.WriteString(buf.String())
}
