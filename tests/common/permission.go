package common

import (
	"fmt"
	"net/http"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/gavv/httpexpect"
)

func permissionTest(e *httpexpect.Expect, sesID *http.Cookie) {

	fmt.Println()
	printlnWithColor("Permission", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Url("/info/permission")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).
		Body().Contains("Dashboard").Contains("All permission")

	// show new form

	printlnWithColor("show new form", "green")
	formBody := e.GET(config.Url("/info/permission/new")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	token := reg.FindStringSubmatch(formBody.Raw())

	// new permission tester

	printlnWithColor("new permission test", "green")
	res := e.POST(config.Url("/new/permission")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("http_method[]", "GET").
		WithForm(map[string]interface{}{
			"name": "tester",
			"slug": "tester",
			"http_path": `/
/admin/info/op`,
			form.PreviousKey: config.Url("/info/permission?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			form.TokenKey:    token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Url("/info/"))
	res.Body().Contains("tester").Contains("GET")

	// show form: without id

	//printlnWithColor("show form: without id", "green")
	//e.GET(config.Url("/info/permission/edit")).
	//	WithCookie(sesID.Name, sesID.Value).
	//	Expect().Status(200).Body().Contains(errors.WrongID)

	// show form

	printlnWithColor("show form", "green")
	formBody = e.GET(config.Url("/info/permission/edit")).
		WithQuery(constant.EditPKKey, "3").
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// edit form

	printlnWithColor("edit form", "green")
	res = e.POST(config.Url("/edit/permission")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("http_method[]", "GET").
		WithFormField("http_method[]", "POST").
		WithForm(map[string]interface{}{
			"name": "tester",
			"slug": "tester",
			"http_path": `/
/admin/info/op`,
			form.PreviousKey: config.Url("/info/permission?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			form.TokenKey:    token[1],
			"id":             "3",
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Url("/info/"))
	res.Body().Contains("tester").Contains("GET,POST")

	// show new form

	printlnWithColor("show new form", "green")
	formBody = e.GET(config.Url("/info/permission/new")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// new tester2

	printlnWithColor("new tester2", "green")
	e.POST(config.Url("/new/permission")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("http_method[]", "GET").
		WithForm(map[string]interface{}{
			"name": "tester2",
			"slug": "tester2",
			"http_path": `/
/admin/info/op`,
			form.PreviousKey: config.Url("/info/permission?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			form.TokenKey:    token[1],
		}).Expect().Status(200)

	// delete tester2

	printlnWithColor("delete permission tester2", "green")
	e.POST(config.Url("/delete/permission")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("id", "4").
		Expect().Status(200).JSON().Object().
		ValueEqual("code", 200).
		ValueEqual("msg", "ok")
}
