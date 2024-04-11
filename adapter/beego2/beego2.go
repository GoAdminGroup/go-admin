package beego2

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/GoAdminGroup/go-admin/adapter"
	gctx "github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type Beego2 struct {
	adapter.BaseAdapter
	ctx *context.Context
	app *web.HttpServer
}

func init() {
	engine.Register(new(Beego2))
}

func (*Beego2) Name() string {
	return "beego2"
}

func (bee2 *Beego2) Use(app interface{}, plugins []plugins.Plugin) error {
	return bee2.GetUse(app, plugins, bee2)
}

func (bee2 *Beego2) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn gctx.NodeProcessor, navButtons ...types.Button) {
	bee2.GetContent(ctx, getPanelFn, bee2, navButtons, fn)
}

func (bee2 *Beego2) User(ctx interface{}) (models.UserModel, bool) {
	return bee2.GetUser(ctx, bee2)
}

func (bee2 *Beego2) AddHandler(method, path string, handlers gctx.Handlers) {
	bee2.app.Handlers.AddMethod(method, path, func(c *context.Context) {
		for key, value := range c.Input.Params() {
			if c.Request.URL.RawQuery == "" {
				c.Request.URL.RawQuery += strings.ReplaceAll(key, ":", "") + "=" + value
			} else {
				c.Request.URL.RawQuery += "&" + strings.ReplaceAll(key, ":", "") + "=" + value
			}
		}
		ctx := gctx.NewContext(c.Request)
		ctx.SetHandlers(handlers).Next()
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

func (bee2 *Beego2) SetApp(app interface{}) error {
	var (
		eng *web.HttpServer
		ok  bool
	)
	if eng, ok = app.(*web.HttpServer); !ok {
		return errors.New("beego2 adapter SetApp: wrong parameter")
	}
	bee2.app = eng
	return nil
}

func (*Beego2) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx *context.Context
		ok  bool
	)
	if ctx, ok = contextInterface.(*context.Context); !ok {
		panic("beego2 adapter SetContext: wrong parameter")
	}
	return &Beego2{ctx: ctx}
}

func (bee2 *Beego2) GetCookie() (string, error) {
	return bee2.ctx.GetCookie(bee2.CookieKey()), nil
}

func (bee2 *Beego2) Lang() string {
	return bee2.ctx.Request.URL.Query().Get("__ga_lang")
}

func (bee2 *Beego2) Path() string {
	return bee2.ctx.Request.URL.Path
}

func (bee2 *Beego2) Method() string {
	return bee2.ctx.Request.Method
}

func (bee2 *Beego2) FormParam() url.Values {
	_ = bee2.ctx.Request.ParseMultipartForm(32 << 20)
	return bee2.ctx.Request.PostForm
}

func (bee2 *Beego2) Query() url.Values {
	return bee2.ctx.Request.URL.Query()
}

func (bee2 *Beego2) IsPjax() bool {
	return bee2.ctx.Request.Header.Get(constant.PjaxHeader) == "true"
}

func (bee2 *Beego2) Redirect() {
	bee2.ctx.Redirect(http.StatusFound, config.Url(config.GetLoginUrl()))
}

func (bee2 *Beego2) SetContentType() {
	bee2.ctx.ResponseWriter.Header().Set("Content-Type", bee2.HTMLContentType())
}

func (bee2 *Beego2) Write(body []byte) {
	_, _ = bee2.ctx.ResponseWriter.Write(body)
}

// Request implements the method Adapter.Request.
func (bee2 *Beego2) Request() *http.Request {
	return bee2.ctx.Request
}