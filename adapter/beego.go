package adapter

import (
	"bytes"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	gctx "github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	template2 "html/template"
	"net/http"
	"strings"
)

type Beego struct {
}

func (bee *Beego) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		engine *beego.App
		ok     bool
	)
	if engine, ok = router.(*beego.App); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy plugins.Plugin
		plugCopy = plug
		for _, req := range plug.GetRequest() {
			engine.Handlers.AddMethod(req.Method, req.URL, func(c *context.Context) {
				for key, value := range c.Input.Params() {
					if c.Request.URL.RawQuery == "" {
						c.Request.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
					} else {
						c.Request.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
					}
				}
				ctx := gctx.NewContext(c.Request)
				plugCopy.GetHandler(c.Request.URL.Path, strings.ToLower(c.Request.Method))(ctx)
				for key, head := range ctx.Response.Header {
					c.ResponseWriter.Header().Add(key, head[0])
				}
				c.ResponseWriter.WriteHeader(ctx.Response.StatusCode)
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					buf.ReadFrom(ctx.Response.Body)
					c.WriteString(buf.String())
				}
			})
		}
	}

	return nil
}

func BeegoContent(ctx *context.Context, c func() types.Panel) {
	globalConfig := config.Get()

	sesKey := ctx.GetCookie("go_admin_session")

	if sesKey == "" {
		ctx.Redirect(http.StatusFound, "/"+globalConfig.PREFIX+"/login")
		return
	}

	userId, ok := auth.Driver.Load(sesKey)["user_id"]

	if !ok {
		ctx.Redirect(http.StatusFound, "/"+globalConfig.PREFIX+"/login")
		return
	}

	user, ok := auth.GetCurUserById(userId.(string))

	if !ok {
		ctx.Redirect(http.StatusFound, "/"+globalConfig.PREFIX+"/login")
		return
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.Request.URL.Path, ctx.Request.Method) {
		alert := template.Get(globalConfig.THEME).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template2.HTML("没有权限")).GetContent()

		panel = types.Panel{
			Content:     alert,
			Description: "Error",
			Title:       "Error",
		}
	} else {
		panel = c()
	}

	tmpl, tmplName := template.Get(globalConfig.THEME).GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

	ctx.ResponseWriter.Header().Add("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: menu.GetGlobalMenu(user),
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: "/" + globalConfig.PREFIX,
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
	})
	ctx.WriteString(buf.String())
}
