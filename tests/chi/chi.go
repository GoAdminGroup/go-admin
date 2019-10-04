package chi

import (
	_ "github.com/chenhg5/go-admin/adapter/chi"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/go-chi/chi"
	"net/http"
	"os"
)

func NewChiHandler() http.Handler {
	r := chi.NewRouter()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)
	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJson(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(r); err != nil {
		panic(err)
	}

	return r
}
