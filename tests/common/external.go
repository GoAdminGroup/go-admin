package common

import (
	"fmt"
	"net/http"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/gavv/httpexpect"
)

func externalTest(e *httpexpect.Expect, sesID *http.Cookie) {
	fmt.Println()
	printlnWithColor("External", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Url("/info/external")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).
		Body().Contains("External").Contains("this is a title").Contains("10")

	// show form: without id

	//printlnWithColor("show form: without id", "green")
	//e.GET(config.Url("/info/external/edit")).
	//	WithCookie(sesID.Name, sesID.Value).
	//	Expect().Status(200).Body().Contains(errors.WrongID)

	// show form

	printlnWithColor("show form", "green")
	e.GET(config.Url("/info/external/edit")).
		WithQuery(constant.EditPKKey, "10").
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	// show new form

	printlnWithColor("show new form", "green")
	e.GET(config.Url("/info/external/new")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()
}
