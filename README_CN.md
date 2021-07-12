<p align="center">
  <a href="https://github.com/GoAdminGroup/go-admin">
    <img width="48%" alt="go-admin" src="http://quick.go-admin.cn/official/assets/imgs/github_logo.png">
  </a>
</p>
<p align="center">
    遗失的Golang编写的数据可视化与管理平台构建框架
</p>
<p align="center">
<a href="http://drone.go-admin.com/GoAdminGroup/go-admin"><img alt="Build Status" src="http://drone.go-admin.com/api/badges/GoAdminGroup/go-admin/status.svg?ref=refs/heads/master"></a>
  <a href="https://goreportcard.com/report/github.com/GoAdminGroup/go-admin"><img alt="Go Report Card" src="https://camo.githubusercontent.com/59eed852617e19c272a4a4764fd09c669957fe75/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f6368656e6867352f676f2d61646d696e"></a>
  <a href="https://goreportcard.com/report/github.com/GoAdminGroup/go-admin"><img alt="golang" src="https://img.shields.io/badge/awesome-golang-blue.svg"></a>
  <a href="https://t.me/joinchat/NlyH6Bch2QARZkArithKvg" rel="nofollow"><img alt="telegram" src="https://img.shields.io/badge/chat%20on-telegram-blue" style="max-width:100%;"></a>
  <a href="https://shang.qq.com/wpa/qunwpa?idkey=ab18729bba609c220f31516a4eb9fce27f582458bd9a865b46523adb5632b873"><img alt="qq群" src="https://img.shields.io/badge/QQ-874825430-yellow.svg"></a>
  <a href="https://godoc.org/github.com/GoAdminGroup/go-admin" rel="nofollow"><img src="https://camo.githubusercontent.com/a9a286d43bdfff9fb41b88b25b35ea8edd2634fc/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f646572656b7061726b65722f64656c76653f7374617475732e737667" alt="GoDoc" data-canonical-src="https://godoc.org/github.com/derekparker/delve?status.svg" style="max-width:100%;"></a>
  <a href="https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/LICENSE" rel="nofollow"><img src="https://img.shields.io/badge/license-Apache2.0-blue.svg" alt="license" data-canonical-src="https://img.shields.io/badge/license-Apache2.0-blue.svg" style="max-width:100%;"></a>
</p>
<p align="center">
    由<a href="https://github.com/z-song/laravel-admin" target="_blank">laravel-admin</a>启发
</p>

## 前言

GoAdmin 可以帮助你的golang应用快速实现数据可视化，搭建一个数据管理平台。

[文档](http://doc.go-admin.cn/zh) | [论坛](http://discuss.go-admin.com) | [Demo](https://demo.go-admin.cn) | [上手例子](https://github.com/GoAdminGroup/example/blob/master/README_CN.md) | [GoAdmin+vue 例子](https://github.com/GoAdminGroup/goadmin-vue-example)


![](http://file.go-admin.cn/introduction/interface_3.png)

## 特征

- 🚀 **高生产效率**: 10分钟内做一个好看的管理后台
- 🎨 **主题**: 默认为adminlte，更多好看的主题正在制作中，欢迎给我们留言
- 🔢 **插件化**: 提供插件使用，真正实现一个插件解决不了问题，那就两个
- ✅ **认证**: 开箱即用的rbac认证系统
- ⚙️ **框架支持**: 支持大部分框架接入，让你更容易去上手和扩展

## 翻译
我们需要您的帮忙： [https://github.com/GoAdminGroup/docs/issues/1](https://github.com/GoAdminGroup/docs/issues/1)

## 谁在使用GoAdmin

[评论这个issue告诉我们](https://github.com/GoAdminGroup/go-admin/issues/71).

## 使用

提示：现在你也可以这样做。

```shell
$ mkdir new_project && cd new_project
$ go install github.com/GoAdminGroup/go-admin/adm
$ adm init -l cn
```

或者：（使用v1.2.16的adm）

```shell
$ mkdir new_project && cd new_project
$ go install github.com/GoAdminGroup/go-admin/adm
$ adm init web -l cn
```


通过以下三步运行：

### 第一步：导入 sql

- [mysql](https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.sql)
- [mssql](https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.mssql)
- [postgresql](https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.pgsql)
- [sqlite](https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.db)

### 第二步：创建 main.go

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
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/GoAdminGroup/go-admin/modules/config"
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
                ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	// 增加 chartjs 组件
	template.AddComp(chartjs.NewChart())
    
    	_ = eng.AddConfig(&cfg).
    		AddGenerators(datamodel.Generators).
    	        // 增加 generator, 第一个参数是对应的访问路由前缀
        	        // 例子:
        	        //
        	        // "user" => http://localhost:9033/admin/info/user
        	        //		
    		AddGenerator("user", datamodel.GetUserTable).
    		Use(r)
    	
    	// 自定义首页
    	eng.HTML("GET", "/admin", datamodel.GetContent)

	_ = r.Run(":9033")
}
```

</p>
</details>

更多框架的例子: [https://github.com/GoAdminGroup/go-admin/tree/master/examples](https://github.com/GoAdminGroup/go-admin/tree/master/examples)

### 第三步：运行

```shell
GO111MODULE=on go run main.go
```

访问：[http://localhost:9033/admin](http://localhost:9033/admin)

账号: admin 密码: admin

更多细节详见 [文档说明](http://doc.go-admin.cn/zh)

[这里一个超级简单上手的例子](https://github.com/GoAdminGroup/example)

## 贡献

[这里有一份贡献指南](CONTRIBUTING_CN.md)

非常欢迎提pr，<strong>这里可以加入开发小组</strong>

<strong>QQ群</strong>：[641768714](https://jq.qq.com/?_wv=1027&k=qn8oXyGC)，记得备注加群来意

这里是[开发计划](https://github.com/GoAdminGroup/go-admin/projects)

<strong>[点击这里申请加微信群（记得备注加群）](http://quick.go-admin.cn/resource/wechat_qrcode_02.jpg)</strong>

<strong>注：在社区中如有问题提问，请务必清晰描述，包括但不限于问题详叙/问题代码/复现方法/已经尝试过的方法，时间生命可贵，请珍惜自己和别人的时间！</strong>

## 十分感谢

inspired by [laravel-admin](https://github.com/z-song/laravel-admin)

## 打赏

留下您的github/gitee用户名，我们将会展示在[捐赠名单](DONATION.md)中。

> 恰饭所需，作者精力时间有限，目前GoAdmin项目捐赠达666元，联系[作者](http://quick.go-admin.cn/resource/wechat_qrcode.jpg)可进vip用户群，vip群中您的问题将得到优先解答，同时也会根据您的需求进行分析和优先安排，vip群也会提供其他关于golang的福利。🙏
>
> 同时您也可以联系我，雇佣我的时间帮助您干活。

<img src="http://quick.go-admin.cn/official/assets/imgs/shoukuan.jpg" width="650" />