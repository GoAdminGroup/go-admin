// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"html/template"
	"io/ioutil"
	"strings"
	"sync"
)

// Database is a type of database connection config.
// Because a little difference of different database driver.
// The Config has multiple options but may be not used.
// Such as the sqlite driver only use the File option which
// can be ignored when the driver is mysql.
type Database struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	User       string `json:"user"`
	Pwd        string `json:"pwd"`
	Name       string `json:"name"`
	MaxIdleCon int    `json:"max_idle_con"`
	MaxOpenCon int    `json:"max_open_con"`
	Driver     string `json:"driver"`
	File       string `json:"file"`
}

type DatabaseList map[string]Database

func (d DatabaseList) GetDefault() Database {
	return d["default"]
}

func (d DatabaseList) Add(key string, db Database) {
	d[key] = db
}

func (d DatabaseList) GroupByDriver() map[string]DatabaseList {
	drivers := make(map[string]DatabaseList)
	for key, item := range d {
		if driverList, ok := drivers[item.Driver]; ok {
			driverList.Add(key, item)
		} else {
			drivers[item.Driver] = make(DatabaseList)
			drivers[item.Driver].Add(key, item)
		}
	}
	return drivers
}

const (
	EnvTest  = "test"
	EnvLocal = "local"
	EnvProd  = "prod"

	DriverMysql      = "mysql"
	DriverSqlite     = "sqlite"
	DriverPostgresql = "postgresql"
	DriverMssql      = "mssql"
)

// Store is the file store config. Path is the local store path.
// and prefix is the url prefix used to visit it.
type Store struct {
	Path   string
	Prefix string
}

// Config type is the global config of goAdmin. It will be
// initialized in the engine.
type Config struct {
	// An map supports multi database connection. The first
	// element of Databases is the default connection. See the
	// file connection.go.
	Databases DatabaseList `json:"database"`

	// The cookie domain used in the auth modules. see
	// the session.go.
	Domain string `json:"domain"`

	// Used to set as the localize language which show in the
	// interface.
	Language string `json:"language"`

	// The global url prefix.
	UrlPrefix string `json:"prefix"`

	// The theme name of template.
	Theme string `json:"theme"`

	// The path where files will be stored into.
	Store Store `json:"store"`

	// The title of web page.
	Title string `json:"title"`

	// Logo is the top text in the sidebar.
	Logo template.HTML `json:"logo"`

	// Mini-logo is the top text in the sidebar when folding.
	MiniLogo template.HTML `json:"mini_logo"`

	// The url redirect to after login.
	IndexUrl string `json:"index"`

	// Debug mode
	Debug bool `json:"debug"`

	// Env is the environment, which maybe local, test, prod.
	Env string `json:"env"`

	// Info log path.
	InfoLogPath string `json:"info_log"`

	// Error log path.
	ErrorLogPath string `json:"error_log"`

	// Access log path.
	AccessLogPath string `json:"access_log"`

	// Sql operator record log switch.
	SqlLog bool `json:"sql_log"`

	AccessLogOff bool `json:"access_log_off"`
	InfoLogOff   bool `json:"info_log_off"`
	ErrorLogOff  bool `json:"error_log_off"`

	// Color scheme.
	ColorScheme string `json:"color_scheme"`

	// Session valid time duration, units are seconds.
	SessionLifeTime int `json:"session_life_time"`

	// Assets visit link.
	AssetUrl string `json:"asset_url"`

	// File upload engine, default "local"
	FileUploadEngine FileUploadEngine `json:"file_upload_engine"`

	// Custom html in the tag head.
	CustomHeadHtml template.HTML `json:"custom_head_html"`

	// Custom html after body.
	CustomFootHtml template.HTML `json:"custom_foot_html"`

	// Login page title
	LoginTitle string `json:"login_title"`

	// Login page logo
	LoginLogo template.HTML `json:"login_logo"`

	prefix string
}

type FileUploadEngine struct {
	Name   string
	Config map[string]interface{}
}

func (c Config) GetIndexUrl() string {
	index := c.Index()
	if index == "/" {
		return c.Prefix()
	}

	return c.Prefix() + index
}

func (c Config) Url(suffix string) string {
	if c.prefix == "/" {
		return suffix
	}
	if suffix == "/" {
		return c.prefix
	}
	return c.prefix + suffix
}

func (c Config) IsTestEnvironment() bool {
	return c.Env == EnvTest
}

func (c Config) IsLocalEnvironment() bool {
	return c.Env == EnvLocal
}

func (c Config) IsProductionEnvironment() bool {
	return c.Env == EnvProd
}

func (c Config) UrlRemovePrefix(u string) string {
	if u == c.prefix {
		return "/"
	}
	return strings.Replace(u, c.Prefix(), "", 1)
}

func (c Config) Index() string {
	if c.IndexUrl == "" {
		return "/"
	}
	if c.IndexUrl[0] != '/' {
		return "/" + c.IndexUrl
	}
	return c.IndexUrl
}

func (c Config) Prefix() string {
	return c.prefix
}

func (c Config) PrefixFixSlash() string {
	if c.UrlPrefix == "/" {
		return ""
	}
	if c.UrlPrefix[0] != '/' {
		return "/" + c.UrlPrefix
	}
	return c.UrlPrefix
}

var (
	globalCfg Config
	mutex     sync.Mutex
	declare   sync.Once
)

func ReadFromJson(path string) Config {
	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonByte, &globalCfg)

	if err != nil {
		panic(err)
	}

	Set(globalCfg)

	return globalCfg
}

// Set sets the config.
func Set(cfg Config) {
	mutex.Lock()
	globalCfg = cfg

	globalCfg.Title = setDefault(globalCfg.Title, "", constant.Title)
	globalCfg.LoginTitle = setDefault(globalCfg.LoginTitle, "", constant.Title)
	globalCfg.Logo = template.HTML(setDefault(string(globalCfg.Logo), "", "<b>Go</b>Admin"))
	globalCfg.MiniLogo = template.HTML(setDefault(string(globalCfg.MiniLogo), "", "<b>G</b>A"))
	globalCfg.Theme = setDefault(globalCfg.Theme, "", "adminlte")
	globalCfg.IndexUrl = setDefault(globalCfg.IndexUrl, "", "/info/manager")
	globalCfg.ColorScheme = setDefault(globalCfg.ColorScheme, "", "skin-black")
	globalCfg.FileUploadEngine.Name = setDefault(globalCfg.FileUploadEngine.Name, "", "local")
	globalCfg.Env = setDefault(globalCfg.Env, "", EnvProd)
	if globalCfg.SessionLifeTime == 0 {
		// default two hours
		globalCfg.SessionLifeTime = 7200
	}

	if globalCfg.UrlPrefix == "" {
		globalCfg.prefix = "/"
	} else if globalCfg.UrlPrefix[0] != '/' {
		globalCfg.prefix = "/" + globalCfg.UrlPrefix
	} else {
		globalCfg.prefix = globalCfg.UrlPrefix
	}

	logger.SetInfoLogger(cfg.InfoLogPath, cfg.Debug, cfg.InfoLogOff)
	logger.SetErrorLogger(cfg.ErrorLogPath, cfg.Debug, cfg.ErrorLogOff)
	logger.SetAccessLogger(cfg.AccessLogPath, cfg.Debug, cfg.AccessLogOff)

	if cfg.SqlLog {
		logger.OpenSqlLog()
	}

	if cfg.Debug {
		declare.Do(func() {
			fmt.Println(`GoAdmin is now running.
Running in "debug" mode. Switch to "release" mode in production.`)
			fmt.Println()
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
