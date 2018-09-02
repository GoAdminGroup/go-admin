package gin

import (
	"github.com/gin-gonic/gin"
	"errors"
	"github.com/chenhg5/go-admin/plugins"
	"strings"
	"github.com/chenhg5/go-admin/context"
	"bytes"
)

type Gin struct {
}

func (gins *Gin) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		engine *gin.Engine
		ok     bool
	)
	if engine, ok = router.(*gin.Engine); !ok {
		return errors.New("错误的参数")
	}

	for _, plug := range plugin {
		var plugCopy plugins.Plugin
		plugCopy = plug
		for _, req := range plug.GetRequest() {
			engine.Handle(strings.ToUpper(req.Method), req.URL, func(c *gin.Context) {
				GetResponse(c, plugCopy)
			})
		} 
	}

	return nil
}

func GetResponse(c *gin.Context, plug plugins.Plugin) {
	ctx := context.NewContext(c.Request)

	for _, param := range c.Params {
		if c.Request.URL.RawQuery == "" {
			c.Request.URL.RawQuery += strings.Replace(param.Key, ":", "", -1) + "=" + param.Value
		} else {
			c.Request.URL.RawQuery += "&" + strings.Replace(param.Key, ":", "", -1) + "=" + param.Value
		}
	}

	plug.GetHandler(c.Request.URL.Path, strings.ToLower(c.Request.Method))(ctx)
	for key, head := range ctx.Response.Header {
		c.Header(key, head[0])
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(ctx.Response.Body)
	c.String(ctx.Response.StatusCode, buf.String())
}
