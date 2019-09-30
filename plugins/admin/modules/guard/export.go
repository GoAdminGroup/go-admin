package guard

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"strings"
)

type ExportParam struct {
	Panel  table.Table
	Id     []string
	Prefix string
}

func Export(ctx *context.Context) {

	prefix := ctx.Query("prefix")
	panel := table.List[prefix]
	if !panel.GetExportable() {
		alert(ctx, panel, "operation not allow")
		ctx.Abort()
		return
	}

	ctx.SetUserValue("export_param", &ExportParam{
		Panel:  panel,
		Id:     strings.Split(ctx.FormValue("id"), ","),
		Prefix: prefix,
	})
	ctx.Next()
}

func GetExportParam(ctx *context.Context) *ExportParam {
	return ctx.UserValue["export_param"].(*ExportParam)
}
