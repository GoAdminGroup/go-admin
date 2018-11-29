# Admin插件使用

admin插件可以帮助你实现快速生成数据库数据表增删改查的Web数据管理平台。

需要如下几步：

- 生成数据表对应的配置文件
- 设置访问路由
- 初始化，并在引擎中加载
- 设置访问菜单

## 生成配置文件

假设你的数据库里面有一个数据表Users，如：

```sql
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gender` tinyint(4) DEFAULT NULL,
  `city` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3635 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

使用自带的命令行工具可以帮助你快速生成配置文件，如：

- 安装

```bash
go install github.com/chenhg5/go-admin/admincli
```

- 生成

```bash
admincli generate -h=127.0.0.1 -p=3306 -u=root -P=root -pa=main -n=goadmin -o=./

-h  host
-p  端口
-u  用户名
-P  密码
-n  数据库名
-pa 包名
-o  文件位置

```

运行完之后，会生成一个文件```users.go```，这个就是对应数据表的配置文件了，关于如何配置，在后面详细介绍。

## 设置访问路由

生成完配置文件后，你需要设置访问这个数据表数据的路由，如：

```go
package datamodel

import "github.com/chenhg5/go-admin/plugins/admin/models"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
var Generators = map[string]models.TableGenerator{
	"user":    GetUserTable,
}
```

其中，```"user"```就是对应的前缀，```GetUserTable```就是配置文件中的方法，只要一一对应即可。

## 初始化，并在引擎中加载

初始化，需要调用```NewAdmin```方法，然后将上面的```Generators```传进去即可。然后再调用引擎的```AddPlugins```方法加载引擎。

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
	eng := engine.Default()
	cfg := config.Config{}

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	
	eng.AddConfig(cfg).
		AddPlugins(adminPlugin).  // 加载插件
		Use(r)

	r.Run(":9033")
}
```

## 设置访问菜单

运行起来后，访问登录网址，进入到菜单管理页面，设置好数据表的管理菜单就可以在侧边栏中进入了。