package iris

import (
	"github.com/chenhg5/go-admin/tests/common"
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func TestIris(t *testing.T) {
	// TODO: BUG: invalid memory address or nil pointer dereference
	common.Test(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(NewIrisHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
