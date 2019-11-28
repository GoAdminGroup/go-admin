package gf

import (
	// add gf adapter
	_ "github.com/GoAdminGroup/go-admin/adapter/gf"
	"net/http"
	// add mysql driver
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	// add adminlte ui theme
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"os"
)

func newHandler() http.Handler {
	s := g.Server(8103)

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators).AddDisplayFilterXssJsFilter()

	template.AddComp(chartjs.NewChart())

	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin).
		Use(s); err != nil {
		panic(err)
	}

	s.BindHandler("GET:/admin", func(ctx *ghttp.Request) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return datamodel.GetContent()
		})
	})

	return new(httpHandler).SetSrv(s)
}

type httpHandler struct {
	srv *ghttp.Server
}

func (hh *httpHandler) SetSrv(s *ghttp.Server) *httpHandler {
	hh.srv = s
	return hh
}

func (hh *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOTE: ╮(╯▽╰)╭
	hh.srv.DefaultHttpHandle(w, r)
}
