package main

import (
	"github.com/labstack/echo"
	"goAdmin/modules/config"
	"goAdmin/adapter"
	"goAdmin/plugins/admin"
	"goAdmin/examples/datamodel"
	"goAdmin/engine"
)

func main() {
	e := echo.New()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: []config.Database{
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
		},
		AUTH_DOMAIN:  "localhost",
		LANGUAGE:     "cn",
		ADMIN_PREFIX: "admin",
	}

	adminPlugin := admin.NewAdmin(datamodel.TableFuncConfig)

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin).AddAdapter(new(adapter.Echo)).Use(e); err != nil {
		panic(err)
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
