package example

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	AssertRootUrl = prefix

	app.GET(prefix + "/example", auth.AuthMiddleware(TestHandler))

	return app
}