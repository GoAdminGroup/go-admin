package gf

import (
	// add gf adapter
	"reflect"

	_ "github.com/ChenSee/go-admin/adapter/gf2"
	"github.com/agiledragon/gomonkey"

	// add mysql driver
	"github.com/ChenSee/go-admin/modules/config"
	_ "github.com/ChenSee/go-admin/modules/db/drivers/mysql"
	"github.com/ChenSee/go-admin/modules/language"

	// add postgresql driver
	_ "github.com/ChenSee/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/ChenSee/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/ChenSee/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"github.com/ChenSee/goAdminThemes/adminlte"

	"net/http"
	"os"

	"github.com/ChenSee/go-admin/engine"
	"github.com/ChenSee/go-admin/plugins/admin"
	"github.com/ChenSee/go-admin/plugins/admin/modules/table"
	"github.com/ChenSee/go-admin/template"
	"github.com/ChenSee/go-admin/template/chartjs"
	"github.com/ChenSee/go-admin/tests/tables"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func internalHandler() http.Handler {
	s := g.Server()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators).AddDisplayFilterXssJsFilter()

	template.AddComp(chartjs.NewChart())

	adminPlugin.AddGenerator("user", tables.GetUserTable)

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin).
		Use(s); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	s.SetPort(8103)

	gomonkey.ApplyMethod(reflect.TypeOf(new(ghttp.Request).Session), "Close",
		func(*ghttp.Session) error {
			return nil
		})

	return s
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {

	s := g.Server(8103)

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(gens)

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
		AddPlugins(adminPlugin).Use(s); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	eng.HTML("GET", "/admin", tables.GetContent)

	return s
}
