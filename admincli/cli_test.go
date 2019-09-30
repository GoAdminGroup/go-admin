package main

import (
	"fmt"
	"github.com/chenhg5/go-admin/modules/system"
	"testing"
)

func TestGetLatestVersion(t *testing.T) {
	fmt.Println(getLatestVersion())
}

func TestCompareVersion(t *testing.T) {
	fmt.Println(isRequireUpdate(system.Version, getLatestVersion()))
}
