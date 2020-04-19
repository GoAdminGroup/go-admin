package controller

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gobuffalo/buffalo/runtime"
	"html/template"
	"os"
)

func (h *Handler) SystemInfo(ctx *context.Context) {

	size := types.Size(6, 6, 6)

	sys := system.GetSysStatus()

	box1 := aBox().
		WithHeadBorder().
		SetHeadColor("#f5f5f5").
		SetHeader("<b>" + lg("application") + "</b>").
		SetBody(srow(lg("app_name"), "GoAdmin") +
			srow(lg("go_admin_version"), system.Version()) +
			srow(lg("theme_name"), aTemplate().Name()) +
			srow(lg("theme_version"), aTemplate().GetVersion())).
		GetContent()

	box2 := aBox().
		WithHeadBorder().
		SetHeadColor("#f5f5f5").
		SetHeader("<b>" + lg("system") + "</b>").
		SetBody(srow(lg("cpu_logical_core"), itos(sys.CpuLogicalCore)) +
			srow(lg("cpu_core"), itos(sys.CpuCore)) +
			`<div><hr></div>` +
			srow(lg("os_platform"), sys.OSPlatform) +
			srow(lg("os_family"), sys.OSFamily) +
			srow(lg("os_version"), sys.OSVersion) +
			`<div><hr></div>` +
			srow(lg("load1"), fmt.Sprintf("%.2f", sys.Load1)) +
			srow(lg("load5"), fmt.Sprintf("%.2f", sys.Load5)) +
			srow(lg("load15"), fmt.Sprintf("%.2f", sys.Load15)) +
			`<div><hr></div>` +
			srow(lg("mem_total"), sys.MemTotal) +
			srow(lg("mem_available"), sys.MemAvailable) +
			srow(lg("mem_used"), sys.MemUsed)).
		GetContent()

	col1 := aCol().SetSize(size).SetContent(box1 + box2).GetContent()

	app := system.GetAppStatus()

	box3 := aBox().
		WithHeadBorder().
		SetHeadColor("#f5f5f5").
		SetHeader("<b>" + lg("application run") + "</b>").
		SetBody(srow(lg("golang_version"), runtime.Version) +
			srow(lg("process_id"), itos(os.Getpid())) +
			srow(lg("server_uptime"), app.Uptime) +
			srow(lg("current_goroutine"), itos(app.NumGoroutine)) +
			`<div><hr></div>` +
			srow(lg("current_memory_usage"), app.MemAllocated) +
			srow(lg("total_memory_allocated"), app.MemTotal) +
			srow(lg("memory_obtained"), app.MemSys) +
			srow(lg("pointer_lookup_times"), itos(app.Lookups)) +
			srow(lg("memory_allocate_times"), itos(app.MemMallocs)) +
			srow(lg("memory_free_times"), itos(app.MemFrees)) +
			`<div><hr></div>` +
			srow(lg("current_heap_usage"), app.HeapAlloc) +
			srow(lg("heap_memory_obtained"), app.HeapSys) +
			srow(lg("heap_memory_idle"), app.HeapIdle) +
			srow(lg("heap_memory_in_use"), app.HeapInuse) +
			srow(lg("heap_memory_released"), app.HeapReleased) +
			srow(lg("heap_objects"), itos(app.HeapObjects)) +
			`<div><hr></div>` +
			srow(lg("bootstrap_stack_usage"), app.StackInuse) +
			srow(lg("stack_memory_obtained"), app.StackSys) +
			srow(lg("mspan_structures_usage"), app.MSpanInuse) +
			srow(lg("mspan_structures_obtained"), app.HeapSys) +
			srow(lg("mcache_structures_usage"), app.MCacheInuse) +
			srow(lg("mcache_structures_obtained"), app.MCacheSys) +
			srow(lg("profiling_bucket_hash_table_obtained"), app.BuckHashSys) +
			srow(lg("gc_metadata_obtained"), app.GCSys) +
			srow(lg("other_system_allocation_obtained"), app.OtherSys) +
			`<div><hr></div>` +
			srow(lg("next_gc_recycle"), app.NextGC) +
			srow(lg("last_gc_time"), app.LastGC) +
			srow(lg("total_gc_pause"), app.PauseTotalNs) +
			srow(lg("last_gc_pause"), app.PauseNs) +
			srow(lg("gc_times"), itos(app.NumGC))).
		GetContent()

	col2 := aCol().SetSize(size).SetContent(box3).GetContent()

	row := aRow().SetContent(col1 + col2).GetContent()

	h.HTML(ctx, auth.Auth(ctx), types.Panel{
		Content:     row,
		Description: language.GetFromHtml("system info", "system"),
		Title:       language.GetFromHtml("system info", "system"),
	})
}

func srow(content1 template.HTML, content2 string) template.HTML {
	size := types.Size(6, 6, 6)
	return aRow().SetContent(aCol().SetSize(size).SetContent(content1).GetContent() +
		aCol().SetSize(size).SetContent(template.HTML(content2)).GetContent()).GetContent()
}

func lg(v template.HTML) template.HTML {
	return language.GetFromHtml(v, "system")
}

func itos(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
