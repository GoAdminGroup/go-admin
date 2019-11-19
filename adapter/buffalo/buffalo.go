// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package buffalo

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
	"github.com/gobuffalo/buffalo"
	"net/http"
	neturl "net/url"
	"regexp"
	"strings"
)

// Buffalo structure value is a Buffalo GoAdmin adapter.
type Buffalo struct {
	adapter.BaseAdapter
}

func init() {
	engine.Register(new(Buffalo))
}

// Use implement WebFrameWork.Use method.
func (bu *Buffalo) Use(router interface{}, plugin []plugins.Plugin) error {

	var (
		eng *buffalo.App
		ok  bool
	)
	if eng, ok = router.(*buffalo.App); !ok {
		return errors.New("wrong parameter")
	}

	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {

			url := req.URL
			url = reg1.ReplaceAllString(url, "{$1}/")
			url = reg2.ReplaceAllString(url, "{$1}")

			getHandleFunc(eng, strings.ToUpper(req.Method))(url, func(c buffalo.Context) error {

				if c.Request().URL.Path[len(c.Request().URL.Path)-1] == '/' {
					c.Request().URL.Path = c.Request().URL.Path[:len(c.Request().URL.Path)-1]
				}

				ctx := context.NewContext(c.Request())

				params := c.Params().(neturl.Values)

				for key, param := range params {
					if c.Request().URL.RawQuery == "" {
						c.Request().URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + param[0]
					} else {
						c.Request().URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + param[0]
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))).Next()
				for key, head := range ctx.Response.Header {
					c.Response().Header().Set(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
					c.Response().WriteHeader(ctx.Response.StatusCode)
					_, _ = c.Response().Write(buf.Bytes())
				} else {
					c.Response().WriteHeader(ctx.Response.StatusCode)
				}
				return nil
			})
		}
	}

	return nil
}

// HandleFun is type of route methods of buffalo.
type HandleFun func(p string, h buffalo.Handler) *buffalo.RouteInfo

func getHandleFunc(eng *buffalo.App, method string) HandleFun {
	switch method {
	case "GET":
		return eng.GET
	case "POST":
		return eng.POST
	case "PUT":
		return eng.PUT
	case "DELETE":
		return eng.DELETE
	case "HEAD":
		return eng.HEAD
	case "OPTIONS":
		return eng.OPTIONS
	case "PATCH":
		return eng.PATCH
	default:
		panic("wrong method")
	}
}

// Content implement WebFrameWork.Content method.
func (bu *Buffalo) Content(contextInterface interface{}, getPanelFn types.GetPanelFn) {

	var (
		ctx buffalo.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(buffalo.Context); !ok {
		panic("wrong parameter")
	}

	sesKey, err := ctx.Cookies().Get(bu.CookieKey())

	if err != nil || sesKey == "" {
		_ = ctx.Redirect(http.StatusFound, config.Get().Url("/login"))
		return
	}

	body, authSuccess, err := bu.GetContent(sesKey, ctx.Request().URL.Path,
		ctx.Request().Method, ctx.Request().Header.Get(constant.PjaxHeader), getPanelFn, ctx)

	if !authSuccess {
		_ = ctx.Redirect(http.StatusFound, config.Get().Url("/login"))
		return
	}

	if err != nil {
		logger.Error("Buffalo Content", err)
	}

	ctx.Response().Header().Set("Content-Type", bu.HTMLContentType())
	ctx.Response().WriteHeader(http.StatusOK)
	_, _ = ctx.Response().Write(body)
}
