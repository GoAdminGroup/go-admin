// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"html/template"
	"sync"
)

// Database is a type of database connection config.
// Because a little difference of different database driver.
// The Config has multiple options but may be not used.
// Such as the sqlite driver only use the FILE option which
// can be ignored when the driver is mysql.
type Database struct {
	HOST         string
	PORT         string
	USER         string
	PWD          string
	NAME         string
	MAX_IDLE_CON int
	MAX_OPEN_CON int
	DRIVER       string

	FILE string
}

// Store is the file store config. Path is the local store path.
// and prefix is the url prefix used to visit it.
type Store struct {
	PATH   string
	PREFIX string
}

// Config type is the global config of goAdmin. It will be
// initialized in the engine.
type Config struct {
	// An array supports multi database connection. The first
	// element of DATABASE is the default connection. See the
	// file connection.go.
	DATABASE []Database

	// The cookie domain used in the auth modules. see
	// the session.go.
	DOMAIN string

	// Used to set as the localize language which show in the
	// interface.
	LANGUAGE string

	// The global url prefix.
	PREFIX string

	// The theme name of template.
	THEME string

	// The path where files will be stored into.
	STORE Store

	// The title of web page.
	TITLE string

	// Logo is the top text in the sidebar.
	LOGO template.HTML

	// Mini-logo is the top text in the sidebar when folding.
	MINILOGO template.HTML

	// The url redirect to after login
	INDEX string
}

var (
	globalCfg Config
	mutex     sync.Mutex
)

// Set sets the config.
func Set(cfg Config) {
	mutex.Lock()
	globalCfg = cfg

	if globalCfg.TITLE == "" {
		globalCfg.TITLE = "GoAdmin"
	}
	if globalCfg.LOGO == "" {
		globalCfg.LOGO = "<b>Go</b>Admin"
	}
	if globalCfg.MINILOGO == "" {
		globalCfg.MINILOGO = "<b>G</b>A"
	}
	if globalCfg.THEME == "" {
		globalCfg.THEME = "adminlte"
	}
	if globalCfg.INDEX == "" {
		globalCfg.INDEX = "/info/manager"
	}
	if globalCfg.INDEX == "/" {
		globalCfg.INDEX = ""
	}

	mutex.Unlock()
}

// Get gets the config.
func Get() Config {
	return globalCfg
}
