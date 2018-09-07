package main

import (
	"github.com/gin-gonic/gin"
	"goAdmin/adapter"
	"goAdmin/plugins/admin"
	"goAdmin/modules/config"
	"goAdmin/plugins/example"
	"goAdmin/examples/datamodel"
	"goAdmin/engine"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: []config.Database{
			//{
			//	FILE:   "./../datamodel/admin.db",
			//	DRIVER: "sqlite",
			//},
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
			{
				IP:           "127.0.0.1",
				PORT:         "5432",
				USER:         "postgres",
				PWD:          "root",
				NAME:         "godmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "postgresql",
			},
		},
		AUTH_DOMAIN:  "localhost",
		LANGUAGE:     "cn",
		ADMIN_PREFIX: "admin",
	}

	adminPlugin := admin.NewAdmin(datamodel.TableFuncConfig)

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).AddAdapter(new(adapter.Gin)).Use(r); err != nil {
		panic(err)
	}

	r.Run(":9033")
}
