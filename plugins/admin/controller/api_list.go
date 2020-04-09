package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
)

func (h *Handler) ApiList(ctx *context.Context) {
	prefix := ctx.Query(constant.PrefixKey)

	panel := h.table(prefix, ctx)

	params := parameter.GetParam(ctx.Request.URL, panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
		panel.GetInfo().GetSort())

	panel, panelInfo, urls, err := h.showTableData(ctx, prefix, params, panel, "api_")
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, map[string]interface{}{
		"panel":  panelInfo,
		"footer": panelInfo.Paginator.GetContent() + panel.GetInfo().FooterHtml,
		"header": aDataTable().GetDataTableHeader() + panel.GetInfo().HeaderHtml,
		"prefix": h.config.PrefixFixSlash(),
		"urls": map[string]string{
			"edit":   urls[0],
			"new":    urls[1],
			"delete": urls[2],
			"export": urls[3],
			"detail": urls[4],
			"info":   urls[5],
			"update": urls[6],
		},
	})
}
