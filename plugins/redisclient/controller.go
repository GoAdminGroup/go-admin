package redisclient

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/modules/page"
	template2 "github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
	"strings"
)

func Show(ctx *context.Context) {
	page.SetPageContent(ctx, func() types.Panel {

		prefix := Config.PREFIX

		editUrl := Config.PREFIX + "/info/" + prefix + "/edit"
		newUrl := Config.PREFIX + "/info/" + prefix + "/new"
		deleteUrl := Config.PREFIX + "/delete/" + prefix

		menu.GlobalMenu.SetActiveClass(strings.Replace(ctx.Path(), Config.PREFIX, "", 1))

		label := template2.Get(Config.THEME).Label().SetContent("list").GetContent()

		infoList := []map[string]template.HTML{
			{
				"Type":  label,
				"Key":   "userinfo:userid:123",
				"Value": "shipped",
			}, {
				"Type":  label,
				"Key":   "userinfo:userid:124",
				"Value": "shipped",
			}, {
				"Type":  label,
				"Key":   "userinfo:userid:125",
				"Value": "shipped",
			},
		}

		thead := []map[string]string{
			{
				"head":     "Type",
				"sortable": "0",
			}, {
				"head":     "Key",
				"sortable": "0",
			}, {
				"head":     "Value",
				"sortable": "0",
			},
		}

		header := `<div class="pull-right">
		<div class="btn-group pull-right" style="margin-right: 10px">

		<a href="/admin/info" class="btn btn-sm btn-success">

		<i class="fa fa-save"></i>&nbsp;&nbsp;新建
		</a>
		</div>
		</div>
		<span>
		<div class="icheckbox_minimal-blue" aria-checked="false" aria-disabled="false" style="position: relative;"><input type="checkbox" class="grid-select-all" style="position: absolute; opacity: 0;"><ins class="iCheck-helper" style="position: absolute; top: 0%; left: 0%; display: block; width: 100%; height: 100%; margin: 0px; padding: 0px; background: rgb(255, 255, 255); border: 0px; opacity: 0;"></ins></div>
		<div class="btn-group">
		<a class="btn btn-sm btn-default">操作</a>
		<button type="button" class="btn btn-sm btn-default dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
		<span class="caret"></span>
		<span class="sr-only">下拉</span>
		</button>
		<ul class="dropdown-menu" role="menu">
		<li><a href="#" class="grid-batch-0">删除</a></li>
		</ul>
		</div>
		<a class="btn btn-sm btn-primary grid-refresh">
		<i class="fa fa-refresh"></i> 刷新
		</a>
		</span><div class="btn-group" style="margin-right: 10px" data-toggle="buttons">
		<label class="btn btn-sm btn-dropbox 5b9857e354690-filter-btn" title="筛选">
		<input type="checkbox"><i class="fa fa-filter"></i><span class="hidden-xs">&nbsp;&nbsp;筛选</span>
		</label>

		<button type="button" class="btn btn-sm btn-dropbox dropdown-toggle" data-toggle="dropdown" aria-expanded="false">

		<span>&nbsp;选择数据库&nbsp;</span>
		<span class="caret"></span>
		<span class="sr-only">Toggle Dropdown</span>
		</button>
		<ul class="dropdown-menu" role="menu">
		<li><a href="/">DB1</a></li>
		<li role="separator" class="divider">
		</li><li><a href="http://demo.laravel-admin.org/posts?id=1&amp;rate_group=0&amp;_pjax=%23pjax-container">取消</a></li>
		</ul>
		</div>`

		dataTable := template2.Get(Config.THEME).DataTable().SetInfoList(infoList).SetThead(thead).
			SetEditUrl(editUrl).SetNewUrl(newUrl).SetDeleteUrl(deleteUrl)
		table := dataTable.GetContent()

		paginator := template2.Get(Config.THEME).Paginator().GetContent()

		box := template2.Get(Config.THEME).Box().
			SetBody(table).
			SetHeader(template.HTML(header)).
			WithHeadBorder(false).
			SetFooter(paginator).
			GetContent()

		return types.Panel{
			Content:     box,
			Title:       "RedisClient",
			Description: "manage redis",
		}
	})
}
