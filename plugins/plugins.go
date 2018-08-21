package plugins

import (
	"goAdmin/plugins/admin"
	"github.com/buaazp/fasthttprouter"
)

var Plugins = map[string]RegisterFunc{
	"admin": admin.RegisterAdmin,
}

type RegisterFunc func(*fasthttprouter.Router) *fasthttprouter.Router

func RegisterService(router *fasthttprouter.Router) *fasthttprouter.Router {
	for _, v := range Plugins {
		router = v(router)
	}
	return router
}
