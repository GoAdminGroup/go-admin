<p align="center">
  <a href="https://github.com/chenhg5/go-admin">
    <img width="50%" alt="go-admin" src="http://www.go-admin.cn/assets/imgs/goadmin_logo.jpg">
  </a>
</p>
<p align="center">
    遗失的Golang语言编写的数据可视化与管理平台构建框架
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
    由<a href="https://github.com/z-song/laravel-admin" target="_blank">laravel-admin</a>启发
</p>

## 前言

goAdmin 可以帮助你的golang应用快速实现数据可视化，搭建一个数据管理平台。

demo: [http://demo.go-admin.cn/admin](http://demo.go-admin.cn/admin)
账号：admin  密码：admin

demo代码： https://github.com/GoAdminGroup/demo

<strong>现在是beta版本，可能存在一些未知的bug。 正式的1.0版本将会在10月8号左右发布</strong>

![](http://www.go-admin.cn/assets/imgs/interface.jpg)

## 特征

- 使用adminlte构建的漂亮的管理界面
- 大量插件供使用（开发中）
- 完善的认证系统
- 支持多个web框架：gin, beego, echo...

## 使用

详见 [文档说明](http://www.go-admin.cn)

[一个超级简单的例子](https://github.com/GoAdminGroup/example)

### 第一步：导入 sql

以mysql为例：

[https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/admin.sql](https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/admin.sql)

### 第二步：创建 main.go

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
		// STORE 必须设置且保证有写权限，否则增加不了新的管理员用户
		Store: config.Store{
		    Path:   "./uploads",
		    Prefix: "uploads",
		},
		Language: language.CN, 
		// 开发模式
                Debug: true,
                // 日志文件位置，需为绝对路径
                InfoLogPath: "/var/logs/info.log",
                AccessLogPath: "/var/logs/access.log",
                ErrorLogPath: "/var/logs/error.log",
	}

    	// Generators： 详见 https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	
	// 增加 generator, 第一个参数是对应的访问路由前缀
	// 例子:
	//
	// "user" => http://localhost:9033/admin/info/user
	//
	// adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	_ = eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	_ = r.Run(":9033")
}
```

</p>
</details>

其他例子: [https://github.com/chenhg5/go-admin/tree/master/examples](https://github.com/chenhg5/go-admin/tree/master/examples)

### 第三步：运行

```shell
GO111MODULE=on go run main.go
```

## 技术支持

- [adminlte](https://adminlte.io/themes/AdminLTE/index2.html)

## 贡献

非常欢迎提pr，<strong>这里可以加入开发小组</strong>

<strong>QQ群</strong>：756664859，记得备注加群来意

这里是[开发计划](https://github.com/chenhg5/go-admin/projects)

<strong>[点击这里加微信群](http://www.go-admin.cn/assets/imgs/qrcode.jpg)</strong>

## 十分感谢

inspired by [laravel-admin](https://github.com/z-song/laravel-admin)

## 打赏

留下您的github/gitee用户名，我们将会展示在[捐赠名单](DONATION.md)中。

<img src="http://www.go-admin.cn/assets/imgs/shoukuan.jpg" width="650" />