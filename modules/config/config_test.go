package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_GetIndexUrl(t *testing.T) {
	Set(Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().GetIndexURL(), "/admin")

	testSetCfg(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().GetIndexURL(), "/admin")

	testSetCfg(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().GetIndexURL(), "/admin")
}

func TestConfig_Index(t *testing.T) {
	testSetCfg(Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Index(), "/")
}

func TestConfig_Prefix(t *testing.T) {
	testSetCfg(Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Prefix(), "/admin")

	testSetCfg(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Prefix(), "/admin")
}

func TestConfig_Url(t *testing.T) {
	testSetCfg(Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Url("/info/user"), "/admin/info/user")

	testSetCfg(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Url("/info/user"), "/admin/info/user")
	assert.Equal(t, Get().Url("/info/user") != "/admin/info/user/", true)
}

func TestConfig_UrlRemovePrefix(t *testing.T) {

	testSetCfg(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().URLRemovePrefix("/admin/info/user"), "/info/user")
}

func TestConfig_PrefixFixSlash(t *testing.T) {

	testSetCfg(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().PrefixFixSlash(), "/admin")

	testSetCfg(Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().PrefixFixSlash(), "/admin")
}

func TestSet(t *testing.T) {
	testSetCfg(Config{Theme: "abc"})
	testSetCfg(Config{Theme: "bcd"})
	assert.Equal(t, Get().Theme, "bcd")
}

func TestStore_URL(t *testing.T) {
	testSetCfg(Config{
		Store: Store{
			Prefix: "/file",
			Path:   "./uploads",
		},
	})

	assert.Equal(t, Get().Store.URL("/xxxxxx.png"), "/file/xxxxxx.png")

	testSetCfg(Config{
		Store: Store{
			Prefix: "http://xxxxx.com/xxxx/file",
			Path:   "./uploads",
		},
	})

	assert.Equal(t, Get().Store.URL("/xxxxxx.png"), "http://xxxxx.com/xxxx/file/xxxxxx.png")

	testSetCfg(Config{
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

func testSetCfg(cfg Config) {
	count = 0
	Set(cfg)
}
