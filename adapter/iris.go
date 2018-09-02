package adapter

import (
	"errors"
	"github.com/chenhg5/go-admin/plugins"
	"strings"
	"github.com/chenhg5/go-admin/context"
	"bytes"
	"github.com/kataras/iris"
)

type Iris struct {
}

func (is *Iris) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		engine *iris.Application
		ok     bool
	)
	if engine, ok = router.(*iris.Application); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy plugins.Plugin
		plugCopy = plug
		for _, req := range plug.GetRequest() {
			engine.Handle(strings.ToUpper(req.Method), req.URL, func(ctx iris.Context) {
				GetIrisResponse(ctx, plugCopy)
			})
		}
	}

	return nil
}

func GetIrisResponse(c iris.Context, plug plugins.Plugin) {
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

	plug.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))(ctx)
	for key, head := range ctx.Response.Header {
		c.Header(key, head[0])
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(ctx.Response.Body)
	c.StatusCode(ctx.Response.StatusCode)
	c.WriteString(buf.String())
}