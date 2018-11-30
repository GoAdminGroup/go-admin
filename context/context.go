// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package context

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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
}

// SetUserValue set the value of user context.
func (ctx *Context) SetUserValue(key string, value interface{}) {
	ctx.UserValue[key] = value
}

// Path return the url path.
func (ctx *Context) Path() string {
	return ctx.Request.URL.Path
}

// Method return the request method.
func (ctx *Context) Method() string {
	return ctx.Request.Method
}

// NewContext used in adapter which return a Context with request
// and slice of UserValue and a default Response.
func NewContext(req *http.Request) *Context {

	return &Context{
		Request:   req,
		UserValue: make(map[string]interface{}, 0),
		Response: &http.Response{
			StatusCode: 200,
			Header:     make(http.Header, 0),
		},
	}
}

// Write save the given status code, header and body string into the response.
func (ctx *Context) Write(code int, Header map[string]string, Body string) {
	ctx.Response.StatusCode = code
	for key, head := range Header {
		ctx.AddHeader(key, head)
	}
	ctx.Response.Body = ioutil.NopCloser(strings.NewReader(Body))
}

// Json serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func (ctx *Context) Json(code int, Body map[string]interface{}) {
	ctx.Response.StatusCode = code
	ctx.AddHeader("Content-Type", "application/json")
	BodyStr, err := json.Marshal(Body)
	if err != nil {
		panic(err)
	}
	ctx.Response.Body = ioutil.NopCloser(bytes.NewReader(BodyStr))
}

// Html output html response.
func (ctx *Context) Html(code int, body string) {
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.SetStatusCode(code)
	ctx.WriteString(body)
}

// Write save the given body string into the response.
func (ctx *Context) WriteString(Body string) {
	ctx.Response.Body = ioutil.NopCloser(strings.NewReader(Body))
}

// SetStatusCode save the given status code into the response.
func (ctx *Context) SetStatusCode(code int) {
	ctx.Response.StatusCode = code
}

// SetStatusCode save the given content type header into the response header.
func (ctx *Context) SetContentType(contentType string) {
	ctx.AddHeader("Content-Type", contentType)
}

// LocalIP return the request client ip.
func (ctx *Context) LocalIP() string {
	return "127.0.0.1"
}

// SetCookie save the given cookie obj into the response Set-Cookie header.
func (ctx *Context) SetCookie(cookie *http.Cookie) {
	if v := cookie.String(); v != "" {
		ctx.AddHeader("Set-Cookie", v)
	}
}

// Query get the query parameter of url.
func (ctx *Context) Query(key string) string {
	return ctx.Request.URL.Query().Get(key)
}

// QueryDefault get the query parameter of url. If it is empty, return the default.
func (ctx *Context) QueryDefault(key, def string) string {
	value := ctx.Query(key)
	if value == "" {
		return def
	}
	return value
}

// Headers get the value of request headers key.
func (ctx *Context) Headers(key string) string {
	return ctx.Request.Header.Get(key)
}

// FormValue get the value of request form key.
func (ctx *Context) FormValue(key string) string {
	return ctx.Request.FormValue(key)
}

// AddHeader adds the key, value pair to the header.
func (ctx *Context) AddHeader(key, value string) {
	ctx.Response.Header.Add(key, value)
}

// User return the current login user.
func (ctx *Context) User() interface{} {
	return ctx.UserValue["user"]
}

// App is the key struct of the package. App as a member of plugin
// entity contains the request and the corresponding handler. Prefix
// is the url prefix and MiddlewareList is for control flow.
type App struct {
	Requests       []Path
	tree           *node
	MiddlewareList []Middleware
	Prefix         string
}

// NewApp return an empty app.
func NewApp() *App {
	return &App{
		Requests:       make([]Path, 0),
		tree:           Tree(),
		Prefix:         "",
		MiddlewareList: make([]Middleware, 0),
	}
}

type Handler func(ctx *Context)

type Middleware func(handler Handler) Handler

// AppendReqAndResp stores the request info and handle into app.
// support the route parameter. The route parameter will be recognized as
// wildcard store into the RegUrl of Path struct. For example:
//
//         /user/:id      => /user/(.*)
//         /user/:id/info => /user/(.*?)/info
//
// The RegUrl will be used to recognize the incoming path and find
// the handler.
func (app *App) AppendReqAndResp(url, method string, handler Handler) {

	app.Requests = append(app.Requests, Path{
		URL:    app.Prefix + url,
		Method: method,
	})

	for _, middleware := range app.MiddlewareList {
		handler = middleware(handler)
	}

	app.tree.addPath(stringToArr(app.Prefix+url), method, handler)
}

// Find is public helper method for findPath of tree.
func (app *App) Find(url, method string) Handler {
	return app.tree.findPath(stringToArr(url), method)
}

// POST is a shortcut for app.AppendReqAndResp(url, "post", handler).
func (app *App) POST(url string, handler Handler) {
	app.AppendReqAndResp(url, "post", handler)
}

// GET is a shortcut for app.AppendReqAndResp(url, "get", handler).
func (app *App) GET(url string, handler Handler) {
	app.AppendReqAndResp(url, "get", handler)
}

// DELETE is a shortcut for app.AppendReqAndResp(url, "delete", handler).
func (app *App) DELETE(url string, handler Handler) {
	app.AppendReqAndResp(url, "delete", handler)
}

// PUT is a shortcut for app.AppendReqAndResp(url, "put", handler).
func (app *App) PUT(url string, handler Handler) {
	app.AppendReqAndResp(url, "put", handler)
}

// OPTIONS is a shortcut for app.AppendReqAndResp(url, "options", handler).
func (app *App) OPTIONS(url string, handler Handler) {
	app.AppendReqAndResp(url, "options", handler)
}

// HEAD is a shortcut for app.AppendReqAndResp(url, "head", handler).
func (app *App) HEAD(url string, handler Handler) {
	app.AppendReqAndResp(url, "head", handler)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, HEAD, OPTIONS, DELETE.
func (app *App) ANY(url string, handler Handler) {
	app.AppendReqAndResp(url, "post", handler)
	app.AppendReqAndResp(url, "get", handler)
	app.AppendReqAndResp(url, "delete", handler)
	app.AppendReqAndResp(url, "put", handler)
	app.AppendReqAndResp(url, "options", handler)
	app.AppendReqAndResp(url, "head", handler)
}

// Group add middlewares and prefix for App.
func (app *App) Group(prefix string, middleware ...Middleware) {
	app.MiddlewareList = append(app.MiddlewareList, middleware...)
	app.Prefix += prefix
}
