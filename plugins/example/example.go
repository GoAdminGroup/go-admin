package example

import (
	"goAdmin/context"
	"goAdmin/plugins"
	"goAdmin/modules/config"
	"goAdmin/modules/menu"
)

type Example struct {
	app *context.App
}

var ExamplePlug = new(Example)

func (example *Example) InitPlugin(cfg config.Config) {
	ExamplePlug.app = InitRouter("/" + cfg.ADMIN_PREFIX)

	menu.InitMenu()
}

func NewExample() *Example {
	return ExamplePlug
}

func (example *Example) GetRequest() []context.Path {
	return example.app.Requests
}

func (example *Example) GetHandler(url, method string) context.Handler {
	return plugins.GetHandler(url, method, &example.app.HandlerList)
}
