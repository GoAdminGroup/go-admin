package common

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/gavv/httpexpect"
	"github.com/mgutz/ansi"
	"regexp"
)

var reg, _ = regexp.Compile("<input type=\"hidden\" name=\"" + form.TokenKey + "\" value='(.*?)'>")

// Test contains unit test sections of the GoAdmin admin plugin.
func Test(e *httpexpect.Expect) {

	cookie := authTest(e)

	// permission check
	permissionTest(e, cookie)
	// role check
	roleTest(e, cookie)
	// manager check
	managerTest(e, cookie)
	// menu check
	menuTest(e, cookie)
	// operation log check
	operationLogTest(e, cookie)
	// get data from outside source check
	externalTest(e, cookie)

	// TODO: selections check
	// TODO: add buttons check
	// TODO: export functions check
	// TODO: filter functions check
}

func printlnWithColor(msg string, color string) {
	fmt.Println(ansi.Color(msg, color))
}
