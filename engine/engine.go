package engine

import (
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/adapter"
	"github.com/chenhg5/go-admin/modules/config"
)

type Engine struct {
	PluginsList []plugins.Plugin
	Adapter     adapter.WebFrameWork
}

func Default() *Engine {
	return new(Engine)
}

func (eng *Engine) Use(router interface{}) error {
	return eng.Adapter.Use(router, eng.PluginsList)
}

func (eng *Engine) AddPlugins(plugs ... plugins.Plugin) *Engine {

	for _, plug := range plugs {
		plug.InitPlugin(config.Get())
	}

	eng.PluginsList = append(eng.PluginsList, plugs...)
	return eng
}

func (eng *Engine) AddConfig(cfg config.Config) *Engine {
	config.Set(cfg)
	return eng
}

func (eng *Engine) AddAdapter(ada adapter.WebFrameWork) {
	eng.Adapter = ada
}
