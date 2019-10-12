package main

import (
	ada "github.com/GoAdminGroup/go-admin/adapter/gorilla"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template/types"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	app := mux.NewRouter()
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
				Driver:     config.DriverMysql,
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

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/GoAdminGroup/go-admin/blob/master/demo/main.go#L30
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// eng.AddConfigFromJson("../datamodel/config.json")

	if err := eng.AddConfig(cfg).AddPlugins(admin.
		NewAdmin(datamodel.Generators).
		AddGenerator("user", datamodel.GetUserTable), examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	app.Handle("/admin", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		engine.Content(ada.Context{Request: request, Response: writer}, func(ctx interface{}) (types.Panel, error) {
			return datamodel.GetContent()
		})
	})).Methods("get")

	log.Println("Listening 9033")
	log.Fatal(http.ListenAndServe(":9033", app))
}
