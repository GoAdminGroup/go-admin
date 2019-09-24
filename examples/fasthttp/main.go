package main

import (
	"github.com/buaazp/fasthttprouter"
	_ "github.com/chenhg5/go-admin/adapter/fasthttp"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()

	eng := engine.Default()

	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       "127.0.0.1",
				Port:       "3306",
				User:       "root",
				Pwd:        "root",
				Name:       "godmin",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     db.DriverMysql,
			},
		},
		Domain:    "localhost",
		UrlPrefix: "admin",
		IndexUrl:  "/",
		Debug:     true,
		Language:  language.CN,
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:8897/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	// customize a plugin

	examplePlugin := example.NewExample()

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/chenhg5/go-admin/blob/master/demo/main.go#L30
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// eng.AddConfigFromJson("../datamodel/config.json")

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(router); err != nil {
		panic(err)
	}

	router.GET("/admin/custom", func(ctx *fasthttp.RequestCtx) {
		engine.Content(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
	})

	var waitChan chan int
	go func() {
		_ = fasthttp.ListenAndServe(":8897", router.Handler)
	}()
	<-waitChan
}
