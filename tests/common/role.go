package common

import (
	"fmt"
	"net/http"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/gavv/httpexpect"
)

func roleTest(e *httpexpect.Expect, sesID *http.Cookie) {

	fmt.Println()
	printlnWithColor("Role", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Url("/info/roles")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).
		Body().Contains("Administrator").Contains("Operator")

	// show new form

	printlnWithColor("show new form", "green")
	formBody := e.GET(config.Url("/info/roles/new")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	token := reg.FindStringSubmatch(formBody.Raw())

	// new roles tester

	printlnWithColor("new roles test", "green")
	res := e.POST(config.Url("/new/roles")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("permission_id[]", "3").
		WithForm(map[string]interface{}{
			"name":           "tester",
			"slug":           "tester",
			form.PreviousKey: config.Url("/info/roles?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			form.TokenKey:    token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Url("/info/"))
	res.Body().Contains("tester")

	// show form: without id

	//printlnWithColor("show form: without id", "green")
	//e.GET(config.Url("/info/roles/edit")).
	//	WithCookie(sesID.Name, sesID.Value).
	//	Expect().Status(200).Body().Contains(errors.WrongID)

	// show form

	printlnWithColor("show form", "green")
	formBody = e.GET(config.Url("/info/roles/edit")).
		WithQuery(constant.EditPKKey, "3").
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// edit form

	printlnWithColor("edit form", "green")
	res = e.POST(config.Url("/edit/roles")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("permission_id[]", "3").
		WithFormField("permission_id[]", "2").
		WithForm(map[string]interface{}{
			"name":           "tester",
			"slug":           "tester",
			form.PreviousKey: config.Url("/info/roles?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			form.TokenKey:    token[1],
			"id":             "3",
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Url("/info/"))
	res.Body().Contains("tester")

	// show new form

	printlnWithColor("show new form", "green")
	formBody = e.GET(config.Url("/info/roles/new")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// new tester2

	printlnWithColor("new tester2", "green")
	e.POST(config.Url("/new/roles")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("permission_id[]", "3").
		WithForm(map[string]interface{}{
			"name":           "tester2",
			"slug":           "tester2",
			form.PreviousKey: config.Url("/info/roles?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			form.TokenKey:    token[1],
		}).Expect().Status(200)

	// delete tester2

	printlnWithColor("delete roles tester2", "green")
	e.POST(config.Url("/delete/roles")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("id", "3").
		Expect().Status(200).JSON().Object().
		ValueEqual("code", 200).
		ValueEqual("msg", "ok")
}
