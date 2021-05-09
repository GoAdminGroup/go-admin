// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package echo

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
	"github.com/labstack/echo/v4"
)

// Echo structure value is an Echo GoAdmin adapter.
type Echo struct {
	adapter.BaseAdapter
	ctx echo.Context
	app *echo.Echo
}

func init() {
	engine.Register(new(Echo))
}

// User implements the method Adapter.User.
func (e *Echo) User(ctx interface{}) (models.UserModel, bool) {
	return e.GetUser(ctx, e)
}

// Use implements the method Adapter.Use.
func (e *Echo) Use(app interface{}, plugs []plugins.Plugin) error {
	return e.GetUse(app, plugs, e)
}

func (e *Echo) Run() error                 { panic("not implement") }
func (e *Echo) DisableLog()                { panic("not implement") }
func (e *Echo) Static(prefix, path string) { panic("not implement") }

// Content implements the method Adapter.Content.
func (e *Echo) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	e.GetContent(ctx, getPanelFn, e, btns, fn)
}

type HandlerFunc func(ctx echo.Context) (types.Panel, error)

func Content(handler HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(echo.Context))
		})
		return nil
	}
}

// SetApp implements the method Adapter.SetApp.
func (e *Echo) SetApp(app interface{}) error {
	var (
		eng *echo.Echo
		ok  bool
	)
	if eng, ok = app.(*echo.Echo); !ok {
		return errors.New("echo adapter SetApp: wrong parameter")
	}
	e.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (e *Echo) AddHandler(method, path string, handlers context.Handlers) {
	e.app.Add(strings.ToUpper(method), path, func(c echo.Context) error {
		ctx := context.NewContext(c.Request())

		for _, key := range c.ParamNames() {
			if c.Request().URL.RawQuery == "" {
				c.Request().URL.RawQuery += strings.ReplaceAll(key, ":", "") + "=" + c.Param(key)
			} else {
				c.Request().URL.RawQuery += "&" + strings.ReplaceAll(key, ":", "") + "=" + c.Param(key)
			}
		}

		ctx.SetHandlers(handlers).Next()
		for key, head := range ctx.Response.Header {
			c.Response().Header().Set(key, head[0])
		}
		if ctx.Response.Body != nil {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(ctx.Response.Body)
			_ = c.String(ctx.Response.StatusCode, buf.String())
		} else {
			c.Response().WriteHeader(ctx.Response.StatusCode)
		}
		return nil
	})
}

// Name implements the method Adapter.Name.
func (e *Echo) Name() string {
	return "echo"
}

// SetContext implements the method Adapter.SetContext.
func (e *Echo) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx echo.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(echo.Context); !ok {
		panic("echo adapter SetContext: wrong parameter")
	}
	return &Echo{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (e *Echo) Redirect() {
	_ = e.ctx.Redirect(http.StatusFound, config.Url(config.GetLoginUrl()))
}

// SetContentType implements the method Adapter.SetContentType.
func (e *Echo) SetContentType() {
	e.ctx.Response().Header().Set("Content-Type", e.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (e *Echo) Write(body []byte) {
	e.ctx.Response().WriteHeader(http.StatusOK)
	_, _ = e.ctx.Response().Write(body)
}

// GetCookie implements the method Adapter.GetCookie.
func (e *Echo) GetCookie() (string, error) {
	cookie, err := e.ctx.Cookie(e.CookieKey())
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

// Lang implements the method Adapter.Lang.
func (e *Echo) Lang() string {
	return e.ctx.Request().URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (e *Echo) Path() string {
	return e.ctx.Request().URL.Path
}

// Method implements the method Adapter.Method.
func (e *Echo) Method() string {
	return e.ctx.Request().Method
}

// FormParam implements the method Adapter.FormParam.
func (e *Echo) FormParam() url.Values {
	_ = e.ctx.Request().ParseMultipartForm(32 << 20)
	return e.ctx.Request().PostForm
}

// IsPjax implements the method Adapter.IsPjax.
func (e *Echo) IsPjax() bool {
	return e.ctx.Request().Header.Get(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (e *Echo) Query() url.Values {
	return e.ctx.Request().URL.Query()
}
