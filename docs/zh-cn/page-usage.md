## 页面类型说明

页面的类型表现形式如下：

```go
type Page struct {
	Content     string
	Title       string
	Description string
}
```

```Content```即为页面内容，```Title```即为页面标题，```Description```即为页面简介

## 简单例子

为了展示一个页面，以仪表盘首页为例。首先添加路由，在```router.go```文件中：

```go
func InitRouter() *fasthttprouter.Router {

	router := fasthttprouter.New()

	// 仪表盘
	router.GET("/", AuthMiddleware(controller.ShowDashboard))

    ...

	return router
}
```

然后实现```ShowDashboard```方法，在控制器文件中：

```go
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
```

在 ```func() models.Page {}``` 这个匿名回调函数中去实现自己自定义数据和页面内容，最终返回一个页面类型即可。

接着需要启动菜单，登入系统，访问：[](http://localhost:4003/menu)，添加一个菜单就可以了。

如图：

![](https://ws2.sinaimg.cn/large/0069RVTdly1ftzu5hinpqj31kw0muafw.jpg)