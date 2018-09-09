// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	"github.com/chenhg5/go-admin/context"
	"strings"
	"regexp"
)

// Plugin as one of the key components of goAdmin has three
// methods. GetRequest return all the path registered in the
// plugin. GetHandler according the url and method return the
// corresponding handler. InitPlugin init the plugin which do
// something like init the database and set the config and register
// the routes. Anyone wants to write the Plugin must implement
// the three methods.
type Plugin interface {
	GetRequest() []context.Path
	GetHandler(url, method string) context.Handler
	InitPlugin()
}

// GetHandler is a help method for Plugin GetHandler.
func GetHandler(url, method string, handleList *map[context.Path]context.Handler) context.Handler {
	for path, handler := range *handleList {
		if path.Method == method {
			if path.RegUrl == "" {
				if path.URL == url || path.URL + "/" == url || path.URL == url + "/" {
					return handler
				}
			} else {
				if strings.Count(path.RegUrl, "/") == strings.Count(url, "/") {
					if ok, err := regexp.MatchString(path.RegUrl, url); ok && err == nil {
						return handler
					}
				}
			}
		}
	}

	panic("handler not found")
}