package echo

import (
	_ "github.com/chenhg5/go-admin/adapter/echo"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func NewEchoHandler() http.Handler {
	e := echo.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJson(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(e); err != nil {
		panic(err)
	}

	return e
}
