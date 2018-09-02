package adapter

import (
	"errors"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	gctx "github.com/chenhg5/go-admin/context"
	"fmt"
	"strings"
	"bytes"
)

type Beego struct {
}

func (bee *Beego) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		engine *beego.App
		ok     bool
	)
	if engine, ok = router.(*beego.App); !ok {
		return errors.New("错误的参数")
	}

	for _, plug := range plugin {
		for _, req := range plug.GetRequest() {
			engine.Handlers.AddMethod(req.Method, req.URL, func(c *context.Context) {
				GetBeegoResponse(c, plug)
			})
		}
	}

	return nil
}

func GetBeegoResponse(c *context.Context, plug plugins.Plugin) {
	fmt.Println("method", c.Request.Method, "URL", c.Request.URL, "params", c.Input.Params())
	for key, value := range c.Input.Params() {
		if c.Request.URL.RawQuery == "" {
			c.Request.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + value
		} else {
			c.Request.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + value
		}
	}
	ctx := gctx.NewContext(c.Request)
	plug.GetHandler(c.Request.URL.Path, strings.ToLower(c.Request.Method))(ctx)
	for key, head := range ctx.Response.Header {
		c.ResponseWriter.Header().Add(key, head[0])
	}
	c.ResponseWriter.WriteHeader(ctx.Response.StatusCode)
	buf := new(bytes.Buffer)
	buf.ReadFrom(ctx.Response.Body)
	c.WriteString(buf.String())
	return
}
