package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestSetDefault(t *testing.T) {
	assert.Equal(t, setDefault("/", "/", "/admin"), "/admin")
	assert.Equal(t, setDefault("/", "/ad", "/admin"), "/")
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

func testSetCfg(cfg Config) {
	count = 0
	Set(cfg)
}
