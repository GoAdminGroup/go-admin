package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/chenhg5/go-admin/adapter"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/template/types"
)

func main() {
	app := beego.NewApp()

	eng := engine.Default()

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
		DOMAIN: "localhost",
		PREFIX: "admin",
		INDEX:  "/",
	}

	// you can custom a plugin like:

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.Generators), examplePlugin).
		AddAdapter(new(adapter.Beego)).
		Use(app); err != nil {
		panic(err)
	}

	// you can custom your pages like:

	app.Handlers.Get("/"+cfg.PREFIX+"/custom", func(ctx *context.Context) {
		adapter.BeegoContent(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
	})

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087
	app.Run()
}
