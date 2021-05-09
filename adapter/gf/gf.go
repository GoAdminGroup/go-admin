// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gf

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
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gogf/gf/net/ghttp"
)

// Gf structure value is a Gf GoAdmin adapter.
type Gf struct {
	adapter.BaseAdapter
	ctx *ghttp.Request
	app *ghttp.Server
}

func init() {
	engine.Register(new(Gf))
}

// User implements the method Adapter.User.
func (gf *Gf) User(ctx interface{}) (models.UserModel, bool) {
	return gf.GetUser(ctx, gf)
}

// Use implements the method Adapter.Use.
func (gf *Gf) Use(app interface{}, plugs []plugins.Plugin) error {
	return gf.GetUse(app, plugs, gf)
}

func (gf *Gf) Run() error                 { panic("not implement") }
func (gf *Gf) DisableLog()                { panic("not implement") }
func (gf *Gf) Static(prefix, path string) { panic("not implement") }

// Content implements the method Adapter.Content.
func (gf *Gf) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	gf.GetContent(ctx, getPanelFn, gf, btns, fn)
}

type HandlerFunc func(ctx *ghttp.Request) (types.Panel, error)

func Content(handler HandlerFunc) ghttp.HandlerFunc {
	return func(ctx *ghttp.Request) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(*ghttp.Request))
		})
	}
}

// SetApp implements the method Adapter.SetApp.
func (gf *Gf) SetApp(app interface{}) error {
	var (
		eng *ghttp.Server
		ok  bool
	)
	if eng, ok = app.(*ghttp.Server); !ok {
		return errors.New("gf adapter SetApp: wrong parameter")
	}
	gf.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (gf *Gf) AddHandler(method, path string, handlers context.Handlers) {
	gf.app.BindHandler(strings.ToUpper(method)+":"+path, func(c *ghttp.Request) {
		ctx := context.NewContext(c.Request)

		newPath := path

		reg1 := regexp.MustCompile(":(.*?)/")
		reg2 := regexp.MustCompile(":(.*?)$")

		params := reg1.FindAllString(newPath, -1)
		newPath = reg1.ReplaceAllString(newPath, "")
		params = append(params, reg2.FindAllString(newPath, -1)...)

		for _, param := range params {
			p := utils.ReplaceAll(param, ":", "", "/", "")

			if c.Request.URL.RawQuery == "" {
				c.Request.URL.RawQuery += p + "=" + c.GetRequestString(p)
			} else {
				c.Request.URL.RawQuery += "&" + p + "=" + c.GetRequestString(p)
			}
		}

		ctx.SetHandlers(handlers).Next()
		for key, head := range ctx.Response.Header {
			c.Response.Header().Add(key, head[0])
		}

		if ctx.Response.Body != nil {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(ctx.Response.Body)
			c.Response.WriteStatus(ctx.Response.StatusCode, buf.Bytes())
		} else {
			c.Response.WriteStatus(ctx.Response.StatusCode)
		}
	})
}

// Name implements the method Adapter.Name.
func (gf *Gf) Name() string {
	return "gf"
}

// SetContext implements the method Adapter.SetContext.
func (gf *Gf) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *ghttp.Request
		ok  bool
	)

	if ctx, ok = contextInterface.(*ghttp.Request); !ok {
		panic("gf adapter SetContext: wrong parameter")
	}
	return &Gf{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (gf *Gf) Redirect() {
	gf.ctx.Response.RedirectTo(config.Url(config.GetLoginUrl()))
}

// SetContentType implements the method Adapter.SetContentType.
func (gf *Gf) SetContentType() {
	gf.ctx.Response.Header().Add("Content-Type", gf.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (gf *Gf) Write(body []byte) {
	gf.ctx.Response.WriteStatus(http.StatusOK, body)
}

// GetCookie implements the method Adapter.GetCookie.
func (gf *Gf) GetCookie() (string, error) {
	return gf.ctx.Cookie.Get(gf.CookieKey()), nil
}

// Lang implements the method Adapter.Lang.
func (gf *Gf) Lang() string {
	return gf.ctx.Request.URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (gf *Gf) Path() string {
	return gf.ctx.URL.Path
}

// Method implements the method Adapter.Method.
func (gf *Gf) Method() string {
	return gf.ctx.Method
}

// FormParam implements the method Adapter.FormParam.
func (gf *Gf) FormParam() url.Values {
	return gf.ctx.Form
}

// IsPjax implements the method Adapter.IsPjax.
func (gf *Gf) IsPjax() bool {
	return gf.ctx.Header.Get(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (gf *Gf) Query() url.Values {
	return gf.ctx.Request.URL.Query()
}
