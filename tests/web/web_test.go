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
	sleep(2)

	contain(t, "Users")
	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[5]/a`).Click(), nil)
	sleep(1)

	popup := `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[6]`

	css(t, page.FindByXPath(popup), "display", "block")

	sleep(1)

	contain(t, "hello world")
	assert.Equal(t, page.FindByButton("Close").Click(), nil)
	sleep(1)

	css(t, page.FindByXPath(popup), "display", "none")

	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[7]/a`).Click(), nil)
	sleep(1)

	contain(t, "Oh li get")

	css(t, page.FindByXPath("/html/body/div[3]"), "display", "block")

	assert.Equal(t, page.FindByButton("OK").Click(), nil)

	sleep(1)

	css(t, page.FindByXPath("/html/body/div[3]"), "display", "none")

	dropDown := `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[2]/div/div/div[1]/div/span/span[1]/span/span[2]`

	assert.Equal(t, page.FindByXPath(dropDown).Click(), nil)

	sleep(1)

	text(t, page.FindByXPath(`/html/body/span/span/span[2]/ul/li[1]`), "men")
	text(t, page.FindByXPath(`/html/body/span/span/span[2]/ul/li[2]`), "women")

	assert.Equal(t, page.FindByXPath(`/html/body/span/span/span[2]/ul/li[2]`).Click(), nil)

	sleep(1)

	liRes, err := page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[2]/div/div/div[1]/div/span/span[1]/span/span[1]`).Attribute("title")
	assert.Equal(t, err, nil)
	assert.Equal(t, liRes, "women")

	multiSelectInput := `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[1]/div/div/div[2]/div/span/span[1]/span/ul/li/input`

	assert.Equal(t, page.FindByXPath(multiSelectInput).Fill(" "), nil)

	sleep(1)

	text(t, page.FindByXPath(`/html/body/span/span/span/ul/li[1]`), "water")
	text(t, page.FindByXPath(`/html/body/span/span/span/ul/li[2]`), "juice")
	text(t, page.FindByXPath(`/html/body/span/span/span/ul/li[3]`), "red bull")

	assert.Equal(t, page.FindByXPath(`/html/body/span/span/span/ul/li[3]`).Click(), nil)

	sleep(1)

	mliRes, err := page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[1]/div/div/div[2]/div/span/span[1]/span/ul/li[1]`).Attribute("title")
	assert.Equal(t, err, nil)
	assert.Equal(t, mliRes, "red bull")

	radio1 := page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[3]/div/div/div[1]/div/div[1]`)
	assert.Equal(t, radio1.Click(), nil)

	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`).Click(), nil)

	sleep(2)

	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[2]/a`).Click(), nil)

	sleep(2)

	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/button`).Click(), nil)
	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[4]/label/div`).Click(), nil)
	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[5]/label/div`).Click(), nil)
	assert.Equal(t, page.FindByButton("submit").Click(), nil)

	sleep(2)

	noContain(t, "guangzhou")

	assert.Equal(t, page.FindByID("filter-btn").Click(), nil)

	sleep(1)

	css(t, page.FindByClass("filter-area"), "display", "none")

	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[1]/div`).Click(), nil)
	assert.Equal(t, page.FindByXPath(`//*[@id="pjax-container"]/section[2]/div/div/div[1]/span/div/button`).Click(), nil)
	assert.Equal(t, page.FindByClass(`grid-batch-1`).Click(), nil)

}
