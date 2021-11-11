package iris

import (
	"github.com/GoAdminGroup/themes/adminlte"
	// add iris adapter
	_ "github.com/digroad/go-admin/adapter/iris"
	"github.com/digroad/go-admin/modules/config"
	"github.com/digroad/go-admin/modules/language"
	"github.com/digroad/go-admin/plugins/admin/modules/table"

	// add mysql driver
	_ "github.com/digroad/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/digroad/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/digroad/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/digroad/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/digroad/go-admin/template"
	"github.com/digroad/go-admin/template/chartjs"

	"net/http"
	"os"

	"github.com/digroad/go-admin/engine"
	"github.com/digroad/go-admin/plugins/admin"
	"github.com/digroad/go-admin/plugins/example"
	"github.com/digroad/go-admin/tests/tables"
	"github.com/kataras/iris/v12"
)

func newHandler() http.Handler {
	app := iris.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	if err := app.Build(); err != nil {
		panic(err)
	}

	return app.Router
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := iris.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)

	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	if err := app.Build(); err != nil {
		panic(err)
	}

	return app.Router
}
