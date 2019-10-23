// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package plugins

import (
	"errors"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"plugin"
)

// Plugin as one of the key components of goAdmin has three
// methods. GetRequest return all the path registered in the
// plugin. GetHandler according the url and method return the
// corresponding handler. InitPlugin init the plugin which do
// something like init the database and set the config and register
// the routes. The Plugin must implement the three methods.
type Plugin interface {
	GetRequest() []context.Path
	GetHandler(url, method string) context.Handlers
	InitPlugin()
}

// GetHandler is a help method for Plugin GetHandler.
func GetHandler(url, method string, app *context.App) context.Handlers {

	handler := app.Find(url, method)

	if len(handler) == 0 {
		panic("handler not found")
	}

	return handler
}

func LoadFromPlugin(mod string) Plugin {

	plug, err := plugin.Open(mod)
	if err != nil {
		logger.Error("LoadFromPlugin err", err)
		panic(err)
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		logger.Error("LoadFromPlugin err", err)
		panic(err)
	}

	var p Plugin
	p, ok := symPlugin.(Plugin)
	if !ok {
		logger.Error("LoadFromPlugin err: unexpected type from module symbol")
		panic(errors.New("LoadFromPlugin err: unexpected type from module symbol"))
	}

	return p
}
