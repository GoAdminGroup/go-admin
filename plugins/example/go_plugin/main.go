package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	c "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins"
	e "github.com/GoAdminGroup/go-admin/plugins/example"
)

type Example struct {
	app *context.App
}

var Plugin Example

var config c.Config

func (example Example) InitPlugin() {
	config = c.Get()
	Plugin.app = e.InitRouter(config.Prefix())
	e.SetConfig(config)
}

func (example Example) GetRequest() []context.Path {
	return example.app.Requests
}

func (example Example) GetHandler(url, method string) context.Handlers {
	return plugins.GetHandler(url, method, example.app)
}
