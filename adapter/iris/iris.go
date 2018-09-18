package iris

import (
	"bytes"
	"errors"
	"github.com/kataras/iris"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/plugins"
	"strings"
	"github.com/chenhg5/go-admin/engine"
)

type Iris struct {
}

func init()  {
	engine.Register(new(Iris))
}

func (is *Iris) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		eng *iris.Application
		ok     bool
	)
	if eng, ok = router.(*iris.Application); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy plugins.Plugin
		plugCopy = plug
		for _, req := range plug.GetRequest() {
			eng.Handle(strings.ToUpper(req.Method), req.URL, func(c iris.Context) {
				ctx := context.NewContext(c.Request())

				var params = map[string]string{}
				c.Params().Visit(func(key string, value string) {
					params[key] = value
				})

				for key, value := range params {
					if c.Request().URL.RawQuery == "" {
						c.Request().URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
					} else {
						c.Request().URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
					}
				}

				plugCopy.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))(ctx)
				for key, head := range ctx.Response.Header {
					c.Header(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					buf.ReadFrom(ctx.Response.Body)
					c.WriteString(buf.String())
				}
				c.StatusCode(ctx.Response.StatusCode)
			})
		}
	}

	return nil
}
