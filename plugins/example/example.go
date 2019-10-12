package example

import (
	"github.com/GoAdminGroup/go-admin/context"
	c "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins"
)

type Example struct {
	app *context.App
}

func NewExample() *Example {
	return Plug
}

var Plug = new(Example)

var config c.Config

func SetConfig(cfg c.Config) {
	config = cfg
}

func (example *Example) InitPlugin() {
	config = c.Get()
	Plug.app = InitRouter(config.Prefix())
}

func (example *Example) GetRequest() []context.Path {
	return example.app.Requests
}

func (example *Example) GetHandler(url, method string) context.Handlers {
	return plugins.GetHandler(url, method, example.app)
}
