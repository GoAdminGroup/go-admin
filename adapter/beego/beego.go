// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package beego

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/GoAdminGroup/go-admin/adapter"
	gctx "github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// Beego structure value is a Beego GoAdmin adapter.
type Beego struct {
	adapter.BaseAdapter
	ctx *context.Context
	app *beego.App
}

func init() {
	engine.Register(new(Beego))
}

// User implements the method Adapter.User.
func (bee *Beego) User(ctx interface{}) (models.UserModel, bool) {
	return bee.GetUser(ctx, bee)
}

// Use implements the method Adapter.Use.
func (bee *Beego) Use(app interface{}, plugs []plugins.Plugin) error {
	return bee.GetUse(app, plugs, bee)
}

func (bee *Beego) Run() error                 { panic("not implement") }
func (bee *Beego) DisableLog()                { panic("not implement") }
func (bee *Beego) Static(prefix, path string) { panic("not implement") }

// Content implements the method Adapter.Content.
func (bee *Beego) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn gctx.NodeProcessor, navButtons ...types.Button) {
	bee.GetContent(ctx, getPanelFn, bee, navButtons, fn)
}

type HandlerFunc func(ctx *context.Context) (types.Panel, error)

func Content(handler HandlerFunc) beego.FilterFunc {
	return func(ctx *context.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(*context.Context))
		})
	}
}

// SetApp implements the method Adapter.SetApp.
func (bee *Beego) SetApp(app interface{}) error {
	var (
		eng *beego.App
		ok  bool
	)
	if eng, ok = app.(*beego.App); !ok {
		return errors.New("beego adapter SetApp: wrong parameter")
	}
	bee.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (bee *Beego) AddHandler(method, path string, handlers gctx.Handlers) {
	bee.app.Handlers.AddMethod(method, path, func(c *context.Context) {
		for key, value := range c.Input.Params() {
			if c.Request.URL.RawQuery == "" {
				c.Request.URL.RawQuery += strings.ReplaceAll(key, ":", "") + "=" + value
			} else {
				c.Request.URL.RawQuery += "&" + strings.ReplaceAll(key, ":", "") + "=" + value
			}
		}
		ctx := gctx.NewContext(c.Request)
		ctx.SetHandlers(handlers).Next()
		for key, head := range ctx.Response.Header {
			c.ResponseWriter.Header().Add(key, head[0])
		}
		c.ResponseWriter.WriteHeader(ctx.Response.StatusCode)
		if ctx.Response.Body != nil {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(ctx.Response.Body)
			c.WriteString(buf.String())
		}
	})
}

// Name implements the method Adapter.Name.
func (bee *Beego) Name() string {
	return "beego"
}

// SetContext implements the method Adapter.SetContext.
func (bee *Beego) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *context.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(*context.Context); !ok {
		panic("beego adapter SetContext: wrong parameter")
	}
	return &Beego{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (bee *Beego) Redirect() {
	bee.ctx.Redirect(http.StatusFound, config.Url(config.GetLoginUrl()))
}

// SetContentType implements the method Adapter.SetContentType.
func (bee *Beego) SetContentType() {
	bee.ctx.ResponseWriter.Header().Set("Content-Type", bee.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (bee *Beego) Write(body []byte) {
	_, _ = bee.ctx.ResponseWriter.Write(body)
}

// GetCookie implements the method Adapter.GetCookie.
func (bee *Beego) GetCookie() (string, error) {
	return bee.ctx.GetCookie(bee.CookieKey()), nil
}

// Lang implements the method Adapter.Lang.
func (bee *Beego) Lang() string {
	return bee.ctx.Request.URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (bee *Beego) Path() string {
	return bee.ctx.Request.URL.Path
}

// Method implements the method Adapter.Method.
func (bee *Beego) Method() string {
	return bee.ctx.Request.Method
}

// FormParam implements the method Adapter.FormParam.
func (bee *Beego) FormParam() url.Values {
	_ = bee.ctx.Request.ParseMultipartForm(32 << 20)
	return bee.ctx.Request.PostForm
}

// IsPjax implements the method Adapter.IsPjax.
func (bee *Beego) IsPjax() bool {
	return bee.ctx.Request.Header.Get(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (bee *Beego) Query() url.Values {
	return bee.ctx.Request.URL.Query()
}
