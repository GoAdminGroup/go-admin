package adapter

import (
	"bytes"
	"errors"
	"github.com/labstack/echo"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/plugins"
	"strings"
)

type Echo struct {
}

func (e *Echo) Use(router interface{}, plugin []plugins.Plugin) error {
	var (
		engine *echo.Echo
		ok     bool
	)
	if engine, ok = router.(*echo.Echo); !ok {
		return errors.New("wrong parameter")
	}

	for _, plug := range plugin {
		var plugCopy plugins.Plugin
		plugCopy = plug
		for _, req := range plug.GetRequest() {
			engine.Add(strings.ToUpper(req.Method), req.URL, func(c echo.Context) error {
				ctx := context.NewContext(c.Request())

				for _, key := range c.ParamNames() {
					if c.Request().URL.RawQuery == "" {
						c.Request().URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + c.Param(key)
					} else {
						c.Request().URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + c.Param(key)
					}
				}

				plugCopy.GetHandler(c.Request().URL.Path, strings.ToLower(c.Request().Method))(ctx)
				for key, head := range ctx.Response.Header {
					c.Response().Header().Set(key, head[0])
				}
				if ctx.Response.Body != nil {
					buf := new(bytes.Buffer)
					buf.ReadFrom(ctx.Response.Body)
					c.String(ctx.Response.StatusCode, buf.String())
				}
				return nil
			})
		}
	}

	return nil
}
