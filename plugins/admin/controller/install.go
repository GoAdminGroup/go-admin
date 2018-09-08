package controller

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/chenhg5/go-admin/context"
)

func ShowInstall(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	buffer := new(bytes.Buffer)
	//template.GetInstallPage(buffer)

	//rs, _ := mysql.Query("show tables;")
	//fmt.Println(rs[0]["Tables_in_godmin"])

	//rs2, _ := mysql.Query("show columns from users")
	//fmt.Println(rs2[0]["Field"])

	ctx.WriteString(buffer.String())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}

func CheckDatabase(ctx *context.Context) {

	ip := ctx.Request.FormValue("h")
	port := ctx.Request.FormValue("po")
	username := ctx.Request.FormValue("u")
	password := ctx.Request.FormValue("pa")
	databaseName := ctx.Request.FormValue("db")

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

		ctx.SetContentType("application/json")
		ctx.WriteString(`{"code":0, "msg":"连接成功", "data": {"list":` + list + `}}`)
	} else {
		fmt.Println(err)
		fmt.Println(err2)
		ctx.SetContentType("application/json")
		ctx.WriteString(`{"code":500, "msg":"请检查参数是否设置正确"}`)
	}
}
