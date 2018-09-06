package controller

import "goAdmin/modules/config"

var Config config.Config

func SetConfig(cfg config.Config)  {
	Config = cfg
}
