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

func TestInfoTablePageOperations(t *testing.T) {

	const (
		popupBtn                = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[5]/a`
		popup                   = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[6]`
		ajaxBtn                 = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[7]/a`
		ajaxAlert               = `/html/body/div[3]`
		selectionDropDown       = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[2]/div/div/div[1]/div/span/span[1]/span/span[2]`
		selectionLi1            = `/html/body/span/span/span[2]/ul/li[1]`
		selectionLi2            = `/html/body/span/span/span[2]/ul/li[2]`
		selectionRes            = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[2]/div/div/div[1]/div/span/span[1]/span/span[1]`
		multiSelectInput        = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[1]/div/div/div[2]/div/span/span[1]/span/ul/li/input`
		multiSelectLi1          = `/html/body/span/span/span/ul/li[1]`
		multiSelectLi2          = `/html/body/span/span/span/ul/li[2]`
		multiSelectLi3          = `/html/body/span/span/span/ul/li[3]`
		multiSelectRes          = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[1]/div/div/div[2]/div/span/span[1]/span/ul/li[1]`
		radio                   = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[3]/div/div/div[1]/div/div[1]`
		searchBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`
		resetBtn                = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[2]/a`
		rowSelector             = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/button`
		rowSelectCityCheckbox   = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[4]/label/div`
		rowSelectAvatarCheckbox = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[5]/label/div`
		actionDropDown          = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[1]/div`
		exportBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/span/div/button`
	)

	assert.Equal(t, page.Navigate(url(config.Get().Url("/info/user"))), nil)
	sleep(2)

	contain(t, "Users")

	// Buttons Check
	// =============================

	printlnWithColor("buttons check", colorBlue)

	click(t, popupBtn)
	sleep(1)

	css(t, page.FindByXPath(popup), "display", "block")

	sleep(1)

	contain(t, "hello world")
	assert.Equal(t, page.FindByButton("Close").Click(), nil)
	sleep(1)

	css(t, page.FindByXPath(popup), "display", "none")

	assert.Equal(t, page.FindByXPath(ajaxBtn).Click(), nil)
	sleep(1)

	contain(t, "Oh li get")

	css(t, page.FindByXPath(ajaxAlert), "display", "block")

	assert.Equal(t, page.FindByButton("OK").Click(), nil)

	sleep(1)

	css(t, page.FindByXPath(ajaxAlert), "display", "none")

	// Filter Area Check
	// =============================

	printlnWithColor("filter area check", colorBlue)

	click(t, selectionDropDown)

	sleep(1)

	text(t, page.FindByXPath(selectionLi1), "men")
	text(t, page.FindByXPath(selectionLi2), "women")

	click(t, selectionLi2)

	sleep(1)

	attr(t, page.FindByXPath(selectionRes), "title", "women")

	assert.Equal(t, page.FindByXPath(multiSelectInput).Fill(" "), nil)

	sleep(1)

	text(t, page.FindByXPath(multiSelectLi1), "water")
	text(t, page.FindByXPath(multiSelectLi2), "juice")
	text(t, page.FindByXPath(multiSelectLi3), "red bull")

	assert.Equal(t, page.FindByXPath(multiSelectLi3).Click(), nil)

	sleep(1)

	attr(t, page.FindByXPath(multiSelectRes), "title", "red bull")

	click(t, radio)

	click(t, searchBtn)

	sleep(2)

	click(t, resetBtn)

	sleep(2)

	// Row Selector Check
	// =============================

	printlnWithColor("row selector check", colorBlue)

	click(t, rowSelector)
	click(t, rowSelectCityCheckbox)
	click(t, rowSelectAvatarCheckbox)
	assert.Equal(t, page.FindByButton("submit").Click(), nil)

	sleep(2)

	noContain(t, "guangzhou")

	assert.Equal(t, page.FindByID("filter-btn").Click(), nil)

	sleep(1)

	css(t, page.FindByClass("filter-area"), "display", "none")

	// Export Check
	// =============================

	printlnWithColor("row selector check", colorBlue)

	assert.Equal(t, page.FindByXPath(actionDropDown).Click(), nil)
	assert.Equal(t, page.FindByXPath(exportBtn).Click(), nil)
	assert.Equal(t, page.FindByClass(`grid-batch-1`).Click(), nil)

}
