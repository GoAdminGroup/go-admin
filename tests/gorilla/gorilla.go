package gorilla

import (
	ada "github.com/GoAdminGroup/go-admin/adapter/gorilla"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template/types"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func NewGorillaHandler() http.Handler {
	app := mux.NewRouter()
	eng := engine.Default()

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJson(os.Args[len(os.Args)-1]).
		AddPlugins(admin.NewAdmin(datamodel.Generators).
			AddGenerator("user", datamodel.GetUserTable), examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	app.Handle("/admin", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		engine.Content(ada.Context{Request: request, Response: writer}, func(ctx interface{}) (types.Panel, error) {
			return datamodel.GetContent()
		})
	})).Methods("get")

	return app
}
