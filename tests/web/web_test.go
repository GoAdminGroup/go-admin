package web

import (
	"io"
	"log"
	"os"
	"testing"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/tests/tables"
	"github.com/gin-gonic/gin"
)

const (
	// Info Table Page

	newPageBtn              = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[3]/a`
	editPageBtn             = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[8]/div/ul/li[1]/a`
	genderActionDropDown    = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[7]/div/span/span[1]/span/span[2]`
	menOptionActionBtn      = `/html/body/span/span/span[2]/ul/li[2]`
	idOrderBtn              = `//*[@id="sort-id"]`
	rowActionDropDown       = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[8]/div/div/a`
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
	filterNameField         = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[1]/div/div/div[1]/div/div/input`
	filterCreatedStart      = `//*[@id="created_at_start__goadmin"]`
	filterCreatedEnd        = `//*[@id="created_at_end__goadmin"]`
	radio                   = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div[3]/div/div/div[1]/div/div[1]`
	searchBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`
	filterResetBtn          = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[2]/a`
	rowSelector             = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/button`
	rowSelectCityCheckbox   = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[4]/label/div`
	rowSelectAvatarCheckbox = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/div/div[1]/ul/li[1]/ul/li[5]/label/div`
	actionDropDown          = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[1]/div`
	exportBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[1]/span/div/button`
	previewAction           = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[8]/div/ul/li[7]/a`
	closePreviewAction      = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[2]/div/div/div[3]/button`
	previewPopup            = `//*[@id="pjax-container"]/section[2]/div/div/div[4]/div[2]`
	rowAjaxAction           = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[8]/div/ul/li[6]/a`
	rowAjaxPopup            = `/html/body/div[3]`
	closeRowAjaxPopup       = `/html/body/div[3]/div[7]/div/button`
	updateNameTd            = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[3]/a`
	updateNameInput         = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[3]/div/div[2]/div/form/div/div[1]/div[1]/input`
	updateNameSaveBtn       = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[3]/div/div[2]/div/form/div/div[1]/div[2]/button[1]`
	updateGenderBtn         = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[4]/div/div/span[1]`
	detailBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[3]/table/tbody/tr[2]/td[10]/a`

	// Form Page

	saveBtn            = `//*[@id="pjax-container"]/section[2]/div/form/div[2]/div[2]/div[1]/button`
	resetBtn           = `//*[@id="pjax-container"]/section[2]/div/form/div[2]/div[2]/div[2]/button`
	nameField          = `//*[@id="tab-form-0"]/div[1]/div/div/input`
	ageField           = `//*[@id="tab-form-0"]/div[2]/div/div/div/input`
	emailField         = `//*[@id="tab-form-0"]/div[4]/div/div/input`
	birthdayField      = `//*[@id="tab-form-0"]/div[5]/div/div/input`
	passwordField      = `//*[@id="tab-form-0"]/div[6]/div/div/input`
	homePageField      = `//*[@id="tab-form-0"]/div[3]/div/div/input`
	ipField            = `//*[@id="tab-form-0"]/div[7]/div/div/input`
	amountField        = `//*[@id="tab-form-0"]/div[9]/div/div/input`
	appleOptField      = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[1]`
	bananaOptField     = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[2]`
	watermelonOptField = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[3]`
	//pearOptField          = `//*[@id="bootstrap-duallistbox-nonselected-list_fruit[]"]/option[4]`
	genderBoyCheckBox     = `//*[@id="tab-form-1"]/div[5]/div/div/div[1]`
	genderGirlCheckBox    = `//*[@id="tab-form-1"]/div[5]/div/div/div[2]`
	experienceDropDown    = `//*[@id="tab-form-1"]/div[7]/div/span/span[1]/span/span[2]`
	twoYearsSelection     = `/html/body/span/span/span[2]/ul/li[1]`
	threeYearsSelection   = `/html/body/span/span/span[2]/ul/li[2]`
	fourYearsSelection    = `/html/body/span/span/span[2]/ul/li[3]`
	fiveYearsSelection    = `/html/body/span/span/span[2]/ul/li[4]`
	inputTab              = `//*[@id="pjax-container"]/section[2]/div/form/div[1]/div/div/ul/li[1]`
	selectTab             = `//*[@id="pjax-container"]/section[2]/div/form/div[1]/div/div/ul/li[2]`
	multiSelectionInput   = `//*[@id="tab-form-1"]/div[6]/div/span/span[1]/span/ul/li[2]/input`
	multiSelectedOpt      = `//*[@id="tab-form-1"]/div[6]/div/span/span[1]/span/ul/li[1]`
	multiBeerOpt          = `/html/body/span/span/span/ul/li[1]`
	multiJuiceOpt         = `/html/body/span/span/span/ul/li[2]`
	multiWaterOpt         = `/html/body/span/span/span/ul/li[3]`
	multiRedBullOpt       = `/html/body/span/span/span/ul/li[4]`
	continueEditCheckBox  = `//*[@id="pjax-container"]/section[2]/div/form/div[2]/div[2]/label/div`
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
	menuNameInput            = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/div/div[2]/form/div[1]/div/div/div[2]/div/div/input`
	menuUriInput             = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/div/div[2]/form/div[1]/div/div/div[5]/div/div/input`
	menuInfoSaveBtn          = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/div/div[2]/form/div[2]/div[2]/div[1]/button`
	testMenuItem             = `//*[@id="tree-model"]/ol/li[2]`
	testMenuDeleteBtn        = `//*[@id="tree-model"]/ol/li[2]/div/span/a[2]`
	testMenuDeleteConfirmBtn = `/html/body/div[3]/div[7]/div/button`
	menuOkBtn                = `/html/body/div[3]/div[7]/div/button`
	userMenuEditBtn          = `//*[@id="tree-model"]/ol/li[3]/div/span/a[1]`
	headFieldInput           = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[4]/div/div/input`
	menuEditSaveBtn          = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`

	managerPageBtn               = `/html/body/div[1]/aside/section/ul/li[2]/ul/li[1]/a`
	rolesPageBtn                 = `/html/body/div[1]/aside/section/ul/li[2]/ul/li[2]/a`
	permissionPageBtn            = `/html/body/div[1]/aside/section/ul/li[2]/ul/li[3]/a`
	operationLogPageBtn          = `/html/body/div[1]/aside/section/ul/li[2]/ul/li[5]/a`
	navLinkBtn                   = `//*[@id="firstnav"]/div[2]/ul/li[1]/a`
	navCloseBtn                  = `//*[@id="firstnav"]/div[2]/ul/li[1]/i`
	userPageBtn                  = `/html/body/div[1]/aside/section/ul/li[3]/a`
	managerEditBtn               = `//*[@id="pjax-container"]/section[2]/div/div[2]/div[2]/table/tbody/tr[3]/td[8]/a[1]`
	operatorEditBtn              = `//*[@id="pjax-container"]/section[2]/div/div[2]/div[2]/table/tbody/tr[2]/td[8]/a[1]`
	managerNameField             = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[2]/div/div/input`
	managerNickNameField         = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[3]/div/div/input`
	managerRoleSelectedOpt       = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[5]/div/span[1]/span[1]/span/ul/li[1]`
	managerPermissionSelectedOpt = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[6]/div/span[1]/span[1]/span/ul/li[1]`
	managerRoleDropDown          = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[5]/div/span[1]/span[1]/span/ul`
	managerRoleOpt2              = `/html/body/span/span/span/ul/li[2]`
	managerPermissionDropDown    = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[6]/div/span[1]/span[1]/span`
	managerPermissionOpt2        = `/html/body/span/span/span/ul/li[2]`
	managerSaveBtn               = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`
	newPermissionBtn             = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[6]/div/span[2]/a`
	managerUserViewSelectOpt     = `/html/body/span/span/span/ul/li[3]`

	permissionNameInput    = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[1]/div/div/input`
	permissionSlugInput    = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[2]/div/div/input`
	permissionMethodSelect = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[3]/div/span[1]/span[1]/span/ul/li/input`
	permissionGetSelectOpt = `/html/body/span/span/span/ul/li[1]`
	permissionPathInput    = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[1]/div/div/div[4]/div/textarea`
	permissionSaveBtn      = `//*[@id="pjax-container"]/section[2]/div/div/div[2]/form/div[2]/div[2]/div[1]/button`

	userNavMenuBtn = `//*[@id="firstnav"]/div[4]/ul/li[5]/a`
	userSettingBtn = `//*[@id="firstnav"]/div[4]/ul/li[5]/ul/li[5]/a`
	userSignOutBtn = `//*[@id="firstnav"]/div[4]/ul/li[5]/ul/li[6]/a`

	loginPageUserNameInput = `//*[@id="username"]`
	loginPagePasswordInput = `//*[@id="password"]`
)

var (
	debugMode  = false
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
)

const (
	port = ":9033"
)

func init() {
	if os.Args[len(os.Args)-1] == "true" {
		debugMode = true
	}
	if !debugMode {
		optionList = append(optionList, "--headless")
	}
}

func startServer(quit chan struct{}) {

	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
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

	if err := eng.AddConfig(&cfg).
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

func TestWeb(t *testing.T) {
	UserAcceptanceTestSuit(t, func(_ *testing.T, page *Page) {
		defer page.Destroy()
		testLogin(page)
		testInfoTablePageOperations(page)
		testNewPageOperations(page)
		testEditPageOperations(page)
		testDetailPageOperations(page)
		testRolePageOperations(page)
		testPermissionPageOperations(page)
		testMenuPageOperations(page)
		testManagerPageOperations(page)
		testPermission(page)
	}, startServer, debugMode, optionList...)
}

func testLogin(page *Page) {
	page.NavigateTo(url("/login"))

	page.Fill(loginPageUserNameInput, "admin")
	page.Fill(loginPagePasswordInput, "admin")
	page.ClickS(page.FindByButton("Login"))

	wait(3)

	page.Contain("main-header")
}

func testInfoTablePageOperations(page *Page) {

	// Nav link Check
	// =============================

	//printPart("nav link check")
	//page.Click(sideBarManageDropDown)
	//page.Click(managerPageBtn)
	//page.Click(rolesPageBtn)
	//page.Click(permissionPageBtn)
	//page.Click(menuPageBtn)
	//page.Click(operationLogPageBtn)
	//page.Click(navLinkBtn)
	//page.Click(navCloseBtn)
	//page.Click(navLinkBtn)
	//page.Click(navCloseBtn)
	//page.Click(navLinkBtn)
	//page.Click(navCloseBtn)
	//page.Click(navLinkBtn)
	//page.Click(navCloseBtn)
	//page.Click(navLinkBtn)
	//page.Click(navCloseBtn)

	page.NavigateTo(url("/info/user"))

	page.Contain("Users")

	// Buttons Check
	// =============================

	printPart("buttons check")

	page.Click(popupBtn)

	page.Display(popup)

	wait(1)

	page.Contain("hello world")
	page.Click(popupCloseBtn)

	page.Nondisplay(popup)

	page.Click(ajaxBtn)

	page.Contain("Oh li get")

	page.Display(ajaxAlert)

	page.ClickS(page.FindByButton("OK"))

	page.Nondisplay(ajaxAlert)

	// Update Check
	// =============================

	printPart("update check")
	page.Click(updateNameTd)
	page.Fill(updateNameInput, "DukeDukeDuke")
	page.Click(updateNameSaveBtn)
	page.Click(updateGenderBtn)
	page.Contain("DukeDukeDuke")

	// Filter Area Check
	// =============================

	printPart("filter area check")

	page.Click(selectionDropDown)

	page.Text(selectionLi1, "men")
	page.Text(selectionLi2, "women")

	page.Click(selectionLi2)

	page.Attr(page.FindByXPath(selectionRes), "title", "women")

	page.Fill(multiSelectInput, " ")

	wait(1)

	page.Text(multiSelectLi1, "water")
	page.Text(multiSelectLi2, "juice")
	page.Text(multiSelectLi3, "red bull")

	page.Click(multiSelectLi3)

	page.Attr(page.FindByXPath(multiSelectRes), "title", "red bull")

	page.Click(radio)

	page.Fill(filterNameField, "Jack")

	//page.Fill(filterCreatedStart, "2020-03-08 15:24:00")
	//page.Click(filterCreatedEnd)

	page.Click(searchBtn, 2)

	page.Click(filterResetBtn, 2)

	// Row Selector Check
	// =============================

	printPart("row selector check")

	page.Click(rowSelector)
	page.Click(rowSelectCityCheckbox)
	page.Click(rowSelectAvatarCheckbox)

	page.ClickS(page.FindByButton("Submit"), 2)

	page.NoContain("guangzhou")

	page.ClickS(page.FindByID("filter-btn"))

	page.CssS(page.FindByClass("filter-area"), "display", "none")

	// Export Check
	// =============================

	printPart("row export check")

	page.ClickS(page.FindByXPath(actionDropDown))
	page.ClickS(page.FindByXPath(exportBtn))
	page.ClickS(page.FindByClass(`grid-batch-1`))

	// Order Check

	printPart("order check")

	page.Click(idOrderBtn)
	page.Click(idOrderBtn)

	// Action Button Check
	// =============================

	printPart("action buttons check")

	page.Click(rowActionDropDown)
	page.Click(previewAction)
	page.Contain("preview content")
	page.Display(previewPopup)

	page.Click(closePreviewAction)

	page.Nondisplay(previewPopup)

	page.Click(rowActionDropDown)
	page.Click(rowAjaxAction)

	page.Display(rowAjaxPopup)

	page.Click(closeRowAjaxPopup)

	page.Nondisplay(rowAjaxPopup)

	wait(2)
}

func testNewPageOperations(page *Page) {

	page.Click(newPageBtn, 2)
	page.Value(homePageField, "http://google.com")

	// Selections Form Item Check
	// =============================

	printPart("selections form items check")

	checkSelectionsInForm(page)

	// Create Error Check
	// =============================

	printPart("create error check")

	page.Click(saveBtn)
	page.Contain("error")

	// Reset Check
	// =============================

	printPart("reset error check")

	fillNewForm(page, "jane", "girl")
	page.Click(resetBtn)

	// Continue Creating Check
	// =============================

	printPart("continue creating check")

	page.Click(inputTab)
	page.Text(ipField, "")
	page.Click(continueEditCheckBox)
	fillNewForm(page, "jane", "girl")
	page.Click(saveBtn)

	// Creating Check
	// =============================

	printPart("creating check")

	fillNewForm(page, "harry", "boy")
	page.Click(saveBtn, 2)

	page.NoContain("harry")
	page.Click(genderActionDropDown)
	page.Click(menOptionActionBtn, 2)
	page.Click(idOrderBtn)
	page.Contain("harry")
}

func fillNewForm(page *Page, name, gender string) {

	page.Fill(nameField, name)
	page.Fill(ageField, "15")
	page.Fill(passwordField, "12345678")
	page.Fill(ipField, "127.0.0.1")
	page.Fill(amountField, "15")
	page.Click(selectTab)
	page.Click(appleOptField)
	if gender == "girl" {
		page.Click(genderGirlCheckBox)
	} else {
		page.Click(genderBoyCheckBox)
	}
	page.Click(experienceDropDown)
	page.Click(twoYearsSelection)
}

func checkSelectionsInForm(page *Page) {
	page.Click(selectTab)
	page.Text(appleOptField, "Apple")
	page.Text(bananaOptField, "Banana")
	page.Text(watermelonOptField, "Watermelon")
	//page.Text(pearOptField, "")
	page.Click(experienceDropDown)
	page.Text(twoYearsSelection, "two years")
	page.Text(threeYearsSelection, "three years")
	page.Text(fourYearsSelection, "four years")
	page.Text(fiveYearsSelection, "five years")
	page.Click(selectTab)
	page.Attr(page.FindByXPath(multiSelectedOpt), "title", "Beer")
	page.Click(multiSelectionInput)
	page.Text(multiBeerOpt, "Beer")
	page.Text(multiJuiceOpt, "Juice")
	page.Text(multiWaterOpt, "Water")
	page.Text(multiRedBullOpt, "Red bull")
	page.Click(inputTab)
}

func testEditPageOperations(page *Page) {
	page.Click(rowActionDropDown)
	page.Click(editPageBtn, 2)

	// Form Field Value Check
	// =============================

	printPart("edit form values check")

	page.Value(nameField, "harry")
	page.Value(homePageField, "http://google.com")
	page.Value(ageField, "15")
	page.Value(emailField, "xxxx@xxx.com")
	page.Value(birthdayField, "2010-09-05 00:00:00")
	page.Value(passwordField, "12345678")
	page.Value(ipField, "127.0.0.1")
	page.Value(amountField, "15")

	page.Click(selectTab)

	page.Text(boxSelectedOpt, "Pear")
	page.Attr(page.FindByXPath(multiSelectedOpt), "title", "Beer")
	page.Text(experienceSelectedOpt, "two years")
}

//  TODO:
//
// [ ] Pagination
// [ ] Join table fields display in table and form

func testDetailPageOperations(_ *Page) {

}

func testRolePageOperations(_ *Page) {

}

func testPermissionPageOperations(_ *Page) {

}

func testMenuPageOperations(page *Page) {

	page.Click(sideBarManageDropDown)
	page.Click(menuPageBtn)

	// ParentIDs Selection Check
	// =============================

	printPart("menu parent ids selection check")

	page.Click(menuParentIdDropDown)
	page.Text(parentIdRootOpt, "ROOT")
	page.Text(parentIdDashboardOpt, "  ┝ Dashboard")
	page.Text(parentIdAdminOpt, "  ┝ Admin")
	page.Text(parentIdUserOpt, "         ┝ Users")
	page.Click(parentIdRootOpt)

	// Roles Selection Check
	// =============================

	printPart("menu roles selection check")

	page.Click(menuRoleDropDown)
	page.Text(menuRoleAdminOpt, "administrator")
	page.Text(menuRoleOperatorOpt, "operator")
	page.Click(menuRoleAdminOpt)

	page.Click(iconPopupBtn)
	page.Display(iconPopup)
	page.Click(iconBtn)
	page.Click(iconPopupBtn)
	page.Nondisplay(iconPopup)

	page.Fill(menuNameInput, "Test")
	page.Fill(menuUriInput, "/info/user")

	// Save Check
	// =============================

	printPart("menu save check")

	page.Click(menuInfoSaveBtn)

	// change order check

	//item := page.FindByXPath(testMenuItem)
	//assert.Equal(t, item.MouseToElement(), nil)
	//assert.Equal(t, page.Click(agouti.HoldClick, agouti.LeftButton), nil)
	//assert.Equal(t, item.ScrollFinger(0, -200), nil)
	//assert.Equal(t, page.Click(agouti.HoldClick, agouti.LeftButton), nil)

	// Delete Check
	// =============================

	printPart("menu delete check")

	page.Click(testMenuDeleteBtn)
	page.Click(testMenuDeleteConfirmBtn)
	page.Click(menuOkBtn)

	page.Click(userMenuEditBtn)
	page.Fill(headFieldInput, "example")
	page.Click(menuEditSaveBtn)
}

func testManagerPageOperations(page *Page) {

	page.Click(managerPageBtn)
	page.Click(managerEditBtn)
	page.Value(managerNameField, "admin")
	page.Value(managerNickNameField, "admin")
	page.Attr(page.FindByXPath(managerRoleSelectedOpt), "title", "administrator")
	page.Attr(page.FindByXPath(managerPermissionSelectedOpt), "title", "*")
	page.Fill(managerNickNameField, "admin1")
	page.Click(managerRoleDropDown)
	page.Text(managerRoleOpt2, "operator")
	page.Click(managerRoleDropDown)
	page.Click(managerPermissionDropDown)
	page.Text(managerPermissionOpt2, "dashboard")
	page.Click(managerPermissionDropDown)
	page.Click(managerSaveBtn)
	page.Contain("admin1")

	page.Click(operatorEditBtn)
	page.Click(newPermissionBtn)
	page.Fill(permissionNameInput, "user_view")
	page.Fill(permissionSlugInput, "user_view")
	page.Click(permissionMethodSelect)
	page.Click(permissionGetSelectOpt)
	page.Fill(permissionPathInput, `/info/user
/info/user/detail`)
	page.Click(permissionSaveBtn)
	page.Click(managerPermissionDropDown)
	page.Click(managerUserViewSelectOpt)
	page.Click(managerSaveBtn)

	page.Click(userNavMenuBtn)
	page.Click(userSettingBtn)
	page.Fill(managerNickNameField, "admin")
	page.Click(managerSaveBtn)
	page.Contain("admin")
	page.Click(userNavMenuBtn)
	page.Click(userSignOutBtn)
}

func testPermission(page *Page) {
	page.Fill(loginPageUserNameInput, "operator")
	page.Fill(loginPagePasswordInput, "admin")
	page.ClickS(page.FindByButton("Login"))
	page.NavigateTo(url("/info/user"))
	page.NoContain("New")
	page.Click(detailBtn)
	page.NoContain("Edit")
}
