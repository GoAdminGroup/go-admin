package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	e := engine.Default()

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
				Driver:     config.DriverMysql,

				//Driver: config.DriverSqlite,
				//File:   "../datamodel/admin.db",
			},
		},
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.CN,
		IndexUrl: "/",
		Debug:    true,
		Animation: config.PageAnimation{
			Type: "fadeInUp",
		},
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	template.AddComp(chartjs.NewChart())

	// customize a plugin

	examplePlugin := example.NewExample()

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/GoAdminGroup/demo.go-admin.cn/blob/master/main.go#L39
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// e.AddConfigFromJSON("../datamodel/config.json")

	if err := e.AddConfig(cfg).
		AddGenerators(datamodel.Generators).
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", datamodel.GetUserTable).
		AddDisplayFilterXssJsFilter().
		AddPlugins(examplePlugin).
		AddNavButtons("Website Info", "", action.PopUp("/website/info", "Website Info",
			func(ctx *context.Context) (success bool, msg string, data interface{}) {
				return true, "ok", `<p>created by <a href="https://github.com/chenhg5">cg33<a/></p>`
			})).
		AddNavButtons("Google", "", action.Jump("https://www.google.com")).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	// customize your pages

	e.HTML("GET", "/admin", datamodel.GetContent)

	go func() {
		_ = r.Run(":9033")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	e.MysqlConnection().Close()
}
