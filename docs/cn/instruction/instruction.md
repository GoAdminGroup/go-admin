# 使用go-admin

go-admin通过各种适配器使得你在各个web framework中使用都十分的方便。

## 例子

先导入```$GOPATH/github.com/chenhg5/go-admin/examples/datamodel/admin.sql```到数据库中。

下面看一个Gin框架的例子：

```go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/chenhg5/go-admin/adapter/gin" // 必须引入，如若不引入，则需要自己定义
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/examples/datamodel"
)

func main() {
	r := gin.Default()

        // 实例化一个go-admin引擎对象
	eng := engine.Default()

	// go-admin全局配置
	cfg := config.Config{
		DATABASE: []config.Database{
		    {
			HOST:         "127.0.0.1",
			PORT:         "3306",
			USER:         "root",
			PWD:          "root",
			NAME:         "godmin",
			MAX_IDLE_CON: 50,
			MAX_OPEN_CON: 150,
			DRIVER:       "mysql",
		    },
        	},
		DOMAIN: "localhost", // 是cookie相关的，访问网站的域名
		PREFIX: "admin",
		// STORE 必须设置且保证有写权限，否则增加不了新的管理员用户
		STORE: config.Store{
		    PATH:   "./uploads",
		    PREFIX: "uploads",
		},
		LANGUAGE: "cn", 
	}

    	// Generators： 详见 https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(datamodel.Generators)

        // 增加配置与插件，使用Use方法挂载到Web框架中
	eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	r.Run(":9033")
}
```

对应的步骤都加上了注释，使用十分简单，只需要：

- 引入适配器
- 设置全局的配置项
- 初始化插件
- 设置插件与配置
- 挂载到Web框架中

更多的例子可以看：[https://github.com/chenhg5/go-admin/tree/master/examples](https://github.com/chenhg5/go-admin/tree/master/examples)

## 配置项

```go
package config

import (
	"html/template"
)

type Database struct {
	HOST         string
	PORT         string
	USER         string
	PWD          string
	NAME         string
	MAX_IDLE_CON int
	MAX_OPEN_CON int
	DRIVER       string

	FILE string
}

type Store struct {
	PATH   string
	PREFIX string
}

type Config struct {
	// 数据库配置
	DATABASE []Database

	// 登录域名
	DOMAIN string

	// 网站语言
	LANGUAGE string

	// 全局的管理前缀
	PREFIX string

	// 主题名
	THEME string

	// 上传文件存储的位置
	STORE Store

	// 网站的标题
	TITLE string

	// 侧边栏上的Logo
	LOGO template.HTML

	// 侧边栏上的Logo缩小版
	MINILOGO template.HTML

	// 登录后跳转的路由
	INDEX string
}

```

[返回目录](https://github.com/chenhg5/go-admin/blob/master/docs/cn/index.md)<br>
[下一页：插件的使用](https://github.com/chenhg5/go-admin/blob/master/docs/cn/instruction/plugins/plugins.md)