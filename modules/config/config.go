// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import "sync"

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
	// DATABASE which is an array supports multi database
	// connection. The first element of DATABASE is the default
	// connection. See the file connection.go.
	DATABASE []Database

	// DOMAIN is the cookie domain used in the auth modules. see
	// the session.go.
	DOMAIN string

	// LANGUAGE is used to set as the localize language which
	// show in the interface.
	LANGUAGE string

	// PREFIX is the global url prefix.
	PREFIX string

	// THEME the theme name of template.
	THEME string

	// STORE the path where files will be stored into.
	STORE Store
}

var (
	globalCfg Config
	mutex     sync.Mutex
)

// Set sets the config.
func Set(cfg Config) {
	mutex.Lock()
	globalCfg = cfg
	mutex.Unlock()
}

// Get gets the config.
func Get() Config {
	return globalCfg
}
