package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/GoAdminGroup/go-admin/modules/utils"

	"github.com/stretchr/testify/assert"
)

func TestConfig_GetIndexUrl(t *testing.T) {
	Initialize(&Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().GetIndexURL(), "/admin")

	testSetCfg(&Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().GetIndexURL(), "/admin")

	testSetCfg(&Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().GetIndexURL(), "/admin")
}

func TestConfig_Index(t *testing.T) {
	testSetCfg(&Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Index(), "/")
}

func TestConfig_Prefix(t *testing.T) {
	testSetCfg(&Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Prefix(), "/admin")

	testSetCfg(&Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Prefix(), "/admin")
}

func TestConfig_Url(t *testing.T) {
	testSetCfg(&Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Url("/info/user"), "/admin/info/user")

	testSetCfg(&Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Url("/info/user"), "/admin/info/user")
	assert.Equal(t, Get().Url("/info/user") != "/admin/info/user/", true)
}

func TestConfig_UrlRemovePrefix(t *testing.T) {

	testSetCfg(&Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().URLRemovePrefix("/admin/info/user"), "/info/user")
}

func TestConfig_PrefixFixSlash(t *testing.T) {

	testSetCfg(&Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().PrefixFixSlash(), "/admin")

	testSetCfg(&Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().PrefixFixSlash(), "/admin")
}

func TestSet(t *testing.T) {
	testSetCfg(&Config{Theme: "abc"})
	testSetCfg(&Config{Theme: "bcd"})
	assert.Equal(t, Get().Theme, "bcd")
}

func TestStore_URL(t *testing.T) {
	testSetCfg(&Config{
		Store: Store{
			Prefix: "/file",
			Path:   "./uploads",
		},
	})

	assert.Equal(t, Get().Store.URL("/xxxxxx.png"), "/file/xxxxxx.png")

	testSetCfg(&Config{
		Store: Store{
			Prefix: "http://xxxxx.com/xxxx/file",
			Path:   "./uploads",
		},
	})

	assert.Equal(t, Get().Store.URL("/xxxxxx.png"), "http://xxxxx.com/xxxx/file/xxxxxx.png")

	testSetCfg(&Config{
		Store: Store{
			Prefix: "/file",
			Path:   "./uploads",
		},
	})

	assert.Equal(t, Get().Store.URL("http://xxxxx.com/xxxx/file/xxxx.png"), "http://xxxxx.com/xxxx/file/xxxx.png")
}

func TestDatabase_ParamStr(t *testing.T) {
	cfg := Database{
		Driver: DriverMysql,
		Params: map[string]string{
			"parseTime": "true",
		},
	}
	assert.Equal(t, cfg.ParamStr(), "?charset=utf8mb4&parseTime=true")
}

func TestReadFromYaml(t *testing.T) {
	cfg := ReadFromYaml("./config.yaml")
	assert.Equal(t, cfg.Databases.GetDefault().Driver, "mssql")
	assert.Equal(t, cfg.Domain, "localhost")
	assert.Equal(t, cfg.UrlPrefix, "admin")
	assert.Equal(t, cfg.Store.Path, "./uploads")
	assert.Equal(t, cfg.IndexUrl, "/")
	assert.Equal(t, cfg.Debug, true)
	assert.Equal(t, cfg.OpenAdminApi, true)
	assert.Equal(t, cfg.ColorScheme, "skin-black")
}

func TestReadFromINI(t *testing.T) {
	cfg := ReadFromINI("./config.ini")
	assert.Equal(t, cfg.Databases.GetDefault().Driver, "postgresql")
	assert.Equal(t, cfg.Domain, "localhost")
	assert.Equal(t, cfg.UrlPrefix, "admin")
	assert.Equal(t, cfg.Store.Path, "./uploads")
	assert.Equal(t, cfg.IndexUrl, "/")
	assert.Equal(t, cfg.Debug, true)
	assert.Equal(t, cfg.OpenAdminApi, true)
	assert.Equal(t, cfg.ColorScheme, "skin-black")
}

func testSetCfg(cfg *Config) {
	count = 0
	Initialize(cfg)
}

func TestUpdate(t *testing.T) {
	m := map[string]string{
		"access_assets_log_off":             `true`,
		"access_log_off":                    `false`,
		"access_log_path":                   "",
		"allow_del_operation_log":           `false`,
		"animation":                         `{"type":"fadeInUp","duration":0,"delay":0}`,
		"animation_delay":                   `0.00`,
		"animation_duration":                `0`,
		"animation_type":                    `fadeInUp`,
		"app_id":                            `70rv3KwjwjXE`,
		"asset_root_path":                   `./public/`,
		"asset_url":                         "",
		"auth_user_table":                   `goadmin_users`,
		"bootstrap_file_path":               `./../datamodel/bootstrap.go`,
		"color_scheme":                      `skin-black`,
		"custom_403_html":                   "",
		"custom_404_html":                   "",
		"custom_500_html":                   "",
		"custom_foot_html":                  "",
		"custom_head_html":                  ` <link rel="icon" type="image/png" sizes="32x32" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-32x32.png">`,
		"databases":                         `{"default":{"host":"127.0.0.1","port":"3306","user":"root","pwd":"root","name":"godmin","max_idle_con":50,"max_open_con":150,"driver":"mysql","file":"","dsn":""}}`,
		"debug":                             `true`,
		"domain":                            "",
		"env":                               `test`,
		"error_log_off":                     `false`,
		"error_log_path":                    "",
		"exclude_theme_components":          `null`,
		"extra":                             "",
		"file_upload_engine":                `{"name":"local","config":null}`,
		"footer_info":                       "",
		"go_mod_file_path":                  "",
		"hide_app_info_entrance":            `false`,
		"hide_config_center_entrance":       `false`,
		"hide_plugin_entrance":              `false`,
		"hide_tool_entrance":                `false`,
		"hide_visitor_user_center_entrance": `false`,
		"index_url":                         `/`,
		"info_log_off":                      `false`,
		"info_log_path":                     "",
		"language":                          `zh`,
		"logger_encoder_caller":             `full`,
		"logger_encoder_caller_key":         `caller`,
		"logger_encoder_duration":           `string`,
		"logger_encoder_encoding":           `console`,
		"logger_encoder_level":              `capitalColor`,
		"logger_encoder_level_key":          `level`,
		"logger_encoder_message_key":        `msg`,
		"logger_encoder_name_key":           `logger`,
		"logger_encoder_stacktrace_key":     `stacktrace`,
		"logger_encoder_time":               `iso8601`,
		"logger_encoder_time_key":           `ts`,
		"logger_level":                      `0`,
		"logger_rotate_compress":            `false`,
		"logger_rotate_max_age":             `30`,
		"logger_rotate_max_backups":         `5`,
		"logger_rotate_max_size":            `10`,
		"login_logo":                        "",
		"login_title":                       `GoAdmin`,
		"login_url":                         `/login`,
		"logo":                              `<b>Go</b>Admin`,
		"mini_logo":                         `<b>G</b>A`,
		"no_limit_login_ip":                 `false`,
		"open_admin_api":                    `false`,
		"operation_log_off":                 `false`,
		"plugin_file_path":                  `/go/src/github.com/GoAdminGroup/go-admin/examples/gin/plugins.go`,
		"session_life_time":                 `7200`,
		"site_off":                          `false`,
		"sql_log":                           `true`,
		"store":                             `{"path":"./uploads","prefix":"uploads"}`,
		"theme":                             `sword`,
		"title":                             `GoAdmin`,
		"url_prefix":                        `admin`,
	}
	c := &Config{}
	c2 := &Config{}
	if c.Update(m) == nil {
		if c.Language != c2.Language ||
			c.Domain != c2.Domain ||
			c.Theme != c2.Theme ||
			c.Title != c2.Title ||
			c.Logo != c2.Logo ||
			c.MiniLogo != c2.MiniLogo ||
			c.Debug != c2.Debug ||
			c.SiteOff != c2.SiteOff ||
			c.AccessLogOff != c2.AccessLogOff ||
			c.InfoLogOff != c2.InfoLogOff ||
			c.ErrorLogOff != c2.ErrorLogOff ||
			c.AccessAssetsLogOff != c2.AccessAssetsLogOff ||
			c.InfoLogPath != c2.InfoLogPath ||
			c.ErrorLogPath != c2.ErrorLogPath ||
			c.AccessLogPath != c2.AccessLogPath ||
			c.SqlLog != c2.SqlLog ||
			c.Logger.Rotate.MaxSize != c2.Logger.Rotate.MaxSize ||
			c.Logger.Rotate.MaxBackups != c2.Logger.Rotate.MaxBackups ||
			c.Logger.Rotate.MaxAge != c2.Logger.Rotate.MaxAge ||
			c.Logger.Rotate.Compress != c2.Logger.Rotate.Compress ||
			c.Logger.Encoder.Encoding != c2.Logger.Encoder.Encoding ||
			c.Logger.Level != c2.Logger.Level ||
			c.Logger.Encoder.TimeKey != c2.Logger.Encoder.TimeKey ||
			c.Logger.Encoder.LevelKey != c2.Logger.Encoder.LevelKey ||
			c.Logger.Encoder.NameKey != c2.Logger.Encoder.NameKey ||
			c.Logger.Encoder.CallerKey != c2.Logger.Encoder.CallerKey ||
			c.Logger.Encoder.MessageKey != c2.Logger.Encoder.MessageKey ||
			c.Logger.Encoder.StacktraceKey != c2.Logger.Encoder.StacktraceKey ||
			c.Logger.Encoder.Level != c2.Logger.Encoder.Level ||
			c.Logger.Encoder.Time != c2.Logger.Encoder.Time ||
			c.Logger.Encoder.Duration != c2.Logger.Encoder.Duration ||
			c.Logger.Encoder.Caller != c2.Logger.Encoder.Caller ||
			c.ColorScheme != c2.ColorScheme ||
			c.SessionLifeTime != c2.SessionLifeTime ||
			c.CustomHeadHtml != c2.CustomHeadHtml ||
			c.CustomFootHtml != c2.CustomFootHtml ||
			c.Custom404HTML != c2.Custom404HTML ||
			c.Custom403HTML != c2.Custom403HTML ||
			c.Custom500HTML != c2.Custom500HTML ||
			c.BootstrapFilePath != c2.BootstrapFilePath ||
			c.GoModFilePath != c2.GoModFilePath ||
			c.FooterInfo != c2.FooterInfo ||
			c.LoginTitle != c2.LoginTitle ||
			c.AssetUrl != c2.AssetUrl ||
			c.LoginLogo != c2.LoginLogo ||
			c.NoLimitLoginIP != c2.NoLimitLoginIP ||
			c.AllowDelOperationLog != c2.AllowDelOperationLog ||
			c.OperationLogOff != c2.OperationLogOff ||
			c.HideConfigCenterEntrance != c2.HideConfigCenterEntrance ||
			c.HideAppInfoEntrance != c2.HideAppInfoEntrance ||
			c.HideToolEntrance != c2.HideToolEntrance ||
			c.HidePluginEntrance != c2.HidePluginEntrance ||
			c.FileUploadEngine.Name != c2.FileUploadEngine.Name ||
			c.Animation.Type != c2.Animation.Type ||
			c.Animation.Duration != c2.Animation.Duration ||
			c.Animation.Delay != c2.Animation.Delay ||
			!reflect.DeepEqual(c.Extra, c2.Extra) {
			panic("c.Extra")
		}
	}
}

func TestToMap(t *testing.T) {
	c := &Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
		MiniLogo:  "<asdfadsf>",
		Animation: PageAnimation{
			Type: "12313213",
		},
		SessionLifeTime:        40,
		ExcludeThemeComponents: []string{"asdfas", "sadfasf"},
	}
	m := c.ToMap()
	fmt.Println(m)
	fmt.Println(m["prefix"], m["animation_type"], m["mini_logo"])

	arr := []string{
		"language", "databases", "domain", "url_prefix", "theme", "store", "title", "logo", "mini_logo", "index_url",
		"site_off", "login_url", "debug", "env", "open_admin_api", "hide_visitor_user_center_entrance",
		"info_log_path", "error_log_path", "access_log_path", "sql_log", "access_log_off", "info_log_off", "error_log_off",
		"access_assets_log_off",
		"logger_rotate_max_size", "logger_rotate_max_backups", "logger_rotate_max_age", "logger_rotate_compress",
		"logger_encoder_time_key", "logger_encoder_level_key", "logger_encoder_name_key", "logger_encoder_caller_key",
		"logger_encoder_message_key", "logger_encoder_stacktrace_key", "logger_encoder_level", "logger_encoder_time",
		"logger_encoder_duration", "logger_encoder_caller", "logger_encoder_encoding", "logger_level",
		"color_scheme", "session_life_time", "asset_url", "file_upload_engine", "custom_head_html", "custom_foot_html",
		"custom_404_html", "custom_403_html", "custom_500_html", "bootstrap_file_path", "go_mod_file_path", "footer_info",
		"app_id", "login_title", "login_logo", "auth_user_table", "exclude_theme_components",
		"extra",
		"animation_type", "animation_duration", "animation_delay",
		"no_limit_login_ip", "allow_del_operation_log", "operation_log_off",
		"hide_config_center_entrance", "hide_app_info_entrance", "hide_tool_entrance", "hide_plugin_entrance",
		"asset_root_path",
	}

	for key := range m {
		if !utils.InArray(arr, key) {
			panic(key)
		}
	}

	fmt.Println(len(arr), len(m))
}
