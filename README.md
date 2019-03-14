<p align="center">
  <a href="https://github.com/chenhg5/go-admin">
    <img width="50%" alt="go-admin" src="https://ws2.sinaimg.cn/large/006tNc79ly1ftvqf8qeddj31bz07e40e.jpg">
  </a>
</p>

<p align="center">
    the missing golang data admin builder tool.
</p>

<p align="center">
    <a href="https://github.com/chenhg5/go-admin/wiki">Documentation</a> | 
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
  <a href="https://raw.githubusercontent.com/chenhg5/go-admin/master/LICENSE" rel="nofollow"><img src="https://camo.githubusercontent.com/e0d5267d60ee425acfe1a1f2d6e6d92a465dcd8f/687474703a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d4d49542d626c75652e737667" alt="license" data-canonical-src="http://img.shields.io/badge/license-MIT-blue.svg" style="max-width:100%;"></a>
</p> 

<p align="center">
    Inspired by <a href="https://github.com/z-song/laravel-admin" target="_blank">laravel-admin</a>
</p>

## Preface

goAdmin is a toolkit help you to build a data visualization and manage platform for your golang app.

demo: [http://demo.go-admin.cn/admin](http://demo.go-admin.cn/admin)
account: admin  password: admin

![](https://ws1.sinaimg.cn/large/0069RVTdly1fv5jpbug82j31ap0pngrr.jpg)

## Feature

- beautiful admin interface builder powerd by adminlte
- many plugins to use
- powerful auth manage system
- support Most of the go web framework

## How to

see the [docs](https://github.com/chenhg5/go-admin/blob/master/docs/cn/index.md) for detail

### install

```go get -v -u github.com/chenhg5/go-admin```

### import sql

[https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/admin.sql](https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/admin.sql)

### generate the data model use cli tool

```
go install github.com/chenhg5/go-admin/admincli

admincli generate -h=127.0.0.1 -p=3306 -P=root -n=godmin -pa=main -o=./model
```

### gin example

```go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/examples/datamodel"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	// global config
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
		DOMAIN: "localhost", // the domain of cookie which be used when visiting your site.
		PREFIX: "admin",
		// STORE is important. And the directory should has permission to write.
		STORE: config.Store{
		    PATH:   "./uploads", 
		    PREFIX: "uploads",
		},
		LANGUAGE: "en",
		// debug mode
		DEBUG: true,
		// log file absolute path
		INFOLOG: "/var/logs/info.log",
		ACCESSLOG: "/var/logs/access.log",
		ERRORLOG: "/var/logs/error.log",
	}

    	// Generators: see https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/tables.go 
	adminPlugin := admin.NewAdmin(datamodel.Generators)

	eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	r.Run(":9033")
}
```

More Examples: [https://github.com/chenhg5/go-admin/tree/master/examples](https://github.com/chenhg5/go-admin/tree/master/examples)

## Powerd by

- [adminlte](https://adminlte.io/themes/AdminLTE/index2.html)

## Contribution

very welcome to pr

<strong>here to join into the develop team</strong>

QQ Group Num: 756664859, remember to add the reason of apply.

## Special thanks

inspired by [laravel-admin](https://github.com/z-song/laravel-admin)
