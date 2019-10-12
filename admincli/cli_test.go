package main

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"path"
	"testing"
)

func TestGetLatestVersion(t *testing.T) {
	fmt.Println(getLatestVersion())
}

func TestCompareVersion(t *testing.T) {
	fmt.Println(isRequireUpdate(system.Version, getLatestVersion()))
}

func TestGetType(t *testing.T) {
	fmt.Println(getType("int(3434)"))
	fmt.Println(path.Ext("sdafdfs.css"))
}

func TestGetAllFiles(t *testing.T) {
	fmt.Println(getAllFiles("."))
}
