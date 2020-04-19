package controller

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"os"
	"runtime"
)

func (h *Handler) SystemInfo(ctx *context.Context) {

	size := types.Size(6, 6, 6)

	sys := system.GetSysStatus()

	box1 := aBox().
		WithHeadBorder().
		SetHeader("<b>" + lg("application") + "</b>").
		SetBody(stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("app_name")},
				"value": types.InfoItem{Content: "GoAdmin"},
			}, {
				"key":   types.InfoItem{Content: lg("go_admin_version")},
				"value": types.InfoItem{Content: template.HTML(system.Version())},
			}, {
				"key":   types.InfoItem{Content: lg("theme_name")},
				"value": types.InfoItem{Content: template.HTML(aTemplate().Name())},
			}, {
				"key":   types.InfoItem{Content: lg("theme_version")},
				"value": types.InfoItem{Content: template.HTML(aTemplate().GetVersion())},
			},
		})).
		GetContent()

	box2 := aBox().
		WithHeadBorder().
		SetHeader("<b>" + lg("system") + "</b>").
		SetBody(stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("cpu_logical_core")},
				"value": types.InfoItem{Content: itos(sys.CpuLogicalCore)},
			}, {
				"key":   types.InfoItem{Content: lg("cpu_core")},
				"value": types.InfoItem{Content: itos(sys.CpuCore)},
			},
		}) + `<div><hr></div>` + stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("os_platform")},
				"value": types.InfoItem{Content: template.HTML(sys.OSPlatform)},
			}, {
				"key":   types.InfoItem{Content: lg("os_family")},
				"value": types.InfoItem{Content: template.HTML(sys.OSFamily)},
			}, {
				"key":   types.InfoItem{Content: lg("os_version")},
				"value": types.InfoItem{Content: template.HTML(sys.OSVersion)},
			},
		}) + `<div><hr></div>` + stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("load1")},
				"value": types.InfoItem{Content: template.HTML(fmt.Sprintf("%.2f", sys.Load1))},
			}, {
				"key":   types.InfoItem{Content: lg("load5")},
				"value": types.InfoItem{Content: template.HTML(fmt.Sprintf("%.2f", sys.Load5))},
			}, {
				"key":   types.InfoItem{Content: lg("load15")},
				"value": types.InfoItem{Content: template.HTML(fmt.Sprintf("%.2f", sys.Load15))},
			},
		}) + `<div><hr></div>` + stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("mem_total")},
				"value": types.InfoItem{Content: template.HTML(sys.MemTotal)},
			}, {
				"key":   types.InfoItem{Content: lg("mem_available")},
				"value": types.InfoItem{Content: template.HTML(sys.MemAvailable)},
			}, {
				"key":   types.InfoItem{Content: lg("mem_used")},
				"value": types.InfoItem{Content: template.HTML(sys.MemUsed)},
			},
		})).
		GetContent()

	app := system.GetAppStatus()

	box3 := aBox().
		WithHeadBorder().
		SetHeader("<b>" + lg("application run") + "</b>").
		SetBody(stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("current_heap_usage")},
				"value": types.InfoItem{Content: template.HTML(app.HeapAlloc)},
			},
			{
				"key":   types.InfoItem{Content: lg("heap_memory_obtained")},
				"value": types.InfoItem{Content: template.HTML(app.HeapSys)},
			},
			{
				"key":   types.InfoItem{Content: lg("heap_memory_idle")},
				"value": types.InfoItem{Content: template.HTML(app.HeapIdle)},
			},
			{
				"key":   types.InfoItem{Content: lg("heap_memory_in_use")},
				"value": types.InfoItem{Content: template.HTML(app.HeapInuse)},
			},
			{
				"key":   types.InfoItem{Content: lg("heap_memory_released")},
				"value": types.InfoItem{Content: template.HTML(app.HeapReleased)},
			},
			{
				"key":   types.InfoItem{Content: lg("heap_objects")},
				"value": types.InfoItem{Content: itos(app.HeapObjects)},
			},
		})).
		GetContent()

	col1 := aCol().SetSize(size).SetContent(box1 + box2 + box3).GetContent()

	box4 := aBox().
		WithHeadBorder().
		SetHeader("<b>" + lg("application run") + "</b>").
		SetBody(stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("golang_version")},
				"value": types.InfoItem{Content: template.HTML(runtime.Version())},
			}, {
				"key":   types.InfoItem{Content: lg("process_id")},
				"value": types.InfoItem{Content: itos(os.Getpid())},
			}, {
				"key":   types.InfoItem{Content: lg("server_uptime")},
				"value": types.InfoItem{Content: template.HTML(app.Uptime)},
			}, {
				"key":   types.InfoItem{Content: lg("current_goroutine")},
				"value": types.InfoItem{Content: itos(app.NumGoroutine)},
			},
		}) + `<div><hr></div>` + stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("current_memory_usage")},
				"value": types.InfoItem{Content: template.HTML(app.MemAllocated)},
			}, {
				"key":   types.InfoItem{Content: lg("total_memory_allocated")},
				"value": types.InfoItem{Content: template.HTML(app.MemTotal)},
			}, {
				"key":   types.InfoItem{Content: lg("memory_obtained")},
				"value": types.InfoItem{Content: itos(app.MemSys)},
			}, {
				"key":   types.InfoItem{Content: lg("pointer_lookup_times")},
				"value": types.InfoItem{Content: itos(app.Lookups)},
			}, {
				"key":   types.InfoItem{Content: lg("memory_allocate_times")},
				"value": types.InfoItem{Content: itos(app.MemMallocs)},
			}, {
				"key":   types.InfoItem{Content: lg("memory_free_times")},
				"value": types.InfoItem{Content: itos(app.MemFrees)},
			},
		}) + `<div><hr></div>` + stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("bootstrap_stack_usage")},
				"value": types.InfoItem{Content: template.HTML(app.StackInuse)},
			}, {
				"key":   types.InfoItem{Content: lg("stack_memory_obtained")},
				"value": types.InfoItem{Content: template.HTML(app.StackSys)},
			}, {
				"key":   types.InfoItem{Content: lg("mspan_structures_usage")},
				"value": types.InfoItem{Content: template.HTML(app.MSpanInuse)},
			}, {
				"key":   types.InfoItem{Content: lg("mspan_structures_obtained")},
				"value": types.InfoItem{Content: template.HTML(app.HeapSys)},
			}, {
				"key":   types.InfoItem{Content: lg("mcache_structures_usage")},
				"value": types.InfoItem{Content: template.HTML(app.MCacheInuse)},
			}, {
				"key":   types.InfoItem{Content: lg("mcache_structures_obtained")},
				"value": types.InfoItem{Content: template.HTML(app.MCacheSys)},
			}, {
				"key":   types.InfoItem{Content: lg("profiling_bucket_hash_table_obtained")},
				"value": types.InfoItem{Content: template.HTML(app.BuckHashSys)},
			}, {
				"key":   types.InfoItem{Content: lg("gc_metadata_obtained")},
				"value": types.InfoItem{Content: template.HTML(app.GCSys)},
			}, {
				"key":   types.InfoItem{Content: lg("other_system_allocation_obtained")},
				"value": types.InfoItem{Content: template.HTML(app.OtherSys)},
			},
		}) + `<div><hr></div>` + stripedTable([]map[string]types.InfoItem{
			{
				"key":   types.InfoItem{Content: lg("next_gc_recycle")},
				"value": types.InfoItem{Content: template.HTML(app.NextGC)},
			}, {
				"key":   types.InfoItem{Content: lg("last_gc_time")},
				"value": types.InfoItem{Content: template.HTML(app.LastGC)},
			}, {
				"key":   types.InfoItem{Content: lg("total_gc_pause")},
				"value": types.InfoItem{Content: template.HTML(app.PauseTotalNs)},
			}, {
				"key":   types.InfoItem{Content: lg("last_gc_pause")},
				"value": types.InfoItem{Content: template.HTML(app.PauseNs)},
			}, {
				"key":   types.InfoItem{Content: lg("gc_times")},
				"value": types.InfoItem{Content: itos(app.NumGC)},
			},
		})).
		GetContent()

	col2 := aCol().SetSize(size).SetContent(box4).GetContent()

	row := aRow().SetContent(col1 + col2).GetContent()

	h.HTML(ctx, auth.Auth(ctx), types.Panel{
		Content:     row,
		Description: language.GetFromHtml("system info", "system"),
		Title:       language.GetFromHtml("system info", "system"),
	})
}

func stripedTable(list []map[string]types.InfoItem) template.HTML {
	return aTable().
		SetStyle("striped").
		SetHideThead().
		SetMinWidth("0.01%").
		SetThead(types.Thead{
			types.TheadItem{Head: "key", Width: "50%"},
			types.TheadItem{Head: "value"},
		}).
		SetInfoList(list).GetContent()
}

func lg(v template.HTML) template.HTML {
	return language.GetFromHtml(v, "system")
}

func itos(i interface{}) template.HTML {
	return template.HTML(fmt.Sprintf("%v", i))
}
