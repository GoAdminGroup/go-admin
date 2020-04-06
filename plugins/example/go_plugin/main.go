package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	c "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
	e "github.com/GoAdminGroup/go-admin/plugins/example"
)

type Example struct {
	app  *context.App
	name string
}

var Plugin = new(Example)

var config c.Config

func (example *Example) InitPlugin(srv service.List) {
	config = c.Get()
	Plugin.app = e.InitRouter(config.Prefix(), srv)
	e.SetConfig(config)
}

func (example *Example) GetRequest() []context.Path {
	return example.app.Requests
}

func (example *Example) GetHandler() context.HandlerMap {
	return plugins.GetHandler(example.app)
}

func (example *Example) Name() string {
	return example.name
}
