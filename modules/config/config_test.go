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

	Set(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().GetIndexURL(), "/admin")

	Set(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().GetIndexURL(), "/admin")
}

func TestConfig_Index(t *testing.T) {
	Set(Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Index(), "/")
}

func TestConfig_Prefix(t *testing.T) {
	Set(Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Prefix(), "/admin")

	Set(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Prefix(), "/admin")
}

func TestConfig_Url(t *testing.T) {
	Set(Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Url("/info/user"), "/admin/info/user")

	Set(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().Url("/info/user"), "/admin/info/user")
	assert.Equal(t, Get().Url("/info/user") != "/admin/info/user/", true)
}

func TestConfig_UrlRemovePrefix(t *testing.T) {

	Set(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().URLRemovePrefix("/admin/info/user"), "/info/user")
}

func TestConfig_PrefixFixSlash(t *testing.T) {

	Set(Config{
		UrlPrefix: "/admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().PrefixFixSlash(), "/admin")

	Set(Config{
		UrlPrefix: "admin",
		IndexUrl:  "/",
	})

	assert.Equal(t, Get().PrefixFixSlash(), "/admin")
}

func TestSetDefault(t *testing.T) {
	assert.Equal(t, setDefault("/", "/", "/admin"), "/admin")
	assert.Equal(t, setDefault("/", "/ad", "/admin"), "/")
}
