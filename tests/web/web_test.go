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
		popup                   = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[3]`
		popupCloseBtn           = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[3]/div/div/div[3]/button`
		ajaxBtn                 = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[6]/a`
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
		rowActionDropDown       = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[7]/div/a`
		previewAction           = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[7]/div/ul/li[7]/a`
		closePreviewAction      = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[2]/div/div/div[3]/button`
		previewPopup            = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[2]`
		rowAjaxAction           = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[7]/div/ul/li[6]/a`
		rowAjaxPopup            = `/html/body/div[3]`
		closeRowAjaxPopup       = `/html/body/div[3]/div[7]/div/button`
	)

	assert.Equal(t, page.Navigate(url(config.Get().Url("/info/user"))), nil)
	sleep(2)

	contain(t, "Users")

	// Buttons Check
	// =============================

	printlnWithColor("> buttons check", colorBlue)

	click(t, popupBtn)
	sleep(1)

	css(t, page.FindByXPath(popup), "display", "block")

	sleep(1)

	contain(t, "hello world")
	click(t, popupCloseBtn)
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

	printlnWithColor("> filter area check", colorBlue)

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

	printlnWithColor("> row selector check", colorBlue)

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

	printlnWithColor("> row export check", colorBlue)

	assert.Equal(t, page.FindByXPath(actionDropDown).Click(), nil)
	assert.Equal(t, page.FindByXPath(exportBtn).Click(), nil)
	assert.Equal(t, page.FindByClass(`grid-batch-1`).Click(), nil)

	// Action Button Check
	// =============================

	click(t, rowActionDropDown)
	sleep(1)
	click(t, previewAction)
	sleep(1)
	contain(t, "preview content")
	sleep(1)
	click(t, closePreviewAction)
	css(t, page.FindByXPath(previewPopup), "display", "block")

	click(t, closePreviewAction)

	sleep(1)

	css(t, page.FindByXPath(previewPopup), "display", "none")

	click(t, rowActionDropDown)
	click(t, rowAjaxAction)

	css(t, page.FindByXPath(rowAjaxPopup), "display", "block")

	click(t, closeRowAjaxPopup)

	sleep(1)

	css(t, page.FindByXPath(rowAjaxPopup), "display", "none")

	sleep(2)
}

func TestNewPageOperations(t *testing.T) {
	const (
		newPageBtn         = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[3]/a`
		saveBtn            = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`
		resetBtn           = `//*[@id="pjax-container"]/section[2]/div/div[2]/div[2]/form/div[2]/div[2]/div[2]/button`
		nameField          = `//*[@id="name"]`
		ageField           = `//*[@id="age"]`
		passwordField      = `//*[@id="password"]`
		ipField            = `//*[@id="ip"]`
		amountField        = `//*[@id="currency"]`
		fruitOptField      = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[1]`
		genderBoyCheckBox  = `//*[@id="tab-form-1"]/div[3]/div/div[1]`
		experienceDropDown = `//*[@id="tab-form-1"]/div[5]/div/span/span[1]/span/span[2]`
		twoYearsSelection  = `/html/body/span/span/span[2]/ul/li[1]`
		inputTab           = `//*[@id="pjax-container"]/section[2]/div/div[2]/div[2]/form/div[1]/div/div/ul/li[1]`
	)

	click(t, newPageBtn)
	sleep(2)

	click(t, saveBtn)
	contain(t, "error")
	fillNewForm(t)
	click(t, resetBtn)
	click(t, inputTab)
	text(t, page.FindByXPath(ipField), "")
	fillNewForm(t)
	click(t, saveBtn)
}

func fillNewForm(t *testing.T) {
	const (
		nameField          = `//*[@id="name"]`
		ageField           = `//*[@id="age"]`
		passwordField      = `//*[@id="password"]`
		ipField            = `//*[@id="ip"]`
		amountField        = `//*[@id="currency"]`
		fruitOptField      = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[1]`
		genderGirlCheckBox = `//*[@id="tab-form-1"]/div[3]/div/div[2]`
		experienceDropDown = `//*[@id="tab-form-1"]/div[5]/div/span/span[1]/span/span[2]`
		twoYearsSelection  = `/html/body/span/span/span[2]/ul/li[1]`
		selectTab          = `//*[@id="pjax-container"]/section[2]/div/div[2]/div[2]/form/div[1]/div/div/ul/li[2]`
	)

	fill(t, nameField, "jane")
	fill(t, ageField, "15")
	fill(t, passwordField, "12345678")
	fill(t, ipField, "127.0.0.1")
	fill(t, amountField, "15")
	click(t, selectTab)
	click(t, fruitOptField)
	click(t, genderGirlCheckBox)
	click(t, experienceDropDown)
	sleep(1)
	click(t, twoYearsSelection)
}

func testDetailPageOperations(t *testing.T) {

}

func testEditPageOperations(t *testing.T) {

}
