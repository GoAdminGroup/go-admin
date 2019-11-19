// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package iris

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/kataras/iris"
	"net/http"
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// Iris structure value is an Iris GoAdmin adapter.
type Iris struct {
	adapter.BaseAdapter
}

func init() {
	engine.Register(new(Iris))
}

// Use implement WebFrameWork.Use method.
func (is *Iris) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *iris.Application
		ok  bool
	)
	if eng, ok = router.(*iris.Application); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Handle(strings.ToUpper(req.Method), req.URL, func(c iris.Context) {
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

				ctx.SetHandlers(plugCopy.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))).Next()
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
	}

	return nil
}

// Content implement WebFrameWork.Content method.
func (is *Iris) Content(contextInterface interface{}, c types.GetPanelFn) {

	var (
		ctx iris.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(iris.Context); !ok {
		panic("wrong parameter")
	}

	body, authSuccess, err := is.GetContent(ctx.GetCookie(is.CookieKey()), ctx.Path(), ctx.Method(),
		ctx.GetHeader(constant.PjaxHeader), c, ctx)

	if !authSuccess {
		ctx.Redirect(config.Get().Url("/login"), http.StatusFound)
		return
	}

	if err != nil {
		logger.Error("Echo Content", err)
	}

	_, _ = ctx.Write(body)
}
