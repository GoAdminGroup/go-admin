// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package buffalo

import (
	"bytes"
	"errors"
	"net/http"
	neturl "net/url"
	"regexp"
	"strings"

	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gobuffalo/buffalo"
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

// User implements the method Adapter.User.
func (bu *Buffalo) User(ctx interface{}) (models.UserModel, bool) {
	return bu.GetUser(ctx, bu)
}

// Use implements the method Adapter.Use.
func (bu *Buffalo) Use(app interface{}, plugs []plugins.Plugin) error {
	return bu.GetUse(app, plugs, bu)
}

func (bu *Buffalo) Run() error                 { panic("not implement") }
func (bu *Buffalo) DisableLog()                { panic("not implement") }
func (bu *Buffalo) Static(prefix, path string) { panic("not implement") }

// Content implements the method Adapter.Content.
func (bu *Buffalo) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	bu.GetContent(ctx, getPanelFn, bu, btns, fn)
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

// SetApp implements the method Adapter.SetApp.
func (bu *Buffalo) SetApp(app interface{}) error {
	var (
		eng *buffalo.App
		ok  bool
	)
	if eng, ok = app.(*buffalo.App); !ok {
		return errors.New("buffalo adapter SetApp: wrong parameter")
	}
	bu.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
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
				c.Request().URL.RawQuery += strings.ReplaceAll(key, ":", "") + "=" + param[0]
			} else {
				c.Request().URL.RawQuery += "&" + strings.ReplaceAll(key, ":", "") + "=" + param[0]
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

// Name implements the method Adapter.Name.
func (bu *Buffalo) Name() string {
	return "buffalo"
}

// SetContext implements the method Adapter.SetContext.
func (bu *Buffalo) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx buffalo.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(buffalo.Context); !ok {
		panic("buffalo adapter SetContext: wrong parameter")
	}
	return &Buffalo{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (bu *Buffalo) Redirect() {
	_ = bu.ctx.Redirect(http.StatusFound, config.Url(config.GetLoginUrl()))
}

// SetContentType implements the method Adapter.SetContentType.
func (bu *Buffalo) SetContentType() {
	bu.ctx.Response().Header().Set("Content-Type", bu.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (bu *Buffalo) Write(body []byte) {
	bu.ctx.Response().WriteHeader(http.StatusOK)
	_, _ = bu.ctx.Response().Write(body)
}

// GetCookie implements the method Adapter.GetCookie.
func (bu *Buffalo) GetCookie() (string, error) {
	return bu.ctx.Cookies().Get(bu.CookieKey())
}

// Lang implements the method Adapter.Lang.
func (bu *Buffalo) Lang() string {
	return bu.ctx.Request().URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (bu *Buffalo) Path() string {
	return bu.ctx.Request().URL.Path
}

// Method implements the method Adapter.Method.
func (bu *Buffalo) Method() string {
	return bu.ctx.Request().Method
}

// FormParam implements the method Adapter.FormParam.
func (bu *Buffalo) FormParam() neturl.Values {
	_ = bu.ctx.Request().ParseMultipartForm(32 << 20)
	return bu.ctx.Request().PostForm
}

// IsPjax implements the method Adapter.IsPjax.
func (bu *Buffalo) IsPjax() bool {
	return bu.ctx.Request().Header.Get(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (bu *Buffalo) Query() neturl.Values {
	return bu.ctx.Request().URL.Query()
}
