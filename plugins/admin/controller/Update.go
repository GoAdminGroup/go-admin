package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"net/http"
)

// Update update the table row of given id.
func Update(ctx *context.Context) {

	param := guard.GetUpdateParam(ctx)

	err := param.Panel.UpdateDataFromDatabase(param.Value)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "ok",
	})
}
