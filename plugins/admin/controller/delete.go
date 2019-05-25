package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"net/http"
)

func DeleteData(ctx *context.Context) {
	prefix := ctx.Query("prefix")
	if !models.TableList[prefix].GetDeletable() {
		ctx.Html(http.StatusNotFound, "page not found")
		return
	}

	//token := ctx.FormValue("_t")
	//
	//if !auth.TokenHelper.CheckToken(token) {
	//	ctx.SetStatusCode(http.StatusBadRequest)
	//	ctx.WriteString(`{"code":400, "msg":"删除失败"}`)
	//	return
	//}

	models.TableList[ctx.Query("prefix")].
		DeleteDataFromDatabase(ctx.FormValue("id"))

	newToken := auth.TokenHelper.AddToken()

	ctx.Json(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "删除成功", // TODO: 配置为根据语言返回内容
		"data": newToken,
	})
	return
}
