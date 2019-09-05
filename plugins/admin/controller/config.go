package controller

import (
	c "github.com/chenhg5/go-admin/modules/config"
)

var config c.Config

func SetConfig(cfg c.Config) {
	config = cfg
}
