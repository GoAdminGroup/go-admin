<p align="center">
  <a href="https://github.com/GoAdminGroup/go-admin">
    <img width="48%" alt="go-admin" src="http://quick.go-admin.cn/official/assets/imgs/github_logo.png">
  </a>
</p>

<p align="center">
    the missing golang data admin panel builder tool.
</p>

<p align="center">
    <a href="https://book.go-admin.cn/en">Documentation</a> | 
	<a href="http://doc.go-admin.cn/zh/">中文文档</a> | 
    <a href="./README_CN.md">中文介绍</a> |
    <a href="https://demo.go-admin.com">DEMO</a> |
    <a href="https://demo.go-admin.cn">中文DEMO</a> |
    <a href="https://twitter.com/cg3365688034">Twitter</a> |
    <a href="http://discuss.go-admin.com">Forum</a>
</p>

<p align="center">
  <a href="http://drone.go-admin.com/GoAdminGroup/go-admin"><img alt="Build Status" src="http://drone.go-admin.com/api/badges/GoAdminGroup/go-admin/status.svg?ref=refs/heads/master"></a>
  <a href="https://goreportcard.com/report/github.com/GoAdminGroup/go-admin"><img alt="Go Report Card" src="https://camo.githubusercontent.com/59eed852617e19c272a4a4764fd09c669957fe75/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f6368656e6867352f676f2d61646d696e"></a>
  <a href="https://goreportcard.com/report/github.com/GoAdminGroup/go-admin"><img alt="golang" src="https://img.shields.io/badge/awesome-golang-blue.svg"></a>
  <a href="https://t.me/joinchat/NlyH6Bch2QARZkArithKvg" rel="nofollow"><img alt="telegram" src="https://img.shields.io/badge/chat%20on-telegram-blue" style="max-width:100%;"></a>
  <a href="https://goadmin.slack.com"><img alt="slack" src="https://img.shields.io/badge/chat on-Slack-yellow.svg"></a>
  <a href="https://godoc.org/github.com/GoAdminGroup/go-admin" rel="nofollow"><img src="https://camo.githubusercontent.com/a9a286d43bdfff9fb41b88b25b35ea8edd2634fc/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f646572656b7061726b65722f64656c76653f7374617475732e737667" alt="GoDoc" data-canonical-src="https://godoc.org/github.com/derekparker/delve?status.svg" style="max-width:100%;"></a>
  <a href="https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/LICENSE" rel="nofollow"><img src="https://img.shields.io/badge/license-Apache2.0-blue.svg" alt="license" data-canonical-src="https://img.shields.io/badge/license-Apache2.0-blue.svg" style="max-width:100%;"></a>
</p> 

<p align="center">
    Inspired by <a href="https://github.com/z-song/laravel-admin" target="_blank">laravel-admin</a>
</p>

## Preface

GoAdmin is a toolkit to help you build a data visualization admin panel for your golang app.

Online demo: [https://demo.go-admin.com](https://demo.go-admin.com)

Quick follow up example: [https://github.com/GoAdminGroup/example](https://github.com/GoAdminGroup/example)

GoAdmin+vue example: [https://github.com/GoAdminGroup/goadmin-vue-example](https://github.com/GoAdminGroup/goadmin-vue-example)

![interface](http://file.go-admin.cn/introduction/interface_en_3.png)

## Features

- 🚀 **Fast**: build a production admin panel app in **ten** minutes.
- 🎨 **Theming**: beautiful ui themes supported(default adminlte, more themes are coming.)
- 🔢 **Plugins**: many plugins to use(more useful and powerful plugins are coming.)
- ✅ **Rbac**: out of box rbac auth system.
- ⚙️ **Frameworks**: support most of the go web frameworks.

## Translation
We need your help: [https://github.com/GoAdminGroup/docs/issues/1](https://github.com/GoAdminGroup/docs/issues/1)

## Who is using

[Comment the issue to tell us](https://github.com/GoAdminGroup/go-admin/issues/71).

## How to

Following three steps to run it.

Note: now you can quickly start by doing like this.

```shell
$ mkdir new_project && cd new_project
$ go install github.com/GoAdminGroup/go-admin/adm
$ adm init
```

Or (use adm whose version higher or equal than v1.2.16)

```shell
$ mkdir new_project && cd new_project
$ go install github.com/GoAdminGroup/go-admin/adm
$ adm init web
```

### Step 1: import sql

- [mysql](https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.sql)
- [mssql](https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.mssql)
- [postgresql](https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.pgsql)
- [sqlite](https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.db)

### Step 2: create main.go

<details><summary>main.go</summary>
<p>

```go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/language"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	// global config
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:         "127.0.0.1",
				Port:         "3306",
				User:         "root",
				Pwd:          "root",
				Name:         "goadmin",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:       "mysql",
			},
        	},
		UrlPrefix: "admin",
		// STORE is important. And the directory should has permission to write.
		Store: config.Store{
		    Path:   "./uploads", 
		    Prefix: "uploads",
		},
		Language: language.EN,
		// debug mode
		Debug: true,
		// log file absolute path
		InfoLogPath: "/var/logs/info.log",
		AccessLogPath: "/var/logs/access.log",
		ErrorLogPath: "/var/logs/error.log",
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	// add component chartjs
	template.AddComp(chartjs.NewChart())

	_ = eng.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
	        // add generator, first parameter is the url prefix of table when visit.
    	        // example:
    	        //
    	        // "user" => http://localhost:9033/admin/info/user
    	        //		
		AddGenerator("user", datamodel.GetUserTable).
		Use(r)
	
	// customize your pages
	eng.HTML("GET", "/admin", datamodel.GetContent)

	_ = r.Run(":9033")
}
```

</p>
</details>

More framework examples: [https://github.com/GoAdminGroup/go-admin/tree/master/examples](https://github.com/GoAdminGroup/go-admin/tree/master/examples)

### Step 3: run

```shell
GO111MODULE=on go run main.go
```

visit: [http://localhost:9033/admin](http://localhost:9033/admin)

account: admin password: admin

[A super simple example here](https://github.com/GoAdminGroup/example)

See the [docs](https://book.go-admin.cn) for more details.

## Backers

 Your support will help me do better! [[Become a backer](https://opencollective.com/go-admin#backer)]
 <a href="https://opencollective.com/go-admin#backers" target="_blank"><img src="https://opencollective.com/go-admin/backers.svg?width=890"></a>

## Contribution

[here for contribution guide](CONTRIBUTING.md)

<strong>here to join into the develop team</strong>

[join telegram](https://t.me/joinchat/NlyH6Bch2QARZkArithKvg)
