package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"goAdmin/template"
)

func ShowInstall(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	buffer := new(bytes.Buffer)
	template.GetInstallPage(buffer)

	//rs, _ := mysql.Query("show tables;")
	//fmt.Println(rs[0]["Tables_in_godmin"])

	//rs2, _ := mysql.Query("show columns from users")
	//fmt.Println(rs2[0]["Field"])

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}
