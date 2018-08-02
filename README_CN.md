![](https://ws2.sinaimg.cn/large/006tNc79ly1ft06s5k7y9j31kw0ogmyx.jpg)

# go-admin

[![Go Report Card](https://goreportcard.com/badge/github.com/chenhg5/go-admin)](https://goreportcard.com/report/github.com/chenhg5/go-admin)

遗失的Golang语言编写的Web管理平台构建框架 

对于一个管理平台来说，有几个东西是重要的：

- 安全性和易于使用
- 独立性，独立于业务系统

![](https://ws3.sinaimg.cn/large/006tNc79ly1ft048byoafj31kw0w847v.jpg)

## 特征

- 使用adminlte构建的漂亮的管理界面
- 可配置的，易于管理数据库数据
- 完善的认证系统
- 使用Go编写
- 可移植的
- 部署简单

## 环境要求

- [GO >= 1.8](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

## 安装

目前只能从源码安装，如下：

```
cd $GOPATH/src
git clone https://github.com/chenhg5/go-admin
mv go-admin goAdmin
```

## 使用 

### 第一步：创建一个数据表的配置文件，参数为要管理的表名 

```
./goman create --table=user
```

windows用户请使用 goman.exe

### 第二步：配置表格显示编辑信息

models/user.go

```
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
		}, {
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			Default:  "",
			Editable: true,
			FormType: "text",
		},
	}

	userTable.Form.Table = "users"
	userTable.Form.Title = "用户表"
	userTable.Form.Description = "用户表"

	return
}
```

### 第三步：路由信息配置

- models/global.go

```
package models

// map下标是路由前缀，对应的值是GlobalTable类型，为表单与表格的数据抽象表示
var GlobalTableList = map[string]GlobalTable{
	"user": GetUserTable(),
}

```

- config/config.go

```
package config

var EnvConfig = map[string]interface{}{

    "SERVER_PORT": ":4003",

	"DATABASE_IP":           "127.0.0.1",
	"DATABASE_PORT":         "3306",
	"DATABASE_USER":         "root",
	"DATABASE_PWD":          "root",
	"DATABASE_NAME":         "goadmin",
	"DATABASE_MAX_IDLE_CON": 50,  // 连接池连接数
	"DATABASE_MAX_OPEN_CON": 150, // 最大连接数

	"REDIS_IP":       "127.0.0.1",
	"REDIS_PORT":     "6379",
	"REDIS_PASSWORD": "",
	"REDIS_DB":       1,
	
	"PORTABLE": false,  // 是否跨平台可移植
	
	"AUTH_DOMAIN": "localhost",
}
```

### 第四步：导入sql文件都数据库中

import admin.sql into database

### 第五步：加载依赖

```
make deps
```

### 第六步：运行 ☕

```
make
```
访问 [http://localhost:4003/login](http://localhost:4003/login)

使用账户名：admin， 密码：admin访问

![](https://ws1.sinaimg.cn/large/006tKfTcly1ft3wwounwjj31kw0w17wl.jpg)

## make 命令

- build
- test
- clean
- run
- restart
- deps : 安装依赖
- cross : 跨平台编译
- pages : 将html文件编译为go文件
- assets : 将静态文件编译为go文件
- fmt

## 技术支持

- [fasthttp](https://github.com/valyala/fasthttp)
- [adminlte](https://adminlte.io/themes/AdminLTE/index2.html)
- [hero](https://github.com/shiyanhui/hero)

## todo

- [x] 增加 [go-bindata](https://github.com/go-bindata/go-bindata) 支持
- [X] 增加更多表格表单组件
- [X] 菜单结构
- [ ] rcba认证
- [ ] 自定义页面
- [ ] 合并优化静态资源
- [ ] 自动安装引擎
- [ ] demo网站的搭建
- [ ] 性能分析

## 贡献

非常欢迎提pr，<strong>这里可以加入开发小组</strong>

QQ群: 756664859

## 十分感谢

inspired by [laravel-admin](https://github.com/z-song/laravel-admin)