// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package chi

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
	cfg "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/go-chi/chi"
)

// Chi structure value is a Chi GoAdmin adapter.
type Chi struct {
	adapter.BaseAdapter
	ctx Context
	app *chi.Mux
}

func init() {
	engine.Register(new(Chi))
}

// User implements the method Adapter.User.
func (ch *Chi) User(ctx interface{}) (models.UserModel, bool) {
	return ch.GetUser(ctx, ch)
}

// Use implements the method Adapter.Use.
func (ch *Chi) Use(app interface{}, plugs []plugins.Plugin) error {
	return ch.GetUse(app, plugs, ch)
}

func (ch *Chi) Run() error                 { panic("not implement") }
func (ch *Chi) DisableLog()                { panic("not implement") }
func (ch *Chi) Static(prefix, path string) { panic("not implement") }

// Content implements the method Adapter.Content.
func (ch *Chi) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	ch.GetContent(ctx, getPanelFn, ch, btns, fn)
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
func (ch *Chi) SetApp(app interface{}) error {
	var (
		eng *chi.Mux
		ok  bool
	)
	if eng, ok = app.(*chi.Mux); !ok {
		return errors.New("chi adapter SetApp: wrong parameter")
	}
	ch.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (ch *Chi) AddHandler(method, path string, handlers context.Handlers) {
	url := path
	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")
	url = reg1.ReplaceAllString(url, "{$1}/")
	url = reg2.ReplaceAllString(url, "{$1}")

	if len(url) > 1 && url[0] == '/' && url[1] == '/' {
		url = url[1:]
	}

	getHandleFunc(ch.app, strings.ToUpper(method))(url, func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path[len(r.URL.Path)-1] == '/' {
			r.URL.Path = r.URL.Path[:len(r.URL.Path)-1]
		}

		ctx := context.NewContext(r)

		params := chi.RouteContext(r.Context()).URLParams

		for i := 0; i < len(params.Values); i++ {
			if r.URL.RawQuery == "" {
				r.URL.RawQuery += strings.ReplaceAll(params.Keys[i], ":", "") + "=" + params.Values[i]
			} else {
				r.URL.RawQuery += "&" + strings.ReplaceAll(params.Keys[i], ":", "") + "=" + params.Values[i]
			}
		}

		ctx.SetHandlers(handlers).Next()
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

// SetContext implements the method Adapter.SetContext.
func (ch *Chi) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx Context
		ok  bool
	)
	if ctx, ok = contextInterface.(Context); !ok {
		panic("chi adapter SetContext: wrong parameter")
	}
	return &Chi{ctx: ctx}
}

// Name implements the method Adapter.Name.
func (ch *Chi) Name() string {
	return "chi"
}

// Redirect implements the method Adapter.Redirect.
func (ch *Chi) Redirect() {
	http.Redirect(ch.ctx.Response, ch.ctx.Request, cfg.Url(cfg.GetLoginUrl()), http.StatusFound)
}

// SetContentType implements the method Adapter.SetContentType.
func (ch *Chi) SetContentType() {
	ch.ctx.Response.Header().Set("Content-Type", ch.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (ch *Chi) Write(body []byte) {
	ch.ctx.Response.WriteHeader(http.StatusOK)
	_, _ = ch.ctx.Response.Write(body)
}

// GetCookie implements the method Adapter.GetCookie.
func (ch *Chi) GetCookie() (string, error) {
	cookie, err := ch.ctx.Request.Cookie(ch.CookieKey())
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

// Lang implements the method Adapter.Lang.
func (ch *Chi) Lang() string {
	return ch.ctx.Request.URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (ch *Chi) Path() string {
	return ch.ctx.Request.URL.Path
}

// Method implements the method Adapter.Method.
func (ch *Chi) Method() string {
	return ch.ctx.Request.Method
}

// FormParam implements the method Adapter.FormParam.
func (ch *Chi) FormParam() url.Values {
	_ = ch.ctx.Request.ParseMultipartForm(32 << 20)
	return ch.ctx.Request.PostForm
}

// IsPjax implements the method Adapter.IsPjax.
func (ch *Chi) IsPjax() bool {
	return ch.ctx.Request.Header.Get(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (ch *Chi) Query() url.Values {
	return ch.ctx.Request.URL.Query()
}
