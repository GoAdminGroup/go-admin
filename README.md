![](https://ws2.sinaimg.cn/large/006tNc79ly1ft06s5k7y9j31kw0ogmyx.jpg)

# go-admin

[![Go Report Card](https://goreportcard.com/badge/github.com/chenhg5/go-admin)](https://goreportcard.com/report/github.com/chenhg5/go-admin)

the missing golang admin builder tool. 

as a admin platform. the following principle is important as i see.

- security and easy to use
- independent of business platform

![](https://ws3.sinaimg.cn/large/006tNc79ly1ft048byoafj31kw0w847v.jpg)

## feature

- beautiful admin interface builder powerd by adminlte
- configurable which help manage your database data easily
- powerful auth manage system
- writed by go
- portable
- easy to deploy

## requirements

- [GO >= 1.8](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

## install

only from source now.

```
cd $GOPATH
git clone https://github.com/chenhg5/go-admin
```

## usage 

### Step 1 : create a table config file 

```
./goman create --table=user
```

### Step 2 : table config

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
			ExcuFun: func(value string) string {
				return value
			},
		},
		{
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			ExcuFun: func(value string) string {
				if value == "1" {
					return "男"
				}
				if value == "2" {
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

### Step 3 : route & database config

- config/router.go

```
package config

import "goAdmin/models"

var GlobalTableList = map[string]models.GlobalTable{
	"user" : models.GetUserTable(),
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
}
```

### Step 4 : import sql

import admin.sql into database

### Step 5 : load dependency

```
make deps
```

### Step 6 : Runit & Enjoy ☕

```
make
```
visit [http://localhost:4003/login](http://localhost:4003/login) by browser

login with username:<strong>admin</strong>, password:<strong>admin</strong>

## make command

- build
- test
- clean
- run
- restart
- deps : install dependency
- cross : cross compile
- pages : compile html into go file
- assets : compile assets into go file

## powerd by

- [fasthttp](https://github.com/valyala/fasthttp)
- [adminlte](https://adminlte.io/themes/AdminLTE/index2.html)
- [hero](https://github.com/shiyanhui/hero)

## todo

- [x] add [go-bindata](https://github.com/go-bindata/go-bindata) support
- [X] add more components
- [ ] rcba auth
- [ ] combine assets
- [ ] auto install engine
- [ ] menu structure
- [ ] performance analysis

## contribution

very welcome to pr

<strong>here to join into the develop team</strong>

QQ Group Num: 756664859

## special thanks

inspired by [laravel-admin](https://github.com/z-song/laravel-admin)