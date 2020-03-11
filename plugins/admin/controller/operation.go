package controller

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"

	"github.com/GoAdminGroup/go-admin/context"
)

// RecordOperationLog record all operation logs, store into database.
func (h *Handler) RecordOperationLog(ctx *context.Context) {
	if user, ok := ctx.UserValue["user"].(models.UserModel); ok {
		var input []byte
		form := ctx.Request.MultipartForm
		if form != nil {
			input, _ = json.Marshal((*form).Value)
		}

		models.OperationLog().SetConn(h.conn).New(user.Id, ctx.Path(), ctx.Method(), ctx.LocalIP(), string(input))
	}
}
