package main

import (
	"goAdmin/config"
	"goAdmin/connections/redis"
	"runtime"
	"goAdmin/menu"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化redis
	redis.InitRedis()

	menu.InitMenu()

	// 开启服务
	InitServer(config.EnvConfig["SERVER_PORT"].(string))

}
