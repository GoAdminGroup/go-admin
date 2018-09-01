package context

import (
	"net/http"
	"strings"
	"io/ioutil"
	"regexp"
)

type Path struct {
	URL    string
	Method string
	RegUrl string
}

type Response struct {
	StatusCode int
	Body       string
	Header     map[string]string
}

type Context struct {
	Request   *http.Request
	Response  *http.Response
	UserValue map[string]interface{}
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
		Response:  &http.Response{
			Header: make(http.Header, 0),
		},
	}
}

type App struct {
	Requests    []Path
	HandlerList map[Path]Handler
}

func NewApp() *App {
	return &App{
		Requests:    make([]Path, 0),
		HandlerList: make(map[Path]Handler, 0),
	}
}

type Handler func(ctx *Context)

func (app *App) FindRequestByUrl(url string) Path {
	for _, req := range app.Requests {
		if req.URL == url {
			return req
		}
	}
	return Path{}
}

func (app *App) AppendReqAndResp(url, method string, handler Handler) {

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

func (app *App) FILE(url string, handler Handler) {
	app.AppendReqAndResp(url, "head", handler)
}

func (ctx *Context) Write(code int, Header map[string]string, Body string) {
	ctx.Response.StatusCode = code
	for key, head := range Header {
		ctx.Response.Header.Add(key, head)
	}
	ctx.Response.Body = ioutil.NopCloser(strings.NewReader(Body))
}

func (ctx *Context) WriteString(Body string) {
	ctx.SetStatusCode(200)
	ctx.Response.Body = ioutil.NopCloser(strings.NewReader(Body))
}

func (ctx *Context) SetStatusCode(code int)  {
	ctx.Response.StatusCode = code
}

func (ctx *Context) SetContentType(contentType string)  {
	ctx.Response.Header.Add("Content-Type", contentType)
}

func (ctx *Context) LocalIP() string {
	return "127.0.0.1"
}

func (ctx *Context) SetCookie(cookie *http.Cookie)  {
	if v := cookie.String(); v != "" {
		ctx.Response.Header.Add("Set-Cookie", v)
	}
}