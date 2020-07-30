package common

import (
	"fmt"
	"net/http"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/gavv/httpexpect"
)

func managerTest(e *httpexpect.Expect, sesID *http.Cookie) {

	fmt.Println()
	printlnWithColor("Manager", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Url("/info/manager")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).
		Body().Contains("Managers").Contains("admin").Contains("1")

	// edit

	printlnWithColor("edit", "green")
	e.POST(config.Url("/edit/manager")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithForm(map[string]interface{}{
			"username":        "admin",
			"name":            "admin1",
			"password":        "admin",
			"password_again":  "admin",
			"role_id[]":       1,
			"permission_id[]": 1,
			form.PreviousKey:  config.Url("/info/manager?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"id":              "1",
			form.TokenKey:     "123",
		}).Expect().Status(200).Body().Contains(errors.EditFailWrongToken)

	// show form: without id

	//printlnWithColor("show form: without id", "green")
	//e.GET(config.Url("/info/manager/edit")).
	//	WithCookie(sesID.Name, sesID.Value).
	//	Expect().Status(200).Body().Contains(errors.WrongID)

	// show form

	printlnWithColor("show form", "green")
	formBody := e.GET(config.Url("/info/manager/edit")).
		WithQuery(constant.EditPKKey, "1").
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	token := reg.FindStringSubmatch(formBody.Raw())

	// edit form

	printlnWithColor("edit form", "green")
	res := e.POST(config.Url("/edit/manager")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithForm(map[string]interface{}{
			"username":            "admin",
			"name":                "admin1",
			"password":            "admin",
			"password_again":      "admin",
			"avatar__delete_flag": "0",
			"role_id[]":           1,
			"permission_id[]":     1,
			form.PreviousKey:      config.Url("/info/manager?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"id":                  "1",
			form.TokenKey:         token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Url("/info/"))
	res.Body().Contains("admin1")

	// show new form

	printlnWithColor("show new form", "green")
	formBody = e.GET(config.Url("/info/manager/new")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// new manager tester

	printlnWithColor("new manager test", "green")
	res = e.POST(config.Url("/new/manager")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithForm(map[string]interface{}{
			"username":            "tester",
			"name":                "tester",
			"password":            "tester",
			"password_again":      "tester",
			"avatar__delete_flag": "0",
			"role_id[]":           1,
			"permission_id[]":     1,
			form.PreviousKey:      config.Url("/info/manager?__page=1&__pageSize=10&__sort=id&__sort_type=desc"),
			"id":                  "1",
			form.TokenKey:         token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Url("/info/"))
	res.Body().Contains("tester")

	// tester login: wrong password

	printlnWithColor("tester login: wrong password", "green")
	e.POST(config.Url("/signin")).WithForm(map[string]string{
		"username": "tester",
		"password": "admin",
	}).Expect().Status(400)

	// tester login success

	printlnWithColor("tester login success", "green")
	e.POST(config.Url("/signin")).WithForm(map[string]string{
		"username": "tester",
		"password": "tester",
	}).Expect().Status(200).JSON().Equal(map[string]interface{}{
		"code": 200,
		"data": map[string]interface{}{
			"url": "/" + config.GetUrlPrefix(),
		},
		"msg": "ok",
	})

	printlnWithColor("delete", "green")
	e.GET("/pong").Expect().Status(404)

}
