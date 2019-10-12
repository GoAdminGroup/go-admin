package common

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/gavv/httpexpect"
	"net/http"
)

func PermissionTest(e *httpexpect.Expect, sesId *http.Cookie) {

	fmt.Println()
	printlnWithColor("Permission", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Get().Url("/info/permission")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().
		Status(200).
		Body().Contains("Dashboard").Contains("All permission")

	// show new form

	printlnWithColor("show new form", "green")
	formBody := e.GET(config.Get().Url("/info/permission/new")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token := reg.FindStringSubmatch(formBody.Raw())

	// new permission tester

	printlnWithColor("new permission tester", "green")
	res := e.POST(config.Get().Url("/new/permission")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("http_method[]", "GET").
		WithForm(map[string]interface{}{
			"name": "tester",
			"slug": "tester",
			"http_path": `/
/admin/info/op`,
			"_previous_": config.Get().Url("/info/permission?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"_t":         token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/info/"))
	res.Body().Contains("tester").Contains("GET")

	// show form: without id

	printlnWithColor("show form: without id", "green")
	e.GET(config.Get().Url("/info/permission/edit")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body().Contains("wrong id")

	// show form

	printlnWithColor("show form", "green")
	formBody = e.GET(config.Get().Url("/info/permission/edit")).
		WithQuery("id", "3").
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// edit form

	printlnWithColor("edit form", "green")
	res = e.POST(config.Get().Url("/edit/permission")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("http_method[]", "GET").
		WithFormField("http_method[]", "POST").
		WithForm(map[string]interface{}{
			"name": "tester",
			"slug": "tester",
			"http_path": `/
/admin/info/op`,
			"_previous_": config.Get().Url("/info/permission?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"_t":         token[1],
			"id":         "3",
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/info/"))
	res.Body().Contains("tester").Contains("GET,POST")

	// show new form

	printlnWithColor("show new form", "green")
	formBody = e.GET(config.Get().Url("/info/permission/new")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// new tester2

	printlnWithColor("new tester2", "green")
	e.POST(config.Get().Url("/new/permission")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("http_method[]", "GET").
		WithForm(map[string]interface{}{
			"name": "tester2",
			"slug": "tester2",
			"http_path": `/
/admin/info/op`,
			"_previous_": config.Get().Url("/info/permission?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"_t":         token[1],
		}).Expect().Status(200)

	// delete tester2

	printlnWithColor("delete permission tester2", "green")
	e.POST(config.Get().Url("/delete/permission")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("id", "4").
		Expect().Status(200).JSON().Object().
		ValueEqual("code", 200).
		ValueEqual("msg", "ok")
}
