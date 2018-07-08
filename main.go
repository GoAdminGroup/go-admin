package main

import (
	"goAdmin/config"
	"goAdmin/connections/mysql"
	"goAdmin/connections/redis"
	"runtime"
	"goAdmin/menu"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化数据库
	mysql.InitDB(config.EnvConfig["DATABASE_USER"].(string),
		config.EnvConfig["DATABASE_PWD"].(string),
		config.EnvConfig["DATABASE_PORT"].(string),
		config.EnvConfig["DATABASE_IP"].(string),
		config.EnvConfig["DATABASE_NAME"].(string))

	// 初始化redis
	redis.InitRedis()

	menu.InitMenu()

	// 开启服务
	InitServer(config.EnvConfig["SERVER_PORT"].(string))

}
