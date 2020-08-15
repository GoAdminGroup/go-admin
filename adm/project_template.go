package main

var projectTemplate = map[string]string{
	"gin": `{{define "project"}}
package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"                    // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/{{.DriverModule}}" // sql driver
	_ "github.com/GoAdminGroup/themes/{{.Theme}}"                       // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"
	
	"{{.Module}}/pages"
	"{{.Module}}/tables"
	{{if ne .Orm ""}}"{{.Module}}/models"{{end}}
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/{{.Prefix}}", pages.GetDashBoard)
	eng.HTMLFile("GET", "/{{.Prefix}}/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	{{if ne .Orm ""}}models.Init(eng.{{title .Driver}}Connection()){{end}}

	_ = r.Run(":{{.Port}}")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.{{title .Driver}}Connection().Close()
}
{{end}}`,

	"beego": `{{define "project"}}
package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/beego"                   // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/{{.DriverModule}}"  // sql driver
	_ "github.com/GoAdminGroup/themes/{{.Theme}}"                        // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/astaxie/beego"

	"{{.Module}}/pages"
	"{{.Module}}/tables"
	{{if ne .Orm ""}}"{{.Module}}/models"{{end}}
)

func main() {
	startServer()
}

func startServer() {
	app := beego.NewApp()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	beego.SetStaticPath("/uploads", "uploads")

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/{{.Prefix}}", pages.GetDashBoard)
	eng.HTMLFile("GET", "/{{.Prefix}}/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	{{if ne .Orm ""}}models.Init(eng.{{title .Driver}}Connection()){{end}}

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = {{.Port}}
	go app.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.{{title .Driver}}Connection().Close()
}
{{end}}`,

	"buffalo": `{{define "project"}}
package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/buffalo"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/{{.DriverModule}}"  // sql driver
	_ "github.com/GoAdminGroup/themes/{{.Theme}}"                        // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gobuffalo/buffalo"	

	"{{.Module}}/pages"
	"{{.Module}}/tables"
	{{if ne .Orm ""}}"{{.Module}}/models"{{end}}
)

func main() {
	startServer()
}

func startServer() {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:{{.Port}}",
	})

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(bu); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/{{.Prefix}}", pages.GetDashBoard)
	eng.HTMLFile("GET", "/{{.Prefix}}/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	{{if ne .Orm ""}}models.Init(eng.{{title .Driver}}Connection()){{end}}

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	go func() {
		_ = bu.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.{{title .Driver}}Connection().Close()
}
{{end}}`,

	"chi": `{{define "project"}}
package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	_ "github.com/GoAdminGroup/go-admin/adapter/chi"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/{{.DriverModule}}"  // sql driver
	_ "github.com/GoAdminGroup/themes/{{.Theme}}"                        // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/go-chi/chi"

	"{{.Module}}/pages"
	"{{.Module}}/tables"
	{{if ne .Orm ""}}"{{.Module}}/models"{{end}}
)

func main() {
	startServer()
}

func startServer() {
	r := chi.NewRouter()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/{{.Prefix}}", pages.GetDashBoard)
	eng.HTMLFile("GET", "/{{.Prefix}}/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	{{if ne .Orm ""}}models.Init(eng.{{title .Driver}}Connection()){{end}}

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "uploads")
	fileServer(r, "/uploads", http.Dir(filesDir))

	go func() {
		_ = http.ListenAndServe(":{{.Port}}", r)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.{{title .Driver}}Connection().Close()
}


// fileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
{{end}}`,

	"echo": `{{define "project"}}
package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/echo"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/{{.DriverModule}}"  // sql driver
	_ "github.com/GoAdminGroup/themes/{{.Theme}}"                        // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/labstack/echo/v4"

	"{{.Module}}/pages"
	"{{.Module}}/tables"
	{{if ne .Orm ""}}"{{.Module}}/models"{{end}}
)

func main() {
	startServer()
}

func startServer() {
	e := echo.New()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(e); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/{{.Prefix}}", pages.GetDashBoard)
	eng.HTMLFile("GET", "/{{.Prefix}}/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	{{if ne .Orm ""}}models.Init(eng.{{title .Driver}}Connection()){{end}}

	e.Static("/uploads", "./uploads")

	go e.Logger.Fatal(e.Start(":{{.Port}}"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.{{title .Driver}}Connection().Close()
}
{{end}}`,

	"fasthttp": `{{define "project"}}
package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/fasthttp"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/{{.DriverModule}}"  // sql driver
	_ "github.com/GoAdminGroup/themes/{{.Theme}}"                        // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"{{.Module}}/pages"
	"{{.Module}}/tables"
	{{if ne .Orm ""}}"{{.Module}}/models"{{end}}
)

func main() {
	startServer()
}

func startServer() {
	router := fasthttprouter.New()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(router); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/{{.Prefix}}", pages.GetDashBoard)
	eng.HTMLFile("GET", "/{{.Prefix}}/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	{{if ne .Orm ""}}models.Init(eng.{{title .Driver}}Connection()){{end}}

	router.ServeFiles("/uploads/*filepath", "./uploads")

	go func() {
		_ = fasthttp.ListenAndServe(":{{.Port}}", router.Handler)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.{{title .Driver}}Connection().Close()
}
{{end}}`,

	"gf": `{{define "project"}}
package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gf"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/{{.DriverModule}}"  // sql driver
	_ "github.com/GoAdminGroup/themes/{{.Theme}}"                        // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gogf/gf/frame/g"

	"{{.Module}}/pages"
	"{{.Module}}/tables"
	{{if ne .Orm ""}}"{{.Module}}/models"{{end}}
)

func main() {
	startServer()
}

func startServer() {
	s := g.Server()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(s); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/{{.Prefix}}", pages.GetDashBoard)
	eng.HTMLFile("GET", "/{{.Prefix}}/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	{{if ne .Orm ""}}models.Init(eng.{{title .Driver}}Connection()){{end}}

	s.AddStaticPath("/uploads", "./uploads")

	s.SetPort({{.Port}})
	go s.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.{{title .Driver}}Connection().Close()
}
{{end}}`,

	"gorilla": `{{define "project"}}
package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gorilla"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/{{.DriverModule}}"  // sql driver
	_ "github.com/GoAdminGroup/themes/{{.Theme}}"                        // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gorilla/mux"

	"{{.Module}}/pages"
	"{{.Module}}/tables"
	{{if ne .Orm ""}}"{{.Module}}/models"{{end}}
)

func main() {
	startServer()
}

func startServer() {
	app := mux.NewRouter()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/{{.Prefix}}", pages.GetDashBoard)
	eng.HTMLFile("GET", "/{{.Prefix}}/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	{{if ne .Orm ""}}models.Init(eng.{{title .Driver}}Connection()){{end}}

	app.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	go func() {
		_ = http.ListenAndServe(":{{.Port}}", app)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.{{title .Driver}}Connection().Close()
}
{{end}}`,

	"iris": `{{define "project"}}
package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/iris"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/{{.DriverModule}}"  // sql driver
	_ "github.com/GoAdminGroup/themes/{{.Theme}}"                        // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/kataras/iris/v12"

	"{{.Module}}/pages"
	"{{.Module}}/tables"
	{{if ne .Orm ""}}"{{.Module}}/models"{{end}}
)

func main() {
	startServer()
}

func startServer() {
	app := iris.Default()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/{{.Prefix}}", pages.GetDashBoard)
	eng.HTMLFile("GET", "/{{.Prefix}}/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	{{if ne .Orm ""}}models.Init(eng.{{title .Driver}}Connection()){{end}}

	app.HandleDir("/uploads", "./uploads", iris.DirOptions{
		IndexName: "/index.html",
		Gzip:      false,
		ShowList:  false,
	})

	go func() {
		_ = app.Run(iris.Addr(":{{.Port}}"))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.{{title .Driver}}Connection().Close()
}
{{end}}`,
}

var swordIndexPage = []byte(`package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/sword/components/card"
	"github.com/GoAdminGroup/themes/sword/components/chart_legend"
	"github.com/GoAdminGroup/themes/sword/components/description"
	"github.com/GoAdminGroup/themes/sword/components/progress_group"
	"html/template"
)

func GetDashBoard(ctx *context.Context) (types.Panel, error) {

	components := template2.Get(config.GetTheme())
	colComp := components.Col()

	/**************************
	 * Info Box
	/**************************/

	cardcard := card.New().
		SetTitle("TOTAL REVENUE").
		SetSubTitle("¥ 113,340").
		SetAction(template.HTML(` + "`" + `<i aria-label="图标: info-circle-o" class="anticon anticon-info-circle-o"><svg viewBox="64 64 896 896" focusable="false" class="" data-icon="info-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M464 336a48 48 0 1 0 96 0 48 48 0 1 0-96 0zm72 112h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V456c0-4.4-3.6-8-8-8z"></path></svg></i>` + "`" + `)).
		SetContent(template.HTML(` + "`" + `<div><div title="" style="margin-right: 16px;"><span><span>Week Compare</span><span style="margin-left: 8px;">12%</span></span><span style="color: #f5222d;margin-left: 4px;top: 1px;"><i style="font-size: 12px;" aria-label="图标: caret-up" class="anticon anticon-caret-up"><svg viewBox="0 0 1024 1024" focusable="false" class="" data-icon="caret-up" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M858.9 689L530.5 308.2c-9.4-10.9-27.5-10.9-37 0L165.1 689c-12.2 14.2-1.2 35 18.5 35h656.8c19.7 0 30.7-20.8 18.5-35z"></path></svg></i></span></div><div class="antd-pro-pages-dashboard-analysis-components-trend-index-trendItem" title=""><span><span>Day Compare</span><span style="margin-left: 8px;">11%</span></span><span style="color: #52c41a;margin-left: 4px;top: 1px;"><i style="font-size: 12px;" aria-label="图标: caret-down" class="anticon anticon-caret-down"><svg viewBox="0 0 1024 1024" focusable="false" class="" data-icon="caret-down" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z"></path></svg></i></span></div></div>` + "`" + `)).
		SetFooter(template.HTML(` + "`" + `TOTAL DAY REVENUE <strong style="margin-left:8px;">$11,325</strong>` + "`" + `))
	infobox := cardcard.GetContent()

	infobox2 := cardcard.GetContent()

	infobox3 := cardcard.GetContent()

	infobox4 := cardcard.GetContent()

	var size = map[string]string{"md": "3", "sm": "6", "xs": "12"}
	infoboxCol1 := colComp.SetSize(size).SetContent(infobox).GetContent()
	infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
	infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
	infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()
	row1 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4).GetContent()

	/**************************
	 * Box
	/**************************/

	lineChart := chartjs.Line().
		SetID("salechart").
		SetHeight(180).
		SetTitle("Sales: 1 Jan, 2019 - 30 Jul, 2019").
		SetLabels([]string{"January", "February", "March", "April", "May", "June", "July"}).
		AddDataSet("Electronics").
		DSData([]float64{65, 59, 80, 81, 56, 55, 40}).
		DSFill(false).
		DSBorderColor("rgb(210, 214, 222)").
		DSLineTension(0.1).
		AddDataSet("Digital Goods").
		DSData([]float64{28, 48, 40, 19, 86, 27, 90}).
		DSFill(false).
		DSBorderColor("rgba(60,141,188,1)").
		DSLineTension(0.1).
		GetContent()

	title := ` + "`" + `<p class="text-center"><strong>Goal Completion</strong></p>` + "`" + `
	progressGroup := progress_group.New().
		SetTitle("Add Products to Cart").
		SetColor("#76b2d4").
		SetDenominator(200).
		SetMolecular(160).
		SetPercent(80).
		GetContent()

	progressGroup1 := progress_group.New().
		SetTitle("Complete Purchase").
		SetColor("#f17c6e").
		SetDenominator(400).
		SetMolecular(310).
		SetPercent(80).
		GetContent()

	progressGroup2 := progress_group.New().
		SetTitle("Visit Premium Page").
		SetColor("#ace0ae").
		SetDenominator(800).
		SetMolecular(490).
		SetPercent(80).
		GetContent()

	progressGroup3 := progress_group.New().
		SetTitle("Send Inquiries").
		SetColor("#fdd698").
		SetDenominator(500).
		SetMolecular(250).
		SetPercent(50).
		GetContent()

	boxInternalCol1 := colComp.SetContent(lineChart).SetSize(types.SizeMD(8)).GetContent()
	boxInternalCol2 := colComp.
		SetContent(template.HTML(title) + progressGroup + progressGroup1 + progressGroup2 + progressGroup3).
		SetSize(types.SizeMD(4)).
		GetContent()

	boxInternalRow := components.Row().SetContent(boxInternalCol1 + boxInternalCol2).GetContent()

	description1 := description.New().
		SetPercent("17").
		SetNumber("¥140,100").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	description2 := description.New().
		SetPercent("2").
		SetNumber("440,560").
		SetTitle("TOTAL REVENUE").
		SetArrow("down").
		SetColor("red").
		SetBorder("right").
		GetContent()

	description3 := description.New().
		SetPercent("12").
		SetNumber("¥140,050").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	description4 := description.New().
		SetPercent("1").
		SetNumber("30943").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		GetContent()

	size2 := map[string]string{"sm": "3", "xs": "6"}
	boxInternalCol3 := colComp.SetContent(description1).SetSize(size2).GetContent()
	boxInternalCol4 := colComp.SetContent(description2).SetSize(size2).GetContent()
	boxInternalCol5 := colComp.SetContent(description3).SetSize(size2).GetContent()
	boxInternalCol6 := colComp.SetContent(description4).SetSize(size2).GetContent()

	boxInternalRow2 := components.Row().SetContent(boxInternalCol3 + boxInternalCol4 + boxInternalCol5 + boxInternalCol6).GetContent()

	box := components.Box().WithHeadBorder().SetHeader("Monthly Recap Report").
		SetBody(boxInternalRow).
		SetFooter(boxInternalRow2).
		GetContent()

	boxcol := colComp.SetContent(box).SetSize(types.SizeMD(12)).GetContent()
	row2 := components.Row().SetContent(boxcol).GetContent()

	/**************************
	 * Pie Chart
	/**************************/

	pie := chartjs.Pie().
		SetHeight(170).
		SetLabels([]string{"Navigator", "Opera", "Safari", "FireFox", "IE", "Chrome"}).
		SetID("pieChart").
		AddDataSet("Chrome").
		DSData([]float64{100, 300, 600, 400, 500, 700}).
		DSBackgroundColor([]chartjs.Color{
			"rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)", "rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)",
		}).
		GetContent()
	legend := chart_legend.New().SetData([]map[string]string{
		{
			"label": " Chrome",
			"color": "red",
		}, {
			"label": " IE",
			"color": "Green",
		}, {
			"label": " FireFox",
			"color": "yellow",
		}, {
			"label": " Sarafri",
			"color": "blue",
		}, {
			"label": " Opera",
			"color": "light-blue",
		}, {
			"label": " Navigator",
			"color": "gray",
		},
	}).GetContent()

	boxDanger := components.Box().SetTheme("danger").WithHeadBorder().SetHeader("Browser Usage").
		SetBody(components.Row().
			SetContent(colComp.SetSize(types.SizeMD(8)).
				SetContent(pie).
				GetContent() + colComp.SetSize(types.SizeMD(4)).
				SetContent(legend).
				GetContent()).GetContent()).
		SetFooter(` + "`" + `<p class="text-center"><a href="javascript:void(0)" class="uppercase">View All Users</a></p>` + "`" + `).
		GetContent()

	tabs := components.Tabs().SetData([]map[string]template.HTML{
		{
			"title": "tabs1",
			"content": template.HTML(` + "`" + `<b>How to use:</b>

<p>Exactly like the original bootstrap tabs except you should use
the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
A wonderful serenity has taken possession of my entire soul,
like these sweet mornings of spring which I enjoy with my whole heart.
I am alone, and feel the charm of existence in this spot,
which was created for the bliss of souls like mine. I am so happy,
my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
that I neglect my talents. I should be incapable of drawing a single stroke
at the present moment; and yet I feel that I never was a greater artist than now.` + "`" + `),
		}, {
			"title": "tabs2",
			"content": template.HTML(` + "`" + `
The European languages are members of the same family. Their separate existence is a myth.
For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
in their grammar, their pronunciation and their most common words. Everyone realizes why a
new common language would be desirable: one could refuse to pay expensive translators. To
achieve this, it would be necessary to have uniform grammar, pronunciation and more common
words. If several languages coalesce, the grammar of the resulting language is more simple
and regular than that of the individual languages.
` + "`" + `),
		}, {
			"title": "tabs3",
			"content": template.HTML(` + "`" + `
Lorem Ipsum is simply dummy text of the printing and typesetting industry.
Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
when an unknown printer took a galley of type and scrambled it to make a type specimen book.
It has survived not only five centuries, but also the leap into electronic typesetting,
remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
like Aldus PageMaker including versions of Lorem Ipsum.
` + "`" + `),
		},
	}).GetContent()

	buttonTest := ` + "`" + `<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#exampleModal" data-whatever="@mdo">Open modal for @mdo</button>` + "`" + `
	popupForm := ` + "`" + `<form>
<div class="form-group">
<label for="recipient-name" class="col-form-label">Recipient:</label>
<input type="text" class="form-control" id="recipient-name">
</div>
<div class="form-group">
<label for="message-text" class="col-form-label">Message:</label>
<textarea class="form-control" id="message-text"></textarea>
</div>
</form>` + "`" + `
	popup := components.Popup().SetID("exampleModal").
		SetFooter("Save Change").
		SetTitle("this is a popup").
		SetBody(template.HTML(popupForm)).
		GetContent()

	col5 := colComp.SetSize(types.SizeMD(8)).SetContent(tabs + template.HTML(buttonTest)).GetContent()
	col6 := colComp.SetSize(types.SizeMD(4)).SetContent(boxDanger + popup).GetContent()

	row4 := components.Row().SetContent(col5 + col6).GetContent()

	return types.Panel{
		Content:     row1 + row2 + row4,
		Title:       "Dashboard",
		Description: "dashboard example",
	}, nil
}
`)

var adminlteIndexPage = []byte(`package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	tmpl "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/components/chart_legend"
	"github.com/GoAdminGroup/themes/adminlte/components/description"
	"github.com/GoAdminGroup/themes/adminlte/components/infobox"
	"github.com/GoAdminGroup/themes/adminlte/components/productlist"
	"github.com/GoAdminGroup/themes/adminlte/components/progress_group"
	"github.com/GoAdminGroup/themes/adminlte/components/smallbox"
	"html/template"
)

func GetDashBoard(ctx *context.Context) (types.Panel, error) {

	components := tmpl.Default()
	colComp := components.Col()

	/**************************
	 * Info Box
	/**************************/

	infobox1 := infobox.New().
		SetText("CPU TRAFFIC").
		SetColor("aqua").
		SetNumber("100").
		SetIcon("ion-ios-gear-outline").
		GetContent()

	infobox2 := infobox.New().
		SetText("Likes").
		SetColor("red").
		SetNumber("1030.00<small>$</small>").
		SetIcon(icon.GooglePlus).
		GetContent()

	infobox3 := infobox.New().
		SetText("Sales").
		SetColor("green").
		SetNumber("760").
		SetIcon("ion-ios-cart-outline").
		GetContent()

	infobox4 := infobox.New().
		SetText("New Members").
		SetColor("yellow").
		SetNumber("2,349").
		SetIcon("ion-ios-people-outline"). // svg is ok
		GetContent()

	var size = types.SizeMD(3).SM(6).XS(12)
	infoboxCol1 := colComp.SetSize(size).SetContent(infobox1).GetContent()
	infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
	infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
	infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()
	row1 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4).GetContent()

	/**************************
	 * Box
	/**************************/

	table := components.Table().SetType("table").SetInfoList([]map[string]types.InfoItem{
		{
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		}, {
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		}, {
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		}, {
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		},
	}).SetThead(types.Thead{
		{Head: "Order ID"},
		{Head: "Item"},
		{Head: "Status"},
		{Head: "Popularity"},
	}).GetContent()

	boxInfo := components.Box().
		WithHeadBorder().
		SetHeader("Latest Orders").
		SetHeadColor("#f7f7f7").
		SetBody(table).
		SetFooter(` + "`" + `<div class="clearfix"><a href="javascript:void(0)" class="btn btn-sm btn-info btn-flat pull-left">处理订单</a><a href="javascript:void(0)" class="btn btn-sm btn-default btn-flat pull-right">查看所有新订单</a> </div>` + "`" + `).
		GetContent()

	tableCol := colComp.SetSize(types.SizeMD(8)).SetContent(row1 + boxInfo).GetContent()

	/**************************
	 * Product List
	/**************************/

	productList := productlist.New().SetData([]map[string]string{
		{
			"img":         "//adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
			"title":       "GoAdmin",
			"has_tabel":   "true",
			"labeltype":   "warning",
			"label":       "free",
			"description": ` + "`" + `a framework help you build the dataviz system` + "`" + `,
		}, {
			"img":         "//adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
			"title":       "GoAdmin",
			"has_tabel":   "true",
			"labeltype":   "warning",
			"label":       "free",
			"description": ` + "`" + `a framework help you build the dataviz system` + "`" + `,
		}, {
			"img":         "//adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
			"title":       "GoAdmin",
			"has_tabel":   "true",
			"labeltype":   "warning",
			"label":       "free",
			"description": ` + "`" + `a framework help you build the dataviz system` + "`" + `,
		}, {
			"img":         "//adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
			"title":       "GoAdmin",
			"has_tabel":   "true",
			"labeltype":   "warning",
			"label":       "free",
			"description": ` + "`" + `a framework help you build the dataviz system` + "`" + `,
		},
	}).GetContent()

	boxWarning := components.Box().SetTheme("warning").WithHeadBorder().SetHeader("Recently Added Products").
		SetBody(productList).
		SetFooter(` + "`" + `<a href="javascript:void(0)" class="uppercase">View All Products</a>` + "`" + `).
		GetContent()

	newsCol := colComp.SetSize(types.SizeMD(4)).SetContent(boxWarning).GetContent()

	row5 := components.Row().SetContent(tableCol + newsCol).GetContent()

	/**************************
	 * Box
	/**************************/

	line := chartjs.Line()

	lineChart := line.
		SetID("salechart").
		SetHeight(180).
		SetTitle("Sales: 1 Jan, 2019 - 30 Jul, 2019").
		SetLabels([]string{"January", "February", "March", "April", "May", "June", "July"}).
		AddDataSet("Electronics").
		DSData([]float64{65, 59, 80, 81, 56, 55, 40}).
		DSFill(false).
		DSBorderColor("rgb(210, 214, 222)").
		DSLineTension(0.1).
		AddDataSet("Digital Goods").
		DSData([]float64{28, 48, 40, 19, 86, 27, 90}).
		DSFill(false).
		DSBorderColor("rgba(60,141,188,1)").
		DSLineTension(0.1).
		GetContent()

	title := ` + "`" + `<p class="text-center"><strong>Goal Completion</strong></p>` + "`" + `
	progressGroup := progress_group.New().
		SetTitle("Add Products to Cart").
		SetColor("#76b2d4").
		SetDenominator(200).
		SetMolecular(160).
		SetPercent(80).
		GetContent()

	progressGroup1 := progress_group.New().
		SetTitle("Complete Purchase").
		SetColor("#f17c6e").
		SetDenominator(400).
		SetMolecular(310).
		SetPercent(80).
		GetContent()

	progressGroup2 := progress_group.New().
		SetTitle("Visit Premium Page").
		SetColor("#ace0ae").
		SetDenominator(800).
		SetMolecular(490).
		SetPercent(80).
		GetContent()

	progressGroup3 := progress_group.New().
		SetTitle("Send Inquiries").
		SetColor("#fdd698").
		SetDenominator(500).
		SetMolecular(250).
		SetPercent(50).
		GetContent()

	boxInternalCol1 := colComp.SetContent(lineChart).SetSize(types.SizeMD(8)).GetContent()
	boxInternalCol2 := colComp.
		SetContent(template.HTML(title) + progressGroup + progressGroup1 + progressGroup2 + progressGroup3).
		SetSize(types.SizeMD(4)).
		GetContent()

	boxInternalRow := components.Row().SetContent(boxInternalCol1 + boxInternalCol2).GetContent()

	description1 := description.New().
		SetPercent("17").
		SetNumber("¥140,100").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	description2 := description.New().
		SetPercent("2").
		SetNumber("440,560").
		SetTitle("TOTAL REVENUE").
		SetArrow("down").
		SetColor("red").
		SetBorder("right").
		GetContent()

	description3 := description.New().
		SetPercent("12").
		SetNumber("¥140,050").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	description4 := description.New().
		SetPercent("1").
		SetNumber("30943").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		GetContent()

	size2 := types.SizeSM(3).XS(6)
	boxInternalCol3 := colComp.SetContent(description1).SetSize(size2).GetContent()
	boxInternalCol4 := colComp.SetContent(description2).SetSize(size2).GetContent()
	boxInternalCol5 := colComp.SetContent(description3).SetSize(size2).GetContent()
	boxInternalCol6 := colComp.SetContent(description4).SetSize(size2).GetContent()

	boxInternalRow2 := components.Row().SetContent(boxInternalCol3 + boxInternalCol4 + boxInternalCol5 + boxInternalCol6).GetContent()

	box := components.Box().WithHeadBorder().SetHeader("Monthly Recap Report").
		SetBody(boxInternalRow).
		SetFooter(boxInternalRow2).
		GetContent()

	boxcol := colComp.SetContent(box).SetSize(types.SizeMD(12)).GetContent()
	row2 := components.Row().SetContent(boxcol).GetContent()

	/**************************
	 * Small Box
	/**************************/

	smallbox1 := smallbox.New().SetColor("blue").SetIcon("ion-ios-gear-outline").SetUrl("/").SetTitle("new users").SetValue("345￥").GetContent()
	smallbox2 := smallbox.New().SetColor("yellow").SetIcon("ion-ios-cart-outline").SetUrl("/").SetTitle("new users").SetValue("80%").GetContent()
	smallbox3 := smallbox.New().SetColor("red").SetIcon("fa-user").SetUrl("/").SetTitle("new users").SetValue("645￥").GetContent()
	smallbox4 := smallbox.New().SetColor("green").SetIcon("ion-ios-cart-outline").SetUrl("/").SetTitle("new users").SetValue("889￥").GetContent()

	col1 := colComp.SetSize(size).SetContent(smallbox1).GetContent()
	col2 := colComp.SetSize(size).SetContent(smallbox2).GetContent()
	col3 := colComp.SetSize(size).SetContent(smallbox3).GetContent()
	col4 := colComp.SetSize(size).SetContent(smallbox4).GetContent()

	row3 := components.Row().SetContent(col1 + col2 + col3 + col4).GetContent()

	/**************************
	 * Pie Chart
	/**************************/

	pie := chartjs.Pie().
		SetHeight(170).
		SetLabels([]string{"Navigator", "Opera", "Safari", "FireFox", "IE", "Chrome"}).
		SetID("pieChart").
		AddDataSet("Chrome").
		DSData([]float64{100, 300, 600, 400, 500, 700}).
		DSBackgroundColor([]chartjs.Color{
			"rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)", "rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)",
		}).
		GetContent()

	legend := chart_legend.New().SetData([]map[string]string{
		{
			"label": " Chrome",
			"color": "red",
		}, {
			"label": " IE",
			"color": "Green",
		}, {
			"label": " FireFox",
			"color": "yellow",
		}, {
			"label": " Sarafri",
			"color": "blue",
		}, {
			"label": " Opera",
			"color": "light-blue",
		}, {
			"label": " Navigator",
			"color": "gray",
		},
	}).GetContent()

	boxDanger := components.Box().SetTheme("danger").WithHeadBorder().SetHeader("Browser Usage").
		SetBody(components.Row().
			SetContent(colComp.SetSize(types.SizeMD(8)).
				SetContent(pie).
				GetContent() + colComp.SetSize(types.SizeMD(4)).
				SetContent(legend).
				GetContent()).GetContent()).
		SetFooter(` + "`" + `<p class="text-center"><a href="javascript:void(0)" class="uppercase">View All Users</a></p>` + "`" + `).
		GetContent()

	tabs := components.Tabs().SetData([]map[string]template.HTML{
		{
			"title": "tabs1",
			"content": template.HTML(` + "`" + `<b>How to use:</b>

<p>Exactly like the original bootstrap tabs except you should use
the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
A wonderful serenity has taken possession of my entire soul,
like these sweet mornings of spring which I enjoy with my whole heart.
I am alone, and feel the charm of existence in this spot,
which was created for the bliss of souls like mine. I am so happy,
my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
that I neglect my talents. I should be incapable of drawing a single stroke
at the present moment; and yet I feel that I never was a greater artist than now.` + "`" + `),
		}, {
			"title": "tabs2",
			"content": template.HTML(` + "`" + `
The European languages are members of the same family. Their separate existence is a myth.
For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
in their grammar, their pronunciation and their most common words. Everyone realizes why a
new common language would be desirable: one could refuse to pay expensive translators. To
achieve this, it would be necessary to have uniform grammar, pronunciation and more common
words. If several languages coalesce, the grammar of the resulting language is more simple
and regular than that of the individual languages.
` + "`" + `),
		}, {
			"title": "tabs3",
			"content": template.HTML(` + "`" + `
Lorem Ipsum is simply dummy text of the printing and typesetting industry.
Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
when an unknown printer took a galley of type and scrambled it to make a type specimen book.
It has survived not only five centuries, but also the leap into electronic typesetting,
remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
like Aldus PageMaker including versions of Lorem Ipsum.
` + "`" + `),
		},
	}).GetContent()

	buttonTest := ` + "`" + `<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#exampleModal" data-whatever="@mdo">Open modal for @mdo</button>` + "`" + `
	popupForm := ` + "`" + `<form>
<div class="form-group">
<label for="recipient-name" class="col-form-label">Recipient:</label>
<input type="text" class="form-control" id="recipient-name">
</div>
<div class="form-group">
<label for="message-text" class="col-form-label">Message:</label>
<textarea class="form-control" id="message-text"></textarea>
</div>
</form>` + "`" + `
	popup := components.Popup().SetID("exampleModal").
		SetFooter("Save Change").
		SetTitle("this is a popup").
		SetBody(template.HTML(popupForm)).
		GetContent()

	col5 := colComp.SetSize(types.SizeMD(8)).SetContent(tabs + template.HTML(buttonTest)).GetContent()
	col6 := colComp.SetSize(types.SizeMD(4)).SetContent(boxDanger + popup).GetContent()

	row4 := components.Row().SetContent(col5 + col6).GetContent()

	return types.Panel{
		Content:     row3 + row2 + row5 + row4,
		Title:       "Dashboard",
		Description: "dashboard example",
	}, nil
}`)

var mainTest = []byte(`package main

import (
	"./tables"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/tests"
	"github.com/GoAdminGroup/go-admin/tests/common"
	"github.com/GoAdminGroup/go-admin/tests/frameworks/gin"
	"github.com/GoAdminGroup/go-admin/tests/web"
	"github.com/gavv/httpexpect"
	"log"
	"testing"
)

// Black box testing
func TestMainBlackBox(t *testing.T) {
	cfg := config.ReadFromJson("./config.json")
	tests.BlackBoxTestSuit(t, gin.NewHandler, cfg.Databases, tables.Generators, func(cfg config.DatabaseList) {
		// Data cleaner of the framework
		tests.Cleaner(cfg)
		// Clean your own data:
		// ...
	}, func(e *httpexpect.Expect) {
		// Test cases of the framework
		common.Test(e)
		// Write your own API test, for example:
		// More usages: https://github.com/gavv/httpexpect
		// e.POST("/signin").Expect().Status(http.StatusOK)
	})
}

// User acceptance testing
func TestMainUserAcceptance(t *testing.T) {
	web.UserAcceptanceTestSuit(t, func(t *testing.T, page *web.Page) {
		// Write test case base on chromedriver, for example:
		// More usages: https://github.com/sclevine/agouti
		page.NavigateTo("http://127.0.0.1:9033/admin")
		//page.Contain("username")
		//page.Click("")
	}, func(quit chan struct{}) {
		// start the server:
		// ....
		go startServer()
		<-quit
		log.Print("test quit")
	}, true) // if local parameter is true, it will not be headless, and window not close when finishing tests.
}`)

var mainTestCN = []byte(`package main

import (
	"./tables"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/tests"
	"github.com/GoAdminGroup/go-admin/tests/common"
	"github.com/GoAdminGroup/go-admin/tests/frameworks/gin"
	"github.com/GoAdminGroup/go-admin/tests/web"
	"github.com/gavv/httpexpect"
	"log"
	"testing"
)

// 黑盒测试
func TestMainBlackBox(t *testing.T) {
	cfg := config.ReadFromJson("./config.json")
	tests.BlackBoxTestSuit(t, gin.NewHandler, cfg.Databases, tables.Generators, func(cfg config.DatabaseList) {
		// 框架自带数据清理
		tests.Cleaner(cfg)
		// 以下清理自己的数据：
		// ...
	}, func(e *httpexpect.Expect) {
		// 框架自带内置表测试
		common.Test(e)
		// 以下写API测试：
		// 更多用法：https://github.com/gavv/httpexpect
		// ...
		// e.POST("/signin").Expect().Status(http.StatusOK)
	})
}

// 浏览器验收测试
func TestMainUserAcceptance(t *testing.T) {
	web.UserAcceptanceTestSuit(t, func(t *testing.T, page *web.Page) {
		// 写浏览器测试，基于chromedriver
		// 更多用法：https://github.com/sclevine/agouti
		// page.NavigateTo("http://127.0.0.1:9033/admin")
		// page.Contain("username")
		// page.Click("")
	}, func(quit chan struct{}) {
		// 启动服务器
		go startServer()
		<-quit
		log.Print("test quit")
	}, true)
}`)

var makefile = []byte(`GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod
GOTEST = $(GOCMD) test
BINARY_NAME = goadmin
CLI = adm

all: serve

init:
	$(GOMOD) init $(module)

install:
	$(GOMOD) tidy

serve:
	$(GOCMD) run .

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -v ./

generate:
	$(CLI) generate -c adm.ini

test: black-box-test user-acceptance-test

black-box-test: ready-for-data
	$(GOTEST) -v -test.run=TestMainBlackBox
	make clean

user-acceptance-test: ready-for-data
	$(GOTEST) -v -test.run=TestMainUserAcceptance
	make clean

ready-for-data:
	cp admin.db admin_test.db

clean:
	rm admin_test.db

.PHONY: all serve build generate test black-box-test user-acceptance-test ready-for-data clean`)

var admINI = `{{define "ini"}}
; default database config 默认数据库配置
[database]{{if eq .DriverName "sqlite"}}
driver = sqlite
file = {{.File}}
{{else}}
driver = {{.DriverName}}
host = {{.Host}}
username = {{.User}} 
port = {{.Port}}
password = {{.Password}}
database = {{.Database}} 
{{end}}
; Here are new tables to generate. 新的待转换的表格
; tables = new_table1,new_table2

; specified connection database config 指定数据库配置
; for example, database config which connection name is mydb.
;[database.mydb]

; table model config 数据模型设置
[model]
package = tables
connection = default
output = ./tables
{{end}}`

var readme = `# GoAdmin Instruction

GoAdmin is a golang framework help gopher quickly build a data visualization platform. 

- [github](https://github.com/GoAdminGroup/go-admin)
- [forum](http://discuss.go-admin.com)
- [document](https://book.go-admin.cn)

## Directories Introduction

` + "```" + `
.
├── Dockerfile          Dockerfile
├── Makefile            Makefile
├── adm.ini             adm config
├── admin.db            sqlite database
├── build               binary build target folder
├── config.json         config file
├── go.mod              go.mod
├── go.sum              go.sum
├── html                frontend html files
├── logs                logs
├── main.go             main.go
├── main_test.go        ci test
├── pages               page controllers
├── tables              table models
└── uploads             upload files
` + "```" + `

## Generate Table Model

### online tool

visit: http://127.0.0.1:%s/info/generate/new

### use adm

` + "```" + `
adm generate
` + "```" + `

`

var readmeCN = `# GoAdmin 介绍

GoAdmin 是一个帮你快速搭建数据可视化管理应用平台的框架。 

- [github](https://github.com/GoAdminGroup/go-admin)
- [论坛](http://discuss.go-admin.com)
- [文档](https://book.go-admin.cn)

## 目录介绍

` + "```" + `
.
├── Dockerfile          Dockerfile
├── Makefile            Makefile
├── adm.ini             adm配置文件
├── admin.db            sqlite数据库
├── build               二进制构建目标文件夹
├── config.json         配置文件
├── go.mod              go.mod
├── go.sum              go.sum
├── html                前端html文件
├── logs                日志
├── main.go             main.go
├── main_test.go        CI测试
├── pages               页面控制器
├── tables              数据模型
└── uploads             上传文件夹
` + "```" + `

## 生成CRUD数据模型

### 在线工具

管理员身份运行后，访问：http://127.0.0.1:%s/info/generate/new

### 使用命令行工具

` + "```" + `
adm generate -l cn -c adm.ini
` + "```" + `

`
