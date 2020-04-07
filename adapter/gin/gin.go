// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gin

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
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

// Gin structure value is a Gin GoAdmin adapter.
type Gin struct {
	adapter.BaseAdapter
	ctx *gin.Context
	app *gin.Engine
}

func init() {
	engine.Register(new(Gin))
}

func (gins *Gin) User(ci interface{}) (models.UserModel, bool) {
	return gins.GetUser(ci, gins)
}

func (gins *Gin) Use(router interface{}, plugs []plugins.Plugin) error {
	return gins.GetUse(router, plugs, gins)
}

func (gins *Gin) Content(ctx interface{}, getPanelFn types.GetPanelFn, btns ...types.Button) {
	gins.GetContent(ctx, getPanelFn, gins, btns)
}

type HandlerFunc func(ctx *gin.Context) (types.Panel, error)

func Content(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(*gin.Context))
		})
	}
}

func (gins *Gin) SetApp(app interface{}) error {
	var (
		eng *gin.Engine
		ok  bool
	)
	if eng, ok = app.(*gin.Engine); !ok {
		return errors.New("wrong parameter")
	}
	gins.app = eng
	return nil
}

func (gins *Gin) AddHandler(method, path string, handlers context.Handlers) {
	gins.app.Handle(strings.ToUpper(method), path, func(c *gin.Context) {
		ctx := context.NewContext(c.Request)

		for _, param := range c.Params {
			if c.Request.URL.RawQuery == "" {
				c.Request.URL.RawQuery += strings.Replace(param.Key, ":", "", -1) + "=" + param.Value
			} else {
				c.Request.URL.RawQuery += "&" + strings.Replace(param.Key, ":", "", -1) + "=" + param.Value
			}
		}

		ctx.SetHandlers(handlers).Next()
		for key, head := range ctx.Response.Header {
			c.Header(key, head[0])
		}
		if ctx.Response.Body != nil {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(ctx.Response.Body)
			c.String(ctx.Response.StatusCode, buf.String())
		} else {
			c.Status(ctx.Response.StatusCode)
		}
	})
}

func (gins *Gin) Name() string {
	return "gin"
}

func (gins *Gin) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *gin.Context
		ok  bool
	)

	if ctx, ok = contextInterface.(*gin.Context); !ok {
		panic("wrong parameter")
	}

	return &Gin{ctx: ctx}
}

func (gins *Gin) Redirect() {
	gins.ctx.Redirect(http.StatusFound, config.Url("/login"))
	gins.ctx.Abort()
}

func (gins *Gin) SetContentType() {
	return
}

func (gins *Gin) Write(body []byte) {
	gins.ctx.Data(http.StatusOK, gins.HTMLContentType(), body)
}

func (gins *Gin) GetCookie() (string, error) {
	return gins.ctx.Cookie(gins.CookieKey())
}

func (gins *Gin) Path() string {
	return gins.ctx.Request.URL.Path
}

func (gins *Gin) Method() string {
	return gins.ctx.Request.Method
}

func (gins *Gin) FormParam() url.Values {
	_ = gins.ctx.Request.ParseMultipartForm(32 << 20)
	return gins.ctx.Request.PostForm
}

func (gins *Gin) IsPjax() bool {
	return gins.ctx.Request.Header.Get(constant.PjaxHeader) == "true"
}
