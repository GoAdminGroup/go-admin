package gorilla

import (
	"bytes"
	"errors"
	template2 "html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
)

type Gorilla struct {
}

func init() {
	engine.Register(new(Gorilla))
}

func (g *Gorilla) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *mux.Router
		ok  bool
	)
	if eng, ok = router.(*mux.Router); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy plugins.Plugin
		plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.HandleFunc(req.URL, func(w http.ResponseWriter, r *http.Request) {
				ctx := context.NewContext(r)

				plugCopy.GetHandler(r.URL.Path, strings.ToLower(r.Method))(ctx)
				for key, head := range ctx.Response.Header {
					w.Header().Add(key, head[0])
				}

				if ctx.Response.Body == nil {
					w.WriteHeader(ctx.Response.StatusCode)
					return
				}

				buf := new(bytes.Buffer)
				_, _ = buf.ReadFrom(ctx.Response.Body)

				_, err := w.Write(buf.Bytes())
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
				w.WriteHeader(ctx.Response.StatusCode)
			}).Methods(strings.ToUpper(req.Method))
		}
	}

	return nil
}

func (g *Gorilla) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx *http.Request
		ok  bool
	)
	if ctx, ok = contextInterface.(*http.Request); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey, err := ctx.Cookie("go_admin_session")

	if err != nil || sesKey == nil {
		ctx.Response.Header.Set("Location", "/"+globalConfig.PREFIX+"/login")
		ctx.Response.StatusCode = http.StatusFound
		return
	}

	userId, ok := auth.Driver.Load(sesKey.Value)["user_id"]

	if !ok {
		ctx.Response.Header.Set("Location", "/"+globalConfig.PREFIX+"/login")
		ctx.Response.StatusCode = http.StatusFound
		return
	}

	user, ok := auth.GetCurUserById(userId.(string))

	if !ok {
		ctx.Response.Header.Set("Location", "/"+globalConfig.PREFIX+"/login")
		ctx.Response.StatusCode = http.StatusFound
		return
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.RequestURI, ctx.Method) {
		alert := template.Get(globalConfig.THEME).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template2.HTML("Permission Denied")).GetContent()

		panel = types.Panel{
			Content:     alert,
			Description: "Error",
			Title:       "Error",
		}
	} else {
		panel = c()
	}

	tmpl, tmplName := template.Get(globalConfig.THEME).GetTemplate(ctx.Header.Get("X-PJAX") == "true")

	ctx.Header.Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	_ = tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *(menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.URL.String(), "/"+globalConfig.PREFIX, "", 1))),
		System: types.SystemInfo{
			Version: "0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: "/" + globalConfig.PREFIX,
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
		ColorScheme:   globalConfig.COLORSCHEME,
	})
}
