package web

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	defer StopDriverOnPanic(t)

	assert.Equal(t, page.Navigate(url("/admin")), nil)
	assert.Equal(t, page.Find("#username").Fill("admin"), nil)
	assert.Equal(t, page.Find("#password").Fill("admin"), nil)
	assert.Equal(t, page.FindByButton("login").Click(), nil)

	time.Sleep(time.Second * 3)

	content, err := page.HTML()
	assert.Equal(t, err, nil)
	assert.Equal(t, strings.Contains(content, "main-header"), true)
}

func TestButtons(t *testing.T) {
	assert.Equal(t, page.Navigate(url(config.Get().Url("/info/user"))), nil)
	time.Sleep(time.Second * 2)

	content, err := page.HTML()
	assert.Equal(t, err, nil)
	assert.Equal(t, strings.Contains(content, "Users"), true)
	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[5]/a`).Click(), nil)
	time.Sleep(time.Second * 1)
	popup := page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[6]`)
	display, err := popup.CSS("display")
	assert.Equal(t, err, nil)
	assert.Equal(t, display, "block")

	time.Sleep(time.Second * 1)

	content, err = page.HTML()
	assert.Equal(t, err, nil)
	assert.Equal(t, strings.Contains(content, "hello world"), true)
	assert.Equal(t, page.FindByButton("Close").Click(), nil)
	time.Sleep(time.Second * 1)
	popup = page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[6]`)
	display, err = popup.CSS("display")
	assert.Equal(t, err, nil)
	assert.Equal(t, display, "none")

	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[7]/a`).Click(), nil)
	time.Sleep(time.Second * 1)

	content, err = page.HTML()
	assert.Equal(t, err, nil)
	assert.Equal(t, strings.Contains(content, "Oh li get"), true)

	alert := page.FindByXPath("/html/body/div[3]")
	display, err = alert.CSS("display")
	assert.Equal(t, err, nil)
	assert.Equal(t, display, "block")

	assert.Equal(t, page.FindByButton("OK").Click(), nil)

	time.Sleep(time.Second * 1)

	alert = page.FindByXPath("/html/body/div[3]")
	display, err = alert.CSS("display")
	assert.Equal(t, err, nil)
	assert.Equal(t, display, "none")

}
