package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"goAdmin/adapter"
	"goAdmin/modules/config"
	"goAdmin/plugins/admin"
	"goAdmin/examples/datamodel"
	"goAdmin/engine"
	"github.com/prometheus/common/log"
)

func main() {
	router := fasthttprouter.New()

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

	log.Fatal(eng.AddConfig(cfg).AddPlugins(adminPlugin).AddAdapter(new(adapter.Fasthttp)).Use(router))

	var waitChan chan int
	go func() {
		fasthttp.ListenAndServe(":8897", router.Handler)
	}()
	<-waitChan
}
