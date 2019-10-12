package common

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/gavv/httpexpect"
	"net/http"
)

func MenuTest(e *httpexpect.Expect, sesId *http.Cookie) {

	fmt.Println()
	printlnWithColor("Menu", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	formBody := e.GET(config.Get().Url("/menu")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().
		Status(200).
		Body().Contains(language.Get("menus manage"))

	token := reg.FindStringSubmatch(formBody.Raw())

	// new menu tester

	printlnWithColor("new menu tester", "green")
	res := e.POST(config.Get().Url("/menu/new")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("roles[]", "1").
		WithForm(map[string]interface{}{
			"parent_id":  0,
			"title":      "test menu",
			"header":     "",
			"icon":       "fa-angellist",
			"uri":        "/example/test",
			"_previous_": "/admin/menu",
			"_t":         token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/menu"))
	res.Body().Contains("test menu").Contains("/example/test")

	// show form: without id

	printlnWithColor("show form: without id", "green")
	e.GET(config.Get().Url("/menu/edit/show")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body().Contains("wrong id")

	// show form

	printlnWithColor("show form", "green")
	formBody = e.GET(config.Get().Url("/menu/edit/show")).
		WithQuery("id", "3").
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// edit form

	printlnWithColor("edit form", "green")
	res = e.POST(config.Get().Url("/menu/edit")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("roles[]", "1").
		WithForm(map[string]interface{}{
			"parent_id":  0,
			"title":      "test2 menu",
			"header":     "",
			"icon":       "fa-angellist",
			"uri":        "/example/test",
			"_previous_": "/admin/menu",
			"_t":         token[1],
			"id":         "3",
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/menu"))
	res.Body().Contains("test2 menu").Contains("/example/test")

	// new tester2

	printlnWithColor("new tester2", "green")
	e.POST(config.Get().Url("/menu/new")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("roles[]", "1").
		WithForm(map[string]interface{}{
			"parent_id":  0,
			"title":      "test2 menu",
			"header":     "",
			"icon":       "fa-angellist",
			"uri":        "/example/test2",
			"_previous_": "/admin/menu",
			"_t":         token[1],
		}).Expect().Status(200)

	// delete tester2

	printlnWithColor("delete menu tester2", "green")
	e.POST(config.Get().Url("/menu/delete")).
		WithQuery("id", "9").
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		Expect().Status(200).JSON().Object().
		ValueEqual("code", 200).
		ValueEqual("msg", "ok")
}
