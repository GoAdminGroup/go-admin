// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package fasthttp

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Fasthttp structure value is a Fasthttp GoAdmin adapter.
type Fasthttp struct {
	adapter.BaseAdapter
}

func init() {
	engine.Register(new(Fasthttp))
}

// Use implement WebFrameWork.Use method.
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
				httpreq := convertCtx(c)
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

func convertCtx(ctx *fasthttp.RequestCtx) *http.Request {
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

// Content implement WebFrameWork.Content method.
func (fast *Fasthttp) Content(contextInterface interface{}, getPanelFn types.GetPanelFn) {

	var (
		ctx *fasthttp.RequestCtx
		ok  bool
	)
	if ctx, ok = contextInterface.(*fasthttp.RequestCtx); !ok {
		panic("wrong parameter")
	}

	body, authSuccess, err := fast.GetContent(string(ctx.Request.Header.Cookie(fast.CookieKey())), string(ctx.Path()),
		string(ctx.Method()), string(ctx.Request.Header.Peek(constant.PjaxHeader)), getPanelFn, ctx)

	if !authSuccess {
		ctx.Redirect(config.Get().Url("/login"), http.StatusFound)
		return
	}

	if err != nil {
		logger.Error("Fasthttp Content", err)
	}

	ctx.Response.Header.Set("Content-Type", fast.HTMLContentType())

	_, _ = ctx.Write(body)
}
