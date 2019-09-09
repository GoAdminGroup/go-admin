// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"github.com/chenhg5/go-admin/modules/logger"
	"github.com/chenhg5/go-admin/plugins/admin/modules/constant"
	"html/template"
	"io/ioutil"
	"strings"
	"sync"
)

// Database is a type of database connection config.
// Because a little difference of different database driver.
// The Config has multiple options but may be not used.
// Such as the sqlite driver only use the FILE option which
// can be ignored when the driver is mysql.
type Database struct {
	HOST         string `json:"host"`
	PORT         string `json:"port"`
	USER         string `json:"user"`
	PWD          string `json:"pwd"`
	NAME         string `json:"name"`
	MAX_IDLE_CON int    `json:"max_idle_con"`
	MAX_OPEN_CON int    `json:"max_open_con"`
	DRIVER       string `json:"driver"`

	FILE string `json:"file"`
}

type DatabaseList map[string]Database

func (d DatabaseList) GetDefault() Database {
	return d["default"]
}

func (d DatabaseList) Add(key string, db Database) {
	d[key] = db
}

func (d DatabaseList) GroupByDriver() map[string]DatabaseList {
	drivers := make(map[string]DatabaseList, 0)
	for key, item := range d {
		if driverList, ok := drivers[item.DRIVER]; ok {
			driverList.Add(key, item)
		} else {
			drivers[item.DRIVER] = make(DatabaseList, 0)
			drivers[item.DRIVER].Add(key, item)
		}
	}
	return drivers
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
	// An map supports multi database connection. The first
	// element of DATABASE is the default connection. See the
	// file connection.go.
	DATABASE DatabaseList `json:"database"`

	// The cookie domain used in the auth modules. see
	// the session.go.
	DOMAIN string `json:"domain"`

	// Used to set as the localize language which show in the
	// interface.
	LANGUAGE string `json:"language"`

	// The global url prefix.
	PREFIX string `json:"prefix"`

	// The theme name of template.
	THEME string `json:"theme"`

	// The path where files will be stored into.
	STORE Store `json:"store"`

	// The title of web page.
	TITLE string `json:"title"`

	// Logo is the top text in the sidebar.
	LOGO template.HTML `json:"logo"`

	// Mini-logo is the top text in the sidebar when folding.
	MINILOGO template.HTML `json:"minilogo"`

	// The url redirect to after login
	INDEX string `json:"index"`

	// DEBUG mode
	DEBUG bool `json:"debug"`

	// Info log path
	INFOLOG string `json:"infolog"`

	// Error log path
	ERRORLOG string `json:"errorlog"`

	// Access log path
	ACCESSLOG string `json:"accesslog"`

	// Sql operator record log switch
	SQLLOG bool `json:"sqllog"`

	// Color scheme
	COLORSCHEME string `json:"colorscheme"`
}

func (c Config) GetIndexUrl() string {
	index := c.Index()
	if index == "/" {
		return c.Prefix()
	}
	return c.Prefix() + index
}

func (c Config) Url(suffix string) string {
	prefix := c.Prefix()
	if prefix == "/" {
		return suffix
	}
	return prefix + suffix
}

func (c Config) UrlRemovePrefix(u string) string {
	return strings.Replace(u, c.Prefix(), "", 1)
}

func (c Config) Index() string {
	if c.INDEX == "" {
		return "/"
	}
	if c.INDEX[0] != '/' {
		return "/" + c.INDEX
	}
	return c.INDEX
}

func (c Config) Prefix() string {
	if c.PREFIX == "" {
		return "/"
	}
	if c.PREFIX[0] != '/' {
		return "/" + c.PREFIX
	}
	return c.PREFIX
}

func (c Config) PrefixFixSlash() string {
	if c.PREFIX == "/" {
		return ""
	}
	if c.PREFIX[0] != '/' {
		return "/" + c.PREFIX
	}
	return c.PREFIX
}

var (
	globalCfg Config
	mutex     sync.Mutex
	declare   sync.Once
)

func ReadFromJson(path string) {
	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonByte, &globalCfg)

	if err != nil {
		panic(err)
	}

	Set(globalCfg)
}

// Set sets the config.
func Set(cfg Config) {
	mutex.Lock()
	globalCfg = cfg

	globalCfg.TITLE = setDefault(globalCfg.TITLE, "", constant.Title)
	globalCfg.LOGO = template.HTML(setDefault(string(globalCfg.LOGO), "", "<b>Go</b>Admin"))
	globalCfg.MINILOGO = template.HTML(setDefault(string(globalCfg.MINILOGO), "", "<b>G</b>A"))
	globalCfg.THEME = setDefault(globalCfg.THEME, "", "adminlte")
	globalCfg.INDEX = setDefault(globalCfg.INDEX, "", "/info/manager")
	globalCfg.INDEX = setDefault(globalCfg.INDEX, "/", "")
	globalCfg.COLORSCHEME = setDefault(globalCfg.COLORSCHEME, "", "skin-black")

	if cfg.INFOLOG != "" {
		logger.SetInfoLogger(cfg.INFOLOG, cfg.DEBUG)
	}

	if cfg.ERRORLOG != "" {
		logger.SetErrorLogger(cfg.ERRORLOG, cfg.DEBUG)
	}

	if cfg.ACCESSLOG != "" {
		logger.SetAccessLogger(cfg.ACCESSLOG, cfg.DEBUG)
	}

	if cfg.SQLLOG {
		logger.OpenSqlLog()
	}

	if cfg.DEBUG {
		declare.Do(func() {
			logger.Info(`go-admin is now running.
Running in "debug" mode. Switch to "release" mode in production.

`)
		})
	}

	mutex.Unlock()
}

// Get gets the config.
func Get() Config {
	return globalCfg
}

func setDefault(value, condition, def string) string {
	if value == condition {
		return def
	}
	return value
}
