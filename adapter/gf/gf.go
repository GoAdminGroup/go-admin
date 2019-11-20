// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gf

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
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"regexp"
	"strings"
)

// Gf structure value is a Gf GoAdmin adapter.
type Gf struct {
	adapter.BaseAdapter
}

func init() {
	engine.Register(new(Gf))
}

// Use implement WebFrameWork.Use method.
func (gf *Gf) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *ghttp.Server
		ok  bool
	)
	if eng, ok = router.(*ghttp.Server); !ok {
		return errors.New("wrong parameter")
	}

	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			pluginReqUrl := req.URL
			eng.BindHandler(strings.ToUpper(req.Method)+":"+req.URL, func(c *ghttp.Request) {
				ctx := context.NewContext(c.Request)

				params := reg1.FindAllString(pluginReqUrl, -1)
				params = append(params, reg2.FindAllString(pluginReqUrl, -1)...)

				for _, param := range params {
					p := strings.Replace(param, ":", "", -1)
					if c.Request.URL.RawQuery == "" {
						c.Request.URL.RawQuery += p + "=" + c.GetRequestString(p)
					} else {
						c.Request.URL.RawQuery += "&" + p + "=" + c.GetRequestString(p)
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(c.Request.URL.Path, strings.ToLower(c.Request.Method))).Next()
				for key, head := range ctx.Response.Header {
					c.Response.Header().Add(key, head[0])
				}

				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
					c.Response.WriteStatus(ctx.Response.StatusCode, buf.Bytes())
				} else {
					c.Response.WriteStatus(ctx.Response.StatusCode)
				}
			})
		}
	}

	return nil
}

// Content implement WebFrameWork.Content method.
func (gf *Gf) Content(contextInterface interface{}, getPanelFn types.GetPanelFn) {

	var (
		ctx *ghttp.Request
		ok  bool
	)

	if ctx, ok = contextInterface.(*ghttp.Request); !ok {
		panic("wrong parameter")
	}

	body, authSuccess, err := gf.GetContent(ctx.Cookie.Get(gf.CookieKey()), ctx.URL.Path,
		ctx.Method, ctx.Header.Get(constant.PjaxHeader), getPanelFn, ctx)

	if !authSuccess {
		ctx.Response.RedirectTo(config.Get().Url("/login"))
		return
	}

	if err != nil {
		logger.Error("Gf Content", err)
	}
	ctx.Response.Header().Add("Content-Type", gf.HTMLContentType())
	ctx.Response.WriteStatus(http.StatusOK, body)
}
