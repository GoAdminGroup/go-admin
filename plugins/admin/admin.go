package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/controller"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// Admin is a GoAdmin plugin.
type Admin struct {
	app       *context.App
	tableList table.GeneratorList
}

// InitPlugin implements Plugin.InitPlugin.
func (admin *Admin) InitPlugin(services service.List) {

	cfg := config.Get()

	generatorList := table.GeneratorList{
		"manager":        table.GetManagerTable,
		"permission":     table.GetPermissionTable,
		"roles":          table.GetRolesTable,
		"op":             table.GetOpTable,
		"menu":           table.GetMenuTable,
		"normal_manager": table.GetNormalManagerTable,
	}.Combine(admin.tableList)

	// Init router
	App.app = initRouter(cfg.Prefix(), services, generatorList)
	admin.tableList.InjectRoutes(App.app, services)

	table.SetServices(services)

	controller.Init(controller.InitConfiguration{
		Config:     cfg,
		RouterMap:  App.app.Routers,
		Services:   services,
		Generators: generatorList,
	})
}

// App is the global Admin plugin.
var App = &Admin{
	tableList: make(table.GeneratorList),
}

// NewAdmin return the global Admin plugin.
func NewAdmin(tableCfg ...table.GeneratorList) *Admin {
	App.tableList.CombineAll(tableCfg)
	return App
}

// SetCaptcha set captcha driver.
func (admin *Admin) SetCaptcha(captcha map[string]string) *Admin {
	controller.SetCaptcha(captcha)
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
func (admin *Admin) GetHandler(url, method string) context.Handlers {
	return plugins.GetHandler(url, method, admin.app)
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
