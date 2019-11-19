// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package beego

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/adapter"
	gctx "github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

// Beego structure value is a Beego GoAdmin adapter.
type Beego struct {
	adapter.BaseAdapter
}

func init() {
	engine.Register(new(Beego))
}

// Use implement WebFrameWork.Use method.
func (bee *Beego) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *beego.App
		ok  bool
	)
	if eng, ok = router.(*beego.App); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Handlers.AddMethod(req.Method, req.URL, func(c *context.Context) {
				for key, value := range c.Input.Params() {
					if c.Request.URL.RawQuery == "" {
						c.Request.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
					} else {
						c.Request.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
					}
				}
				ctx := gctx.NewContext(c.Request)
				ctx.SetHandlers(plugCopy.GetHandler(c.Request.URL.Path, strings.ToLower(c.Request.Method))).Next()
				for key, head := range ctx.Response.Header {
					c.ResponseWriter.Header().Add(key, head[0])
				}
				c.ResponseWriter.WriteHeader(ctx.Response.StatusCode)
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
					c.WriteString(buf.String())
				}
			})
		}
	}

	return nil
}

// Content implement WebFrameWork.Content method.
func (bee *Beego) Content(contextInterface interface{}, getPanelFn types.GetPanelFn) {

	var (
		ctx *context.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(*context.Context); !ok {
		panic("wrong parameter")
	}

	body, authSuccess, err := bee.GetContent(ctx.GetCookie(bee.CookieKey()), ctx.Request.URL.Path,
		ctx.Request.Method, ctx.Request.Header.Get(constant.PjaxHeader), getPanelFn, ctx)

	if !authSuccess {
		ctx.Redirect(http.StatusFound, config.Get().Url("/login"))
		return
	}

	if err != nil {
		logger.Error("Beego Content", err)
	}
	ctx.ResponseWriter.Header().Set("Content-Type", bee.HTMLContentType())
	_, _ = ctx.ResponseWriter.Write(body)
}
