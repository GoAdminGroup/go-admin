package common

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/gavv/httpexpect"
	"net/http"
)

func ManagerTest(e *httpexpect.Expect, sesId *http.Cookie) {

	fmt.Println()
	printlnWithColor("Manager", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Get().Url("/info/manager")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().
		Status(200).
		Body().Contains("Managers").Contains("admin").Contains("1")

	// edit

	printlnWithColor("edit", "green")
	e.POST(config.Get().Url("/edit/manager")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithForm(map[string]interface{}{
			"username":        "admin",
			"name":            "admin1",
			"password":        "admin",
			"role_id[]":       1,
			"permission_id[]": 1,
			"_previous_":      config.Get().Url("/info/manager?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"id":              "1",
			"_t":              "123",
		}).Expect().Status(200).Body().Contains("edit fail, wrong token")

	// show form: without id

	printlnWithColor("show form: without id", "green")
	e.GET(config.Get().Url("/info/manager/edit")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body().Contains("wrong id")

	// show form

	printlnWithColor("show form", "green")
	formBody := e.GET(config.Get().Url("/info/manager/edit")).
		WithQuery("id", "1").
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token := reg.FindStringSubmatch(formBody.Raw())

	// edit form

	printlnWithColor("edit form", "green")
	res := e.POST(config.Get().Url("/edit/manager")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithForm(map[string]interface{}{
			"username":        "admin",
			"name":            "admin1",
			"password":        "admin",
			"avatar":          "",
			"role_id[]":       1,
			"permission_id[]": 1,
			"_previous_":      config.Get().Url("/info/manager?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"id":              "1",
			"_t":              token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/info/"))
	res.Body().Contains("admin1")

	// show new form

	printlnWithColor("show new form", "green")
	formBody = e.GET(config.Get().Url("/info/manager/new")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// new manager tester

	printlnWithColor("new manager tester", "green")
	res = e.POST(config.Get().Url("/new/manager")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithForm(map[string]interface{}{
			"username":        "tester",
			"name":            "tester",
			"password":        "tester",
			"avatar":          "",
			"role_id[]":       1,
			"permission_id[]": 1,
			"_previous_":      config.Get().Url("/info/manager?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"id":              "1",
			"_t":              token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/info/"))
	res.Body().Contains("tester")

	// tester login: wrong password

	printlnWithColor("tester login: wrong password", "green")
	e.POST(config.Get().Url("/signin")).WithForm(map[string]string{
		"username": "tester",
		"password": "admin",
	}).Expect().Status(400)

	// tester login success

	printlnWithColor("tester login success", "green")
	e.POST(config.Get().Url("/signin")).WithForm(map[string]string{
		"username": "tester",
		"password": "tester",
	}).Expect().Status(200).JSON().Equal(map[string]interface{}{
		"code": 200,
		"data": map[string]interface{}{
			"url": "/" + config.Get().UrlPrefix,
		},
		"msg": "ok",
	})

	printlnWithColor("delete", "green")
	e.GET("/pong").Expect().Status(404)

}
