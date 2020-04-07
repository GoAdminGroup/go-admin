// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package echo

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
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"strings"
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

func (e *Echo) User(ci interface{}) (models.UserModel, bool) {
	return e.GetUser(ci, e)
}

func (e *Echo) Use(router interface{}, plugs []plugins.Plugin) error {
	return e.GetUse(router, plugs, e)
}

func (e *Echo) Content(ctx interface{}, getPanelFn types.GetPanelFn, btns ...types.Button) {
	e.GetContent(ctx, getPanelFn, e, btns)
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

func (e *Echo) SetApp(app interface{}) error {
	var (
		eng *echo.Echo
		ok  bool
	)
	if eng, ok = app.(*echo.Echo); !ok {
		return errors.New("wrong parameter")
	}
	e.app = eng
	return nil
}

func (e *Echo) AddHandler(method, path string, handlers context.Handlers) {
	e.app.Add(strings.ToUpper(method), path, func(c echo.Context) error {
		ctx := context.NewContext(c.Request())

		for _, key := range c.ParamNames() {
			if c.Request().URL.RawQuery == "" {
				c.Request().URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + c.Param(key)
			} else {
				c.Request().URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + c.Param(key)
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
		}
		return nil
	})
}

func (e *Echo) Name() string {
	return "echo"
}

func (e *Echo) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx echo.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(echo.Context); !ok {
		panic("wrong parameter")
	}
	return &Echo{ctx: ctx}
}

func (e *Echo) Redirect() {
	_ = e.ctx.Redirect(http.StatusFound, config.Url("/login"))
}

func (e *Echo) SetContentType() {
	e.ctx.Response().Header().Set("Content-Type", e.HTMLContentType())
}

func (e *Echo) Write(body []byte) {
	e.ctx.Response().WriteHeader(http.StatusOK)
	_, _ = e.ctx.Response().Write(body)
}

func (e *Echo) GetCookie() (string, error) {
	cookie, err := e.ctx.Cookie(e.CookieKey())
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

func (e *Echo) Path() string {
	return e.ctx.Request().URL.Path
}

func (e *Echo) Method() string {
	return e.ctx.Request().Method
}

func (e *Echo) FormParam() url.Values {
	_ = e.ctx.Request().ParseMultipartForm(32 << 20)
	return e.ctx.Request().PostForm
}

func (e *Echo) IsPjax() bool {
	return e.ctx.Request().Header.Get(constant.PjaxHeader) == "true"
}
