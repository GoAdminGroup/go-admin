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
	editPageBtn             = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[10]/div/ul/li[1]/a`
	genderActionDropDown    = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[7]/div/span/span[1]/span/span[2]`
	menOptionActionBtn      = `/html/body/span/span/span[2]/ul/li[2]`
	idOrderBtn              = `//*[@id="sort-id"]`
	rowActionDropDown       = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[10]/div/a`
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
	filterNameField         = `//*[@id="name"]`
	radio                   = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[3]/div/div/div[1]/div/div[1]`
	searchBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`
	filterResetBtn          = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[2]/a`
	rowSelector             = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/button`
	rowSelectCityCheckbox   = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[4]/label/div`
	rowSelectAvatarCheckbox = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[5]/label/div`
	actionDropDown          = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[1]/div`
	exportBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/span/div/button`
	previewAction           = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[10]/div/ul/li[7]/a`
	closePreviewAction      = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[2]/div/div/div[3]/button`
	previewPopup            = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[2]`
	rowAjaxAction           = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[10]/div/ul/li[6]/a`
	rowAjaxPopup            = `/html/body/div[3]`
	closeRowAjaxPopup       = `/html/body/div[3]/div[7]/div/button`

	// Form Page

	saveBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`
	resetBtn              = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[2]/button`
	nameField             = `//*[@id="name"]`
	ageField              = `//*[@id="age"]`
	emailField            = `//*[@id="email"]`
	birthdayField         = `//*[@id="birthday"]`
	passwordField         = `//*[@id="password"]`
	homePageField         = `//*[@id="homepage"]`
	ipField               = `//*[@id="ip"]`
	amountField           = `//*[@id="money"]`
	appleOptField         = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[1]`
	bananaOptField        = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[2]`
	watermelonOptField    = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[3]`
	pearOptField          = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[4]`
	genderBoyCheckBox     = `//*[@id="tab-form-1"]/div[5]/div/div[1]`
	genderGirlCheckBox    = `//*[@id="tab-form-1"]/div[5]/div/div[2]`
	experienceDropDown    = `//*[@id="tab-form-1"]/div[7]/div/span/span[1]/span/span[2]`
	twoYearsSelection     = `/html/body/span/span/span[2]/ul/li[1]`
	threeYearsSelection   = `/html/body/span/span/span[2]/ul/li[2]`
	fourYearsSelection    = `/html/body/span/span/span[2]/ul/li[3]`
	fiveYearsSelection    = `/html/body/span/span/span[2]/ul/li[4]`
	inputTab              = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/ul/li[1]`
	selectTab             = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/ul/li[2]`
	multiSelectionInput   = `//*[@id="tab-form-1"]/div[6]/div/span/span[1]/span/ul/li[2]/input`
	multiSelectedOpt      = `//*[@id="tab-form-1"]/div[6]/div/span/span[1]/span/ul/li[1]`
	multiBeerOpt          = `/html/body/span/span/span/ul/li[1]`
	multiJuiceOpt         = `/html/body/span/span/span/ul/li[2]`
	multiWaterOpt         = `/html/body/span/span/span/ul/li[3]`
	multiRedBullOpt       = `/html/body/span/span/span/ul/li[4]`
	continueEditCheckBox  = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/label/div`
	boxSelectedOpt        = `//*[@id="bootstrap-duallistbox-selected-list_fruit[]"]/option`
	experienceSelectedOpt = `//*[@id="tab-form-1"]/div[7]/div/span/span[1]/span/span[1]`

	sideBarManageDropDown    = `/html/body/div[1]/aside/section/ul/li[2]/a/span[2]`
	menuPageBtn              = `/html/body/div[1]/aside/section/ul/li[2]/ul/li[4]/a`
	menuParentIdDropDown     = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/div/div[2]/form/div[1]/div/div/div[1]/div/span/span[1]/span/span[2]`
	parentIdRootOpt          = `/html/body/span/span/span[2]/ul/li[1]`
	parentIdDashboardOpt     = `/html/body/span/span/span[2]/ul/li[2]`
	parentIdAdminOpt         = `/html/body/span/span/span[2]/ul/li[3]`
	parentIdUserOpt          = `/html/body/span/span/span[2]/ul/li[4]`
	menuRoleDropDown         = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/div/div[2]/form/div[1]/div/div/div[6]/div/span/span[1]/span`
	menuRoleAdminOpt         = `/html/body/span/span/span/ul/li[1]`
	menuRoleOperatorOpt      = `/html/body/span/span/span/ul/li[2]`
	iconPopupBtn             = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/div/div[2]/form/div[1]/div/div/div[4]/div/div[1]/span`
	iconPopup                = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/div/div[2]/form/div[1]/div/div/div[4]/div/div[2]`
	iconBtn                  = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/div/div[2]/form/div[1]/div/div/div[4]/div/div[2]/div[3]/div/div/a[5]`
	menuNameInput            = `//*[@id="title"]`
	menuUriInput             = `//*[@id="uri"]`
	menuInfoSaveBtn          = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/div/div[2]/form/div[2]/div[2]/div[1]/button`
	testMenuItem             = `//*[@id="tree-model"]/ol/li[2]`
	testMenuDeleteBtn        = `//*[@id="tree-model"]/ol/li[2]/div/span/a[2]`
	testMenuDeleteConfirmBtn = `/html/body/div[3]/div[7]/div/button`
	menuOkBtn                = `/html/body/div[3]/div[7]/div/button`

	managerPageBtn               = `/html/body/div[1]/aside/section/ul/li[2]/ul/li[1]/a`
	managerEditBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[3]/td[8]/a[1]`
	managerNameField             = `//*[@id="username"]`
	managerNickNameField         = `//*[@id="name"]`
	managerRoleSelectedOpt       = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[5]/div/span[1]/span[1]/span/ul/li[1]`
	managerPermissionSelectedOpt = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[6]/div/span[1]/span[1]/span/ul/li[1]`
	managerRoleDropDown          = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[5]/div/span[1]/span[1]/span/ul`
	managerRoleOpt2              = `/html/body/span/span/span/ul/li[2]`
	managerPermissionDropDown    = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[6]/div/span[1]/span[1]/span`
	managerPermissionOpt2        = `/html/body/span/span/span/ul/li[2]`
	managerSaveBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`
)

func TestLogin(t *testing.T) {
	defer StopDriverOnPanic(t)

	assert.Equal(t, page.Navigate(url("/admin")), nil)

	//wait(5)

	assert.Equal(t, page.Find("#username").Fill("admin"), nil)
	assert.Equal(t, page.Find("#password").Fill("admin"), nil)
	clickS(t, page.FindByButton("login"))

	time.Sleep(time.Second * 3)

	content, err := page.HTML()
	assert.Equal(t, err, nil)
	assert.Equal(t, strings.Contains(content, "main-header"), true)
}

func TestInfoTablePageOperations(t *testing.T) {
	defer StopDriverOnPanic(t)

	navigate(t, "/info/user")

	contain(t, "Users")

	// Buttons Check
	// =============================

	printPart("buttons check")

	click(t, popupBtn)

	display(t, popup)

	wait(1)

	contain(t, "hello world")
	click(t, popupCloseBtn)

	nondisplay(t, popup)

	click(t, ajaxBtn)

	contain(t, "Oh li get")

	display(t, ajaxAlert)

	clickS(t, page.FindByButton("OK"))

	nondisplay(t, ajaxAlert)

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

	fill(t, filterNameField, "Jack")

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

	cssS(t, page.FindByClass("filter-area"), "display", "none")

	// Export Check
	// =============================

	printPart("row export check")

	clickS(t, page.FindByXPath(actionDropDown))
	clickS(t, page.FindByXPath(exportBtn))
	clickS(t, page.FindByClass(`grid-batch-1`))

	// Order Check

	printPart("order check")

	click(t, idOrderBtn)
	click(t, idOrderBtn)

	// Action Button Check
	// =============================

	printPart("action buttons check")

	click(t, rowActionDropDown)
	click(t, previewAction)
	contain(t, "preview content")
	display(t, previewPopup)

	click(t, closePreviewAction)

	nondisplay(t, previewPopup)

	click(t, rowActionDropDown)
	click(t, rowAjaxAction)

	display(t, rowAjaxPopup)

	click(t, closeRowAjaxPopup)

	nondisplay(t, rowAjaxPopup)

	wait(2)
}

func TestNewPageOperations(t *testing.T) {
	defer StopDriverOnPanic(t)

	click(t, newPageBtn, 2)
	value(t, homePageField, "http://google.com")

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
	click(t, idOrderBtn)
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
	click(t, experienceDropDown)
	text(t, twoYearsSelection, "two years")
	text(t, threeYearsSelection, "three years")
	text(t, fourYearsSelection, "four years")
	text(t, fiveYearsSelection, "five years")
	click(t, selectTab)
	attr(t, page.FindByXPath(multiSelectedOpt), "title", "Beer")
	click(t, multiSelectionInput)
	text(t, multiBeerOpt, "Beer")
	text(t, multiJuiceOpt, "Juice")
	text(t, multiWaterOpt, "Water")
	text(t, multiRedBullOpt, "Red bull")
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
	text(t, experienceSelectedOpt, "two years")
}

//  TODO:
//
// [ ] Pagination
// [ ] Join table fields display in table and form

func testDetailPageOperations(t *testing.T) {

}

func testRolePageOperations(t *testing.T) {

}

func testPermissionPageOperations(t *testing.T) {

}

func TestMenuPageOperations(t *testing.T) {
	defer StopDriverOnPanic(t)

	click(t, sideBarManageDropDown)
	click(t, menuPageBtn)

	// ParentIDs Selection Check
	// =============================

	printPart("menu parent ids selection check")

	click(t, menuParentIdDropDown)
	text(t, parentIdRootOpt, "root")
	text(t, parentIdDashboardOpt, "Dashboard")
	text(t, parentIdAdminOpt, "Admin")
	text(t, parentIdUserOpt, "User")
	click(t, parentIdRootOpt)

	// Roles Selection Check
	// =============================

	printPart("menu roles selection check")

	click(t, menuRoleDropDown)
	text(t, menuRoleAdminOpt, "administrator")
	text(t, menuRoleOperatorOpt, "operator")
	click(t, menuRoleAdminOpt)

	click(t, iconPopupBtn)
	display(t, iconPopup)
	click(t, iconBtn)
	click(t, iconPopupBtn)
	nondisplay(t, iconPopup)

	fill(t, menuNameInput, "Test")
	fill(t, menuUriInput, "/info/user")

	// Save Check
	// =============================

	printPart("menu save check")

	click(t, menuInfoSaveBtn)

	// change order check

	//item := page.FindByXPath(testMenuItem)
	//assert.Equal(t, item.MouseToElement(), nil)
	//assert.Equal(t, page.Click(agouti.HoldClick, agouti.LeftButton), nil)
	//assert.Equal(t, item.ScrollFinger(0, -200), nil)
	//assert.Equal(t, page.Click(agouti.HoldClick, agouti.LeftButton), nil)

	// Delete Check
	// =============================

	printPart("menu delete check")

	click(t, testMenuDeleteBtn)
	click(t, testMenuDeleteConfirmBtn)
	click(t, menuOkBtn)
}

func TestManagerPageOperations(t *testing.T) {
	defer StopDriverOnPanic(t)

	click(t, managerPageBtn)
	click(t, managerEditBtn)
	value(t, managerNameField, "admin")
	value(t, managerNickNameField, "admin")
	attr(t, page.FindByXPath(managerRoleSelectedOpt), "title", "administrator")
	attr(t, page.FindByXPath(managerPermissionSelectedOpt), "title", "*")
	fill(t, managerNickNameField, "admin1")
	click(t, managerRoleDropDown)
	text(t, managerRoleOpt2, "operator")
	click(t, managerRoleDropDown)
	click(t, managerPermissionDropDown)
	text(t, managerPermissionOpt2, "dashboard")
	click(t, managerPermissionDropDown)
	click(t, managerSaveBtn)
	contain(t, "admin1")
}
