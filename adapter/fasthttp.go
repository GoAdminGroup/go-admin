package adapter

import (
	"errors"
	"github.com/chenhg5/go-admin/plugins"
	"strings"
	"github.com/chenhg5/go-admin/context"
	"bytes"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
	"io"
)

type Fasthttp struct {
}

func (fast *Fasthttp) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		engine *fasthttprouter.Router
		ok     bool
	)
	if engine, ok = router.(*fasthttprouter.Router); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy plugins.Plugin
		plugCopy = plug
		for _, req := range plug.GetRequest() {
			engine.Handle(strings.ToUpper(req.Method), req.URL, func(c *fasthttp.RequestCtx) {
				httpreq := Convertor(c)
				ctx := context.NewContext(httpreq)

				var params map[string]string
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

				plugCopy.GetHandler(string(c.Path()), strings.ToLower(string(c.Method())))(ctx)
				for key, head := range ctx.Response.Header {
					c.Response.Header.Set(key, head[0])
				}
				buf := new(bytes.Buffer)
				buf.ReadFrom(ctx.Response.Body)
				c.Response.SetStatusCode(ctx.Response.StatusCode)
				c.WriteString(buf.String())
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