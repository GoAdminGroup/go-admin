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
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Gin structure value is a Gin GoAdmin adapter.
type Gin struct {
	adapter.BaseAdapter
}

func init() {
	engine.Register(new(Gin))
}

// Use implement WebFrameWork.Use method.
func (gins *Gin) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *gin.Engine
		ok  bool
	)
	if eng, ok = router.(*gin.Engine); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Handle(strings.ToUpper(req.Method), req.URL, func(c *gin.Context) {
				ctx := context.NewContext(c.Request)

				for _, param := range c.Params {
					if c.Request.URL.RawQuery == "" {
						c.Request.URL.RawQuery += strings.Replace(param.Key, ":", "", -1) + "=" + param.Value
					} else {
						c.Request.URL.RawQuery += "&" + strings.Replace(param.Key, ":", "", -1) + "=" + param.Value
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(c.Request.URL.Path, strings.ToLower(c.Request.Method))).Next()
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
	}

	return nil
}

// Content implement WebFrameWork.Content method.
func (gins *Gin) Content(contextInterface interface{}, getPanelFn types.GetPanelFn) {

	var (
		ctx *gin.Context
		ok  bool
	)

	if ctx, ok = contextInterface.(*gin.Context); !ok {
		panic("wrong parameter")
	}

	sesKey, err := ctx.Cookie(gins.CookieKey())

	if err != nil || sesKey == "" {
		ctx.Redirect(http.StatusFound, config.Get().Url("/login"))
		ctx.Abort()
		return
	}

	body, authSuccess, err := gins.GetContent(sesKey, ctx.Request.URL.Path,
		ctx.Request.Method, ctx.Request.Header.Get(constant.PjaxHeader), getPanelFn, ctx)

	if !authSuccess {
		ctx.Redirect(http.StatusFound, config.Get().Url("/login"))
		ctx.Abort()
		return
	}

	if err != nil {
		logger.Error("Gin Content", err)
	}
	ctx.Data(http.StatusOK, gins.HTMLContentType(), body)
}
