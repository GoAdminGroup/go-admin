// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package chi

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	cfg "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/go-chi/chi"
	"net/http"
	"regexp"
	"strings"
)

// Chi structure value is a Chi GoAdmin adapter.
type Chi struct {
	adapter.BaseAdapter
}

func init() {
	engine.Register(new(Chi))
}

// Use implement WebFrameWork.Use method.
func (ch *Chi) Use(router interface{}, plugin []plugins.Plugin) error {

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

// HandleFun is type of route methods of chi.
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

// Context wraps the Request and Response object of Chi.
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

// Content implement WebFrameWork.Content method.
func (ch *Chi) Content(contextInterface interface{}, getPanelFn types.GetPanelFn) {

	var (
		ctx Context
		ok  bool
	)
	if ctx, ok = contextInterface.(Context); !ok {
		panic("wrong parameter")
	}

	sesKey, err := ctx.Request.Cookie(ch.CookieKey())

	if err != nil || sesKey == nil {
		http.Redirect(ctx.Response, ctx.Request, cfg.Get().Url("/login"), http.StatusFound)
		return
	}

	body, authSuccess, err := ch.GetContent(sesKey.Value, ctx.Request.URL.Path,
		ctx.Request.Method, ctx.Request.Header.Get(constant.PjaxHeader), getPanelFn, ctx)

	if !authSuccess {
		http.Redirect(ctx.Response, ctx.Request, cfg.Get().Url("/login"), http.StatusFound)
		return
	}

	if err != nil {
		logger.Error("Chi Content", err)
	}

	ctx.Response.Header().Set("Content-Type", ch.HTMLContentType())
	ctx.Response.WriteHeader(http.StatusOK)
	_, _ = ctx.Response.Write(body)
}
