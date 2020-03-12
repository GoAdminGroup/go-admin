package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/controller"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// Admin is a GoAdmin plugin.
type Admin struct {
	app       *context.App
	tableList table.GeneratorList
	services  service.List
	conn      db.Connection
	guardian  *guard.Guard
	handler   *controller.Handler
}

// InitPlugin implements Plugin.InitPlugin.
func (admin *Admin) InitPlugin(services service.List) {

	cfg := config.Get()

	// TODO: find a better way to manage the dependencies

	admin.services = services
	admin.conn = db.GetConnection(admin.services)
	st := table.NewSystemTable(admin.conn)
	admin.tableList.Combine(table.GeneratorList{
		"manager":        st.GetManagerTable,
		"permission":     st.GetPermissionTable,
		"roles":          st.GetRolesTable,
		"op":             st.GetOpTable,
		"menu":           st.GetMenuTable,
		"normal_manager": st.GetNormalManagerTable,
	})
	admin.guardian = guard.New(admin.services, admin.conn, admin.tableList)
	admin.handler = controller.New(controller.Config{
		Config:     cfg,
		Services:   services,
		Generators: admin.tableList,
		Connection: admin.conn,
	})
	admin.initRouter(cfg.Prefix())
	admin.handler.SetRoutes(admin.app.Routers)
	admin.tableList.InjectRoutes(admin.app, admin.services)

	table.SetServices(services)
}

// NewAdmin return the global Admin plugin.
func NewAdmin(tableCfg ...table.GeneratorList) *Admin {
	return &Admin{
		tableList: make(table.GeneratorList).CombineAll(tableCfg),
	}
}

// SetCaptcha set captcha driver.
func (admin *Admin) SetCaptcha(captcha map[string]string) *Admin {
	admin.handler.SetCaptcha(captcha)
	return admin
}

// AddGenerator add table model generator.
func (admin *Admin) AddGenerator(key string, g table.Generator) *Admin {
	admin.tableList.Add(key, g)
	return admin
}

// GetRequest implements Plugin.GetRequest.
func (admin *Admin) GetRequest() []context.Path {
	return admin.app.Requests
}

// GetHandler implements Plugin.GetHandler.
func (admin *Admin) GetHandler() context.HandlerMap {
	return plugins.GetHandler(admin.app)
}

// AddGlobalDisplayProcessFn call types.AddGlobalDisplayProcessFn
func (admin *Admin) AddGlobalDisplayProcessFn(f types.DisplayProcessFn) *Admin {
	types.AddGlobalDisplayProcessFn(f)
	return admin
}

// AddDisplayFilterLimit call types.AddDisplayFilterLimit
func (admin *Admin) AddDisplayFilterLimit(limit int) *Admin {
	types.AddLimit(limit)
	return admin
}

// AddDisplayFilterTrimSpace call types.AddDisplayFilterTrimSpace
func (admin *Admin) AddDisplayFilterTrimSpace() *Admin {
	types.AddTrimSpace()
	return admin
}

// AddDisplayFilterSubstr call types.AddDisplayFilterSubstr
func (admin *Admin) AddDisplayFilterSubstr(start int, end int) *Admin {
	types.AddSubstr(start, end)
	return admin
}

// AddDisplayFilterToTitle call types.AddDisplayFilterToTitle
func (admin *Admin) AddDisplayFilterToTitle() *Admin {
	types.AddToTitle()
	return admin
}

// AddDisplayFilterToUpper call types.AddDisplayFilterToUpper
func (admin *Admin) AddDisplayFilterToUpper() *Admin {
	types.AddToUpper()
	return admin
}

// AddDisplayFilterToLower call types.AddDisplayFilterToLower
func (admin *Admin) AddDisplayFilterToLower() *Admin {
	types.AddToUpper()
	return admin
}

// AddDisplayFilterXssFilter call types.AddDisplayFilterXssFilter
func (admin *Admin) AddDisplayFilterXssFilter() *Admin {
	types.AddXssFilter()
	return admin
}

// AddDisplayFilterXssJsFilter call types.AddDisplayFilterXssJsFilter
func (admin *Admin) AddDisplayFilterXssJsFilter() *Admin {
	types.AddXssJsFilter()
	return admin
}
