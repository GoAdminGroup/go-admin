// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package chi

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	cfg "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/go-chi/chi"
	template2 "html/template"
	"net/http"
	"regexp"
	"strings"
)

type Chi struct {
}

func init() {
	engine.Register(new(Chi))
}

func (bu *Chi) Use(router interface{}, plugin []plugins.Plugin) error {

	var (
		eng *chi.Mux
		ok  bool
	)
	if eng, ok = router.(*chi.Mux); !ok {
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

			if len(url) > 1 && url[0] == '/' && url[1] == '/' {
				url = url[1:]
			}

			getHandleFunc(eng, strings.ToUpper(req.Method))(url, func(w http.ResponseWriter, r *http.Request) {

				if r.URL.Path[len(r.URL.Path)-1] == '/' {
					r.URL.Path = r.URL.Path[:len(r.URL.Path)-1]
				}

				ctx := context.NewContext(r)

				params := chi.RouteContext(r.Context()).URLParams

				for i := 0; i < len(params.Values); i++ {
					if r.URL.RawQuery == "" {
						r.URL.RawQuery += strings.Replace(params.Keys[i], ":", "", -1) + "=" + params.Values[i]
					} else {
						r.URL.RawQuery += "&" + strings.Replace(params.Keys[i], ":", "", -1) + "=" + params.Values[i]
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(r.URL.Path, strings.ToLower(r.Method))).Next()
				for key, head := range ctx.Response.Header {
					w.Header().Set(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
					w.WriteHeader(ctx.Response.StatusCode)
					_, _ = w.Write(buf.Bytes())
				} else {
					w.WriteHeader(ctx.Response.StatusCode)
				}
			})
		}
	}

	return nil
}

type HandleFun func(pattern string, handlerFn http.HandlerFunc)

func getHandleFunc(eng *chi.Mux, method string) HandleFun {
	switch method {
	case "GET":
		return eng.Get
	case "POST":
		return eng.Post
	case "PUT":
		return eng.Put
	case "DELETE":
		return eng.Delete
	case "HEAD":
		return eng.Head
	case "OPTIONS":
		return eng.Options
	case "PATCH":
		return eng.Patch
	default:
		panic("wrong method")
	}
}

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func (bu *Chi) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx Context
		ok  bool
	)
	if ctx, ok = contextInterface.(Context); !ok {
		panic("wrong parameter")
	}

	config := cfg.Get()

	sesKey, err := ctx.Request.Cookie("go_admin_session")

	if err != nil || sesKey == nil {
		http.Redirect(ctx.Response, ctx.Request, config.Url("/login"), http.StatusFound)
		return
	}

	userId, ok := auth.Driver.Load(sesKey.Value)["user_id"]

	if !ok {
		http.Redirect(ctx.Response, ctx.Request, config.Url("/login"), http.StatusFound)
		return
	}

	user, ok := auth.GetCurUserById(int64(userId.(float64)))

	if !ok {
		http.Redirect(ctx.Response, ctx.Request, config.Url("/login"), http.StatusFound)
		return
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.Request.URL.Path, ctx.Request.Method) {
		alert := template.Get(config.Theme).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").SetContent(template2.HTML("no permission")).GetContent()

		panel = types.Panel{
			Content:     alert,
			Description: language.Get("error"),
			Title:       language.Get("error"),
		}
	} else {
		panel, err = c(ctx)
		if err != nil {
			alert := template.Get(config.Theme).
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

	tmpl, tmplName := template.Get(config.Theme).GetTemplate(ctx.Request.Header.Get(constant.PjaxHeader) == "true")

	ctx.Response.Header().Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user,
		*(menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Request.URL.String()))),
		panel, config))
	if err != nil {
		logger.Error("Chi Content", err)
	}
	ctx.Response.WriteHeader(http.StatusOK)
	_, _ = ctx.Response.Write(buf.Bytes())
}
