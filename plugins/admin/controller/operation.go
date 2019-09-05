package controller

import (
	"encoding/json"
	"github.com/chenhg5/go-admin/plugins/admin/models"

	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
)

func RecordOperationLog(ctx *context.Context) {
	if user, ok := ctx.UserValue["user"].(auth.User); ok {
		var input []byte
		form := ctx.Request.MultipartForm
		if form != nil {
			input, _ = json.Marshal((*form).Value)
		}

		models.OperationLog().New(user.ID, ctx.Path(), ctx.Method(), ctx.LocalIP(), string(input))
	}
}
