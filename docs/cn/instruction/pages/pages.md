# 如何自定义页面

调用引擎的```Content```方法，如下：

```go
package main

import (
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	cfg := config.Config{}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	// 这样子去自定义一个页面：

	r.GET("/"+cfg.PREFIX+"/custom", func(ctx *gin.Context) {
		engine.Content(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
	})

	r.Run(":9033")
}
```

```Content```方法会将内容写入到框架的```context```中。

```GetContent```方法代码如下：

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

[返回目录](https://github.com/chenhg5/go-admin/blob/master/docs/cn/index.md)<br>
[上一页：admin插件的使用](https://github.com/chenhg5/go-admin/blob/master/docs/cn/instruction/plugins/admin.md)<br>
[下一页：介绍页面模块](https://github.com/chenhg5/go-admin/blob/master/docs/cn/instruction/pages/modules.md)