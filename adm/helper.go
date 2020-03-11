package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/mgutz/ansi"
)

func cliInfo() {
	fmt.Println("GoAdmin CLI " + system.Version() + compareVersion(system.Version()))
	fmt.Println()
}

func exitWithError(msg string) {
	fmt.Println()
	fmt.Println(ansi.Color("go-admin cli error: "+msg, "red"))
	fmt.Println()
	os.Exit(-1)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getLatestVersion() string {
	http.DefaultClient.Timeout = 3 * time.Second
	res, err := http.Get("https://goproxy.cn/github.com/!go!admin!group/go-admin/@v/list")

	if err != nil || res.Body == nil {
		return ""
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return ""
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil || body == nil {
		return ""
	}

	versionsArr := strings.Split(string(body), "\n")

	return versionsArr[len(versionsArr)-1]
}

func isRequireUpdate(srcVersion, toCompareVersion string) bool {
	if toCompareVersion == "" {
		return false
	}

	exp, _ := regexp.Compile(`-(.*)`)
	srcVersion = exp.ReplaceAllString(srcVersion, "")
	toCompareVersion = exp.ReplaceAllString(toCompareVersion, "")

	srcVersionArr := strings.Split(strings.Replace(srcVersion, "v", "", -1), ".")
	toCompareVersionArr := strings.Split(strings.Replace(toCompareVersion, "v", "", -1), ".")

	for i := len(srcVersionArr) - 1; i > -1; i-- {
		v, err := strconv.Atoi(srcVersionArr[i])
		if err != nil {
			return false
		}
		vv, err := strconv.Atoi(toCompareVersionArr[i])
		if err != nil {
			return false
		}
		if v < vv {
			return true
		} else if v > vv {
			return false
		} else {
			continue
		}
	}

	return false
}

func compareVersion(srcVersion string) string {
	toCompareVersion := getLatestVersion()
	if isRequireUpdate(srcVersion, toCompareVersion) {
		return ", the latest version is " + toCompareVersion + " now."
	}
	return ""
}
