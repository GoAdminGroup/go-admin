// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gorilla

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strings"
)

// Gorilla structure value is a Gorilla GoAdmin adapter.
type Gorilla struct {
	adapter.BaseAdapter
}

func init() {
	engine.Register(new(Gorilla))
}

// Use implement WebFrameWork.Use method.
func (g *Gorilla) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *mux.Router
		ok  bool
	)
	if eng, ok = router.(*mux.Router); !ok {
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

			eng.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
				ctx := context.NewContext(r)

				params := mux.Vars(r)

				for key, param := range params {
					if r.URL.RawQuery == "" {
						r.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + param
					} else {
						r.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + param
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(r.URL.Path, strings.ToLower(r.Method))).Next()
				for key, head := range ctx.Response.Header {
					w.Header().Add(key, head[0])
				}

				if ctx.Response.Body == nil {
					w.WriteHeader(ctx.Response.StatusCode)
					return
				}

				w.WriteHeader(ctx.Response.StatusCode)

				buf := new(bytes.Buffer)
				_, _ = buf.ReadFrom(ctx.Response.Body)

				_, err := w.Write(buf.Bytes())
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}).Methods(strings.ToUpper(req.Method))
		}
	}

	return nil
}

// Context wraps the Request and Response object of Gorilla.
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

// Content implement WebFrameWork.Content method.
func (g *Gorilla) Content(contextInterface interface{}, getPanelFn types.GetPanelFn) {

	var (
		ctx Context
		ok  bool
	)
	if ctx, ok = contextInterface.(Context); !ok {
		panic("wrong parameter")
	}

	sesKey, err := ctx.Request.Cookie(g.CookieKey())

	if err != nil || sesKey == nil {
		http.Redirect(ctx.Response, ctx.Request, config.Get().Url("/login"), http.StatusFound)
		return
	}

	body, authSuccess, err := g.GetContent(sesKey.Value, ctx.Request.RequestURI,
		ctx.Request.Method, ctx.Request.Header.Get(constant.PjaxHeader), getPanelFn, ctx)

	if !authSuccess {
		http.Redirect(ctx.Response, ctx.Request, config.Get().Url("/login"), http.StatusFound)
		return
	}

	if err != nil {
		logger.Error("Echo Content", err)
	}

	ctx.Response.WriteHeader(http.StatusOK)
	_, _ = ctx.Response.Write(body)
}
