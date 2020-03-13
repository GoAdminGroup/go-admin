package web

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

const (
	// Info Table Page

	newPageBtn              = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[3]/a`
	editPageBtn             = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[7]/div/ul/li[1]/a`
	genderActionDropDown    = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[7]/div/span/span[1]/span/span[2]`
	menOptionActionBtn      = `/html/body/span/span/span[2]/ul/li[2]`
	rowActionDropDown       = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[7]/div/a`
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
	filterResetBtn          = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[2]/a`
	rowSelector             = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/button`
	rowSelectCityCheckbox   = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[4]/label/div`
	rowSelectAvatarCheckbox = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[5]/label/div`
	actionDropDown          = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[1]/div`
	exportBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/span/div/button`
	previewAction           = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[7]/div/ul/li[7]/a`
	closePreviewAction      = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[2]/div/div/div[3]/button`
	previewPopup            = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[2]`
	rowAjaxAction           = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[7]/div/ul/li[6]/a`
	rowAjaxPopup            = `/html/body/div[3]`
	closeRowAjaxPopup       = `/html/body/div[3]/div[7]/div/button`

	// Form Page

	saveBtn              = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`
	resetBtn             = `//*[@id="pjax-container"]/section[2]/div/div[2]/div[2]/form/div[2]/div[2]/div[2]/button`
	nameField            = `//*[@id="name"]`
	ageField             = `//*[@id="age"]`
	emailField           = `//*[@id="email"]`
	birthdayField        = `//*[@id="birthday"]`
	passwordField        = `//*[@id="password"]`
	homePageField        = `//*[@id="homepage"]`
	ipField              = `//*[@id="ip"]`
	amountField          = `//*[@id="money"]`
	appleOptField        = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[1]`
	bananaOptField       = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[2]`
	watermelonOptField   = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[3]`
	pearOptField         = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[4]`
	genderBoyCheckBox    = `//*[@id="tab-form-1"]/div[3]/div/div[1]`
	genderGirlCheckBox   = `//*[@id="tab-form-1"]/div[3]/div/div[2]`
	experienceDropDown   = `//*[@id="tab-form-1"]/div[5]/div/span/span[1]/span/span[2]`
	twoYearsSelection    = `/html/body/span/span/span[2]/ul/li[1]`
	threeYearsSelection  = `/html/body/span/span/span[2]/ul/li[2]`
	fourYearsSelection   = `/html/body/span/span/span[2]/ul/li[3]`
	fiveYearsSelection   = `/html/body/span/span/span[2]/ul/li[4]`
	inputTab             = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/ul/li[1]`
	selectTab            = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/ul/li[2]`
	multiSelectionInput  = `//*[@id="tab-form-1"]/div[4]/div/span/span[1]/span/ul/li[2]/input`
	multiSelectedOpt     = `//*[@id="tab-form-1"]/div[4]/div/span/span[1]/span/ul/li[1]`
	multiBeerOpt         = `/html/body/span/span/span/ul/li[1]`
	multiJuiceOpt        = `/html/body/span/span/span/ul/li[2]`
	multiWaterOpt        = `/html/body/span/span/span/ul/li[3]`
	multiRedBullOpt      = `/html/body/span/span/span/ul/li[4]`
	continueEditCheckBox = `//*[@id="pjax-container"]/section[2]/div/div[2]/div[2]/form/div[2]/div[2]/label/div/ins`
	boxSelectedOpt       = `//*[@id="bootstrap-duallistbox-selected-list_fruit[]"]/option`
	singleSelectedOpt    = `//*[@id="tab-form-1"]/div[5]/div/span/span[1]/span/span[1]`
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

	navigate(t, "/info/user")

	contain(t, "Users")

	// Buttons Check
	// =============================

	printPart("buttons check")

	click(t, popupBtn)

	css(t, page.FindByXPath(popup), "display", "block")

	wait(1)

	contain(t, "hello world")
	click(t, popupCloseBtn)

	css(t, page.FindByXPath(popup), "display", "none")

	click(t, ajaxBtn)

	contain(t, "Oh li get")

	css(t, page.FindByXPath(ajaxAlert), "display", "block")

	clickS(t, page.FindByButton("OK"))

	css(t, page.FindByXPath(ajaxAlert), "display", "none")

	// Filter Area Check
	// =============================

	printPart("filter area check")

	click(t, selectionDropDown)

	text(t, selectionLi1, "men")
	text(t, selectionLi2, "women")

	click(t, selectionLi2)

	attr(t, page.FindByXPath(selectionRes), "title", "women")

	fill(t, multiSelectInput, " ")

	wait(1)

	text(t, multiSelectLi1, "water")
	text(t, multiSelectLi2, "juice")
	text(t, multiSelectLi3, "red bull")

	click(t, multiSelectLi3)

	attr(t, page.FindByXPath(multiSelectRes), "title", "red bull")

	click(t, radio)

	click(t, searchBtn, 2)

	click(t, filterResetBtn, 2)

	// Row Selector Check
	// =============================

	printPart("row selector check")

	click(t, rowSelector)
	click(t, rowSelectCityCheckbox)
	click(t, rowSelectAvatarCheckbox)

	clickS(t, page.FindByButton("submit"), 2)

	noContain(t, "guangzhou")

	clickS(t, page.FindByID("filter-btn"))

	css(t, page.FindByClass("filter-area"), "display", "none")

	// Export Check
	// =============================

	printPart("row export check")

	assert.Equal(t, page.FindByXPath(actionDropDown).Click(), nil)
	assert.Equal(t, page.FindByXPath(exportBtn).Click(), nil)
	assert.Equal(t, page.FindByClass(`grid-batch-1`).Click(), nil)

	// Action Button Check
	// =============================

	printPart("action buttons check")

	click(t, rowActionDropDown)
	click(t, previewAction)
	contain(t, "preview content")
	css(t, page.FindByXPath(previewPopup), "display", "block")

	click(t, closePreviewAction)

	css(t, page.FindByXPath(previewPopup), "display", "none")

	click(t, rowActionDropDown)
	click(t, rowAjaxAction)

	css(t, page.FindByXPath(rowAjaxPopup), "display", "block")

	click(t, closeRowAjaxPopup)

	css(t, page.FindByXPath(rowAjaxPopup), "display", "none")

	wait(2)
}

func TestNewPageOperations(t *testing.T) {

	click(t, newPageBtn, 2)
	attr(t, page.FindByXPath(homePageField), "value", "http://google.com")

	// Selections Form Item Check
	// =============================

	printPart("selections form items check")

	checkSelectionsInForm(t)

	// Create Error Check
	// =============================

	printPart("create error check")

	click(t, saveBtn)
	contain(t, "error")

	// Reset Check
	// =============================

	printPart("reset error check")

	fillNewForm(t, "jane", "girl")
	click(t, resetBtn)

	// Continue Creating Check
	// =============================

	printPart("continue creating check")

	click(t, inputTab)
	text(t, ipField, "")
	click(t, continueEditCheckBox)
	fillNewForm(t, "jane", "girl")
	click(t, saveBtn)

	// Creating Check
	// =============================

	printPart("creating check")

	fillNewForm(t, "harry", "boy")
	click(t, saveBtn, 2)

	noContain(t, "harry")
	click(t, genderActionDropDown)
	click(t, menOptionActionBtn, 2)
	contain(t, "harry")
}

func fillNewForm(t *testing.T, name, gender string) {

	fill(t, nameField, name)
	fill(t, ageField, "15")
	fill(t, passwordField, "12345678")
	fill(t, ipField, "127.0.0.1")
	fill(t, amountField, "15")
	click(t, selectTab)
	click(t, appleOptField)
	if gender == "girl" {
		click(t, genderGirlCheckBox)
	} else {
		click(t, genderBoyCheckBox)
	}
	click(t, experienceDropDown)
	click(t, twoYearsSelection)
}

func checkSelectionsInForm(t *testing.T) {
	click(t, selectTab)
	text(t, appleOptField, "Apple")
	text(t, bananaOptField, "Banana")
	text(t, watermelonOptField, "Watermelon")
	text(t, pearOptField, "Pear")
	attr(t, page.FindByXPath(multiSelectedOpt), "title", "Beer")
	click(t, multiSelectionInput)
	text(t, multiBeerOpt, "Beer")
	text(t, multiJuiceOpt, "Juice")
	text(t, multiWaterOpt, "Water")
	text(t, multiRedBullOpt, "Red bull")
	click(t, experienceDropDown)
	text(t, twoYearsSelection, "two years")
	text(t, threeYearsSelection, "three years")
	text(t, fourYearsSelection, "four years")
	text(t, fiveYearsSelection, "five years")
	click(t, inputTab)
}

func TestEditPageOperations(t *testing.T) {
	click(t, rowActionDropDown)
	click(t, editPageBtn, 2)

	// Form Field Value Check
	// =============================

	printPart("edit form values check")

	value(t, nameField, "harry")
	value(t, homePageField, "http://google.com")
	value(t, ageField, "15")
	value(t, emailField, "xxxx@xxx.com")
	value(t, birthdayField, "2010-09-05 00:00:00")
	value(t, passwordField, "12345678")
	value(t, ipField, "127.0.0.1")
	value(t, amountField, "15.00")

	click(t, selectTab)

	text(t, boxSelectedOpt, "Pear")
	attr(t, page.FindByXPath(multiSelectedOpt), "title", "Beer")
	text(t, singleSelectedOpt, "two years")
}

func testDetailPageOperations(t *testing.T) {

}

func testManagerPageOperations(t *testing.T) {

}

func testRolePageOperations(t *testing.T) {

}

func testPermissionPageOperations(t *testing.T) {

}

func testMenuPageOperations(t *testing.T) {

}
