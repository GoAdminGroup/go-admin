package main

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"net/http"
)

func TestIndexApi(t *testing.T) {

	serv, _ := GetTestServer()

	req, _ := http.NewRequest("GET", "/", nil)
	resp, _ := serv.SendRequest(req)

	assert.Equal(t, 302, resp.StatusCode())
}
