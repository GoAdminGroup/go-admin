// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package iris

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/kataras/iris/v12"
	"net/http"
	"net/url"
	"strings"

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

func (is *Iris) User(ci interface{}) (models.UserModel, bool) {
	return is.GetUser(ci, is)
}

func (is *Iris) Use(router interface{}, plugs []plugins.Plugin) error {
	return is.GetUse(router, plugs, is)
}

func (is *Iris) Content(ctx interface{}, getPanelFn types.GetPanelFn, btns ...types.Button) {
	is.GetContent(ctx, getPanelFn, is, btns)
}

type HandlerFunc func(ctx iris.Context) (types.Panel, error)

func Content(handler HandlerFunc) iris.Handler {
	return func(ctx iris.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(iris.Context))
		})
	}
}

func (is *Iris) SetApp(app interface{}) error {
	var (
		eng *iris.Application
		ok  bool
	)
	if eng, ok = app.(*iris.Application); !ok {
		return errors.New("wrong parameter")
	}
	is.app = eng
	return nil
}

func (is *Iris) AddHandler(method, path string, handlers context.Handlers) {
	is.app.Handle(strings.ToUpper(method), path, func(c iris.Context) {
		ctx := context.NewContext(c.Request())

		var params = map[string]string{}
		c.Params().Visit(func(key string, value string) {
			params[key] = value
		})

		for key, value := range params {
			if c.Request().URL.RawQuery == "" {
				c.Request().URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
			} else {
				c.Request().URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
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

func (is *Iris) Name() string {
	return "iris"
}

func (is *Iris) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx iris.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(iris.Context); !ok {
		panic("wrong parameter")
	}

	return &Iris{ctx: ctx}
}

func (is *Iris) Redirect() {
	is.ctx.Redirect(config.Url("/login"), http.StatusFound)
}

func (is *Iris) SetContentType() {
	is.ctx.Header("Content-Type", is.HTMLContentType())
}

func (is *Iris) Write(body []byte) {
	_, _ = is.ctx.Write(body)
}

func (is *Iris) GetCookie() (string, error) {
	return is.ctx.GetCookie(is.CookieKey()), nil
}

func (is *Iris) Path() string {
	return is.ctx.Path()
}

func (is *Iris) Method() string {
	return is.ctx.Method()
}

func (is *Iris) FormParam() url.Values {
	return is.ctx.FormValues()
}

func (is *Iris) IsPjax() bool {
	return is.ctx.GetHeader(constant.PjaxHeader) == "true"
}
