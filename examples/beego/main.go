package main

import (
	"github.com/astaxie/beego"
	beegoFw "github.com/chenhg5/go-admin/framework/beego"
	"goAdmin"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
)

func main() {
	app := beego.NewApp()

	ad := goAdmin.Default()

	cfg := config.Config{
		DATABASE_IP:           "127.0.0.1",
		DATABASE_PORT:         "3306",
		DATABASE_USER:         "root",
		DATABASE_PWD:          "root",
		DATABASE_NAME:         "godmin",
		DATABASE_MAX_IDLE_CON: "50",
		DATABASE_MAX_OPEN_CON: "150",

		AUTH_DOMAIN: "localhost",
		LANGUAGE: "cn",
		ADMIN_PREFIX: "admin_goal",
	}

	ad.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.TableFuncConfig)).Use(new(beegoFw.Beego), app)

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087
	app.Run()
}

