package controller

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIsInfoUrl(t *testing.T) {
	u := "https://localhost:8098/admin/info/user?id=sdfs"
	assert.Equal(t, true, isInfoUrl(u))
}

func TestIsNewUrl(t *testing.T) {
	u := "https://localhost:8098/admin/info/user/new?id=sdfs"
	assert.Equal(t, true, isNewUrl(u, "user"))
}
