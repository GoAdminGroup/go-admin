package web

import (
	"fmt"
	"testing"
	"time"

	"github.com/mgutz/ansi"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/sclevine/agouti"
)

type Testers func(t *testing.T, page *Page)
type ServerStarter func(quit chan struct{})

// UserAcceptanceTestSuit make sure the chromedriver version Is corresponding to the
// chrome version. Using the following link to get the latest version of Chrome and ChromeDriver.
// https://googlechromelabs.github.io/chrome-for-testing/
func UserAcceptanceTestSuit(t *testing.T, testers Testers, serverStarter ServerStarter, local bool, options ...string) {
	var quit = make(chan struct{})
	go serverStarter(quit)

	if len(options) == 0 {
		options = []string{
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36",
			"--window-size=1500,900",
			"--incognito",
			"--blink-settings=imagesEnabled=true",
			"--no-default-browser-check",
			"--ignore-ssl-errors=true",
			"--ssl-protocol=any",
			"--no-sandbox",
			"--disable-breakpad",
			"--disable-gpu",
			"--disable-logging",
			"--no-zygote",
			"--allow-running-insecure-content",
		}
		if !local {
			options = append(options, "--headless")
		}
	}

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", options),
		agouti.Desired(
			agouti.Capabilities{
				"loggingPrefs": map[string]string{
					"performance": "ALL",
				},
				"acceptSslCerts":      true,
				"acceptInsecureCerts": true,
			},
		))
	err := driver.Start()
	if err != nil {
		panic("failed to start driver, error: " + err.Error())
	}

	page, err := driver.NewPage()
	if err != nil {
		panic("failed to open page, error: " + err.Error())
	}

	fmt.Println()
	fmt.Println("============================================")
	printlnWithColor("User Acceptance Testing", "blue")
	fmt.Println("============================================")
	fmt.Println()

	testers(t, &Page{T: t, Page: page, Driver: driver, Quit: quit})

	wait(2)

	if !local {
		err = page.CloseWindow()
		if err != nil {
			fmt.Println("failed to close page, error: ", err)
		}

		err = page.Destroy()
		if err != nil {
			fmt.Println("failed to destroy page, error: ", err)
		}

		err = driver.Stop()
		if err != nil {
			fmt.Println("failed to stop driver, error: ", err)
		}
	}

	quit <- struct{}{}
}

func printlnWithColor(msg string, color string) {
	fmt.Println(ansi.Color(msg, color))
}

func printPart(part string) {
	printlnWithColor("> "+part, colorBlue)
}

func wait(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}

const basePath = "http://localhost:9033"

func url(suffix string) string {
	if suffix == "/" {
		suffix = ""
	}
	return basePath + "/admin" + suffix
}

const (
	colorBlue  = "blue"
	colorGreen = "green"
)
