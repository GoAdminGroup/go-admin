/***
# File Name: ../../adapter/gear/gear.go
# Author: eavesmy
# Email: eavesmy@gmail.com
# Created Time: 2021年06月03日 星期四 19时05分06秒
***/

package gear

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/teambition/gear"
)

// Gear structure value is a Gin GoAdmin adapter.
type Gear struct {
	adapter.BaseAdapter
	ctx    *gear.Context
	app    *gear.App
	router *gear.Router
}

func init() {
	engine.Register(new(Gear))
}

// User implements the method Adapter.User.
func (gears *Gear) User(ctx interface{}) (models.UserModel, bool) {
	return gears.GetUser(ctx, gears)
}

// Use implements the method Adapter.Use.
func (gears *Gear) Use(app interface{}, plugs []plugins.Plugin) error {
	return gears.GetUse(app, plugs, gears)
}

// Content implements the method Adapter.Content.
func (gears *Gear) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	gears.GetContent(ctx, getPanelFn, gears, btns, fn)
}

type HandlerFunc func(ctx *gear.Context) (types.Panel, error)

func Content(handler HandlerFunc) gear.Middleware {
	return func(ctx *gear.Context) error {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(*gear.Context))
		})
		return nil
	}
}

// SetApp implements the method Adapter.SetApp.
func (gears *Gear) SetApp(app interface{}) error {
	gears.app = app.(*gear.App)
	gears.router = gear.NewRouter()
	var (
		eng *gear.App
		ok  bool
	)
	if eng, ok = app.(*gear.App); !ok {
		return errors.New("beego adapter SetApp: wrong parameter")
	}
	gears.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (gears *Gear) AddHandler(method, path string, handlers context.Handlers) {

	if gears.router == nil {
		gears.router = gear.NewRouter()
	}

	gears.router.Handle(strings.ToUpper(method), path, func(c *gear.Context) error {

		ctx := context.NewContext(c.Req)

		newPath := path

		reg1 := regexp.MustCompile(":(.*?)/")
		reg2 := regexp.MustCompile(":(.*?)$")

		params := reg1.FindAllString(newPath, -1)
		newPath = reg1.ReplaceAllString(newPath, "")
		params = append(params, reg2.FindAllString(newPath, -1)...)

		for _, param := range params {
			p := utils.ReplaceAll(param, ":", "", "/", "")

			if c.Req.URL.RawQuery == "" {
				c.Req.URL.RawQuery += p + "=" + c.Param(p)
			} else {
				c.Req.URL.RawQuery += "&" + p + "=" + c.Param(p)
			}
		}

		ctx.SetHandlers(handlers).Next()

		for key, head := range ctx.Response.Header {
			c.Res.Header().Add(key, head[0])
		}

		if ctx.Response.Body != nil {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(ctx.Response.Body)

			return c.End(ctx.Response.StatusCode, buf.Bytes())
		}

		c.Status(ctx.Response.StatusCode)
		return nil
	})

	gears.app.UseHandler(gears.router)
}

// Name implements the method Adapter.Name.
func (*Gear) Name() string {
	return "gear"
}

// SetContext implements the method Adapter.SetContext.
func (*Gear) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *gear.Context
		ok  bool
	)

	if ctx, ok = contextInterface.(*gear.Context); !ok {
		panic("gear adapter SetContext: wrong parameter")
	}

	return &Gear{ctx: ctx}
}

// Redirect implements the method Adapter.Redirect.
func (gears *Gear) Redirect() {
	gears.ctx.Redirect(config.Url(config.GetLoginUrl()))
}

// SetContentType implements the method Adapter.SetContentType.
func (gears *Gear) SetContentType() {
	gears.ctx.Res.Header().Set("Content-Type", gears.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (gears *Gear) Write(body []byte) {
	gears.ctx.End(http.StatusOK, body)
}

// GetCookie implements the method Adapter.GetCookie.
func (gears *Gear) GetCookie() (string, error) {
	return gears.ctx.Cookies.Get(gears.CookieKey())
}

// Lang implements the method Adapter.Lang.
func (gears *Gear) Lang() string {
	return gears.ctx.Req.URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (gears *Gear) Path() string {
	return gears.ctx.Req.URL.Path
}

// Method implements the method Adapter.Method.
func (gears *Gear) Method() string {
	return gears.ctx.Req.Method
}

// FormParam implements the method Adapter.FormParam.
func (gears *Gear) FormParam() url.Values {
	_ = gears.ctx.Req.ParseMultipartForm(32 << 20)
	return gears.ctx.Req.PostForm
}

// IsPjax implements the method Adapter.IsPjax.
func (gears *Gear) IsPjax() bool {
	return gears.ctx.Req.Header.Get(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (gears *Gear) Query() url.Values {
	return gears.ctx.Req.URL.Query()
}

// Request implements the method Adapter.Request.
func (gears *Gear) Request() *http.Request {
	return gears.ctx.Req
}