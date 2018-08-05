# 安装

### 下载源码

```bash
git clone https://github.com/chenhg5/go-admin $GOPATH/src/goAdmin
```

### 配置数据库连接信息

- config/config.go

```go
package config

var EnvConfig = map[string]interface{}{

    "SERVER_PORT": ":4003",

	"DATABASE_IP":           "127.0.0.1",
	"DATABASE_PORT":         "3306",
	"DATABASE_USER":         "root",
	"DATABASE_PWD":          "root",
	"DATABASE_NAME":         "goadmin",
	"DATABASE_MAX_IDLE_CON": 50,  // 连接池连接数
	"DATABASE_MAX_OPEN_CON": 150, // 最大连接数

	"REDIS_IP":       "127.0.0.1",
	"REDIS_PORT":     "6379",
	"REDIS_PASSWORD": "",
	"REDIS_DB":       1,
	
	"PORTABLE": false,  // 是否跨平台可移植
	
	"AUTH_DOMAIN": "localhost",
}
```

### 导入数据库

将 admin.sql 导入到业务数据库中

<strong>到这里就安装完成了！~~撒花<strong>