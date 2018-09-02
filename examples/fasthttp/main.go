package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	fastFw "github.com/chenhg5/go-admin/framework/fasthttp"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/engine"
)

func main() {
	router := fasthttprouter.New()

	eng := engine.Default()

	// goAdmin 全局配置
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

	// 增删改查管理后台插件
	adminPlugin := admin.NewAdmin(datamodel.TableFuncConfig)

	eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(new(fastFw.Fasthttp), router)

	var waitChan chan int
	go func() {
		fasthttp.ListenAndServe(":8897", router.Handler)
	}()
	<-waitChan
}
