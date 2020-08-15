package controller

import (
	"bytes"
	"database/sql"
	"net/http"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
)

// ShowInstall show install page.
func (h *Handler) ShowInstall(ctx *context.Context) {

	buffer := new(bytes.Buffer)
	//template.GetInstallPage(buffer)

	//rs, _ := mysql.Query("show tables;")
	//fmt.Println(rs[0]["Tables_in_godmin"])

	//rs2, _ := mysql.Query("show columns from users")
	//fmt.Println(rs2[0]["Field"])

	ctx.HTML(http.StatusOK, buffer.String())
}

// CheckDatabase check the database connection.
func (h *Handler) CheckDatabase(ctx *context.Context) {

	ip := ctx.FormValue("h")
	port := ctx.FormValue("po")
	username := ctx.FormValue("u")
	password := ctx.FormValue("pa")
	databaseName := ctx.FormValue("db")

	SqlDB, err := sql.Open("mysql", username+":"+password+"@tcp("+ip+":"+port+")/"+databaseName+"?charset=utf8mb4")
	if SqlDB != nil {
		if SqlDB.Ping() != nil {
			response.Error(ctx, "请检查参数是否设置正确")
			return
		}
	}

	defer func() {
		_ = SqlDB.Close()
	}()

	if err != nil {
		response.Error(ctx, "请检查参数是否设置正确")
		return

	}

	//db.InitDB(username, password, port, ip, databaseName, 100, 100)

	tables := make([]map[string]interface{}, 0)

	list := "["

	for i := 0; i < len(tables); i++ {
		if i != len(tables)-1 {
			list += `"` + tables[i]["Tables_in_godmin"].(string) + `",`
		} else {
			list += `"` + tables[i]["Tables_in_godmin"].(string) + `"`
		}
	}
	list += "]"

	response.OkWithData(ctx, map[string]interface{}{
		"list": list,
	})
}
