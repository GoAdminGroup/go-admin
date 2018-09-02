package main

import (
	"github.com/kataras/iris"
	"github.com/chenhg5/go-admin"
	"github.com/chenhg5/go-admin/modules/config"
	irisFw "github.com/chenhg5/go-admin/framework/iris"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
)

func main() {
	app := iris.Default()

	engine := goAdmin.DefaultEngine()

	// goAdmin 全局配置
	cfg := config.Config{
		DATABASE: config.Database{
			IP:           "127.0.0.1",
			PORT:         "3306",
			USER:         "root",
			PWD:          "root",
			NAME:         "godmin",
			MAX_IDLE_CON: 50,
			MAX_OPEN_CON: 150,
			DRIVER:       "mysql",
		},
		AUTH_DOMAIN:  "localhost",
		LANGUAGE:     "cn",         // 语言
		ADMIN_PREFIX: "admin", // 前缀
	}

	// 增删改查管理后台插件
	adminPlugin := admin.NewAdmin(datamodel.TableFuncConfig)

	engine.AddConfig(cfg).AddPlugins(adminPlugin).Use(new(irisFw.Iris), app)

	app.Run(iris.Addr(":8099"))
}