package adapter

import (
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/template/types"
)

type WebFrameWork interface {
	Use(interface{}, []plugins.Plugin) error
	Content(interface{}, types.GetPanel)
}
