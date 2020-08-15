package common

import (
	"fmt"
	"net/http"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/gavv/httpexpect"
)

func normalTest(e *httpexpect.Expect, sesID *http.Cookie) {

	fmt.Println()
	printlnWithColor("Normal Table", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Url("/info/user")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).
		Body().Contains("Users")

	// export

	printlnWithColor("export test", "green")
	e.POST(config.Url("/export/user")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("id", "1").
		Expect().Status(200)

	// show form: without id

	//printlnWithColor("show form: without id", "green")
	//e.GET(config.Url("/info/user/edit")).
	//	WithCookie(sesID.Name, sesID.Value).
	//	Expect().Status(200).Body().Contains(errors.WrongID)

	// show form

	//printlnWithColor("show form", "green")
	//e.GET(config.Url("/info/user/edit")).
	//	WithQuery(constant.EditPKKey, "362").
	//	WithCookie(sesID.Name, sesID.Value).
	//	Expect().Status(200).Body()

	// show new form

	printlnWithColor("show new form", "green")
	e.GET(config.Url("/info/user/new")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

}
