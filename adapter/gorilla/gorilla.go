// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gorilla

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gorilla/mux"
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

// User implements the method Adapter.User.
func (g *Gorilla) User(ctx interface{}) (models.UserModel, bool) {
	return g.GetUser(ctx, g)
}

// Use implements the method Adapter.Use.
func (g *Gorilla) Use(app interface{}, plugs []plugins.Plugin) error {
	return g.GetUse(app, plugs, g)
}

func (g *Gorilla) Run() error                 { panic("not implement") }
func (g *Gorilla) DisableLog()                { panic("not implement") }
func (g *Gorilla) Static(prefix, path string) { panic("not implement") }

// Content implements the method Adapter.Content.
func (g *Gorilla) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	g.GetContent(ctx, getPanelFn, g, btns, fn)
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

// SetApp implements the method Adapter.SetApp.
func (g *Gorilla) SetApp(app interface{}) error {
	var (
		eng *mux.Router
		ok  bool
	)
	if eng, ok = app.(*mux.Router); !ok {
		return errors.New("gorilla adapter SetApp: wrong parameter")
	}
	g.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (g *Gorilla) AddHandler(method, path string, handlers context.Handlers) {

	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")

	u := path
	u = reg1.ReplaceAllString(u, "{$1}/")
	u = reg2.ReplaceAllString(u, "{$1}")

	g.app.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewContext(r)

		params := mux.Vars(r)

		for key, param := range params {
			if r.URL.RawQuery == "" {
				r.URL.RawQuery += strings.ReplaceAll(key, ":", "") + "=" + param
			} else {
				r.URL.RawQuery += "&" + strings.ReplaceAll(key, ":", "") + "=" + param
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

// Name implements the method Adapter.Name.
func (g *Gorilla) Name() string {
	return "gorilla"
}

// SetContext implements the method Adapter.SetContext.
func (g *Gorilla) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx Context
		ok  bool
	)
	if ctx, ok = contextInterface.(Context); !ok {
		panic("gorilla adapter SetContext: wrong parameter")
	}

	return &Gorilla{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (g *Gorilla) Redirect() {
	http.Redirect(g.ctx.Response, g.ctx.Request, config.Url(config.GetLoginUrl()), http.StatusFound)
}

// SetContentType implements the method Adapter.SetContentType.
func (g *Gorilla) SetContentType() {
	g.ctx.Response.Header().Set("Content-Type", g.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (g *Gorilla) Write(body []byte) {
	_, _ = g.ctx.Response.Write(body)
}

// GetCookie implements the method Adapter.GetCookie.
func (g *Gorilla) GetCookie() (string, error) {
	cookie, err := g.ctx.Request.Cookie(g.CookieKey())
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

// Lang implements the method Adapter.Lang.
func (g *Gorilla) Lang() string {
	return g.ctx.Request.URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (g *Gorilla) Path() string {
	return g.ctx.Request.RequestURI
}

// Method implements the method Adapter.Method.
func (g *Gorilla) Method() string {
	return g.ctx.Request.Method
}

// FormParam implements the method Adapter.FormParam.
func (g *Gorilla) FormParam() url.Values {
	_ = g.ctx.Request.ParseMultipartForm(32 << 20)
	return g.ctx.Request.PostForm
}

// IsPjax implements the method Adapter.IsPjax.
func (g *Gorilla) IsPjax() bool {
	return g.ctx.Request.Header.Get(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (g *Gorilla) Query() url.Values {
	return g.ctx.Request.URL.Query()
}
