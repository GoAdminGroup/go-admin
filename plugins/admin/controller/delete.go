package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
)

// Delete delete the row from database.
func (h *Handler) Delete(ctx *context.Context) {

	param := guard.GetDeleteParam(ctx)

	//token := ctx.FormValue("_t")
	//
	//if !auth.TokenHelper.CheckToken(token) {
	//	ctx.SetStatusCode(http.StatusBadRequest)
	//	ctx.WriteString(`{"code":400, "msg":"delete fail"}`)
	//	return
	//}

	if err := h.table(param.Prefix, ctx).DeleteData(param.Id); err != nil {
		logger.Error(err)
		response.Error(ctx, "删除失败")
		return
	}

	newToken := h.authSrv().AddToken()

	response.OkWithData(ctx, map[string]interface{}{
		"token": newToken,
	})
}
