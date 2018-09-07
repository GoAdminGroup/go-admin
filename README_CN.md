<p align="center">
  <a href="https://github.com/chenhg5/go-admin">
    <img width="50%" alt="go-admin" src="https://ws2.sinaimg.cn/large/006tNc79ly1ftvqf8qeddj31bz07e40e.jpg">
  </a>
</p>
<p align="center">
    遗失的Golang语言编写的Web管理平台构建框架
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
    由<a href="https://github.com/z-song/laravel-admin" target="_blank">laravel-admin</a>启发
</p>

## 前言

对于一个管理平台来说，有几个东西是重要的：

- 安全性和易于使用
- 独立性，独立于业务系统

![](https://cloud.githubusercontent.com/assets/1479100/19625297/3b3deb64-9947-11e6-807c-cffa999004be.jpg)

## 特征

- 使用adminlte构建的漂亮的管理界面
- 大量插件供使用
- 完善的认证系统
- 支持多个web框架：gin, beego, echo...

## 环境要求

- [GO >= 1.8](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

## 使用

详见 [wiki](https://github.com/chenhg5/go-admin/wiki)

### 安装

```go get -v -u github.com/chenhg5/go-admin```

### Gin 例子

```go
package main

import (
	"github.com/gin-gonic/gin"
	"goAdmin/adapter"
	"goAdmin/engine"
	"goAdmin/plugins/admin"
	"goAdmin/modules/config"
	"goAdmin/examples/datamodel"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	// global config
	cfg := config.Config{
		DATABASE: []config.Database{
            {
                IP:           "127.0.0.1",
                PORT:         "3306",
                USER:         "root",
                PWD:          "root",
                NAME:         "godmin",
                MAX_IDLE_CON: 50,
                MAX_OPEN_CON: 150,
                DRIVER:       "mysql",
            },
        },
		AUTH_DOMAIN:  "localhost",
		LANGUAGE:     "cn",         
		ADMIN_PREFIX: "admin", 
	}

	adminPlugin := admin.NewAdmin(datamodel.TableFuncConfig)

	eng.AddConfig(cfg).AddPlugins(adminPlugin).AddAdapter(new(adapter.Gin)).Use(r)

	r.Run(":9033")
}
```

其他例子: [https://github.com/chenhg5/go-admin/tree/master/examples](https://github.com/chenhg5/go-admin/tree/master/examples)

## 技术支持

- [adminlte](https://adminlte.io/themes/AdminLTE/index2.html)

## 贡献

非常欢迎提pr，<strong>这里可以加入开发小组</strong>

QQ群: 756664859

## 十分感谢

inspired by [laravel-admin](https://github.com/z-song/laravel-admin)