package gorilla

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gorilla/mux"
	template2 "html/template"
	"net/http"
	"regexp"
	"strings"
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

	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {

			url := req.URL
			url = reg1.ReplaceAllString(url, "{$1}/")
			url = reg2.ReplaceAllString(url, "{$1}")

			eng.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
				ctx := context.NewContext(r)

				params := mux.Vars(r)

				for key, param := range params {
					if r.URL.RawQuery == "" {
						r.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + param
					} else {
						r.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + param
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(r.URL.Path, strings.ToLower(r.Method))).Next()
				for key, head := range ctx.Response.Header {
					w.Header().Add(key, head[0])
				}

				if ctx.Response.Body == nil {
					w.WriteHeader(ctx.Response.StatusCode)
					return
				}

				w.WriteHeader(ctx.Response.StatusCode)

				buf := new(bytes.Buffer)
				_, _ = buf.ReadFrom(ctx.Response.Body)

				_, err := w.Write(buf.Bytes())
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}).Methods(strings.ToUpper(req.Method))
		}
	}

	return nil
}

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func (g *Gorilla) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx Context
		ok  bool
	)
	if ctx, ok = contextInterface.(Context); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey, err := ctx.Request.Cookie("go_admin_session")

	if err != nil || sesKey == nil {
		http.Redirect(ctx.Response, ctx.Request, globalConfig.Url("/login"), http.StatusFound)
		return
	}

	userId, ok := auth.Driver.Load(sesKey.Value)["user_id"]

	if !ok {
		http.Redirect(ctx.Response, ctx.Request, globalConfig.Url("/login"), http.StatusFound)
		return
	}

	user, ok := auth.GetCurUserById(int64(userId.(float64)))

	if !ok {
		http.Redirect(ctx.Response, ctx.Request, globalConfig.Url("/login"), http.StatusFound)
		return
	}

	var panel types.Panel

	if !auth.CheckPermissions(user, ctx.Request.RequestURI, ctx.Request.Method) {
		alert := template.Get(globalConfig.Theme).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").SetContent(template2.HTML("Permission Denied")).GetContent()

		panel = types.Panel{
			Content:     alert,
			Description: language.Get("error"),
			Title:       language.Get("error"),
		}
	} else {
		panel, err = c(ctx)
		if err != nil {
			alert := template.Get(globalConfig.Theme).
				Alert().
				SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
				SetTheme("warning").SetContent(template2.HTML(err.Error())).GetContent()
			panel = types.Panel{
				Content:     alert,
				Description: language.Get("error"),
				Title:       language.Get("error"),
			}
		}
	}

	tmpl, tmplName := template.Get(globalConfig.Theme).GetTemplate(ctx.Request.Header.Get(constant.PjaxHeader) == "true")

	ctx.Response.Header().Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user, *(menu.GetGlobalMenu(user).SetActiveClass(globalConfig.UrlRemovePrefix(ctx.Request.URL.String()))), panel, globalConfig))
	if err != nil {
		logger.Error("Gorilla Content", err)
	}
	ctx.Response.WriteHeader(http.StatusOK)
	_, _ = ctx.Response.Write(buf.Bytes())
}
