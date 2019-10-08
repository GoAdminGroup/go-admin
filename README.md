<p align="center">
  <a href="https://github.com/chenhg5/go-admin">
    <img width="50%" alt="go-admin" src="http://www.go-admin.cn/assets/imgs/goadmin_logo.jpg">
  </a>
</p>

<p align="center">
    the missing golang data admin builder tool.
</p>

<p align="center">
    <a href="http://www.go-admin.cn/en">Documentation</a> | 
    <a href="./README_CN.md">中文文档</a> |
    <a href="http://demo.go-admin.cn/admin">DEMO</a>
</p>

<p align="center">
  <a href="https://api.travis-ci.org/chenhg5/go-admin"><img alt="Go Report Card" src="https://api.travis-ci.org/chenhg5/go-admin.svg?branch=master"></a>
  <a href="https://goreportcard.com/report/github.com/chenhg5/go-admin"><img alt="Go Report Card" src="https://camo.githubusercontent.com/59eed852617e19c272a4a4764fd09c669957fe75/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f6368656e6867352f676f2d61646d696e"></a>
  <a href="https://goreportcard.com/report/github.com/chenhg5/go-admin"><img alt="golang" src="https://img.shields.io/badge/awesome-golang-blue.svg"></a>
  <a href="https://gitter.im/golangadmin/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link" rel="nofollow"><img alt="gitter" src="https://camo.githubusercontent.com/6bb364d591efcfeebc1b9eefaf18a4bdb3fc5158/68747470733a2f2f696d672e736869656c64732e696f2f6769747465722f726f6f6d2f646f63736966796a732f646f63736966792e7376673f7374796c653d666c61742d737175617265" style="max-width:100%;"></a>
  <a href="https://jq.qq.com/?_wv=1027&k=5L3e3kS"><img alt="qq群" src="https://img.shields.io/badge/QQ-756664859-yellow.svg"></a>
  <a href="https://godoc.org/github.com/chenhg5/go-admin" rel="nofollow"><img src="https://camo.githubusercontent.com/a9a286d43bdfff9fb41b88b25b35ea8edd2634fc/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f646572656b7061726b65722f64656c76653f7374617475732e737667" alt="GoDoc" data-canonical-src="https://godoc.org/github.com/derekparker/delve?status.svg" style="max-width:100%;"></a>
  <a href="https://raw.githubusercontent.com/chenhg5/go-admin/master/LICENSE" rel="nofollow"><img src="https://img.shields.io/badge/license-Apache2.0-blue.svg" alt="license" data-canonical-src="https://img.shields.io/badge/license-Apache2.0-blue.svg" style="max-width:100%;"></a>
</p> 

<p align="center">
    Inspired by <a href="https://github.com/z-song/laravel-admin" target="_blank">laravel-admin</a>
</p>

## Preface

goAdmin is a toolkit help you to build a data visualization and manage platform for your golang app.

<strong>Now is the beta version. It means that you may meet some unpredictable bugs. The v1.0 will be released about October 8th.</strong>

demo: [http://demo.go-admin.cn/admin](http://demo.go-admin.cn/admin)
account: admin  password: admin

demo source code: https://github.com/GoAdminGroup/demo

![](http://www.go-admin.cn/assets/imgs/interface.jpg)

## Feature

- beautiful admin interface builder powerd by adminlte
- many plugins to use(working on it)
- powerful auth manage system
- support most of the go web framework

## How to

see the [docs](http://www.go-admin.cn/en) for detail

[a super simple example here](https://github.com/GoAdminGroup/example)

### Step 1: import sql

[https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/admin.sql](https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/admin.sql)

### Step 2: create main.go

<details><summary>main.go</summary>
<p>

```go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/language"
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
				Name:         "godmin",
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
	}

    	// Generators: see https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/tables.go 
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	
	// add generator, first parameter is the url prefix of table when visit.
    	// example:
    	//
    	// "user" => http://localhost:9033/admin/info/user
    	//
    	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	_ = eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	_ = r.Run(":9033")
}
```

</p>
</details>


More Examples: [https://github.com/chenhg5/go-admin/tree/master/examples](https://github.com/chenhg5/go-admin/tree/master/examples)

### Step 3: run

```shell
GO111MODULE=on go run main.go
```

## Powered by

- [adminlte](https://adminlte.io/themes/AdminLTE/index2.html)

## Backers

 Your support will help me do better! [[Become a backer](https://opencollective.com/go-admin#backer)]
 <a href="https://opencollective.com/go-admin#backers" target="_blank"><img src="https://opencollective.com/go-admin/backers.svg?width=890"></a>

## Code Contribution

very welcome to pr.

<strong>here to join into the develop team</strong>

[join slack](https://app.slack.com/client/T029RQSE6/CME0MBX38/thread/C3MSAFD40-1565569187.323500)