package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/mgutz/ansi"

	"github.com/GoAdminGroup/go-admin/modules/system"
)

func cliInfo() {
	fmt.Println("GoAdmin CLI " + system.Version() + compareVersion(system.Version()))
	fmt.Println()
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

func compareVersion(srcVersion string) string {
	toCompareVersion := getLatestVersion()
	if utils.CompareVersion(srcVersion, toCompareVersion) {
		return ", the latest version is " + toCompareVersion + " now."
	}
	return ""
}

func printSuccessInfo(msg string) {
	fmt.Println()
	fmt.Println()
	fmt.Println(ansi.Color(getWord(msg), "green"))
	fmt.Println()
	fmt.Println()
}

func newError(msg string) error {
	return errors.New(getWord(msg))
}
