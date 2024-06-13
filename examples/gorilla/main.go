package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/ChenSee/go-admin/adapter/gorilla"
	_ "github.com/ChenSee/go-admin/modules/db/drivers/mysql"
	_ "github.com/ChenSee/goAdminThemes/adminlte"

	"github.com/ChenSee/go-admin/engine"
	"github.com/ChenSee/go-admin/examples/datamodel"
	"github.com/ChenSee/go-admin/modules/config"
	"github.com/ChenSee/go-admin/modules/language"
	"github.com/ChenSee/go-admin/plugins/example"
	"github.com/ChenSee/go-admin/template"
	"github.com/ChenSee/go-admin/template/chartjs"
	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
	eng := engine.Default()

	cfg := config.Config{
		Env: config.EnvLocal,
		Databases: config.DatabaseList{
			"default": {
				Host:            "127.0.0.1",
				Port:            "3306",
				User:            "root",
				Pwd:             "root",
				Name:            "godmin",
				MaxIdleConns:    50,
				MaxOpenConns:    150,
				ConnMaxLifetime: time.Hour,
				Driver:          config.DriverMysql,
			},
		},
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		UrlPrefix: "admin",
		IndexUrl:  "/",
		Debug:     true,
		Language:  language.EN,
	}

	// customize a plugin

	examplePlugin := example.NewExample()

	template.AddComp(chartjs.NewChart())

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/GoAdminGroup/demo.go-admin.cn/blob/master/main.go#L39
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// eng.AddConfigFromJSON("../datamodel/config.json")

	if err := eng.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
		AddDisplayFilterXssJsFilter().
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", datamodel.GetUserTable).
		AddPlugins(examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	app.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	eng.HTML("GET", "/admin", datamodel.GetContent)

	go func() {
		_ = http.ListenAndServe(":9033", app)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
