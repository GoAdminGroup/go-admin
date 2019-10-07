package common

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/mgutz/ansi"
	"regexp"
)

var reg, _ = regexp.Compile("<input type=\"hidden\" name=\"_t\" value='(.*?)'>")

func Test(e *httpexpect.Expect) {

	cookie := AuthTest(e)

	PermissionTest(e, cookie)
	RoleTest(e, cookie)
	ManagerTest(e, cookie)
	MenuTest(e, cookie)
	OperationLogTest(e, cookie)
	UserTest(e, cookie)
}

func printlnWithColor(msg string, color string) {
	fmt.Println(ansi.Color(msg, color))
}
