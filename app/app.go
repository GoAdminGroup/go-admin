package app

import (
	"goAdmin/components/adminlte"
)

type App struct {
	Theme   adminlte.AdminlteStruct
}

var app App

func GetApp() App {
	return app
}

func GetComponents() adminlte.AdminlteComponents {
	return app.Theme.Components
}