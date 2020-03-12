package common

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/gavv/httpexpect"
	"net/http"
)

func normalTest(e *httpexpect.Expect, sesID *http.Cookie) {

	fmt.Println()
	printlnWithColor("Normal Table", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Get().Url("/info/user")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).
		Body().Contains("Users")

	// show form: without id

	printlnWithColor("show form: without id", "green")
	e.GET(config.Get().Url("/info/user/edit")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body().Contains("wrong id")

	// show form

	printlnWithColor("show form", "green")
	e.GET(config.Get().Url("/info/user/edit")).
		WithQuery(constant.EditPKKey, "10").
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	// show new form

	printlnWithColor("show new form", "green")
	e.GET(config.Get().Url("/info/user/new")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

}
