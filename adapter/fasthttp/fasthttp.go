// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package fasthttp

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// Fasthttp structure value is a Fasthttp GoAdmin adapter.
type Fasthttp struct {
	adapter.BaseAdapter
	ctx *fasthttp.RequestCtx
	app *fasthttprouter.Router
}

func init() {
	engine.Register(new(Fasthttp))
}

// User implements the method Adapter.User.
func (fast *Fasthttp) User(ctx interface{}) (models.UserModel, bool) {
	return fast.GetUser(ctx, fast)
}

// Use implements the method Adapter.Use.
func (fast *Fasthttp) Use(app interface{}, plugs []plugins.Plugin) error {
	return fast.GetUse(app, plugs, fast)
}

func (fast *Fasthttp) Run() error                 { panic("not implement") }
func (fast *Fasthttp) DisableLog()                { panic("not implement") }
func (fast *Fasthttp) Static(prefix, path string) { panic("not implement") }

// Content implements the method Adapter.Content.
func (fast *Fasthttp) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	fast.GetContent(ctx, getPanelFn, fast, btns, fn)
}

type HandlerFunc func(ctx *fasthttp.RequestCtx) (types.Panel, error)

func Content(handler HandlerFunc) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(*fasthttp.RequestCtx))
		})
	}
}

// SetApp implements the method Adapter.SetApp.
func (fast *Fasthttp) SetApp(app interface{}) error {
	var (
		eng *fasthttprouter.Router
		ok  bool
	)
	if eng, ok = app.(*fasthttprouter.Router); !ok {
		return errors.New("fasthttp adapter SetApp: wrong parameter")
	}

	fast.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (fast *Fasthttp) AddHandler(method, path string, handlers context.Handlers) {
	fast.app.Handle(strings.ToUpper(method), path, func(c *fasthttp.RequestCtx) {
		httpreq := convertCtx(c)
		ctx := context.NewContext(httpreq)

		var params = make(map[string]string)
		c.VisitUserValues(func(i []byte, i2 interface{}) {
			if value, ok := i2.(string); ok {
				params[string(i)] = value
			}
		})

		for key, value := range params {
			if httpreq.URL.RawQuery == "" {
				httpreq.URL.RawQuery += strings.ReplaceAll(key, ":", "") + "=" + value
			} else {
				httpreq.URL.RawQuery += "&" + strings.ReplaceAll(key, ":", "") + "=" + value
			}
		}

		ctx.SetHandlers(handlers).Next()
		for key, head := range ctx.Response.Header {
			c.Response.Header.Set(key, head[0])
		}
		if ctx.Response.Body != nil {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(ctx.Response.Body)
			_, _ = c.WriteString(buf.String())
		}
		c.Response.SetStatusCode(ctx.Response.StatusCode)
	})
}

func convertCtx(ctx *fasthttp.RequestCtx) *http.Request {
	var r http.Request

	body := ctx.PostBody()
	r.Method = string(ctx.Method())
	r.Proto = "HTTP/1.1"
	r.ProtoMajor = 1
	r.ProtoMinor = 1
	r.RequestURI = string(ctx.RequestURI())
	r.ContentLength = int64(len(body))
	r.Host = string(ctx.Host())
	r.RemoteAddr = ctx.RemoteAddr().String()

	hdr := make(http.Header)
	ctx.Request.Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		switch sk {
		case "Transfer-Encoding":
			r.TransferEncoding = append(r.TransferEncoding, sv)
		default:
			hdr.Set(sk, sv)
		}
	})
	r.Header = hdr
	r.Body = &netHTTPBody{body}
	rURL, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		ctx.Logger().Printf("cannot parse requestURI %q: %s", r.RequestURI, err)
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return &r
	}
	r.URL = rURL
	return &r
}

type netHTTPBody struct {
	b []byte
}

func (r *netHTTPBody) Read(p []byte) (int, error) {
	if len(r.b) == 0 {
		return 0, io.EOF
	}
	n := copy(p, r.b)
	r.b = r.b[n:]
	return n, nil
}

func (r *netHTTPBody) Close() error {
	r.b = r.b[:0]
	return nil
}

// Name implements the method Adapter.Name.
func (fast *Fasthttp) Name() string {
	return "fasthttp"
}

// SetContext implements the method Adapter.SetContext.
func (fast *Fasthttp) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *fasthttp.RequestCtx
		ok  bool
	)
	if ctx, ok = contextInterface.(*fasthttp.RequestCtx); !ok {
		panic("fasthttp adapter SetContext: wrong parameter")
	}
	return &Fasthttp{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (fast *Fasthttp) Redirect() {
	fast.ctx.Redirect(config.Url(config.GetLoginUrl()), http.StatusFound)
}

// SetContentType implements the method Adapter.SetContentType.
func (fast *Fasthttp) SetContentType() {
	fast.ctx.Response.Header.Set("Content-Type", fast.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (fast *Fasthttp) Write(body []byte) {
	_, _ = fast.ctx.Write(body)
}

// GetCookie implements the method Adapter.GetCookie.
func (fast *Fasthttp) GetCookie() (string, error) {
	return string(fast.ctx.Request.Header.Cookie(fast.CookieKey())), nil
}

// Lang implements the method Adapter.Lang.
func (fast *Fasthttp) Lang() string {
	return string(fast.ctx.Request.URI().QueryArgs().Peek("__ga_lang"))
}

// Path implements the method Adapter.Path.
func (fast *Fasthttp) Path() string {
	return string(fast.ctx.Path())
}

// Method implements the method Adapter.Method.
func (fast *Fasthttp) Method() string {
	return string(fast.ctx.Method())
}

// FormParam implements the method Adapter.FormParam.
func (fast *Fasthttp) FormParam() url.Values {
	f, _ := fast.ctx.MultipartForm()
	if f != nil {
		return f.Value
	}
	return url.Values{}
}

// IsPjax implements the method Adapter.IsPjax.
func (fast *Fasthttp) IsPjax() bool {
	return string(fast.ctx.Request.Header.Peek(constant.PjaxHeader)) == "true"
}

// Query implements the method Adapter.Query.
func (fast *Fasthttp) Query() url.Values {
	queryStr := fast.ctx.URI().QueryString()
	queryObj, err := url.Parse(string(queryStr))

	if err != nil {
		return url.Values{}
	}

	return queryObj.Query()
}
