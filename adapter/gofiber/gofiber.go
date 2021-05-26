// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gofiber

import (
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
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// Fasthttp structure value is a Fasthttp GoAdmin adapter.
type Gofiber struct {
	adapter.BaseAdapter
	ctx *fiber.Ctx
	app *fiber.App
}

func init() {
	engine.Register(new(Gofiber))
}

// User implements the method Adapter.User.
func (gof *Gofiber) User(ctx interface{}) (models.UserModel, bool) {
	return gof.GetUser(ctx, gof)
}

// Use implements the method Adapter.Use.
func (gof *Gofiber) Use(app interface{}, plugs []plugins.Plugin) error {
	return gof.GetUse(app, plugs, gof)
}

func (fagof *Gofiber) Run() error               { panic("not implement") }
func (gof *Gofiber) DisableLog()                { panic("not implement") }
func (gof *Gofiber) Static(prefix, path string) { panic("not implement") }

// Content implements the method Adapter.Content.
func (gof *Gofiber) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	gof.GetContent(ctx, getPanelFn, gof, btns, fn)
}

// SetApp implements the method Adapter.SetApp.
func (gof *Gofiber) SetApp(app interface{}) error {
	var (
		eng *fiber.App
		ok  bool
	)
	if eng, ok = app.(*fiber.App); !ok {
		return errors.New("fiber.App adapter SetApp: wrong parameter")
	}

	gof.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (gof *Gofiber) AddHandler(method, path string, handlers context.Handlers) {

	gof.app.Add(strings.ToUpper(method), path, func(c *fiber.Ctx) error {

		httpreq := convertCtx(c.Context())
		ctx := context.NewContext(httpreq)

		for _, key := range c.Route().Params {
			if httpreq.URL.RawQuery == "" {
				httpreq.URL.RawQuery += strings.ReplaceAll(key, ":", "") + "=" + c.Params(key)
			} else {
				httpreq.URL.RawQuery += "&" + strings.ReplaceAll(key, ":", "") + "=" + c.Params(key)
			}

		}

		ctx.SetHandlers(handlers).Next()
		for key, head := range ctx.Response.Header {
			c.Set(key, head[0])

		}

		return c.Status(ctx.Response.StatusCode).SendStream(ctx.Response.Body)

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
func (gof *Gofiber) Name() string {
	return "gofiber"
}

// SetContext implements the method Adapter.SetContext.
func (gof *Gofiber) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *fiber.Ctx
		ok  bool
	)
	if ctx, ok = contextInterface.(*fiber.Ctx); !ok {
		panic("gofiber adapter SetContext: wrong parameter")
	}
	return &Gofiber{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (gof *Gofiber) Redirect() {
	_ = gof.ctx.Redirect(config.Url(config.GetLoginUrl()), http.StatusFound)
}

// SetContentType implements the method Adapter.SetContentType.
func (gof *Gofiber) SetContentType() {
	gof.ctx.Response().Header.Set("Content-Type", gof.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (gof *Gofiber) Write(body []byte) {
	_, _ = gof.ctx.Write(body)
}

// GetCookie implements the method Adapter.GetCookie.
func (gof *Gofiber) GetCookie() (string, error) {
	return string(gof.ctx.Cookies(gof.CookieKey())), nil
}

// Lang implements the method Adapter.Lang.
func (gof *Gofiber) Lang() string {
	return string(gof.ctx.Request().URI().QueryArgs().Peek("__ga_lang"))
}

// Path implements the method Adapter.Path.
func (gof *Gofiber) Path() string {
	return string(gof.ctx.Path())
}

// Method implements the method Adapter.Method.
func (gof *Gofiber) Method() string {
	return string(gof.ctx.Method())
}

// FormParam implements the method Adapter.FormParam.
func (gof *Gofiber) FormParam() url.Values {
	f, _ := gof.ctx.MultipartForm()
	if f != nil {
		return f.Value
	}
	return url.Values{}
}

// IsPjax implements the method Adapter.IsPjax.
func (gof *Gofiber) IsPjax() bool {
	return string(gof.ctx.Request().Header.Peek(constant.PjaxHeader)) == "true"
}

// Query implements the method Adapter.Query.
func (gof *Gofiber) Query() url.Values {
	queryStr := gof.ctx.Context().QueryArgs().QueryString()
	queryObj, err := url.Parse(string(queryStr))

	if err != nil {
		return url.Values{}
	}

	return queryObj.Query()
}
