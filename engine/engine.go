// Copyright 2018 ChenHonggui.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package engine

import (
	"github.com/chenhg5/go-admin/adapter"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins"
)

// Engine is the core components of goAdmin. It has two attributes.
// PluginList is an array of plugin. Adapter is the adapter of
// web framework context and goAdmin context. The relationship of adapter and
// plugin is that the adapter use the plugin which contains routers and
// controller methods to inject into the framework entity and make it work.
type Engine struct {
	PluginList []plugins.Plugin
	Adapter    adapter.WebFrameWork
}

// Default return the default engine instance.
func Default() *Engine {
	return new(Engine)
}

// Use enable the adapter.
func (eng *Engine) Use(router interface{}) error {
	return eng.Adapter.Use(router, eng.PluginList)
}

// AddPlugins add the plugins and initialize them.
func (eng *Engine) AddPlugins(plugs ...plugins.Plugin) *Engine {

	for _, plug := range plugs {
		plug.InitPlugin()
	}

	eng.PluginList = append(eng.PluginList, plugs...)
	return eng
}

// AddConfig set the global config.
func (eng *Engine) AddConfig(cfg config.Config) *Engine {
	config.Set(cfg)
	return eng
}

// AddAdapter add the adapter of engine.
func (eng *Engine) AddAdapter(ada adapter.WebFrameWork) *Engine {
	eng.Adapter = ada
	return eng
}
