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
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
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
	ctx buffalo.Context
	app *buffalo.App
}

func init() {
	engine.Register(new(Buffalo))
}

func (bu *Buffalo) User(ci interface{}) (models.UserModel, bool) {
	return bu.GetUser(ci, bu)
}

func (bu *Buffalo) Use(router interface{}, plugs []plugins.Plugin) error {
	return bu.GetUse(router, plugs, bu)
}

func (bu *Buffalo) Content(ctx interface{}, getPanelFn types.GetPanelFn, btns ...types.Button) {
	bu.GetContent(ctx, getPanelFn, bu, btns)
}

type HandlerFunc func(ctx buffalo.Context) (types.Panel, error)

func Content(handler HandlerFunc) buffalo.Handler {
	return func(ctx buffalo.Context) error {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(buffalo.Context))
		})
		return nil
	}
}

func (bu *Buffalo) SetApp(app interface{}) error {
	var (
		eng *buffalo.App
		ok  bool
	)
	if eng, ok = app.(*buffalo.App); !ok {
		return errors.New("wrong parameter")
	}
	bu.app = eng
	return nil
}

func (bu *Buffalo) AddHandler(method, path string, handlers context.Handlers) {
	url := path
	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")
	url = reg1.ReplaceAllString(url, "{$1}/")
	url = reg2.ReplaceAllString(url, "{$1}")

	getHandleFunc(bu.app, strings.ToUpper(method))(url, func(c buffalo.Context) error {

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

		ctx.SetHandlers(handlers).Next()
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

func (bu *Buffalo) Name() string {
	return "buffalo"
}

func (bu *Buffalo) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx buffalo.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(buffalo.Context); !ok {
		panic("wrong parameter")
	}
	return &Buffalo{ctx: ctx}
}

func (bu *Buffalo) Redirect() {
	_ = bu.ctx.Redirect(http.StatusFound, config.Url("/login"))
}

func (bu *Buffalo) SetContentType() {
	bu.ctx.Response().Header().Set("Content-Type", bu.HTMLContentType())
}

func (bu *Buffalo) Write(body []byte) {
	bu.ctx.Response().WriteHeader(http.StatusOK)
	_, _ = bu.ctx.Response().Write(body)
}

func (bu *Buffalo) GetCookie() (string, error) {
	return bu.ctx.Cookies().Get(bu.CookieKey())
}

func (bu *Buffalo) Path() string {
	return bu.ctx.Request().URL.Path
}

func (bu *Buffalo) Method() string {
	return bu.ctx.Request().Method
}

func (bu *Buffalo) FormParam() neturl.Values {
	_ = bu.ctx.Request().ParseMultipartForm(32 << 20)
	return bu.ctx.Request().PostForm
}

func (bu *Buffalo) IsPjax() bool {
	return bu.ctx.Request().Header.Get(constant.PjaxHeader) == "true"
}
