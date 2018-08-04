package main

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"bufio"
	"net/http"
	"fmt"
	"strings"
	"net"
)

func TestIndexApi(t *testing.T) {

	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	router := InitRouter()
	go fasthttp.Serve(ln, router.Handler)

	c, err := ln.Dial()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	req, _ := http.NewRequest("GET", "/", nil)
	resp, _ := SendRequest(&c, req)

	assert.Equal(t, 302, resp.StatusCode())
}

func SendRequest(c *net.Conn, req *http.Request) (resp fasthttp.Response, err error) {
	req.Host = "127.0.0.1"

	if _, err = (*c).Write([]byte(FormatRequest(req))); err != nil {
		return resp, err
	}
	br := bufio.NewReader(*c)
	if err = resp.Read(br); err != nil {
		return resp, err
	}
	return resp, nil
}

func FormatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
	r.ParseForm()
	request = append(request, "\n")
	request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n") + "\r\n\r\n"
}