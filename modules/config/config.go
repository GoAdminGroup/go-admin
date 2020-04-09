// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"strconv"
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

func (d DatabaseList) JSON() string {
	return utils.JSON(d)
}

func GetDatabaseListFromJSON(m string) DatabaseList {
	var d = make(DatabaseList, 0)
	if m == "" {
		panic("wrong config")
	}
	_ = json.Unmarshal([]byte(m), &d)
	return d
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
	Path   string `json:"path",yaml:"path",ini:"path"`
	Prefix string `json:"prefix",yaml:"prefix",ini:"prefix"`
}

func (s Store) URL(suffix string) string {
	if len(suffix) > 4 && suffix[:4] == "http" {
		return suffix
	}
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
		if len(s.Prefix) > 4 && s.Prefix[:4] == "http" {
			return s.Prefix + suffix
		}
		return "/" + s.Prefix + suffix
	}
	if len(s.Prefix) > 4 && s.Prefix[:4] == "http" {
		return s.Prefix + "/" + suffix
	}
	return "/" + s.Prefix + "/" + suffix
}

func (s Store) JSON() string {
	if s.Path == "" && s.Prefix == "" {
		return ""
	}
	return utils.JSON(s)
}

func GetStoreFromJSON(m string) Store {
	var s Store
	if m == "" {
		return s
	}
	_ = json.Unmarshal([]byte(m), &s)
	return s
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

	// Login page URL
	LoginUrl string `json:"login_url",yaml:"login_url",ini:"login_url"`

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

	// Footer Info html
	FooterInfo template.HTML `json:"footer_info",yaml:"footer_info",ini:"footer_info"`

	// Login page title
	LoginTitle string `json:"login_title",yaml:"login_title",ini:"login_title"`

	// Login page logo
	LoginLogo template.HTML `json:"login_logo",yaml:"login_logo",ini:"login_logo"`

	// Auth user table
	AuthUserTable string `json:"auth_user_table",yaml:"auth_user_table",ini:"auth_user_table"`

	// Extra config info
	Extra ExtraInfo `json:"extra",yaml:"extra",ini:"extra"`

	// Page animation
	Animation PageAnimation `json:"animation",yaml:"animation",ini:"animation"`

	// Limit login with different IPs
	NoLimitLoginIP bool `json:"no_limit_login_ip",yaml:"no_limit_login_ip",ini:"no_limit_login_ip"`

	// When site off is true, website will be closed
	SiteOff bool `json:"site_off",yaml:"site_off",ini:"site_off"`

	// Hide config center entrance flag
	HideConfigCenterEntrance bool `json:"hide_config_center_entrance",yaml:"hide_config_center_entrance",ini:"hide_config_center_entrance"`

	// Update Process Function
	UpdateProcessFn UpdateConfigProcessFn `json:"-",yaml:"-",ini:"-"`

	// Favicon string `json:"favicon",yaml:"favicon",ini:"favicon"`

	// Is open admin plugin json api
	OpenAdminApi bool `json:"open_admin_api",yaml:"open_admin_api",ini:"open_admin_api"`

	prefix string
}

type ExtraInfo map[string]interface{}

type UpdateConfigProcessFn func(values form.Values) (form.Values, error)

// see more: https://daneden.github.io/animate.css/
type PageAnimation struct {
	Type     string  `json:"type",yaml:"type",ini:"type"`
	Duration float32 `json:"duration",yaml:"duration",ini:"duration"`
	Delay    float32 `json:"delay",yaml:"delay",ini:"delay"`
}

func (p PageAnimation) JSON() string {
	if p.Type == "" {
		return ""
	}
	return utils.JSON(p)
}

func GetPageAnimationFromJSON(m string) PageAnimation {
	var p PageAnimation
	if m == "" {
		return p
	}
	_ = json.Unmarshal([]byte(m), &p)
	return p
}

// FileUploadEngine is a file upload engine.
type FileUploadEngine struct {
	Name   string                 `json:"name",yaml:"name",ini:"name"`
	Config map[string]interface{} `json:"config",yaml:"config",ini:"config"`
}

func (f FileUploadEngine) JSON() string {
	if f.Name == "" {
		return ""
	}
	if len(f.Config) == 0 {
		f.Config = nil
	}
	return utils.JSON(f)
}

func GetFileUploadEngineFromJSON(m string) FileUploadEngine {
	var f FileUploadEngine
	if m == "" {
		return f
	}
	_ = json.Unmarshal([]byte(m), &f)
	return f
}

// GetIndexURL get the index url with prefix.
func (c *Config) GetIndexURL() string {
	index := c.Index()
	if index == "/" {
		return c.Prefix()
	}

	return c.Prefix() + index
}

// Url get url with the given suffix.
func (c *Config) Url(suffix string) string {
	if c.prefix == "/" {
		return suffix
	}
	if suffix == "/" {
		return c.prefix
	}
	return c.prefix + suffix
}

// IsTestEnvironment check the environment if it is test.
func (c *Config) IsTestEnvironment() bool {
	return c.Env == EnvTest
}

// IsLocalEnvironment check the environment if it is local.
func (c *Config) IsLocalEnvironment() bool {
	return c.Env == EnvLocal
}

// IsProductionEnvironment check the environment if it is production.
func (c *Config) IsProductionEnvironment() bool {
	return c.Env == EnvProd
}

// URLRemovePrefix remove prefix from the given url.
func (c *Config) URLRemovePrefix(url string) string {
	if url == c.prefix {
		return "/"
	}
	if c.prefix == "/" {
		return url
	}
	return strings.Replace(url, c.prefix, "", 1)
}

// Index return the index url without prefix.
func (c *Config) Index() string {
	if c.IndexUrl == "" {
		return "/"
	}
	if c.IndexUrl[0] != '/' {
		return "/" + c.IndexUrl
	}
	return c.IndexUrl
}

// Prefix return the prefix.
func (c *Config) Prefix() string {
	return c.prefix
}

// AssertPrefix return the prefix of assert.
func (c *Config) AssertPrefix() string {
	if c.prefix == "/" {
		return ""
	}
	return c.prefix
}

func (c *Config) AddUpdateProcessFn(fn UpdateConfigProcessFn) *Config {
	c.UpdateProcessFn = fn
	return c
}

// PrefixFixSlash return the prefix fix the slash error.
func (c *Config) PrefixFixSlash() string {
	if c.UrlPrefix == "/" {
		return ""
	}
	if c.UrlPrefix != "" && c.UrlPrefix[0] != '/' {
		return "/" + c.UrlPrefix
	}
	return c.UrlPrefix
}

func (c *Config) Copy() *Config {
	return &Config{
		Databases:        c.Databases,
		Domain:           c.Domain,
		Language:         c.Language,
		UrlPrefix:        c.UrlPrefix,
		Theme:            c.Theme,
		Store:            c.Store,
		Title:            c.Title,
		Logo:             c.Logo,
		MiniLogo:         c.MiniLogo,
		IndexUrl:         c.IndexUrl,
		LoginUrl:         c.LoginUrl,
		Debug:            c.Debug,
		Env:              c.Env,
		InfoLogPath:      c.InfoLogPath,
		ErrorLogPath:     c.ErrorLogPath,
		AccessLogPath:    c.AccessLogPath,
		SqlLog:           c.SqlLog,
		AccessLogOff:     c.AccessLogOff,
		InfoLogOff:       c.InfoLogOff,
		ErrorLogOff:      c.ErrorLogOff,
		ColorScheme:      c.ColorScheme,
		SessionLifeTime:  c.SessionLifeTime,
		AssetUrl:         c.AssetUrl,
		FileUploadEngine: c.FileUploadEngine,
		CustomHeadHtml:   c.CustomHeadHtml,
		CustomFootHtml:   c.CustomFootHtml,
		FooterInfo:       c.FooterInfo,
		LoginTitle:       c.LoginTitle,
		LoginLogo:        c.LoginLogo,
		AuthUserTable:    c.AuthUserTable,
		Extra:            c.Extra,
		Animation:        c.Animation,
		NoLimitLoginIP:   c.NoLimitLoginIP,
		prefix:           c.prefix,
	}
}

func (c *Config) ToMap() map[string]string {
	var m = make(map[string]string, 0)
	m["language"] = c.Language
	m["databases"] = c.Databases.JSON()
	m["domain"] = c.Domain
	m["url_prefix"] = c.UrlPrefix
	m["theme"] = c.Theme
	m["store"] = c.Store.JSON()
	m["title"] = c.Title
	m["logo"] = string(c.Logo)
	m["mini_logo"] = string(c.MiniLogo)
	m["index_url"] = c.IndexUrl
	m["site_off"] = strconv.FormatBool(c.SiteOff)
	m["login_url"] = c.LoginUrl
	m["debug"] = strconv.FormatBool(c.Debug)
	m["env"] = c.Env
	m["info_log_path"] = c.InfoLogPath
	m["error_log_path"] = c.ErrorLogPath
	m["access_log_path"] = c.AccessLogPath
	m["sql_log"] = strconv.FormatBool(c.SqlLog)
	m["access_log_off"] = strconv.FormatBool(c.AccessLogOff)
	m["info_log_off"] = strconv.FormatBool(c.InfoLogOff)
	m["error_log_off"] = strconv.FormatBool(c.ErrorLogOff)
	m["color_scheme"] = c.ColorScheme
	m["session_life_time"] = strconv.Itoa(c.SessionLifeTime)
	m["asset_url"] = c.AssetUrl
	m["file_upload_engine"] = c.FileUploadEngine.JSON()
	m["custom_head_html"] = string(c.CustomHeadHtml)
	m["custom_foot_Html"] = string(c.CustomFootHtml)
	m["footer_info"] = string(c.FooterInfo)
	m["login_title"] = c.LoginTitle
	m["login_logo"] = string(c.LoginLogo)
	m["auth_user_table"] = c.AuthUserTable
	if len(c.Extra) == 0 {
		m["extra"] = ""
	} else {
		m["extra"] = utils.JSON(c.Extra)
	}
	m["animation"] = c.Animation.JSON()
	m["no_limit_login_ip"] = strconv.FormatBool(c.NoLimitLoginIP)
	return m
}

func (c *Config) Update(m map[string]string) error {
	updateLock.Lock()
	defer updateLock.Unlock()
	c.Language = m["language"]
	c.Domain = m["domain"]
	c.Theme = m["theme"]
	c.Title = m["title"]
	c.Logo = template.HTML(m["logo"])
	c.MiniLogo = template.HTML(m["mini_logo"])
	c.Debug = utils.ParseBool(m["debug"])
	c.Env = m["env"]
	c.SiteOff = utils.ParseBool(m["site_off"])

	c.AccessLogOff = utils.ParseBool(m["access_log_off"])
	c.InfoLogOff = utils.ParseBool(m["info_log_off"])
	c.ErrorLogOff = utils.ParseBool(m["error_log_off"])

	if c.InfoLogPath != m["info_log_path"] {
		c.InfoLogPath = m["info_log_path"]
		logger.SetInfoLogger(c.InfoLogPath, c.Debug, c.InfoLogOff)
	}
	if c.ErrorLogPath != m["error_log_path"] {
		c.ErrorLogPath = m["error_log_path"]
		logger.SetErrorLogger(c.ErrorLogPath, c.Debug, c.ErrorLogOff)
	}
	if c.AccessLogPath != m["access_log_path"] {
		c.AccessLogPath = m["access_log_path"]
		logger.SetAccessLogger(c.AccessLogPath, c.Debug, c.AccessLogOff)
	}

	c.SqlLog = utils.ParseBool(m["sql_log"])
	if c.Theme == "adminlte" {
		c.ColorScheme = m["color_scheme"]
	}
	ses, _ := strconv.Atoi(m["session_life_time"])
	if ses != 0 {
		c.SessionLifeTime = ses
	}
	c.CustomHeadHtml = template.HTML(m["custom_head_html"])
	c.CustomFootHtml = template.HTML(m["custom_foot_Html"])
	c.FooterInfo = template.HTML(m["footer_info"])
	c.LoginTitle = m["login_title"]
	c.AssetUrl = m["asset_url"]
	c.LoginLogo = template.HTML(m["login_logo"])
	c.NoLimitLoginIP = utils.ParseBool(m["no_limit_login_ip"])

	c.FileUploadEngine = GetFileUploadEngineFromJSON(m["file_upload_engine"])
	c.Animation = GetPageAnimationFromJSON(m["animation"])
	if m["extra"] != "" {
		var extra = make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(m["extra"]), &extra)
		c.Extra = extra
	}

	return nil
}

// eraseSens erase sensitive info.
func (c *Config) EraseSens() *Config {
	for key := range c.Databases {
		c.Databases[key] = Database{
			Driver: c.Databases[key].Driver,
		}
	}
	return c
}

var (
	globalCfg  = new(Config)
	declare    sync.Once
	updateLock sync.Mutex
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
func Set(cfg Config) *Config {

	lock.Lock()
	defer lock.Unlock()

	if atomic.LoadUint32(&count) != 0 {
		panic("can not set config twice")
	}
	atomic.StoreUint32(&count, 1)

	cfg.Title = setDefault(cfg.Title, "", "GoAdmin")
	cfg.LoginTitle = setDefault(cfg.LoginTitle, "", "GoAdmin")
	cfg.Logo = template.HTML(setDefault(string(cfg.Logo), "", "<b>Go</b>Admin"))
	cfg.MiniLogo = template.HTML(setDefault(string(cfg.MiniLogo), "", "<b>G</b>A"))
	cfg.Theme = setDefault(cfg.Theme, "", "adminlte")
	cfg.IndexUrl = setDefault(cfg.IndexUrl, "", "/info/manager")
	cfg.LoginUrl = setDefault(cfg.LoginUrl, "", "/login")
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

	globalCfg = &cfg

	return globalCfg
}

// AssertPrefix return the prefix of assert.
func AssertPrefix() string {
	return globalCfg.AssertPrefix()
}

// GetIndexURL get the index url with prefix.
func GetIndexURL() string {
	return globalCfg.GetIndexURL()
}

// IsProductionEnvironment check the environment if it is production.
func IsProductionEnvironment() bool {
	return globalCfg.IsProductionEnvironment()
}

// URLRemovePrefix remove prefix from the given url.
func URLRemovePrefix(url string) string {
	return globalCfg.URLRemovePrefix(url)
}

func Url(suffix string) string {
	return globalCfg.Url(suffix)
}

// Prefix return the prefix.
func Prefix() string {
	return globalCfg.prefix
}

// PrefixFixSlash return the prefix fix the slash error.
func PrefixFixSlash() string {
	return globalCfg.PrefixFixSlash()
}

// Get gets the config.
func Get() Config {
	c := globalCfg.Copy().EraseSens()
	return *c
}

func setDefault(value, condition, def string) string {
	if value == condition {
		return def
	}
	return value
}

// Getter methods
// ============================

func GetDatabases() DatabaseList {
	var list = make(DatabaseList, len(globalCfg.Databases))
	for key := range globalCfg.Databases {
		list[key] = Database{
			Driver: globalCfg.Databases[key].Driver,
		}
	}
	return list
}

func GetDomain() string {
	return globalCfg.Domain
}

func GetLanguage() string {
	return globalCfg.Language
}

func GetUrlPrefix() string {
	return globalCfg.UrlPrefix
}

func GetOpenAdminApi() bool {
	return globalCfg.OpenAdminApi
}

func GetTheme() string {
	return globalCfg.Theme
}

func GetStore() Store {
	return globalCfg.Store
}

func GetTitle() string {
	return globalCfg.Title
}

func GetLogo() template.HTML {
	return globalCfg.Logo
}

func GetSiteOff() bool {
	return globalCfg.SiteOff
}

func GetMiniLogo() template.HTML {
	return globalCfg.MiniLogo
}

func GetIndexUrl() string {
	return globalCfg.IndexUrl
}

func GetLoginUrl() string {
	return globalCfg.LoginUrl
}

func GetDebug() bool {
	return globalCfg.Debug
}

func GetEnv() string {
	return globalCfg.Env
}

func GetInfoLogPath() string {
	return globalCfg.InfoLogPath
}

func GetErrorLogPath() string {
	return globalCfg.ErrorLogPath
}

func GetAccessLogPath() string {
	return globalCfg.AccessLogPath
}

func GetSqlLog() bool {
	return globalCfg.SqlLog
}

func GetAccessLogOff() bool {
	return globalCfg.AccessLogOff
}
func GetInfoLogOff() bool {
	return globalCfg.InfoLogOff
}
func GetErrorLogOff() bool {
	return globalCfg.ErrorLogOff
}

func GetColorScheme() string {
	return globalCfg.ColorScheme
}

func GetSessionLifeTime() int {
	return globalCfg.SessionLifeTime
}

func GetAssetUrl() string {
	return globalCfg.AssetUrl
}

func GetFileUploadEngine() FileUploadEngine {
	return globalCfg.FileUploadEngine
}

func GetCustomHeadHtml() template.HTML {
	return globalCfg.CustomHeadHtml
}

func GetCustomFootHtml() template.HTML {
	return globalCfg.CustomFootHtml
}

func GetFooterInfo() template.HTML {
	return globalCfg.FooterInfo
}

func GetLoginTitle() string {
	return globalCfg.LoginTitle
}

func GetLoginLogo() template.HTML {
	return globalCfg.LoginLogo
}

func GetAuthUserTable() string {
	return globalCfg.AuthUserTable
}

func GetExtra() map[string]interface{} {
	return globalCfg.Extra
}

func GetAnimation() PageAnimation {
	return globalCfg.Animation
}

func GetNoLimitLoginIP() bool {
	return globalCfg.NoLimitLoginIP
}

type Service struct {
	C *Config
}

func (s *Service) Name() string {
	return "config"
}

func SrvWithConfig(c *Config) *Service {
	return &Service{c}
}

func GetService(s interface{}) *Config {
	if srv, ok := s.(*Service); ok {
		return srv.C
	}
	panic("wrong service")
}
