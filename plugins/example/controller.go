package example

import (
	"goAdmin/context"
	"goAdmin/modules/page"
	"goAdmin/template/adminlte/components"
	"html/template"
)

var AssertRootUrl = ""

func TestHandler(ctx *context.Context) {
	page.SetPageContent(AssertRootUrl, ctx, func() components.Panel {
		box := components.InfoBox().SetUrl("/").SetTitle("例子数据").SetValue("1000").GetContent()

		col1 := components.Col().SetContent(box).GetContent()
		col2 := components.Col().SetContent(box).GetContent()
		col3 := components.Col().SetContent(box).GetContent()
		col4 := components.Col().SetContent(box).GetContent()

		row := components.Row().SetContent(col1 + col2 + col3 + col4).GetContent()

		return components.Panel{
			Content:     template.HTML(row),
			Title:       "这是一个插件例子",
			Description: "这是一个插件例子",
		}
	})
}