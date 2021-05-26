// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package iris

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/kataras/iris/v12"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
)

// Iris structure value is an Iris GoAdmin adapter.
type Iris struct {
	adapter.BaseAdapter
	ctx iris.Context
	app *iris.Application
}

func init() {
	engine.Register(new(Iris))
}

// User implements the method Adapter.User.
func (is *Iris) User(ctx interface{}) (models.UserModel, bool) {
	return is.GetUser(ctx, is)
}

// Use implements the method Adapter.Use.
func (is *Iris) Use(app interface{}, plugs []plugins.Plugin) error {
	return is.GetUse(app, plugs, is)
}

func (is *Iris) Run() error                 { panic("not implement") }
func (is *Iris) DisableLog()                { panic("not implement") }
func (is *Iris) Static(prefix, path string) { panic("not implement") }

// Content implements the method Adapter.Content.
func (is *Iris) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	is.GetContent(ctx, getPanelFn, is, btns, fn)
}

type HandlerFunc func(ctx iris.Context) (types.Panel, error)

func Content(handler HandlerFunc) iris.Handler {
	return func(ctx iris.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(iris.Context))
		})
	}
}

// SetApp implements the method Adapter.SetApp.
func (is *Iris) SetApp(app interface{}) error {
	var (
		eng *iris.Application
		ok  bool
	)
	if eng, ok = app.(*iris.Application); !ok {
		return errors.New("iris adapter SetApp: wrong parameter")
	}
	is.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (is *Iris) AddHandler(method, path string, handlers context.Handlers) {
	is.app.Handle(strings.ToUpper(method), path, func(c iris.Context) {
		ctx := context.NewContext(c.Request())

		var params = map[string]string{}
		c.Params().Visit(func(key string, value string) {
			params[key] = value
		})

		for key, value := range params {
			if c.Request().URL.RawQuery == "" {
				c.Request().URL.RawQuery += strings.ReplaceAll(key, ":", "") + "=" + value
			} else {
				c.Request().URL.RawQuery += "&" + strings.ReplaceAll(key, ":", "") + "=" + value
			}
		}

		ctx.SetHandlers(handlers).Next()
		for key, head := range ctx.Response.Header {
			c.Header(key, head[0])
		}
		c.StatusCode(ctx.Response.StatusCode)
		if ctx.Response.Body != nil {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(ctx.Response.Body)
			_, _ = c.WriteString(buf.String())
		}
	})
}

// Name implements the method Adapter.Name.
func (is *Iris) Name() string {
	return "iris"
}

// SetContext implements the method Adapter.SetContext.
func (is *Iris) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx iris.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(iris.Context); !ok {
		panic("iris adapter SetContext: wrong parameter")
	}

	return &Iris{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (is *Iris) Redirect() {
	is.ctx.Redirect(config.Url(config.GetLoginUrl()), http.StatusFound)
}

// SetContentType implements the method Adapter.SetContentType.
func (is *Iris) SetContentType() {
	is.ctx.Header("Content-Type", is.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (is *Iris) Write(body []byte) {
	_, _ = is.ctx.Write(body)
}

// GetCookie implements the method Adapter.GetCookie.
func (is *Iris) GetCookie() (string, error) {
	return is.ctx.GetCookie(is.CookieKey()), nil
}

// Lang implements the method Adapter.Lang.
func (is *Iris) Lang() string {
	return is.ctx.Request().URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (is *Iris) Path() string {
	return is.ctx.Path()
}

// Method implements the method Adapter.Method.
func (is *Iris) Method() string {
	return is.ctx.Method()
}

// FormParam implements the method Adapter.FormParam.
func (is *Iris) FormParam() url.Values {
	return is.ctx.FormValues()
}

// IsPjax implements the method Adapter.IsPjax.
func (is *Iris) IsPjax() bool {
	return is.ctx.GetHeader(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (is *Iris) Query() url.Values {
	return is.ctx.Request().URL.Query()
}
