package example

import (
	"goAdmin/context"
	"goAdmin/modules/auth"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	AssertRootUrl = prefix

	app.GET(prefix + "/example", auth.AuthMiddleware(TestHandler))

	return app
}