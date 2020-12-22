package language

import (
	"html/template"
	"testing"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	Add("cn", map[string]string{})
}

func TestGetWithScope(t *testing.T) {
	config.Initialize(&config.Config{
		Language: CN,
	})
	cn["foo"] = "bar"
	assert.Equal(t, GetWithScope("foo"), "bar")
	cn["user.table.foo2"] = "bar"
	assert.Equal(t, GetWithScope("foo2"), "foo2")
	assert.Equal(t, GetWithScope("foo2", "user"), "foo2")
	assert.Equal(t, GetWithScope("foo2", "user", "table"), "bar")
}

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{
		Language: CN,
	})
	cn["foo"] = "bar"
	assert.Equal(t, Get("foo"), "bar")
}

func TestWithScopes(t *testing.T) {
	assert.Equal(t, WithScopes("foo", "user", "table"), "user.table.foo")
}

func TestGetFromHtml(t *testing.T) {
	config.Initialize(&config.Config{
		Language: CN,
	})
	cn["user.table.foo"] = "bar"
	assert.Equal(t, GetFromHtml("foo", "user", "table"), template.HTML("bar"))
}
