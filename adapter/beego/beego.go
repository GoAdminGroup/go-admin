// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package beego

import (
	"bytes"
	"errors"
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
	"net/http"
	"net/url"
	"strings"
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

func (bee *Beego) User(ci interface{}) (models.UserModel, bool) {
	return bee.GetUser(ci, bee)
}

func (bee *Beego) Use(router interface{}, plugs []plugins.Plugin) error {
	return bee.GetUse(router, plugs, bee)
}

func (bee *Beego) Content(ctx interface{}, getPanelFn types.GetPanelFn, btns ...types.Button) {
	bee.GetContent(ctx, getPanelFn, bee, btns)
}

type HandlerFunc func(ctx *context.Context) (types.Panel, error)

func Content(handler HandlerFunc) beego.FilterFunc {
	return func(ctx *context.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(*context.Context))
		})
	}
}

func (bee *Beego) SetApp(app interface{}) error {
	var (
		eng *beego.App
		ok  bool
	)
	if eng, ok = app.(*beego.App); !ok {
		return errors.New("wrong parameter")
	}
	bee.app = eng
	return nil
}

func (bee *Beego) AddHandler(method, path string, handlers gctx.Handlers) {
	bee.app.Handlers.AddMethod(method, path, func(c *context.Context) {
		for key, value := range c.Input.Params() {
			if c.Request.URL.RawQuery == "" {
				c.Request.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
			} else {
				c.Request.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
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

func (bee *Beego) Name() string {
	return "beego"
}

func (bee *Beego) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *context.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(*context.Context); !ok {
		panic("wrong parameter")
	}
	return &Beego{ctx: ctx}
}

func (bee *Beego) Redirect() {
	bee.ctx.Redirect(http.StatusFound, config.Url("/login"))
}

func (bee *Beego) SetContentType() {
	bee.ctx.ResponseWriter.Header().Set("Content-Type", bee.HTMLContentType())
}

func (bee *Beego) Write(body []byte) {
	_, _ = bee.ctx.ResponseWriter.Write(body)
}

func (bee *Beego) GetCookie() (string, error) {
	return bee.ctx.GetCookie(bee.CookieKey()), nil
}

func (bee *Beego) Path() string {
	return bee.ctx.Request.URL.Path
}

func (bee *Beego) Method() string {
	return bee.ctx.Request.Method
}

func (bee *Beego) FormParam() url.Values {
	_ = bee.ctx.Request.ParseMultipartForm(32 << 20)
	return bee.ctx.Request.PostForm
}

func (bee *Beego) IsPjax() bool {
	return bee.ctx.Request.Header.Get(constant.PjaxHeader) == "true"
}
