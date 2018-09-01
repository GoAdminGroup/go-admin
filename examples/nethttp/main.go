package main

import (
	"net/http"
	"github.com/chenhg5/go-admin/framework/nethttp"
	"github.com/chenhg5/go-admin"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
)

func main() {
	mux := http.NewServeMux()

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

	ad.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.TableFuncConfig)).Use(new(nethttp.Http), mux)

	http.ListenAndServe(":9002", mux)
}
