// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gin

import (
	"bytes"
	"errors"
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
	"github.com/gin-gonic/gin"
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

// User implements the method Adapter.User.
func (gins *Gin) User(ctx interface{}) (models.UserModel, bool) {
	return gins.GetUser(ctx, gins)
}

// Use implements the method Adapter.Use.
func (gins *Gin) Use(app interface{}, plugs []plugins.Plugin) error {
	return gins.GetUse(app, plugs, gins)
}

// Content implements the method Adapter.Content.
func (gins *Gin) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	gins.GetContent(ctx, getPanelFn, gins, btns, fn)
}

type HandlerFunc func(ctx *gin.Context) (types.Panel, error)

func Content(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(*gin.Context))
		})
	}
}

func (gins *Gin) Run() error                 { panic("not implement") }
func (gins *Gin) DisableLog()                { panic("not implement") }
func (gins *Gin) Static(prefix, path string) { panic("not implement") }

// SetApp implements the method Adapter.SetApp.
func (gins *Gin) SetApp(app interface{}) error {
	var (
		eng *gin.Engine
		ok  bool
	)
	if eng, ok = app.(*gin.Engine); !ok {
		return errors.New("gin adapter SetApp: wrong parameter")
	}
	gins.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (gins *Gin) AddHandler(method, path string, handlers context.Handlers) {
	gins.app.Handle(strings.ToUpper(method), path, func(c *gin.Context) {
		ctx := context.NewContext(c.Request)

		for _, param := range c.Params {
			if c.Request.URL.RawQuery == "" {
				c.Request.URL.RawQuery += strings.ReplaceAll(param.Key, ":", "") + "=" + param.Value
			} else {
				c.Request.URL.RawQuery += "&" + strings.ReplaceAll(param.Key, ":", "") + "=" + param.Value
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

// Name implements the method Adapter.Name.
func (gins *Gin) Name() string {
	return "gin"
}

// SetContext implements the method Adapter.SetContext.
func (gins *Gin) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *gin.Context
		ok  bool
	)

	if ctx, ok = contextInterface.(*gin.Context); !ok {
		panic("gin adapter SetContext: wrong parameter")
	}

	return &Gin{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (gins *Gin) Redirect() {
	gins.ctx.Redirect(http.StatusFound, config.Url(config.GetLoginUrl()))
	gins.ctx.Abort()
}

// SetContentType implements the method Adapter.SetContentType.
func (gins *Gin) SetContentType() {
}

// Write implements the method Adapter.Write.
func (gins *Gin) Write(body []byte) {
	gins.ctx.Data(http.StatusOK, gins.HTMLContentType(), body)
}

// GetCookie implements the method Adapter.GetCookie.
func (gins *Gin) GetCookie() (string, error) {
	return gins.ctx.Cookie(gins.CookieKey())
}

// Lang implements the method Adapter.Lang.
func (gins *Gin) Lang() string {
	return gins.ctx.Request.URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (gins *Gin) Path() string {
	return gins.ctx.Request.URL.Path
}

// Method implements the method Adapter.Method.
func (gins *Gin) Method() string {
	return gins.ctx.Request.Method
}

// FormParam implements the method Adapter.FormParam.
func (gins *Gin) FormParam() url.Values {
	_ = gins.ctx.Request.ParseMultipartForm(32 << 20)
	return gins.ctx.Request.PostForm
}

// IsPjax implements the method Adapter.IsPjax.
func (gins *Gin) IsPjax() bool {
	return gins.ctx.Request.Header.Get(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (gins *Gin) Query() url.Values {
	return gins.ctx.Request.URL.Query()
}
