// Copyright 2018 ChenHonggui.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package context

import (
	"net/http"
	"strings"
	"io/ioutil"
	"regexp"
	"encoding/json"
	"bytes"
)

// Context is the simplify version of web framework context.
// But it is important which will be used in plugins to custom
// the request and response. And adapter will help to transform
// the Context to the web framework`s context. It has three attributes.
// Request and response are belongs to net/http package. UserValue
// is the custom key-value store of context.
type Context struct {
	Request   *http.Request
	Response  *http.Response
	UserValue map[string]interface{}
}

// Path is used in the matching of request and response. URL stores the
// raw register url. RegUrl contains the wildcard which on behalf of
// the route params.
type Path struct {
	URL    string
	Method string
	RegUrl string
}

func (ctx *Context) SetUserValue(key string, value interface{}) {
	ctx.UserValue[key] = value
}

func (ctx *Context) Path() string {
	return ctx.Request.URL.Path
}

func (ctx *Context) Method() string {
	return ctx.Request.Method
}

func NewContext(req *http.Request) *Context {

	return &Context{
		Request:   req,
		UserValue: make(map[string]interface{}, 0),
		Response: &http.Response{
			Header: make(http.Header, 0),
		},
	}
}

type App struct {
	Requests       []Path
	HandlerList    map[Path]Handler
	MiddlewareList []Middleware
	Prefix         string
}

func NewApp() *App {
	return &App{
		Requests:    make([]Path, 0),
		HandlerList: make(map[Path]Handler, 0),
		Prefix: "",
		MiddlewareList: make([]Middleware, 0),
	}
}

type Handler func(ctx *Context)

type Middleware func(handler Handler) Handler

func (app *App) FindRequestByUrl(url string) Path {
	for _, req := range app.Requests {
		if req.URL == url {
			return req
		}
	}
	return Path{}
}

func (app *App) AppendReqAndResp(url, method string, handler Handler) {

	if url == "/" {
		if app.Prefix != "" {
			url = app.Prefix
		}
	} else {
		url = app.Prefix + url
	}

	for _, middleware := range app.MiddlewareList {
		handler = middleware(handler)
	}

	regUrl := ""

	if strings.Contains(url, ":") {
		r, _ := regexp.Compile(":(.*?)/")

		regUrl = r.ReplaceAllString(url, "(.*?)/")

		r, _ = regexp.Compile(":(.*)")

		regUrl = r.ReplaceAllString(regUrl, "(.*)")
	}

	app.Requests = append(app.Requests, Path{
		URL:    url,
		Method: method,
		RegUrl: regUrl,
	})

	app.HandlerList[Path{
		URL:    url,
		Method: method,
		RegUrl: regUrl,
	}] = handler
}

func (app *App) POST(url string, handler Handler) {
	app.AppendReqAndResp(url, "post", handler)
}

func (app *App) GET(url string, handler Handler) {
	app.AppendReqAndResp(url, "get", handler)
}

func (app *App) DELETE(url string, handler Handler) {
	app.AppendReqAndResp(url, "delete", handler)
}

func (app *App) PUT(url string, handler Handler) {
	app.AppendReqAndResp(url, "put", handler)
}

func (app *App) OPTIONS(url string, handler Handler) {
	app.AppendReqAndResp(url, "options", handler)
}

func (app *App) HEAD(url string, handler Handler) {
	app.AppendReqAndResp(url, "head", handler)
}

func (app *App) Group(prefix string, middleware ...Middleware) {
	app.MiddlewareList = append(app.MiddlewareList, middleware...)
	app.Prefix += prefix
}

func (ctx *Context) Write(code int, Header map[string]string, Body string) {
	ctx.Response.StatusCode = code
	for key, head := range Header {
		ctx.Response.Header.Add(key, head)
	}
	ctx.Response.Body = ioutil.NopCloser(strings.NewReader(Body))
}

func (ctx *Context) Json(code int, Body map[string]interface{}) {
	ctx.Response.StatusCode = code
	ctx.Response.Header.Add("Content-Type", "application/json")
	BodyStr, _ := json.Marshal(Body)
	ctx.Response.Body = ioutil.NopCloser(bytes.NewReader(BodyStr))
}

func (ctx *Context) WriteString(Body string) {
	ctx.SetStatusCode(200)
	ctx.Response.Body = ioutil.NopCloser(strings.NewReader(Body))
}

func (ctx *Context) SetStatusCode(code int) {
	ctx.Response.StatusCode = code
}

func (ctx *Context) SetContentType(contentType string) {
	ctx.Response.Header.Add("Content-Type", contentType)
}

func (ctx *Context) LocalIP() string {
	return "127.0.0.1"
}

func (ctx *Context) SetCookie(cookie *http.Cookie) {
	if v := cookie.String(); v != "" {
		ctx.Response.Header.Add("Set-Cookie", v)
	}
}
