package main

import (
	"net/http"
	"github.com/chenhg5/go-admin/framework/nethttp"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/engine"
)

func main() {
	mux := http.NewServeMux()

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

	eng.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.TableFuncConfig)).Use(new(nethttp.Http), mux)

	http.ListenAndServe(":9002", mux)
}
