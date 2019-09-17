package plugins

import "testing"

func TestLoadFromPlugin(t *testing.T) {
	LoadFromPlugin("./example/go_plugin/plugin.so")
}
