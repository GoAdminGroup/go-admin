package main

import (
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetLatestVersion(t *testing.T) {
	assert.Equal(t, getLatestVersion(), system.Version())
}
