package common

import (
	"fmt"
	"net/http"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/gavv/httpexpect"
)

func menuTest(e *httpexpect.Expect, sesID *http.Cookie) {

	fmt.Println()
	printlnWithColor("Menu", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	formBody := e.GET(config.Url("/menu")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).
		Body().Contains(language.Get("menus manage"))

	token := reg.FindStringSubmatch(formBody.Raw())

	// new menu tester

	printlnWithColor("new menu test", "green")
	res := e.POST(config.Url("/menu/new")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("roles[]", "1").
		WithForm(map[string]interface{}{
			"parent_id":      0,
			"title":          "test menu",
			"header":         "",
			"icon":           "fa-angellist",
			"uri":            "/example/test",
			form.PreviousKey: "/admin/menu",
			form.TokenKey:    token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Url("/menu"))
	res.Body().Contains("test menu").Contains("/example/test")

	// show form: without id

	//printlnWithColor("show form: without id", "green")
	//e.GET(config.Url("/menu/edit/show")).
	//	WithCookie(sesID.Name, sesID.Value).
	//	Expect().Status(200).Body().Contains(errors.WrongID)

	// show form

	printlnWithColor("show form", "green")
	formBody = e.GET(config.Url("/menu/edit/show")).
		WithQuery(constant.EditPKKey, "3").
		WithCookie(sesID.Name, sesID.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// edit form

	printlnWithColor("edit form", "green")
	res = e.POST(config.Url("/menu/edit")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("roles[]", "1").
		WithForm(map[string]interface{}{
			"parent_id":      0,
			"title":          "test2 menu",
			"header":         "",
			"icon":           "fa-angellist",
			"uri":            "/example/test",
			form.PreviousKey: "/admin/menu",
			form.TokenKey:    token[1],
			"id":             "3",
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Url("/menu"))
	res.Body().Contains("test2 menu").Contains("/example/test")

	// new tester2

	printlnWithColor("new tester2", "green")
	e.POST(config.Url("/menu/new")).
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		WithFormField("roles[]", "1").
		WithForm(map[string]interface{}{
			"parent_id":      0,
			"title":          "test2 menu",
			"header":         "",
			"icon":           "fa-angellist",
			"uri":            "/example/test2",
			form.PreviousKey: "/admin/menu",
			form.TokenKey:    token[1],
		}).Expect().Status(200)

	// delete tester2

	printlnWithColor("delete menu tester2", "green")
	e.POST(config.Url("/menu/delete")).
		WithQuery("id", "9").
		WithCookie(sesID.Name, sesID.Value).
		WithMultipart().
		Expect().Status(200).JSON().Object().
		ValueEqual("code", 200).
		ValueEqual("msg", "delete succeed")
}
