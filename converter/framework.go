package converter

import (
	"github.com/chenhg5/go-admin/plugins"
)

type WebFrameWork interface {
	Use(interface{}, []plugins.Plugin) error
}