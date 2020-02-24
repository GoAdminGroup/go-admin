package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
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

		prefix := ctx.Query(constant.PrefixKey)
		panel := table.Get(prefix, ctx)
		if !panel.GetExportable() {
			alert(ctx, panel, "operation not allow", conn)
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
}

func GetExportParam(ctx *context.Context) *ExportParam {
	return ctx.UserValue["export_param"].(*ExportParam)
}
