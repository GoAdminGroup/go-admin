package common

import (
	"fmt"
	"net/http"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/gavv/httpexpect"
)

func apiTest(e *httpexpect.Expect, sesID *http.Cookie) {

	fmt.Println()
	printlnWithColor("Api", "blue")
	fmt.Println("============================")

	printlnWithColor("show", "green")
	e.GET(config.Url("/api/list/manager")).
		WithHeader("Accept", "application/json, text/plain, */*").
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).JSON().Object().ValueEqual("code", 200)

	//printlnWithColor("update form without id", "green")
	//e.GET(config.Url("/api/edit/form/manager")).
	//	WithHeader("Accept", "application/json, text/plain, */*").
	//	WithCookie(sesID.Name, sesID.Value).
	//	Expect().
	//	Status(400).JSON().Object().ValueEqual("code", 400)

	printlnWithColor("update form", "green")
	e.GET(config.Url("/api/edit/form/manager")).
		WithHeader("Accept", "application/json, text/plain, */*").
		WithQuery(constant.EditPKKey, "1").
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).JSON().Object().ValueEqual("code", 200)

	printlnWithColor("create form", "green")
	e.GET(config.Url("/api/create/form/manager")).
		WithHeader("Accept", "application/json, text/plain, */*").
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).JSON().Object().ValueEqual("code", 200)
}
