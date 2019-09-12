package main

import (
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/plugins/example"
)

var Plugin plugins.Plugin = example.NewExample()
