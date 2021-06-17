package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gear"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/themes/sword"
	"github.com/teambition/gear"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/teambition/gear/middleware/static"
)

func main() {

	app := gear.New()

	e := engine.Default()

	cfg := config.Config{
		Env: config.EnvLocal,
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
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		UrlPrefix:          "admin",
		Language:           language.CN,
		IndexUrl:           "/",
		Debug:              true,
		AccessAssetsLogOff: true,
		Animation: config.PageAnimation{
			Type: "fadeInUp",
		},
		ColorScheme:       adminlte.ColorschemeSkinBlack,
		BootstrapFilePath: "./../datamodel/bootstrap.go",
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

	app.Use(static.New(static.Options{Root: "./uploads", Prefix: "uploads"}))

	if err := e.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", datamodel.GetUserTable).
		AddDisplayFilterXssJsFilter().
		AddPlugins(examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	// customize your pages

	e.HTML("GET", "/admin", datamodel.GetContent)

	go func() {
		app.Start(":8099")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	e.MysqlConnection().Close()
}
