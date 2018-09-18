// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	"github.com/chenhg5/go-admin/context"
)

// Plugin as one of the key components of goAdmin has three
// methods. GetRequest return all the path registered in the
// plugin. GetHandler according the url and method return the
// corresponding handler. InitPlugin init the plugin which do
// something like init the database and set the config and register
// the routes. The Plugin must implement the three methods.
type Plugin interface {
	GetRequest() []context.Path
	GetHandler(url, method string) context.Handler
	InitPlugin()
}

// GetHandler is a help method for Plugin GetHandler.
func GetHandler(url, method string, app *context.App) context.Handler {

	handler := app.Find(url, method)

	if handler == nil {
		panic("handler not found")
	}

	return handler
}
