# Quick Start

Suppose there is a users table in our business database that needs to be managed, as follows:

```sql
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nickname` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `wx_unionid` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `sex` tinyint(4) DEFAULT NULL,
  `city` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `province` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `label` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unionid` (`wx_unionid`)
) ENGINE=InnoDB AUTO_INCREMENT=3131 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### Generate a corresponding configuration file for the data table

```bash
./goman_mac create --table=users
```

After the execution is completed, the corresponding file ```users.go`` will be generated in the ```models``` folder.

### Modify configuration file


models/user.go

```go
package models

func GetUserTable() (userTable GlobalTable) {
    
    // 列显示配置
	userTable.Info.FieldList = []FieldStruct{
		{
			Head:     "姓名",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
                return model.Value
            },
		},
		{
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			ExcuFun: func(model RowModel) string {
				if model.Value == "1" {
					return "男"
				}
				if model.Value == "2" {
					return "女"
				}
				return "未知"
			},
		},
	}

	userTable.Info.Table = "users"
	userTable.Info.Title = "用户表"
	userTable.Info.Description = "用户表"

    // 表单显示配置
	userTable.Form.FormList = []FormStruct{
		{
			Head:     "姓名",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) string {
                return model.Value
            },
		}, {
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) string {
                return model.Value
            },
		},
	}

	userTable.Form.Table = "users"
	userTable.Form.Title = "用户表"
	userTable.Form.Description = "用户表"

	return
}
```

### Add route

- models/global.go

```go
package models

type GetTableDataFunc func() GlobalTable

var TableFuncConfig = map[string]GetTableDataFunc{
	// 管理员管理部分
	"manager":    GetManagerTable,    // 管理员管理
	"permission": GetPermissionTable, // 权限管理
	"roles":      GetRolesTable,      // 角色管理
	"op":         GetOpTable,         // 操作日志管理

	// 自定义管理部分
	"user": GetUserTable,
}
```

### Runit

```bash
make
```
visit [http://localhost:4003/login](http://localhost:4003/login)

use account：admin， password：admin

![](https://ws1.sinaimg.cn/large/006tKfTcly1ft3wwounwjj31kw0w17wl.jpg)