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
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Gf structure value is a Gf GoAdmin adapter.
type Gf struct {
	adapter.BaseAdapter
	ctx *ghttp.Request
	app *ghttp.Server
}

func init() {
	engine.Register(new(Gf))
}

func (gf *Gf) User(ci interface{}) (models.UserModel, bool) {
	return gf.GetUser(ci, gf)
}

func (gf *Gf) Use(router interface{}, plugs []plugins.Plugin) error {
	return gf.GetUse(router, plugs, gf)
}

func (gf *Gf) Content(ctx interface{}, getPanelFn types.GetPanelFn, btns ...types.Button) {
	gf.GetContent(ctx, getPanelFn, gf, btns)
}

type HandlerFunc func(ctx *ghttp.Request) (types.Panel, error)

func Content(handler HandlerFunc) ghttp.HandlerFunc {
	return func(ctx *ghttp.Request) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(*ghttp.Request))
		})
	}
}

func (gf *Gf) SetApp(app interface{}) error {
	var (
		eng *ghttp.Server
		ok  bool
	)
	if eng, ok = app.(*ghttp.Server); !ok {
		return errors.New("wrong parameter")
	}
	gf.app = eng
	return nil
}

func (gf *Gf) AddHandler(method, path string, handlers context.Handlers) {
	gf.app.BindHandler(strings.ToUpper(method)+":"+path, func(c *ghttp.Request) {
		ctx := context.NewContext(c.Request)

		newPath := path

		reg1 := regexp.MustCompile(":(.*?)/")
		reg2 := regexp.MustCompile(":(.*?)$")

		params := reg1.FindAllString(newPath, -1)
		newPath = reg1.ReplaceAllString(newPath, "")
		params = append(params, reg2.FindAllString(newPath, -1)...)

		for _, param := range params {
			p := strings.Replace(param, ":", "", -1)
			p = strings.Replace(p, "/", "", -1)

			if c.Request.URL.RawQuery == "" {
				c.Request.URL.RawQuery += p + "=" + c.GetRequestString(p)
			} else {
				c.Request.URL.RawQuery += "&" + p + "=" + c.GetRequestString(p)
			}
		}

		ctx.SetHandlers(handlers).Next()
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

func (gf *Gf) Name() string {
	return "gf"
}

func (gf *Gf) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *ghttp.Request
		ok  bool
	)

	if ctx, ok = contextInterface.(*ghttp.Request); !ok {
		panic("wrong parameter")
	}
	return &Gf{ctx: ctx}
}

func (gf *Gf) Redirect() {
	gf.ctx.Response.RedirectTo(config.Url("/login"))
}

func (gf *Gf) SetContentType() {
	gf.ctx.Response.Header().Add("Content-Type", gf.HTMLContentType())
}

func (gf *Gf) Write(body []byte) {
	gf.ctx.Response.WriteStatus(http.StatusOK, body)
}

func (gf *Gf) GetCookie() (string, error) {
	return gf.ctx.Cookie.Get(gf.CookieKey()), nil
}

func (gf *Gf) Path() string {
	return gf.ctx.URL.Path
}

func (gf *Gf) Method() string {
	return gf.ctx.Method
}

func (gf *Gf) FormParam() url.Values {
	return gf.ctx.Form
}

func (gf *Gf) IsPjax() bool {
	return gf.ctx.Header.Get(constant.PjaxHeader) == "true"
}
