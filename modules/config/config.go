// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"strings"
	"sync"
	"sync/atomic"
)

// Database is a type of database connection config.
//
// Because a little difference of different database driver.
// The Config has multiple options but may not be used.
// Such as the sqlite driver only use the File option which
// can be ignored when the driver is mysql.
//
// If the Dsn is configured, when driver is mysql/postgresql/
// mssql, the other configurations will be ignored, except for
// MaxIdleCon and MaxOpenCon.
type Database struct {
	Host       string `json:"host",yaml:"host",ini:"host"`
	Port       string `json:"port",yaml:"port",ini:"port"`
	User       string `json:"user",yaml:"user",ini:"user"`
	Pwd        string `json:"pwd",yaml:"pwd",ini:"pwd"`
	Name       string `json:"name",yaml:"name",ini:"name"`
	MaxIdleCon int    `json:"max_idle_con",yaml:"max_idle_con",ini:"max_idle_con"`
	MaxOpenCon int    `json:"max_open_con",yaml:"max_open_con",ini:"max_open_con"`
	Driver     string `json:"driver",yaml:"driver",ini:"driver"`
	File       string `json:"file",yaml:"file",ini:"file"`
	Dsn        string `json:"dsn",yaml:"dsn",ini:"dsn"`
}

// DatabaseList is a map of Database.
type DatabaseList map[string]Database

// GetDefault get the default Database.
func (d DatabaseList) GetDefault() Database {
	return d["default"]
}

// Add add a Database to the DatabaseList.
func (d DatabaseList) Add(key string, db Database) {
	d[key] = db
}

// GroupByDriver group the Databases with the drivers.
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
	// EnvTest is a const value of test environment.
	EnvTest = "test"
	// EnvLocal is a const value of local environment.
	EnvLocal = "local"
	// EnvProd is a const value of production environment.
	EnvProd = "prod"

	// DriverMysql is a const value of mysql driver.
	DriverMysql = "mysql"
	// DriverSqlite is a const value of sqlite driver.
	DriverSqlite = "sqlite"
	// DriverPostgresql is a const value of postgresql driver.
	DriverPostgresql = "postgresql"
	// DriverMssql is a const value of mssql driver.
	DriverMssql = "mssql"
)

// Store is the file store config. Path is the local store path.
// and prefix is the url prefix used to visit it.
type Store struct {
	Path   string
	Prefix string
}

func (s Store) URL(suffix string) string {
	if s.Prefix == "" {
		if suffix[0] == '/' {
			return suffix
		}
		return "/" + suffix
	}
	if s.Prefix[0] == '/' {
		if suffix[0] == '/' {
			return s.Prefix + suffix
		}
		return s.Prefix + "/" + suffix
	}
	if suffix[0] == '/' {
		return "/" + s.Prefix + suffix
	}
	return "/" + s.Prefix + "/" + suffix
}

// Config type is the global config of goAdmin. It will be
// initialized in the engine.
type Config struct {
	// An map supports multi database connection. The first
	// element of Databases is the default connection. See the
	// file connection.go.
	Databases DatabaseList `json:"database",yaml:"database",ini:"database"`

	// The cookie domain used in the auth modules. see
	// the session.go.
	Domain string `json:"domain",yaml:"domain",ini:"domain"`

	// Used to set as the localize language which show in the
	// interface.
	Language string `json:"language",yaml:"language",ini:"language"`

	// The global url prefix.
	UrlPrefix string `json:"prefix",yaml:"prefix",ini:"prefix"`

	// The theme name of template.
	Theme string `json:"theme",yaml:"theme",ini:"theme"`

	// The path where files will be stored into.
	Store Store `json:"store",yaml:"store",ini:"store"`

	// The title of web page.
	Title string `json:"title",yaml:"title",ini:"title"`

	// Logo is the top text in the sidebar.
	Logo template.HTML `json:"logo",yaml:"logo",ini:"logo"`

	// Mini-logo is the top text in the sidebar when folding.
	MiniLogo template.HTML `json:"mini_logo",yaml:"mini_logo",ini:"mini_logo"`

	// The url redirect to after login.
	IndexUrl string `json:"index",yaml:"index",ini:"index"`

	// Debug mode
	Debug bool `json:"debug",yaml:"debug",ini:"debug"`

	// Env is the environment,which maybe local,test,prod.
	Env string `json:"env",yaml:"env",ini:"env"`

	// Info log path.
	InfoLogPath string `json:"info_log",yaml:"info_log",ini:"info_log"`

	// Error log path.
	ErrorLogPath string `json:"error_log",yaml:"error_log",ini:"error_log"`

	// Access log path.
	AccessLogPath string `json:"access_log",yaml:"access_log",ini:"access_log"`

	// Sql operator record log switch.
	SqlLog bool `json:"sql_log",yaml:"sql_log",ini:"sql_log"`

	AccessLogOff bool `json:"access_log_off",yaml:"access_log_off",ini:"access_log_off"`
	InfoLogOff   bool `json:"info_log_off",yaml:"info_log_off",ini:"info_log_off"`
	ErrorLogOff  bool `json:"error_log_off",yaml:"error_log_off",ini:"error_log_off"`

	// Color scheme.
	ColorScheme string `json:"color_scheme",yaml:"color_scheme",ini:"color_scheme"`

	// Session valid time duration,units are seconds. Default 7200.
	SessionLifeTime int `json:"session_life_time",yaml:"session_life_time",ini:"session_life_time"`

	// Assets visit link.
	AssetUrl string `json:"asset_url",yaml:"asset_url",ini:"asset_url"`

	// File upload engine,default "local"
	FileUploadEngine FileUploadEngine `json:"file_upload_engine",yaml:"file_upload_engine",ini:"file_upload_engine"`

	// Custom html in the tag head.
	CustomHeadHtml template.HTML `json:"custom_head_html",yaml:"custom_head_html",ini:"custom_head_html"`

	// Custom html after body.
	CustomFootHtml template.HTML `json:"custom_foot_html",yaml:"custom_foot_html",ini:"custom_foot_html"`

	// Login page title
	LoginTitle string `json:"login_title",yaml:"login_title",ini:"login_title"`

	// Login page logo
	LoginLogo template.HTML `json:"login_logo",yaml:"login_logo",ini:"login_logo"`

	// Auth user table
	AuthUserTable string `json:"auth_user_table",yaml:"auth_user_table",ini:"auth_user_table"`

	// Extra config info
	Extra map[string]interface{} `json:"extra",yaml:"extra",ini:"extra"`

	// Page animation
	Animation PageAnimation `json:"animation",yaml:"animation",ini:"animation"`

	prefix string
}

// UserConfig type is the user config of goAdmin.
type UserConfig struct {
	// user id
	UserId int64 `json:"userid",yaml:"userid",ini:"userid"`

	// Used to set as the user language which show in the
	// interface.
	Language string `json:"language",yaml:"language",ini:"language"`

	// Extend
	// ...
}

var userConfig []UserConfig

// Set SetUserConfig the config.
func SetUserConfig(uConf UserConfig) {
	// insert or update to database, If there is a database
	for i := 0;i < len(userConfig);i++ {
		if userConfig[i].UserId == uConf.UserId {
			userConfig[i] = uConf
			return
		}
	}
	userConfig = append(userConfig, uConf)
}

// Get GetUserConf the config.
func GetUserConf(uId int64) *UserConfig {
	for i := 0;i < len(userConfig);i++ {
		if userConfig[i].UserId == uId {
			return &userConfig[i]
		}
	}
	SetUserConfig(UserConfig{
		UserId:   uId,
		Language: globalCfg.Language,
	})
	return &userConfig[len(userConfig)-1]
}

// see more: https://daneden.github.io/animate.css/
type PageAnimation struct {
	Type     string  `json:"type",yaml:"type",ini:"type"`
	Duration float32 `json:"duration",yaml:"duration",ini:"duration"`
	Delay    float32 `json:"delay",yaml:"delay",ini:"delay"`
}

// FileUploadEngine is a file upload engine.
type FileUploadEngine struct {
	Name   string
	Config map[string]interface{}
}

// GetIndexURL get the index url with prefix.
func (c Config) GetIndexURL() string {
	index := c.Index()
	if index == "/" {
		return c.Prefix()
	}

	return c.Prefix() + index
}

// Url get url with the given suffix.
func (c Config) Url(suffix string) string {
	if c.prefix == "/" {
		return suffix
	}
	if suffix == "/" {
		return c.prefix
	}
	return c.prefix + suffix
}

// IsTestEnvironment check the environment if it is test.
func (c Config) IsTestEnvironment() bool {
	return c.Env == EnvTest
}

// IsLocalEnvironment check the environment if it is local.
func (c Config) IsLocalEnvironment() bool {
	return c.Env == EnvLocal
}

// IsProductionEnvironment check the environment if it is production.
func (c Config) IsProductionEnvironment() bool {
	return c.Env == EnvProd
}

// URLRemovePrefix remove prefix from the given url.
func (c Config) URLRemovePrefix(url string) string {
	if url == c.prefix {
		return "/"
	}
	if c.prefix == "/" {
		return url
	}
	return strings.Replace(url, c.prefix, "", 1)
}

// Index return the index url without prefix.
func (c Config) Index() string {
	if c.IndexUrl == "" {
		return "/"
	}
	if c.IndexUrl[0] != '/' {
		return "/" + c.IndexUrl
	}
	return c.IndexUrl
}

// Prefix return the prefix.
func (c Config) Prefix() string {
	return c.prefix
}

// AssertPrefix return the prefix of assert.
func (c Config) AssertPrefix() string {
	if c.prefix == "/" {
		return ""
	}
	return c.prefix
}

// PrefixFixSlash return the prefix fix the slash error.
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
	declare   sync.Once
)

// ReadFromJson read the Config from a JSON file.
func ReadFromJson(path string) Config {
	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var cfg Config

	err = json.Unmarshal(jsonByte, &cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}

// ReadFromYaml read the Config from a YAML file.
func ReadFromYaml(path string) Config {
	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var cfg Config

	err = yaml.Unmarshal(jsonByte, &cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}

// ReadFromINI read the Config from a INI file.
func ReadFromINI(path string) Config {
	iniCfg, err := ini.Load(path)

	if err != nil {
		panic(err)
	}

	var cfg Config

	err = iniCfg.MapTo(&cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}

var (
	count uint32
	lock  sync.Mutex
)

// Set sets the config.
func Set(cfg Config) Config {

	lock.Lock()
	defer lock.Unlock()

	if atomic.LoadUint32(&count) != 0 {
		panic("can not set config twice")
	}
	atomic.StoreUint32(&count, 1)

	cfg.Title = setDefault(cfg.Title, "", constant.Title)
	cfg.LoginTitle = setDefault(cfg.LoginTitle, "", constant.Title)
	cfg.Logo = template.HTML(setDefault(string(cfg.Logo), "", "<b>Go</b>Admin"))
	cfg.MiniLogo = template.HTML(setDefault(string(cfg.MiniLogo), "", "<b>G</b>A"))
	cfg.Theme = setDefault(cfg.Theme, "", "adminlte")
	cfg.IndexUrl = setDefault(cfg.IndexUrl, "", "/info/manager")
	cfg.AuthUserTable = setDefault(cfg.AuthUserTable, "", "goadmin_users")
	cfg.ColorScheme = setDefault(cfg.ColorScheme, "", "skin-black")
	cfg.FileUploadEngine.Name = setDefault(cfg.FileUploadEngine.Name, "", "local")
	cfg.Env = setDefault(cfg.Env, "", EnvProd)
	if cfg.SessionLifeTime == 0 {
		// default two hours
		cfg.SessionLifeTime = 7200
	}

	if cfg.UrlPrefix == "" {
		cfg.prefix = "/"
	} else if cfg.UrlPrefix[0] != '/' {
		cfg.prefix = "/" + cfg.UrlPrefix
	} else {
		cfg.prefix = cfg.UrlPrefix
	}

	logger.SetInfoLogger(cfg.InfoLogPath, cfg.Debug, cfg.InfoLogOff)
	logger.SetErrorLogger(cfg.ErrorLogPath, cfg.Debug, cfg.ErrorLogOff)
	logger.SetAccessLogger(cfg.AccessLogPath, cfg.Debug, cfg.AccessLogOff)

	if cfg.SqlLog {
		logger.OpenSQLLog()
	}

	if cfg.Debug {
		declare.Do(func() {
			fmt.Println(`GoAdmin is now running.
Running in "debug" mode. Switch to "release" mode in production.`)
			fmt.Println()
		})
	}

	globalCfg = cfg
	eraseSens()

	return cfg
}

// Get gets the config.
func Get() Config {
	return globalCfg
}

// eraseSens erase sensitive info.
func eraseSens() {
	for _, d := range globalCfg.Databases {
		d.Host = ""
		d.Port = ""
		d.User = ""
		d.Pwd = ""
		d.Name = ""
		d.MaxIdleCon = 0
		d.MaxOpenCon = 0
		d.File = ""
		d.Dsn = ""
	}
}

func setDefault(value, condition, def string) string {
	if value == condition {
		return def
	}
	return value
}
