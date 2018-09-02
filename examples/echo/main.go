package main

import (
	"github.com/labstack/echo"
	"github.com/chenhg5/go-admin/modules/config"
	echoFw "github.com/chenhg5/go-admin/framework/echo"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/engine"
)

func main() {
	e := echo.New()

	eng := engine.Default()

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

	eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(new(echoFw.Echo), e)


	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
