package controller

import (
	"encoding/json"

	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/db/dialect"
)

func RecordOperationLog(ctx *context.Context) {
	if user, ok := ctx.UserValue["user"].(auth.User); ok {
		var input []byte
		form := ctx.Request.MultipartForm
		if form != nil {
			input, _ = json.Marshal((*form).Value)
		}

		_, _ = db.Table("goadmin_operation_log").Insert(dialect.H{
			"user_id": user.ID,
			"path":    ctx.Path(),
			"method":  ctx.Method(),
			"ip":      ctx.LocalIP(),
			"input":   string(input),
		})
	}
}
