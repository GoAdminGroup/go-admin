// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package http

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
	template2 "html/template"
	"net/http"
	"strings"
)

type Http struct {
}

func init() {
	engine.Register(new(Http))
}

func (ht *Http) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *http.ServeMux
		ok  bool
	)
	if eng, ok = router.(*http.ServeMux); !ok {
		return errors.New("wrong parameter")
	}

	var reqs map[string][]context.Path
	for _, plug := range plugin {
		reqs = ConstructNetHttpRequest(plug.GetRequest())
		for basicUrl, reqlist := range reqs {
			eng.HandleFunc(basicUrl, func(httpWriter http.ResponseWriter, httpRequest *http.Request) {
				for _, req := range reqlist {
					if httpRequest.Method == strings.ToUpper(req.Method) {
						ctx := context.NewContext(httpRequest)
						plug.GetHandler(req.URL, req.Method)(context.NewContext(httpRequest))
						httpWriter.WriteHeader(ctx.Response.StatusCode)
						if ctx.Response.Body != nil {
							buf := new(bytes.Buffer)
							buf.ReadFrom(ctx.Response.Body)
							httpWriter.Write(buf.Bytes())
						}
					}
				}
			})
		}
	}

	return nil
}

func ConstructNetHttpRequest(reqs []context.Path) map[string][]context.Path {
	var (
		NetHttpRequest = make(map[string][]context.Path, 0)
		usedUrl        []string
	)
	for _, req := range reqs {
		for _, url := range usedUrl {
			if url == req.URL {
				continue
			}
		}
		usedUrl = append(usedUrl, req.URL)
		NetHttpRequest[req.URL] = append(NetHttpRequest[req.URL], req)
	}
	return NetHttpRequest
}

func (ht *Http) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx *http.Request
		ok  bool
	)
	if ctx, ok = contextInterface.(*http.Request); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey, err := ctx.Cookie("go_admin_session")

	if err != nil || sesKey == nil {
		ctx.Response.Header.Set("Location", "/"+globalConfig.PREFIX+"/login")
		ctx.Response.StatusCode = http.StatusFound
		return
	}

	userId, ok := auth.Driver.Load(sesKey.Value)["user_id"]

	if !ok {
		ctx.Response.Header.Set("Location", "/"+globalConfig.PREFIX+"/login")
		ctx.Response.StatusCode = http.StatusFound
		return
	}

	user, ok := auth.GetCurUserById(userId.(string))

	if !ok {
		ctx.Response.Header.Set("Location", "/"+globalConfig.PREFIX+"/login")
		ctx.Response.StatusCode = http.StatusFound
		return
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.RequestURI, ctx.Method) {
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

	tmpl, tmplName := template.Get(globalConfig.THEME).GetTemplate(ctx.Header.Get("X-PJAX") == "true")

	ctx.Header.Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *(menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.URL.String(), "/"+globalConfig.PREFIX, "", 1))),
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: "/" + globalConfig.PREFIX,
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
	})
	//ctx.Body.(buf.String())
}
