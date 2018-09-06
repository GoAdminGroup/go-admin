package adapter

import (
	"goAdmin/plugins"
)

type WebFrameWork interface {
	Use(interface{}, []plugins.Plugin) error
}