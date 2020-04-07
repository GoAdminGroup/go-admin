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
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Gorilla structure value is a Gorilla GoAdmin adapter.
type Gorilla struct {
	adapter.BaseAdapter
	ctx Context
	app *mux.Router
}

func init() {
	engine.Register(new(Gorilla))
}

func (g *Gorilla) User(ci interface{}) (models.UserModel, bool) {
	return g.GetUser(ci, g)
}

func (g *Gorilla) Use(router interface{}, plugs []plugins.Plugin) error {
	return g.GetUse(router, plugs, g)
}

func (g *Gorilla) Content(ctx interface{}, getPanelFn types.GetPanelFn, btns ...types.Button) {
	g.GetContent(ctx, getPanelFn, g, btns)
}

type HandlerFunc func(ctx Context) (types.Panel, error)

func Content(handler HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := Context{
			Request:  request,
			Response: writer,
		}
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(Context))
		})
	}
}

func (g *Gorilla) SetApp(app interface{}) error {
	var (
		eng *mux.Router
		ok  bool
	)
	if eng, ok = app.(*mux.Router); !ok {
		return errors.New("wrong parameter")
	}
	g.app = eng
	return nil
}

func (g *Gorilla) AddHandler(method, path string, handlers context.Handlers) {

	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")

	url := path
	url = reg1.ReplaceAllString(url, "{$1}/")
	url = reg2.ReplaceAllString(url, "{$1}")

	g.app.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewContext(r)

		params := mux.Vars(r)

		for key, param := range params {
			if r.URL.RawQuery == "" {
				r.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + param
			} else {
				r.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + param
			}
		}

		ctx.SetHandlers(handlers).Next()
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
	}).Methods(strings.ToUpper(method))
}

// Context wraps the Request and Response object of Gorilla.
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func (g *Gorilla) Name() string {
	return "gorilla"
}

func (g *Gorilla) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx Context
		ok  bool
	)
	if ctx, ok = contextInterface.(Context); !ok {
		panic("wrong parameter")
	}

	return &Gorilla{ctx: ctx}
}

func (g *Gorilla) Redirect() {
	http.Redirect(g.ctx.Response, g.ctx.Request, config.Url("/login"), http.StatusFound)
}

func (g *Gorilla) SetContentType() {
	g.ctx.Response.Header().Set("Content-Type", g.HTMLContentType())
}

func (g *Gorilla) Write(body []byte) {
	_, _ = g.ctx.Response.Write(body)
}

func (g *Gorilla) GetCookie() (string, error) {
	cookie, err := g.ctx.Request.Cookie(g.CookieKey())
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

func (g *Gorilla) Path() string {
	return g.ctx.Request.RequestURI
}

func (g *Gorilla) Method() string {
	return g.ctx.Request.Method
}

func (g *Gorilla) FormParam() url.Values {
	_ = g.ctx.Request.ParseMultipartForm(32 << 20)
	return g.ctx.Request.PostForm
}

func (g *Gorilla) IsPjax() bool {
	return g.ctx.Request.Header.Get(constant.PjaxHeader) == "true"
}
