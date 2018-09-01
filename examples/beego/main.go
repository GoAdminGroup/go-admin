package main

import (
	"github.com/astaxie/beego"
	beegoFw "github.com/chenhg5/go-admin/framework/beego"
	"github.com/chenhg5/go-admin"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
)

func main() {
	app := beego.NewApp()

	ad := goAdmin.Default()

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
		ADMIN_PREFIX: "admin_goal", // 前缀
	}

	ad.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.TableFuncConfig)).Use(new(beegoFw.Beego), app)

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087
	app.Run()
}

