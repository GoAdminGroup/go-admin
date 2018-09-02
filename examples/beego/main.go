package main

import (
	"github.com/astaxie/beego"
	beegoFw "github.com/chenhg5/go-admin/framework/beego"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/engine"
)

func main() {
	app := beego.NewApp()

	eng := engine.Default()

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

	eng.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.TableFuncConfig)).Use(new(beegoFw.Beego), app)

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087
	app.Run()
}

