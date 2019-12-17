package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"strings"
)

type ExportParam struct {
	Panel  table.Table
	Id     []string
	Prefix string
	IsAll  bool
}

func Export(conn db.Connection) context.Handler {
	return func(ctx *context.Context) {

		prefix := ctx.Query("__prefix")
		panel := table.Get(prefix)
		if !panel.GetExportable() {
			alert(ctx, panel, "operation not allow", conn)
			ctx.Abort()
			return
		}

		ctx.SetUserValue("export_param", &ExportParam{
			Panel:  panel,
			Id:     strings.Split(ctx.FormValue("id"), ","),
			Prefix: prefix,
			IsAll:  ctx.FormValue("is_all") == "true",
		})
		ctx.Next()
	}
}

func GetExportParam(ctx *context.Context) *ExportParam {
	return ctx.UserValue["export_param"].(*ExportParam)
}
