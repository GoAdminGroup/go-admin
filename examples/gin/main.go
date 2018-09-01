package main

import (
	"github.com/gin-gonic/gin"
	ginFw "github.com/chenhg5/go-admin/framework/gin"
	"goAdmin"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/examples/datamodel"
)

func main() {
	r := gin.Default()

	ad := goAdmin.Default()

	// goAdmin 全局配置
	cfg := config.Config{
		DATABASE_IP:           "127.0.0.1",
		DATABASE_PORT:         "3306",
		DATABASE_USER:         "root",
		DATABASE_PWD:          "root",
		DATABASE_NAME:         "godmin",
		DATABASE_MAX_IDLE_CON: "50",
		DATABASE_MAX_OPEN_CON: "150",
		DATABASE_DRIVER:       "mysql",

		AUTH_DOMAIN:  "localhost",
		LANGUAGE:     "cn",         // 语言
		ADMIN_PREFIX: "admin_goal", // 前缀
	}

	// 增删改查管理后台插件
	adminPlugin := admin.NewAdmin(datamodel.TableFuncConfig)
	// 后台插件例子
	examplePlugin := example.NewExample()

	ad.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(new(ginFw.Gin), r)

	r.Run(":9033")
}


