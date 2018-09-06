package main

import (
	"net/http"
	"goAdmin/adapter"
	"goAdmin/plugins/admin"
	"goAdmin/examples/datamodel"
	"goAdmin/modules/config"
	"goAdmin/engine"
)

func main() {
	mux := http.NewServeMux()

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

	if err := eng.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.TableFuncConfig)).AddAdapter(new(adapter.Http)).Use(mux); err != nil {
		panic(err)
	}

	http.ListenAndServe(":9002", mux)
}
