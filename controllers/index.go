package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/components"
	"goAdmin/models"
)

// 显示仪表盘
func ShowDashboard(ctx *fasthttp.RequestCtx) {
	defer GlobalDeferHandler(ctx)

	SetPageContent(ctx, func() models.Page {
		box := components.GetBox().SetUrl("/").SetTitle("用户总数").SetValue("1000").GetContent()

		col1 := components.Col.GetContent(box)
		col2 := components.Col.GetContent(box)
		col3 := components.Col.GetContent(box)
		col4 := components.Col.GetContent(box)

		row := components.Row.GetContent(col1 + col2 + col3 + col4)

		return models.Page{
			Content:row,
			Title:"仪表盘",
			Description:"仪表盘",
		}
	})
}