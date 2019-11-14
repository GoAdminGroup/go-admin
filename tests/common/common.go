package common

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/mgutz/ansi"
	"regexp"
)

var reg, _ = regexp.Compile("<input type=\"hidden\" name=\"_t\" value='(.*?)'>")

// Test contains unit test sections of the GoAdmin admin plugin.
func Test(e *httpexpect.Expect) {

	cookie := authTest(e)

	permissionTest(e, cookie)
	roleTest(e, cookie)
	managerTest(e, cookie)
	menuTest(e, cookie)
	operationLogTest(e, cookie)
	userTest(e, cookie)
}

func printlnWithColor(msg string, color string) {
	fmt.Println(ansi.Color(msg, color))
}
