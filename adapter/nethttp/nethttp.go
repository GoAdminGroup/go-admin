package nethttp

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	cfg "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type NetHTTP struct {
	adapter.BaseAdapter
	ctx Context
	app *http.ServeMux
}

func init() {
	engine.Register(new(NetHTTP))
}

// User implements the method Adapter.User.
func (nh *NetHTTP) User(ctx interface{}) (models.UserModel, bool) {
	return nh.GetUser(ctx, nh)
}

// Use implements the method Adapter.Use.
func (nh *NetHTTP) Use(app interface{}, plugs []plugins.Plugin) error {
	return nh.GetUse(app, plugs, nh)
}

// Content implements the method Adapter.Content.
func (nh *NetHTTP) Content(ctx interface{}, getPanelFn types.GetPanelFn, fn context.NodeProcessor, btns ...types.Button) {
	nh.GetContent(ctx, getPanelFn, nh, btns, fn)
}

type HandlerFunc func(ctx Context) (types.Panel, error)

func Content(handler HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := Context{
			Request:  request,
			Response: writer,
		}
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return handler(ctx.(Context))
		})
	}
}

// SetApp implements the method Adapter.SetApp.
func (nh *NetHTTP) SetApp(app interface{}) error {
	var (
		eng *http.ServeMux
		ok  bool
	)
	if eng, ok = app.(*http.ServeMux); !ok {
		return errors.New("net/http adapter SetApp: wrong parameter")
	}
	nh.app = eng
	return nil
}

// AddHandler implements the method Adapter.AddHandler.
func (nh *NetHTTP) AddHandler(method, path string, handlers context.Handlers) {
	url := path
	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")
	url = reg1.ReplaceAllString(url, "{$1}/")
	url = reg2.ReplaceAllString(url, "{$1}")

	if len(url) > 1 && url[0] == '/' && url[1] == '/' {
		url = url[1:]
	}

	pattern := fmt.Sprintf("%s %s", strings.ToUpper(method), url)

	nh.app.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path[len(r.URL.Path)-1] == '/' {
			r.URL.Path = r.URL.Path[:len(r.URL.Path)-1]
		}

		ctx := context.NewContext(r)

		params := getPathParams(url, r.URL.Path)
		for key, value := range params {
			if r.URL.RawQuery == "" {
				r.URL.RawQuery += strings.ReplaceAll(key, ":", "") + "=" + value
			} else {
				r.URL.RawQuery += "&" + strings.ReplaceAll(key, ":", "") + "=" + value
			}
		}

		ctx.SetHandlers(handlers).Next()
		for key, head := range ctx.Response.Header {
			w.Header().Set(key, head[0])
		}
		if ctx.Response.Body != nil {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(ctx.Response.Body)
			w.WriteHeader(ctx.Response.StatusCode)
			_, _ = w.Write(buf.Bytes())
		} else {
			w.WriteHeader(ctx.Response.StatusCode)
		}
	})
}

// HandleFun is type of route methods of chi.
type HandleFun func(pattern string, handlerFn http.HandlerFunc)

// Context wraps the Request and Response object of Chi.
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

// SetContext implements the method Adapter.SetContext.
func (*NetHTTP) SetContext(contextInterface interface{}) adapter.WebFrameWork {
	var (
		ctx Context
		ok  bool
	)
	if ctx, ok = contextInterface.(Context); !ok {
		panic("net/http adapter SetContext: wrong parameter")
	}
	return &NetHTTP{ctx: ctx}
}

// Name implements the method Adapter.Name.
func (*NetHTTP) Name() string {
	return "net/http"
}

// Redirect implements the method Adapter.Redirect.
func (nh *NetHTTP) Redirect() {
	http.Redirect(nh.ctx.Response, nh.ctx.Request, cfg.Url(cfg.GetLoginUrl()), http.StatusFound)
}

// SetContentType implements the method Adapter.SetContentType.
func (nh *NetHTTP) SetContentType() {
	nh.ctx.Response.Header().Set("Content-Type", nh.HTMLContentType())
}

// Write implements the method Adapter.Write.
func (nh *NetHTTP) Write(body []byte) {
	nh.ctx.Response.WriteHeader(http.StatusOK)
	_, _ = nh.ctx.Response.Write(body)
}

// GetCookie implements the method Adapter.GetCookie.
func (nh *NetHTTP) GetCookie() (string, error) {
	cookie, err := nh.ctx.Request.Cookie(nh.CookieKey())
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

// Lang implements the method Adapter.Lang.
func (nh *NetHTTP) Lang() string {
	return nh.ctx.Request.URL.Query().Get("__ga_lang")
}

// Path implements the method Adapter.Path.
func (nh *NetHTTP) Path() string {
	return nh.ctx.Request.URL.Path
}

// Method implements the method Adapter.Method.
func (nh *NetHTTP) Method() string {
	return nh.ctx.Request.Method
}

// FormParam implements the method Adapter.FormParam.
func (nh *NetHTTP) FormParam() url.Values {
	_ = nh.ctx.Request.ParseMultipartForm(32 << 20)
	return nh.ctx.Request.PostForm
}

// IsPjax implements the method Adapter.IsPjax.
func (nh *NetHTTP) IsPjax() bool {
	return nh.ctx.Request.Header.Get(constant.PjaxHeader) == "true"
}

// Query implements the method Adapter.Query.
func (nh *NetHTTP) Query() url.Values {
	return nh.ctx.Request.URL.Query()
}

// Request implements the method Adapter.Request.
func (nh *NetHTTP) Request() *http.Request {
	return nh.ctx.Request
}

// getPathParams extracts path parameters from a URL based on a given pattern.
func getPathParams(pattern, url string) map[string]string {
	params := make(map[string]string)

	// Convert pattern to regex
	placeholderRegex := regexp.MustCompile(`\{(\w+)\}`)
	regexPattern := "^" + placeholderRegex.ReplaceAllStringFunc(pattern, func(s string) string {
		return `(?P<` + s[1:len(s)-1] + `>\w+)`
	}) + `$`

	// Compile regex
	regex := regexp.MustCompile(regexPattern)

	// Match the URL against the regex
	match := regex.FindStringSubmatch(url)
	if match == nil {
		return nil
	}

	// Extract named groups
	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" { // Ignore the whole match at index 0
			params[name] = match[i]
		}
	}

	return params
}
