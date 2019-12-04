package main

import (
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/magiconair/properties/assert"
	"path"
	"testing"
)

func TestGetLatestVersion(t *testing.T) {
	assert.Equal(t, getLatestVersion(), system.Version())
}

func TestCompareVersion(t *testing.T) {
	assert.Equal(t, isRequireUpdate(system.Version(), getLatestVersion()), false)
}

func TestGetType(t *testing.T) {
	assert.Equal(t, getType("int(3434)"), "Int")
	assert.Equal(t, path.Ext("sdafdfs.css"), ".css")
}

func TestCamelcase(t *testing.T) {
	assert.Equal(t, camelcase("goadmin_menu"), "goadminMenu")
	assert.Equal(t, camelcase("goadmin"), "goadmin")
}
