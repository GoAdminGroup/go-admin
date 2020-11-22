package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/controller"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	_ "github.com/GoAdminGroup/go-admin/template/types/display"
)

// Admin is a GoAdmin plugin.
type Admin struct {
	*plugins.Base
	tableList table.GeneratorList
	guardian  *guard.Guard
	handler   *controller.Handler
}

// InitPlugin implements Plugin.InitPlugin.
// TODO: find a better way to manage the dependencies
func (admin *Admin) InitPlugin(services service.List) {

	// DO NOT DELETE
	admin.InitBase(services, "")

	c := config.GetService(services.Get("config"))
	st := table.NewSystemTable(admin.Conn, c)
	genList := table.GeneratorList{
		"manager":        st.GetManagerTable,
		"permission":     st.GetPermissionTable,
		"roles":          st.GetRolesTable,
		"op":             st.GetOpTable,
		"menu":           st.GetMenuTable,
		"normal_manager": st.GetNormalManagerTable,
	}
	if c.IsAllowConfigModification() {
		genList.Add("site", st.GetSiteTable)
	}
	if c.IsNotProductionEnvironment() {
		genList.Add("generate", st.GetGenerateForm)
	}
	admin.tableList.Combine(genList)
	admin.guardian = guard.New(admin.Services, admin.Conn, admin.tableList, admin.UI.NavButtons)
	handlerCfg := controller.Config{
		Config:     c,
		Services:   services,
		Generators: admin.tableList,
		Connection: admin.Conn,
	}
	admin.handler.UpdateCfg(handlerCfg)
	admin.initRouter()
	admin.handler.SetRoutes(admin.App.Routers)
	admin.handler.AddNavButton(admin.UI.NavButtons)

	table.SetServices(services)

	action.InitOperationHandlerSetter(admin.GetAddOperationFn())
}

func (admin *Admin) GetIndexURL() string {
	return config.GetIndexURL()
}

func (admin *Admin) GetInfo() plugins.Info {
	return plugins.Info{
		Title:       "Basic Admin",
		Website:     "https://www.go-admin.cn",
		Description: "A built-in plugins of GoAdmin which help you to build a crud manager platform quickly.",
		Author:      "official",
		Version:     system.Version(),
		CreateDate:  utils.ParseTime("2018-07-08 00:00:00"),
		UpdateDate:  utils.ParseTime("2020-06-28 00:00:00"),
	}
}

func (admin *Admin) IsInstalled() bool {
	return true
}

// NewAdmin return the global Admin plugin.
func NewAdmin(tableCfg ...table.GeneratorList) *Admin {
	return &Admin{
		tableList: make(table.GeneratorList).CombineAll(tableCfg),
		Base:      &plugins.Base{PlugName: "admin"},
		handler:   controller.New(),
	}
}

func (admin *Admin) GetAddOperationFn() context.NodeProcessor {
	return admin.handler.AddOperation
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

// AddGenerators add table model generators.
func (admin *Admin) AddGenerators(gen ...table.GeneratorList) *Admin {
	admin.tableList.CombineAll(gen)
	return admin
}

// AddGlobalDisplayProcessFn call types.AddGlobalDisplayProcessFn
func (admin *Admin) AddGlobalDisplayProcessFn(f types.FieldFilterFn) *Admin {
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
