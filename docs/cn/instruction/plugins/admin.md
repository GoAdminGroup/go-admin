# Admin插件使用

admin插件可以帮助你实现快速生成数据库数据表增删改查的Web数据管理平台。

## 快速开始

需要如下几步：

- 生成数据表对应的配置文件
- 设置访问路由
- 初始化，并在引擎中加载
- 设置访问菜单

### 生成配置文件

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

### 设置访问路由

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

### 初始化，并在引擎中加载

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

### 设置访问菜单

运行起来后，访问登录网址，进入到菜单管理页面，设置好数据表的管理菜单就可以在侧边栏中进入了。

## 配置文件介绍

配置文件，如下：

```go
package main

import (
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func GetUsersTable() (usersTable models.Table) {

	usersTable.Info.FieldList = []types.Field{}

	usersTable.Info.Table = "users"
	usersTable.Info.Title = "Users"
	usersTable.Info.Description = "Users"

	usersTable.Form.FormList = []types.Form{}

	usersTable.Form.Table = "users"
	usersTable.Form.Title = "Users"
	usersTable.Form.Description = "Users"

	usersTable.ConnectionDriver = "mysql"

	return
}
```

是一个函数，返回了```models.Table```这个类型对象。以下是```models.Table```的定义：

```go
type Table struct {
	Info             types.InfoPanel
	Form             types.FormPanel
	ConnectionDriver string
}
```

包括了```Info```和```Form```，这两种类型对应的ui就是显示数据的表格和编辑新建数据的表单，截图展示如下：

- 此为```Info```表格

![](https://ws4.sinaimg.cn/large/006tNbRwly1fxoy26qnc5j31y60u0q91.jpg)

- 此为```Form```表单

![](https://ws1.sinaimg.cn/large/006tNbRwly1fxoy2w3cobj318k0ooabv.jpg)

### Info表格

```go
type InfoPanel struct {
	FieldList   []Field  // 字段类型
	Table       string         // 表格
	Title       string         // 标题
	Description string         // 描述
}

type Field struct {
	ExcuFun  FieldValueFun     // 过滤函数
	Field    string            // 字段名
	TypeName string            // 字段类型名
	Head     string            // 标题
	Sortable bool              // 是否可以排序
	Filter   bool              // 是否可以筛选
}
```

### Form表单

```go
type FormPanel struct {
	FormList   []Form    // 字段类型
	Table       string         // 表格
	Title       string         // 标题
	Description string         // 描述
}

type Form struct {
	Field    string                // 字段名
	TypeName string                // 字段类型名
	Head     string                // 标题
	Default  string                // 默认
	Editable bool                  // 是否可编辑
	FormType string                // 表单类型
	Value    string                // 表单默认值
	Options  []map[string]string   // 表单选项
	ExcuFun  FieldValueFun         // 过滤函数
}
```

[返回目录](https://github.com/chenhg5/go-admin/blob/master/docs/cn/index.md)<br>
[上一页：插件的使用](https://github.com/chenhg5/go-admin/blob/master/docs/cn/instruction/plugins/plugins.md)<br>
[下一页：自定义页面](https://github.com/chenhg5/go-admin/blob/master/docs/cn/instruction/pages/pages.md)