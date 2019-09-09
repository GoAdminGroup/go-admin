package main

import (
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: config.DatabaseList{
			"default": {
				HOST:         database.HOST,
				PORT:         database.PORT,
				USER:         database.USER,
				PWD:          database.PWD,
				NAME:         database.NAME,
				MAX_IDLE_CON: database.MAX_IDLE_CON,
				MAX_OPEN_CON: database.MAX_OPEN_CON,
				DRIVER:       database.DRIVER,
			},
		},
		DOMAIN: "demo.go-admin.cn",
		PREFIX: "admin",
		STORE: config.Store{
			PATH:   "/data/www/go-admin/uploads",
			PREFIX: "uploads",
		},
		LANGUAGE: language.EN,
		INDEX:    "/",
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:9033/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	// you can custom a plugin like:

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "/data/www/go-admin/uploads")

	// you can custom your pages like:

	r.GET("/"+cfg.PREFIX+"/custom", func(ctx *gin.Context) {
		engine.Content(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
	})

	_ = r.Run(":9033")
}
