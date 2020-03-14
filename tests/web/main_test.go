package web

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/mgutz/ansi"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/tests/tables"
	"github.com/gin-gonic/gin"
	"github.com/sclevine/agouti"
)

var (
	debugMode  = false
	driver     *agouti.WebDriver
	page       *agouti.Page
	optionList = []string{
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
	quit = make(chan uint8)
)

const (
	baseURL = "http://localhost:9033"
	port    = ":9033"
)

func init() {
	if os.Args[len(os.Args)-1] == "true" {
		debugMode = true
	}
	if !debugMode {
		optionList = append(optionList, "--headless")
	}
}

func startServer() {

	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	r := gin.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	template.AddComp(chartjs.NewChart())

	cfg := config.ReadFromJson("./config.json")
	if debugMode {
		cfg.SqlLog = true
		cfg.Debug = true
		cfg.AccessLogOff = false
	}

	if err := eng.AddConfig(cfg).
		AddPlugins(adminPlugin).
		Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	r.Static("/uploads", "./uploads")

	go func() {
		_ = r.Run(port)
	}()

	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}

func TestMain(m *testing.M) {

	go startServer()

	var err error

	driver = agouti.ChromeDriver(
		agouti.ChromeOptions("args", optionList),
		agouti.Desired(
			agouti.Capabilities{
				"loggingPrefs": map[string]string{
					"performance": "ALL",
				},
				"acceptSslCerts":      true,
				"acceptInsecureCerts": true,
			},
		))
	err = driver.Start()
	if err != nil {
		panic("failed to start driver, error: " + err.Error())
	}

	page, err = driver.NewPage()
	if err != nil {
		panic("failed to open page, error: " + err.Error())
	}

	fmt.Println()
	fmt.Println("============================================")
	printlnWithColor("User Acceptance Testing", "blue")
	fmt.Println("============================================")
	fmt.Println()

	test := m.Run()

	wait(2)

	if !debugMode {
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

	quit <- 0
	os.Exit(test)
}

func StopDriverOnPanic(t *testing.T) {
	if r := recover(); r != nil {
		debug.PrintStack()
		fmt.Println("Recovered in f", r)
		_ = page.Destroy()
		_ = driver.Stop()
		t.Fail()
		quit <- 0
	}
}

func url(suffix string) string {
	return baseURL + suffix
}

func wait(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}

func contain(t *testing.T, s string) {
	content, err := page.HTML()
	assert.Equal(t, err, nil)
	assert.Equal(t, strings.Contains(content, s), true)
}

func noContain(t *testing.T, s string) {
	content, err := page.HTML()
	assert.Equal(t, err, nil)
	assert.Equal(t, strings.Contains(content, s), false)
}

func css(t *testing.T, xpath, css, res string) {
	style, err := page.FindByXPath(xpath).CSS(css)
	assert.Equal(t, err, nil)
	assert.Equal(t, style, res)
}

func cssS(t *testing.T, s *agouti.Selection, css, res string) {
	style, err := s.CSS(css)
	assert.Equal(t, err, nil)
	assert.Equal(t, style, res)
}

func text(t *testing.T, xpath, text string) {
	mli1, err := page.FindByXPath(xpath).Text()
	assert.Equal(t, err, nil)
	assert.Equal(t, mli1, text)
}

func display(t *testing.T, xpath string) {
	css(t, xpath, "display", "block")
}

func nondisplay(t *testing.T, xpath string) {
	css(t, xpath, "display", "none")
}

func value(t *testing.T, xpath, value string) {
	val, err := page.FindByXPath(xpath).Attribute("value")
	assert.Equal(t, err, nil)
	assert.Equal(t, val, value)
}

func click(t *testing.T, xpath string, intervals ...int) {
	assert.Equal(t, page.FindByXPath(xpath).Click(), nil)
	interval := 1
	if len(intervals) > 0 {
		interval = intervals[0]
	}
	wait(interval)
}

func clickS(t *testing.T, s *agouti.Selection, intervals ...int) {
	assert.Equal(t, s.Click(), nil)
	interval := 1
	if len(intervals) > 0 {
		interval = intervals[0]
	}
	wait(interval)
}

func attr(t *testing.T, s *agouti.Selection, attr, res string) {
	style, err := s.Attribute(attr)
	assert.Equal(t, err, nil)
	assert.Equal(t, style, res)
}

func printlnWithColor(msg string, color string) {
	fmt.Println(ansi.Color(msg, color))
}

func fill(t *testing.T, xpath, content string) {
	assert.Equal(t, page.FindByXPath(xpath).Fill(content), nil)
}

func navigate(t *testing.T, path string) {
	assert.Equal(t, page.Navigate(url(config.Get().Url(path))), nil)
	wait(2)
}

func printPart(part string) {
	printlnWithColor("> "+part, colorBlue)
}

const (
	colorBlue  = "blue"
	colorGreen = "green"
)
