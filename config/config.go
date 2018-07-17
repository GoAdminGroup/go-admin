package config

var EnvConfig = map[string]interface{}{

	"SERVER_PORT": ":4003",

	"DATABASE_IP":           "127.0.0.1",
	"DATABASE_PORT":         "3306",
	"DATABASE_USER":         "root",
	"DATABASE_PWD":          "root",
	"DATABASE_NAME":         "godmin",
	"DATABASE_MAX_IDLE_CON": 50,  // 连接池连接数
	"DATABASE_MAX_OPEN_CON": 150, // 最大连接数

	"REDIS_IP":       "127.0.0.1",
	"REDIS_PORT":     "6379",
	"REDIS_PASSWORD": "",
	"REDIS_DB":       1,

	"PORTABLE": false,
}
