package gorilla

import (
	// add gorilla adapter
	_ "github.com/GoAdminGroup/go-admin/adapter/gorilla"
	// add mysql driver
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	// add adminlte ui theme
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/tests/tables"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func newGorillaHandler() http.Handler {
	app := mux.NewRouter()
	eng := engine.Default()

	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(admin.NewAdmin(tables.Generators).
			AddGenerator("user", tables.GetUserTable), examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}
