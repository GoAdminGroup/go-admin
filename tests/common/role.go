package common

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/gavv/httpexpect"
	"net/http"
)

func RoleTest(e *httpexpect.Expect, sesId *http.Cookie) {

	fmt.Println()
	printlnWithColor("Role", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Get().Url("/info/roles")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().
		Status(200).
		Body().Contains("Administrator").Contains("Operator")

	// show new form

	printlnWithColor("show new form", "green")
	formBody := e.GET(config.Get().Url("/info/roles/new")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token := reg.FindStringSubmatch(formBody.Raw())

	// new roles tester

	printlnWithColor("new roles tester", "green")
	res := e.POST(config.Get().Url("/new/roles")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("permission_id[]", "3").
		WithForm(map[string]interface{}{
			"name":       "tester",
			"slug":       "tester",
			"_previous_": config.Get().Url("/info/roles?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"_t":         token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/info/"))
	res.Body().Contains("tester")

	// show form: without id

	printlnWithColor("show form: without id", "green")
	e.GET(config.Get().Url("/info/roles/edit")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body().Contains("wrong id")

	// show form

	printlnWithColor("show form", "green")
	formBody = e.GET(config.Get().Url("/info/roles/edit")).
		WithQuery("id", "3").
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// edit form

	printlnWithColor("edit form", "green")
	res = e.POST(config.Get().Url("/edit/roles")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("permission_id[]", "3").
		WithFormField("permission_id[]", "2").
		WithForm(map[string]interface{}{
			"name":       "tester",
			"slug":       "tester",
			"_previous_": config.Get().Url("/info/roles?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"_t":         token[1],
			"id":         "3",
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/info/"))
	res.Body().Contains("tester")

	// show new form

	printlnWithColor("show new form", "green")
	formBody = e.GET(config.Get().Url("/info/roles/new")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// new tester2

	printlnWithColor("new tester2", "green")
	e.POST(config.Get().Url("/new/roles")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("permission_id[]", "3").
		WithForm(map[string]interface{}{
			"name":       "tester2",
			"slug":       "tester2",
			"_previous_": config.Get().Url("/info/roles?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"_t":         token[1],
		}).Expect().Status(200)

	// delete tester2

	printlnWithColor("delete roles tester2", "green")
	e.POST(config.Get().Url("/delete/roles")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("id", "3").
		Expect().Status(200).JSON().Object().
		ValueEqual("code", 200).
		ValueEqual("msg", "ok")
}
