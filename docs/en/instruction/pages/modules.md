# Page Modularity

Call the engine's ```Content``` method, that returns an object ```types.Panel``` for page customization

```types.Panel``` definition：

```go
type Panel struct {
	Content     template.HTML // page content
	Title       string        // page title
	Description string        // page description
	Url         string
}
```

![](https://ws3.sinaimg.cn/large/006tNbRwly1fxoz5bm02oj31ek0u0wtz.jpg)

## Usage

```go
package datamodel

import (
	"github.com/chenhg5/go-admin/modules/config"
	template2 "github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

func GetContent() types.Panel {

	components := template2.Get(config.Get().THEME)
	colComp := components.Col()

	infobox := components.InfoBox().
		SetText("CPU TRAFFIC").
		SetColor("blue").
		SetNumber("41,410").
		SetIcon("ion-ios-gear-outline").
		GetContent()

	var size = map[string]string{"md": "3", "sm": "6", "xs": "12"}
	infoboxCol1 := colComp.SetSize(size).SetContent(infobox).GetContent()
	row1 := components.Row().SetContent(infoboxCol1).GetContent()

	return types.Panel{
		Content:     row1,
		Title:       "Dashboard",
		Description: "this is a example",
	}
}
```

## Column

Column is represented by the ```ColAttribute``` interface：

```go
type ColAttribute interface {
	SetSize(value map[string]string) ColAttribute
	SetContent(value template.HTML) ColAttribute
	GetContent() template.HTML
}
```

```size``` looks like: ```map[string]string{"md": "3", "sm": "6", "xs": "12"}```

## Row

Row is represented by the ```RowAttribute``` interface:

```go
type RowAttribute interface {
	SetContent(value template.HTML) RowAttribute
	GetContent() template.HTML
}
```

[Back to Contents](https://github.com/chenhg5/go-admin/blob/master/docs/en/index.md)<br>
[Previous: Page Customization](https://github.com/chenhg5/go-admin/blob/master/docs/en/instruction/pages/pages.md)<br>
[Next: Page Components](https://github.com/chenhg5/go-admin/blob/master/docs/en/instruction/pages/components.md)
