package gorilla

import (
	_ "github.com/chenhg5/go-admin/adapter/gorilla"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
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

	return app
}
