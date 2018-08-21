package controller

import (
	"github.com/valyala/fasthttp"
	"html/template"
	"goAdmin/app"
	"goAdmin/components"
)

// 显示仪表盘
func ShowDashboard(ctx *fasthttp.RequestCtx) {
	defer GlobalDeferHandler(ctx)

	SetPageContent(ctx, func() components.Panel {
		box := app.GetApp().Theme.Components.Box().SetUrl("/").SetTitle("用户总数").SetValue("1000").GetContent()

		col1 := app.GetApp().Theme.Components.Col().SetContent(string(box)).GetContent()
		col2 := app.GetApp().Theme.Components.Col().SetContent(string(box)).GetContent()
		col3 := app.GetApp().Theme.Components.Col().SetContent(string(box)).GetContent()
		col4 := app.GetApp().Theme.Components.Col().SetContent(string(box)).GetContent()

		row := app.GetApp().Theme.Components.Row().SetContent(string(col1) + string(col2) + string(col3) + string(col4)).GetContent()

		return components.Panel{
			Content:     template.HTML(row),
			Title:       "仪表盘",
			Description: "仪表盘",
		}
	})
}
