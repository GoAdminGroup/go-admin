package iris

import (
	_ "github.com/chenhg5/go-admin/adapter/iris"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/kataras/iris"
	"net/http"
	"os"
)

func NewIrisHandler() http.Handler {
	app := iris.Default()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)
	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJson(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	return app
}
