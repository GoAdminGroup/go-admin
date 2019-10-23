// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package fasthttp

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
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	template2 "html/template"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Fasthttp struct {
}

func init() {
	engine.Register(new(Fasthttp))
}

func (fast *Fasthttp) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *fasthttprouter.Router
		ok  bool
	)
	if eng, ok = router.(*fasthttprouter.Router); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Handle(strings.ToUpper(req.Method), req.URL, func(c *fasthttp.RequestCtx) {
				httpreq := Convertor(c)
				ctx := context.NewContext(httpreq)

				var params = make(map[string]string)
				c.VisitUserValues(func(i []byte, i2 interface{}) {
					if value, ok := i2.(string); ok {
						params[string(i)] = value
					}
				})

				for key, value := range params {
					if httpreq.URL.RawQuery == "" {
						httpreq.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
					} else {
						httpreq.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
					}
				}

				ctx.SetHandlers(plugCopy.GetHandler(string(c.Path()), strings.ToLower(string(c.Method())))).Next()
				for key, head := range ctx.Response.Header {
					c.Response.Header.Set(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(ctx.Response.Body)
					_, _ = c.WriteString(buf.String())
				}
				c.Response.SetStatusCode(ctx.Response.StatusCode)
			})
		}
	}

	return nil
}

func Convertor(ctx *fasthttp.RequestCtx) *http.Request {
	var r http.Request

	body := ctx.PostBody()
	r.Method = string(ctx.Method())
	r.Proto = "HTTP/1.1"
	r.ProtoMajor = 1
	r.ProtoMinor = 1
	r.RequestURI = string(ctx.RequestURI())
	r.ContentLength = int64(len(body))
	r.Host = string(ctx.Host())
	r.RemoteAddr = ctx.RemoteAddr().String()

	hdr := make(http.Header)
	ctx.Request.Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		switch sk {
		case "Transfer-Encoding":
			r.TransferEncoding = append(r.TransferEncoding, sv)
		default:
			hdr.Set(sk, sv)
		}
	})
	r.Header = hdr
	r.Body = &netHTTPBody{body}
	rURL, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		ctx.Logger().Printf("cannot parse requestURI %q: %s", r.RequestURI, err)
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return &r
	}
	r.URL = rURL
	return &r
}

type netHTTPBody struct {
	b []byte
}

func (r *netHTTPBody) Read(p []byte) (int, error) {
	if len(r.b) == 0 {
		return 0, io.EOF
	}
	n := copy(p, r.b)
	r.b = r.b[n:]
	return n, nil
}

func (r *netHTTPBody) Close() error {
	r.b = r.b[:0]
	return nil
}

func (fast *Fasthttp) Content(contextInterface interface{}, c types.GetPanel) {

	var (
		ctx *fasthttp.RequestCtx
		ok  bool
	)
	if ctx, ok = contextInterface.(*fasthttp.RequestCtx); !ok {
		panic("wrong parameter")
	}

	globalConfig := config.Get()

	sesKey := string(ctx.Request.Header.Cookie("go_admin_session"))

	if sesKey == "" {
		ctx.Redirect(globalConfig.Url("/login"), http.StatusFound)
		return
	}

	userId, ok := auth.Driver.Load(sesKey)["user_id"]

	if !ok {
		ctx.Redirect(globalConfig.Url("/login"), http.StatusFound)
		return
	}

	user, ok := auth.GetCurUserById(int64(userId.(float64)))

	if !ok {
		ctx.Redirect(globalConfig.Url("/login"), http.StatusFound)
		return
	}

	var (
		panel types.Panel
		err   error
	)

	if !auth.CheckPermissions(user, string(ctx.Path()), string(ctx.Method())) {
		alert := template.Get(globalConfig.Theme).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").SetContent(template2.HTML("no permission")).GetContent()

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

	tmpl, tmplName := template.Get(globalConfig.Theme).GetTemplate(string(ctx.Request.Header.Peek(constant.PjaxHeader)) == "true")

	ctx.Response.Header.Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName,
		types.NewPage(user, *(menu.GetGlobalMenu(user).SetActiveClass(globalConfig.UrlRemovePrefix(ctx.Request.URI().String()))),
			panel, globalConfig))
	if err != nil {
		logger.Error("Fasthttp Content", err)
	}
	_, _ = ctx.WriteString(buf.String())
}
