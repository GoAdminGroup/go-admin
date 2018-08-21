package main

import (
	"goAdmin/config"
	"goAdmin/components/menu"
)

func main() {

	menu.InitMenu()

	// 开启服务
	InitServer(config.EnvConfig["SERVER_PORT"].(string))

}
