package controller

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/chenhg5/go-admin/context"
	"net/http"
)

func ShowInstall(ctx *context.Context) {

	buffer := new(bytes.Buffer)
	//template.GetInstallPage(buffer)

	//rs, _ := mysql.Query("show tables;")
	//fmt.Println(rs[0]["Tables_in_godmin"])

	//rs2, _ := mysql.Query("show columns from users")
	//fmt.Println(rs2[0]["Field"])

	ctx.Html(http.StatusOK, buffer.String())
}

func CheckDatabase(ctx *context.Context) {

	ip := ctx.FormValue("h")
	port := ctx.FormValue("po")
	username := ctx.FormValue("u")
	password := ctx.FormValue("pa")
	databaseName := ctx.FormValue("db")

	SqlDB, err := sql.Open("mysql", username+":"+password+"@tcp("+ip+":"+port+")/"+databaseName+"?charset=utf8mb4")
	err2 := SqlDB.Ping()
	defer SqlDB.Close()

	if err == nil && err2 == nil {

		//db.InitDB(username, password, port, ip, databaseName, 100, 100)
		//tables, _ := db.Query("show tables")

		tables := []map[string]interface{}{}

		list := "["

		for i := 0; i < len(tables); i++ {
			if i != len(tables)-1 {
				list += `"` + tables[i]["Tables_in_godmin"].(string) + `",`
			} else {
				list += `"` + tables[i]["Tables_in_godmin"].(string) + `"`
			}
		}
		list += "]"

		fmt.Println(list)

		ctx.Json(http.StatusOK, map[string]interface{}{
			"code": 0,
			"msg":  "连接成功",
			"data": map[string]interface{}{
				"list": list,
			},
		})

	} else {
		fmt.Println(err)
		fmt.Println(err2)
		ctx.Json(http.StatusInternalServerError, map[string]interface{}{
			"code": 500,
			"msg":  "请检查参数是否设置正确",
		})
	}
}
