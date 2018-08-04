package main

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"net/http"
)

func TestIndexApi(t *testing.T) {

	serv := GetTestServer()

	req, _ := http.NewRequest("GET", "/", nil)
	resp := serv.SendRequest(req)

	assert.Equal(t, 302, resp.StatusCode())
}