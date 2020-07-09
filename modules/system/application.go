package system

import (
	"fmt"
	"runtime"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/utils"
)

var (
	startTime = time.Now()
)

type AppStatus struct {
	Uptime       string
	NumGoroutine int

	// General statistics.
	MemAllocated string // bytes allocated and still in use
	MemTotal     string // bytes allocated (even if freed)
	MemSys       string // bytes obtained from system (sum of XxxSys below)
	Lookups      uint64 // number of pointer lookups
	MemMallocs   uint64 // number of mallocs
	MemFrees     uint64 // number of frees

	// Main allocation heap statistics.
	HeapAlloc    string // bytes allocated and still in use
	HeapSys      string // bytes obtained from system
	HeapIdle     string // bytes in idle spans
	HeapInuse    string // bytes in non-idle span
	HeapReleased string // bytes released to the OS
	HeapObjects  uint64 // total number of allocated objects

	// Low-level fixed-size structure allocator statistics.
	//	Inuse is bytes used now.
	//	Sys is bytes obtained from system.
	StackInuse  string // bootstrap stacks
	StackSys    string
	MSpanInuse  string // mspan structures
	MSpanSys    string
	MCacheInuse string // mcache structures
	MCacheSys   string
	BuckHashSys string // profiling bucket hash table
	GCSys       string // GC metadata
	OtherSys    string // other system allocations

	// Garbage collector statistics.
	NextGC       string // next run in HeapAlloc time (bytes)
	LastGC       string // last run in absolute time (ns)
	PauseTotalNs string
	PauseNs      string // circular buffer of recent GC pause times, most recent at [(NumGC+255)%256]
	NumGC        uint32
}

func GetAppStatus() AppStatus {
	var app AppStatus
	app.Uptime = utils.TimeSincePro(startTime, language.Lang[config.GetLanguage()])

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	app.NumGoroutine = runtime.NumGoroutine()

	app.MemAllocated = utils.FileSize(m.Alloc)
	app.MemTotal = utils.FileSize(m.TotalAlloc)
	app.MemSys = utils.FileSize(m.Sys)
	app.Lookups = m.Lookups
	app.MemMallocs = m.Mallocs
	app.MemFrees = m.Frees

	app.HeapAlloc = utils.FileSize(m.HeapAlloc)
	app.HeapSys = utils.FileSize(m.HeapSys)
	app.HeapIdle = utils.FileSize(m.HeapIdle)
	app.HeapInuse = utils.FileSize(m.HeapInuse)
	app.HeapReleased = utils.FileSize(m.HeapReleased)
	app.HeapObjects = m.HeapObjects

	app.StackInuse = utils.FileSize(m.StackInuse)
	app.StackSys = utils.FileSize(m.StackSys)
	app.MSpanInuse = utils.FileSize(m.MSpanInuse)
	app.MSpanSys = utils.FileSize(m.MSpanSys)
	app.MCacheInuse = utils.FileSize(m.MCacheInuse)
	app.MCacheSys = utils.FileSize(m.MCacheSys)
	app.BuckHashSys = utils.FileSize(m.BuckHashSys)
	app.GCSys = utils.FileSize(m.GCSys)
	app.OtherSys = utils.FileSize(m.OtherSys)

	app.NextGC = utils.FileSize(m.NextGC)
	app.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	app.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	app.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	app.NumGC = m.NumGC

	return app
}

type SysStatus struct {
	CpuLogicalCore int
	CpuCore        int
	OSPlatform     string
	OSFamily       string
	OSVersion      string
	Load1          float64
	Load5          float64
	Load15         float64
	MemTotal       string
	MemAvailable   string
	MemUsed        string
}
