package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/system"

	"github.com/GoAdminGroup/go-admin/modules/logger"

	"github.com/GoAdminGroup/go-admin/modules/config"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/remote_server"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/GoAdminGroup/html"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Plugins(ctx *context.Context) {
	list := plugins.Get()
	size := types.Size(6, 3, 2)
	rows := template.HTML("")
	if h.config.IsNotProductionEnvironment() {
		getMoreCover := config.Url("/assets/dist/img/plugin_more.png")
		list = list.Add(plugins.NewBasePluginWithInfoAndIndexURL(plugins.Info{
			Title:     "get more plugins",
			Name:      "",
			MiniCover: getMoreCover,
			Cover:     getMoreCover,
		}, config.Url("/plugins/store"), true))
	}
	for i := 0; i < len(list); i += 6 {
		box1 := aBox().
			SetBody(h.pluginBox(GetPluginBoxParamFromPlug(list[i]))).
			GetContent()
		content := aCol().SetSize(size).SetContent(box1).GetContent()
		offset := len(list) - i
		if offset > 6 {
			offset = 6
		}
		for j := i + 1; j < offset; j++ {
			box2 := aBox().
				SetBody(h.pluginBox(GetPluginBoxParamFromPlug(list[j]))).
				GetContent()
			content += aCol().SetSize(size).SetContent(box2).GetContent()
		}
		rows += aRow().SetContent(content).GetContent()
	}
	h.HTML(ctx, auth.Auth(ctx), types.Panel{
		Content:     rows,
		CSS:         pluginsPageCSS,
		Description: language.GetFromHtml("plugins"),
		Title:       language.GetFromHtml("plugins"),
	})
}

func (h *Handler) PluginStore(ctx *context.Context) {
	var (
		size       = types.Size(12, 6, 4)
		list, page = plugins.GetAll(
			remote_server.GetOnlineReq{
				Page:       ctx.Query("page"),
				Free:       ctx.Query("free"),
				PageSize:   ctx.Query("page_size"),
				Filter:     ctx.Query("filter"),
				Order:      ctx.Query("order"),
				Lang:       h.config.Language,
				Version:    system.Version(),
				CategoryId: ctx.Query("category_id"),
			}, ctx.Cookie(remote_server.TokenKey))
		rows = template.HTML(page.HTML)
	)

	if ctx.Query("page") == "" && len(list) == 0 {
		h.HTML(ctx, auth.Auth(ctx), types.Panel{
			Content: pluginStore404(),
			CSS: template.CSS(`.plugin-store-404-content {
    margin: auto;
    width: 80%;
    text-align: center;
    color: #9e9e9e;
    font-size: 17px;
    height: 250px;
    line-height: 250px;
}`),
			Description: language.GetFromHtml("plugin store"),
			Title:       language.GetFromHtml("plugin store"),
		})
		return
	}

	for i := 0; i < len(list); i += 3 {
		box1 := aBox().
			SetBody(h.pluginStoreBox(GetPluginBoxParamFromPlug(list[i]))).
			GetContent()
		col1 := aCol().SetSize(size).SetContent(box1).GetContent()
		box2, col2, box3, col3 := template.HTML(""), template.HTML(""), template.HTML(""), template.HTML("")
		if i+1 < len(list) {
			box2 = aBox().
				SetBody(h.pluginStoreBox(GetPluginBoxParamFromPlug(list[i+1]))).
				GetContent()
			col2 = aCol().SetSize(size).SetContent(box2).GetContent()
			if i+2 < len(list) {
				box3 = aBox().
					SetBody(h.pluginStoreBox(GetPluginBoxParamFromPlug(list[i+2]))).
					GetContent()
				col3 = aCol().SetSize(size).SetContent(box3).GetContent()
			}
		}
		rows += aRow().SetContent(col1 + col2 + col3).GetContent()
	}

	detailPopupModal := template2.Default().Popup().SetID("detail-popup-modal").
		SetTitle(plugWordHTML("plugin detail")).
		SetBody(pluginsPageDetailPopupBody()).
		SetWidth("730px").
		SetHeight("400px").
		SetFooter("1").
		GetContent()

	buyPopupModal := template2.Default().Popup().SetID("buy-popup-modal").
		SetTitle(plugWordHTML("plugin detail")).
		SetWidth("730px").
		SetHeight("400px").
		SetFooter("1").
		GetContent()

	loginPopupModal := template2.Default().Popup().SetID("login-popup-modal").
		SetTitle(plugWordHTML("login to goadmin member system")).
		SetBody(aForm().SetContent(types.FormFields{
			{Field: "name", Head: plugWord("account"), FormType: form.Text, Editable: true},
			{Field: "password", Head: plugWord("password"), FormType: form.Password, Editable: true,
				HelpMsg: template.HTML(fmt.Sprintf(plugWord("no account? click %s here %s to register."),
					"<a target='_blank' href='http://www.go-admin.cn/register'>", "</a>"))},
		}).GetContent()).
		SetWidth("540px").
		SetHeight("250px").
		SetFooterHTML(template.HTML(`<button type="button" class="btn btn-primary" onclick="login()">` +
			plugWord("login") + `</button>`)).
		GetContent()

	h.HTML(ctx, auth.Auth(ctx), types.Panel{
		Content:     rows + detailPopupModal + buyPopupModal + loginPopupModal,
		CSS:         pluginsStorePageCSS + template.CSS(page.CSS),
		JS:          template.JS(page.JS) + GetPluginsPageJS(PluginsPageJSData{Prefix: h.config.Prefix()}),
		Description: language.GetFromHtml("plugin store"),
		Title:       language.GetFromHtml("plugin store"),
	})
}

func (h *Handler) PluginDetail(ctx *context.Context) {

	name := ctx.Query("name")

	plug, exist := plugins.FindByNameAll(name)
	if !exist {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	info := plug.GetInfo()

	if info.MiniCover == "" {
		info.MiniCover = config.Url("/assets/dist/img/plugin_default.png")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"mini_cover":      info.MiniCover,
			"title":           language.GetWithScope(info.Title, name),
			"author":          fmt.Sprintf(plugWord("provided by %s"), language.GetWithScope(info.Author, name)),
			"introduction":    language.GetWithScope(info.Description, name),
			"website":         language.GetWithScope(info.Website, name),
			"version":         language.GetWithScope(info.Version, name),
			"created_at":      language.GetWithScope(info.CreateDate.Format("2006-01-02"), name),
			"updated_at":      language.GetWithScope(info.UpdateDate.Format("2006-01-02"), name),
			"downloaded":      info.Downloaded,
			"download_reboot": plugins.Exist(plug),
			"skip":            info.SkipInstallation,
			"uuid":            info.Uuid,
			"upgrade":         info.CanUpdate,
			"install":         plug.IsInstalled(),
			"free":            info.IsFree(),
		},
	})
}

type PluginBoxParam struct {
	Info           plugins.Info
	Install        bool
	Upgrade        bool
	Skip           bool
	DownloadReboot bool
	Name           string
	IndexURL       string
}

func GetPluginBoxParamFromPlug(plug plugins.Plugin) PluginBoxParam {
	return PluginBoxParam{
		Info:           plug.GetInfo(),
		Install:        plug.IsInstalled(),
		Upgrade:        plug.GetInfo().CanUpdate,
		Skip:           plug.GetInfo().SkipInstallation,
		DownloadReboot: plugins.Exist(plug),
		Name:           plug.Name(),
		IndexURL:       plug.GetIndexURL(),
	}
}

func (h *Handler) pluginStoreBox(param PluginBoxParam) template.HTML {
	cover := template2.HTML(param.Info.MiniCover)
	if cover == template2.HTML("") {
		cover = template2.HTML(config.Url("/assets/dist/img/plugin_default.png"))
	}
	col1 := html.DivEl().SetClass("plugin-store-item-img").
		SetContent(aImage().
			SetSrc(cover).
			SetHeight("110px").
			SetWidth("110px").
			GetContent()).
		Get()
	footer := html.ButtonEl().SetClass(pluginBtnClass("plugin-info")...).
		SetAttr("onclick", `pluginDetail('`+param.Name+`','`+param.Info.Uuid+`')`).
		SetContent(plugWordHTML("info")).
		Get()
	if param.Install {
		if param.Upgrade {
			footer += html.ButtonEl().SetClass(pluginBtnClass("installation")...).
				SetAttr("onclick", `pluginDownload('`+param.Name+`', this)`).
				SetContent(plugWordHTML("upgrade")).
				Get()
		}
	} else {
		if param.Info.Downloaded {
			if param.DownloadReboot {
				if !param.Skip && !param.Install {
					footer += html.AEl().SetAttr("href", h.config.Url(`/info/plugin_`+param.Name+`/new`)).
						SetContent(
							html.ButtonEl().SetClass(pluginBtnClass("installation")...).
								SetAttr("onclick", `pluginInstall('`+param.Name+`')`).
								SetContent(plugWordHTML("install")).
								Get(),
						).Get()
				}
			} else {
				footer += html.ButtonEl().SetClass(pluginBtnClass("installation")...).
					SetAttr("onclick", `pluginRebootInstall()`).
					SetContent(plugWordHTML("install")).
					Get()
			}
		} else {
			if param.Info.IsFree() || param.Info.HasBought {
				footer += html.ButtonEl().SetClass(pluginBtnClass("installation")...).
					SetAttr("onclick", `pluginDownload('`+param.Name+`', this)`).
					SetContent(plugWordHTML("download")).
					Get()
			} else {
				footer += html.ButtonEl().SetClass(pluginBtnClass("installation")...).
					SetAttr("onclick", `pluginBuy('`+param.Name+`', '`+param.Info.Uuid+`')`).
					SetContent(plugWordHTML("buy")).
					Get()
			}
		}
	}

	col2 := html.DivEl().SetClass("plugin-item-content").SetContent(
		html.DivEl().SetClass("plugin-item-content-title").
			SetContent(language.GetFromHtml(template.HTML(param.Info.Title), param.Name)).
			Get() +
			html.DivEl().SetClass("plugin-item-content-description").
				SetContent(language.GetFromHtml(template.HTML(param.Info.Description), param.Name)).
				Get() +
			footer,
	).Get()

	return html.Div(col1+col2, html.M{"clear": "both"})
}

func (h *Handler) pluginBox(param PluginBoxParam) template.HTML {
	cover := template2.HTML(param.Info.MiniCover)
	if cover == template2.HTML("") {
		cover = "/admin/assets/dist/img/plugin_default.png"
	}

	jump := param.IndexURL
	label := template.HTML("")
	if !param.Install {
		jump = h.config.Url("/info/plugin_" + param.Name + "/new")
		label = html.SpanEl().SetClass("plugin-item-label").SetContent(language.GetFromHtml("uninstalled")).Get()
	}
	col1 := html.AEl().SetContent(html.DivEl().SetClass("plugin-item-img").
		SetContent(aImage().
			SetSrc(cover).
			GetContent()+
			html.PEl().SetContent(language.GetFromHtml(template.HTML(param.Info.Title), param.Name)).
				SetClass("plugin-item-title").Get()).
		Get()+label).SetAttr("href", jump).Get()
	return col1
}

func (h *Handler) PluginDownload(ctx *context.Context) {

	if !h.config.Debug {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 400,
			"msg":  plugWord("change to debug mode first"),
		})
		return
	}

	name := ctx.FormValue("name")

	if name == "" {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 400,
			"msg":  plugWord("download fail, wrong name"),
		})
		return
	}

	plug, exist := plugins.FindByNameAll(name)

	if !exist {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 400,
			"msg":  plugWord("download fail, plugin not exist"),
		})
		return
	}

	if !plug.GetInfo().IsFree() && !plug.GetInfo().HasBought {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 400,
			"msg":  plugWord("download fail, plugin has not been bought"),
		})
		return
	}

	downloadURL := plug.GetInfo().Url
	extraDownloadURL := plug.GetInfo().ExtraDownloadUrl

	if !plug.GetInfo().IsFree() {
		var err error
		downloadURL, extraDownloadURL, err = remote_server.GetDownloadURL(plug.GetInfo().Uuid, ctx.Cookie(remote_server.TokenKey))
		if err != nil {
			logger.Error("download plugins error", err)
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code": 500,
				"msg":  plugWord("download fail"),
			})
			return
		}
	}

	tempFile := "./temp-" + utils.Uuid(10) + ".zip"

	err := utils.DownloadTo(downloadURL, tempFile)

	if err != nil {
		logger.Error("download plugins error", map[string]interface{}{
			"error":       err,
			"downloadURL": downloadURL,
		})
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500,
			"msg":  plugWord("download fail"),
		})
		return
	}

	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500,
			"msg":  plugWord("golang develop environment does not exist"),
		})
		return
	}

	gomodule := os.Getenv("GO111MODULE")
	base := filepath.Dir(plug.GetInfo().ModulePath)
	installPath := ""

	if gomodule == "off" {
		installPath = filepath.ToSlash(gopath + "/src/" + base)
	} else {
		installPath = filepath.ToSlash(gopath + "/pkg/mod/" + base)
	}

	err = utils.UnzipDir(tempFile, installPath)

	if err != nil {
		logger.Error("download plugins, unzip error", map[string]interface{}{
			"error":       err,
			"installPath": installPath,
		})
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500,
			"msg":  plugWord("download fail"),
		})
		return
	}

	_ = os.Remove(tempFile)

	if len(downloadURL) > 18 && downloadURL[:18] == "https://github.com" {
		name := filepath.Base(plug.GetInfo().ModulePath)
		version := strings.ReplaceAll(plug.GetInfo().Version, "v", "")
		rawPath := installPath + "/" + name
		nowPath := rawPath + "-" + version
		if gomodule == "off" {
			err = os.Rename(nowPath, rawPath)
		} else {
			err = os.Rename(nowPath, rawPath+"@"+plug.GetInfo().Version)
		}
		if err != nil {
			logger.Error("download plugins, rename error", map[string]interface{}{
				"error":   err,
				"nowPath": nowPath,
				"rawPath": rawPath,
			})
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code": 500,
				"msg":  plugWord("download fail"),
			})
			return
		}
	} else if gomodule != "off" {
		rawPath := installPath + "/" + name
		err = os.Rename(rawPath, rawPath+"@"+plug.GetInfo().Version)
		if err != nil {
			logger.Error("download plugins, rename error", map[string]interface{}{
				"error":   err,
				"rawPath": rawPath,
			})
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code": 500,
				"msg":  plugWord("download fail"),
			})
			return
		}
	}

	if h.config.BootstrapFilePath != "" && utils.FileExist(h.config.BootstrapFilePath) {
		content, err := ioutil.ReadFile(h.config.BootstrapFilePath)
		if err != nil {
			logger.Error("read bootstrap file error: ", err)
		} else {
			err = ioutil.WriteFile(h.config.BootstrapFilePath, []byte(string(content)+`
import _ "`+plug.GetInfo().ModulePath+`"`), 0644)
			if err != nil {
				logger.Error("write bootstrap file error: ", err)
			}
		}
	}

	if h.config.GoModFilePath != "" && utils.FileExist(h.config.GoModFilePath) &&
		plug.GetInfo().CanUpdate && plug.GetInfo().OldVersion != "" {
		content, _ := ioutil.ReadFile(h.config.BootstrapFilePath)
		src := plug.GetInfo().ModulePath + " " + plug.GetInfo().OldVersion
		dist := plug.GetInfo().ModulePath + " " + plug.GetInfo().Version
		content = bytes.ReplaceAll(content, []byte(src), []byte(dist))
		_ = ioutil.WriteFile(h.config.BootstrapFilePath, content, 0644)
	}

	// TODO: 实现运行环境与编译环境隔离

	if plug.GetInfo().ExtraDownloadUrl != "" {
		err = utils.DownloadTo(extraDownloadURL, "./"+plug.Name()+"_extra_"+
			fmt.Sprintf("%d", time.Now().Unix())+".zip")
		if err != nil {
			logger.Error("failed to download "+plug.Name()+" extra data: ", err)
		}
	}

	plug.(*plugins.BasePlugin).Info.Downloaded = true
	plug.(*plugins.BasePlugin).Info.CanUpdate = false

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  plugWord("download success, restart to install"),
	})
}

func (h *Handler) ServerLogin(ctx *context.Context) {
	param := guard.GetServerLoginParam(ctx)
	res := remote_server.Login(param.Account, param.Password)
	if res.Code == 0 && res.Data.Token != "" {
		ctx.SetCookie(&http.Cookie{
			Name:     remote_server.TokenKey,
			Value:    res.Data.Token,
			Expires:  time.Now().Add(time.Second * time.Duration(res.Data.Expire/1000)),
			HttpOnly: true,
			Path:     "/",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": res.Code,
		"data": res.Data,
		"msg":  res.Msg,
	})
}

func pluginBtnClass(class ...string) []string {
	return append([]string{"btn", "btn-primary"}, class...)
}

func plugWord(word string) string {
	return language.GetWithScope(word, "plugin")
}

func plugWordHTML(word template.HTML) template.HTML {
	return language.GetFromHtml(word, "plugin")
}
