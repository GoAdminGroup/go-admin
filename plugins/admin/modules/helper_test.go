package modules

import (
	"github.com/magiconair/properties/assert"
	"regexp"
	"testing"
)

func TestInArray(t *testing.T) {
	assert.Equal(t, isFormURL("/admin/info/profile/new"), true)
}

func isFormURL(s string) bool {
	reg, _ := regexp.Compile("(.*?)info/(.*)/(new|edit)(.*?)")
	return reg.MatchString(s)
}
