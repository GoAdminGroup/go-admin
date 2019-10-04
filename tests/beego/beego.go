package beego

import (
	"github.com/astaxie/beego"
	_ "github.com/chenhg5/go-admin/adapter/beego"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"net/http"
	"os"
)

func NewBeegoHandler() http.Handler {

	app := beego.NewApp()

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJson(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}
	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087

	return app.Handlers
}
