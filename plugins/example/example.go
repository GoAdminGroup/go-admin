package example

import (
	"github.com/chenhg5/go-admin/context"
	c "github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins"
)

type Example struct {
	app *context.App
}

var Plug = new(Example)

var config c.Config

func (example *Example) InitPlugin() {
	config = c.Get()
	Plug.app = InitRouter(config.Prefix())
}

func NewExample() *Example {
	return Plug
}

func (example *Example) GetRequest() []context.Path {
	return example.app.Requests
}

func (example *Example) GetHandler(url, method string) context.Handlers {
	return plugins.GetHandler(url, method, example.app)
}
