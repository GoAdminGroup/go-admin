# 插件的使用

插件的使用，只需要调用引擎的```AddPlugins```方法即可。

如：

```go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/chenhg5/go-admin/adapter/gin" // 必须引入，如若不引入，则需要自己定义
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/examples/datamodel"
)

func main() {
	r := gin.Default()
	eng := engine.Default()
	cfg := config.Config{}

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	examplePlugin := example.NewExample()
	
	eng.AddConfig(cfg).
		AddPlugins(adminPlugin, examplePlugin).  // 加载插件
		Use(r)

	r.Run(":9033")
}
```

[返回目录](https://github.com/chenhg5/go-admin/blob/master/docs/cn/index.md)<br>
[上一页：go-admin介绍](https://github.com/chenhg5/go-admin/blob/master/docs/cn/instruction/instruction.md)<br>
[下一页：admin插件的使用](https://github.com/chenhg5/go-admin/blob/master/docs/cn/instruction/plugins/admin.md)