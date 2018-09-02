package example

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/page"
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
	template2 "github.com/chenhg5/go-admin/template"
)

var AssertRootUrl = ""

func TestHandler(ctx *context.Context) {
	page.SetPageContent(AssertRootUrl, ctx, func() types.Panel {
		box := template2.Get(Config.THEME).InfoBox().SetUrl("/").SetTitle("例子数据").SetValue("1000").GetContent()

		col1 := template2.Get(Config.THEME).Col().SetContent(box).GetContent()
		col2 := template2.Get(Config.THEME).Col().SetContent(box).GetContent()
		col3 := template2.Get(Config.THEME).Col().SetContent(box).GetContent()
		col4 := template2.Get(Config.THEME).Col().SetContent(box).GetContent()

		row := template2.Get(Config.THEME).Row().SetContent(col1 + col2 + col3 + col4).GetContent()

		return types.Panel{
			Content:     template.HTML(row),
			Title:       "这是一个插件例子",
			Description: "这是一个插件例子",
		}
	})
}