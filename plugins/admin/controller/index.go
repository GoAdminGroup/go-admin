package controller

import (
	"html/template"
	"goAdmin/template/adminlte/components"
	"goAdmin/context"
	"goAdmin/modules/page"
)

var AssertRootUrl = ""

func ShowDashboard(ctx *context.Context) {
	page.SetPageContent(AssertRootUrl, ctx, func() components.Panel {
		box := components.InfoBox().SetUrl("/").SetTitle("用户总数").SetValue("1000").GetContent()

		col1 := components.Col().SetContent(box).GetContent()
		col2 := components.Col().SetContent(box).GetContent()
		col3 := components.Col().SetContent(box).GetContent()
		col4 := components.Col().SetContent(box).GetContent()

		row := components.Row().SetContent(col1 + col2 + col3 + col4).GetContent()

		return components.Panel{
			Content:     template.HTML(row),
			Title:       "仪表盘",
			Description: "仪表盘",
		}
	})
}
