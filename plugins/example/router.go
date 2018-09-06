package example

import (
	"goAdmin/context"
	"goAdmin/modules/auth"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	app.GET(prefix + "/example", auth.SetPrefix(prefix).Middleware(TestHandler))

	return app
}