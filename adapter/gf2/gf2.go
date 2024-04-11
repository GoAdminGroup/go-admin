package gf2

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
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gogf/gf/v2/net/ghttp"
)

type GF2 struct {
	adapter.BaseAdapter
	ctx *ghttp.Request
	app *ghttp.Server
}

func init() {
	engine.Register(new(GF2))
}

func (*GF2) Name() string {
	return "gf2"
}

func (gf2 *GF2) Use(app interface{}, plugins []plugins.Plugin) error {
	return gf2.GetUse(app, plugins, gf2)
}

func (gf2 *GF2) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	gf2.GetContent(ctx, getPanelFn, gf2, btns, fn)
}

func (gf2 *GF2) User(ctx interface{}) (models.UserModel, bool) {
	return gf2.GetUser(ctx, gf2)
}

func (gf2 *GF2) AddHandler(method, path string, handlers context.Handlers) {
	gf2.app.BindHandler(strings.ToUpper(method)+":"+path, func(c *ghttp.Request) {
		ctx := context.NewContext(c.Request)

		newPath := path

		reg1 := regexp.MustCompile(":(.*?)/")
		reg2 := regexp.MustCompile(":(.*?)$")

		params := reg1.FindAllString(newPath, -1)
		newPath = reg1.ReplaceAllString(newPath, "")
		params = append(params, reg2.FindAllString(newPath, -1)...)

		for _, param := range params {
			p := utils.ReplaceAll(param, ":", "", "/", "")

			if c.Request.URL.RawQuery == "" {
				c.Request.URL.RawQuery += p + "=" + c.GetRequest(p).String()
			} else {
				c.Request.URL.RawQuery += "&" + p + "=" + c.GetRequest(p).String()
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

func (gf2 *GF2) SetApp(app interface{}) error {
	var (
		eng *ghttp.Server
		ok  bool
	)
	if eng, ok = app.(*ghttp.Server); !ok {
		return errors.New("gf2 adapter SetApp: wrong parameter")
	}
	gf2.app = eng
	return nil
}

func (*GF2) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *ghttp.Request
		ok  bool
	)
	if ctx, ok = contextInterface.(*ghttp.Request); !ok {
		panic("gf2 adapter SetContext: wrong parameter")
	}
	return &GF2{ctx: ctx}
}

func (gf2 *GF2) GetCookie() (string, error) {
	return gf2.ctx.Cookie.Get(gf2.CookieKey()).String(), nil
}

func (gf2 *GF2) Lang() string {
	return gf2.ctx.Request.URL.Query().Get("__ga_lang")
}

func (gf2 *GF2) Path() string {
	return gf2.ctx.URL.Path
}

func (gf2 *GF2) Method() string {
	return gf2.ctx.Method
}

func (gf2 *GF2) FormParam() url.Values {
	return gf2.ctx.Form
}

func (gf2 *GF2) Query() url.Values {
	return gf2.ctx.Request.URL.Query()
}

func (gf2 *GF2) IsPjax() bool {
	return gf2.ctx.Header.Get(constant.PjaxHeader) == "true"
}

func (gf2 *GF2) Redirect() {
	gf2.ctx.Response.RedirectTo(config.Url(config.GetLoginUrl()))
}

func (gf2 *GF2) SetContentType() {
	gf2.ctx.Response.Header().Add("Content-Type", gf2.HTMLContentType())
}

func (gf2 *GF2) Write(body []byte) {
	gf2.ctx.Response.WriteStatus(http.StatusOK, body)
}

// Request implements the method Adapter.Request.
func (gf2 *GF2) Request() *http.Request {
	return gf2.ctx.Request
}