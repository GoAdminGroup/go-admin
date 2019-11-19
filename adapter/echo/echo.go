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
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

// Echo structure value is an Echo GoAdmin adapter.
type Echo struct {
	adapter.BaseAdapter
}

func init() {
	engine.Register(new(Echo))
}

// Use implement WebFrameWork.Use method.
func (e *Echo) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *echo.Echo
		ok  bool
	)
	if eng, ok = router.(*echo.Echo); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Add(strings.ToUpper(req.Method), req.URL, func(c echo.Context) error {
				ctx := context.NewContext(c.Request())

				for _, key := range c.ParamNames() {
					if c.Request().URL.RawQuery == "" {
						c.Request().URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + c.Param(key)
					} else {
						c.Request().URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + c.Param(key)
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))).Next()
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
	}

	return nil
}

// Content implement WebFrameWork.Content method.
func (e *Echo) Content(contextInterface interface{}, getPanelFn types.GetPanelFn) {

	var (
		ctx echo.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(echo.Context); !ok {
		panic("wrong parameter")
	}

	sesKey, err := ctx.Cookie(e.CookieKey())

	if err != nil || sesKey == nil {
		_ = ctx.Redirect(http.StatusFound, config.Get().Url("/login"))
		return
	}

	body, authSuccess, err := e.GetContent(sesKey.Value, ctx.Path(),
		ctx.Request().Method, ctx.Request().Header.Get(constant.PjaxHeader), getPanelFn, ctx)

	if !authSuccess {
		_ = ctx.Redirect(http.StatusFound, config.Get().Url("/login"))
		return
	}

	if err != nil {
		logger.Error("Echo Content", err)
	}

	_ = ctx.Blob(http.StatusOK, e.HTMLContentType(), body)
}
