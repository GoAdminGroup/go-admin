package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"strings"
)

type ExportParam struct {
	Panel  table.Table
	Id     []string
	Prefix string
	IsAll  bool
}

func (g *Guard) Export(ctx *context.Context) {
	panel, prefix := g.table(ctx)
	if !panel.GetExportable() {
		alert(ctx, panel, "operation not allow", g.conn)
		ctx.Abort()
		return
	}

	idStr := make([]string, 0)
	ids := ctx.FormValue("id")
	if ids != "" {
		idStr = strings.Split(ctx.FormValue("id"), ",")
	}

	ctx.SetUserValue("export_param", &ExportParam{
		Panel:  panel,
		Id:     idStr,
		Prefix: prefix,
		IsAll:  ctx.FormValue("is_all") == "true",
	})
	ctx.Next()
}

func GetExportParam(ctx *context.Context) *ExportParam {
	return ctx.UserValue["export_param"].(*ExportParam)
}
