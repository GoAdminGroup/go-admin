package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/page"
	template2 "github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

func ShowDashboard(ctx *context.Context) {
	page.SetPageContent(ctx, func() types.Panel {

		components := template2.Get(Config.THEME)
		colComp := components.Col()

		/**************************
		 * Info Box
		/**************************/

		infobox := components.InfoBox().
			SetText("CPU TRAFFIC").
			SetColor("blue").SetNumber("41,410").
			SetIcon("ion-ios-gear-outline").
			GetContent()

		infobox2 := components.InfoBox().
			SetText("Likes").
			SetColor("red").SetNumber("90<small>%</small>").
			SetIcon("fa-google-plus").
			GetContent()

		infobox3 := components.InfoBox().
			SetText("Sales").
			SetColor("green").SetNumber("760").
			SetIcon("ion-ios-cart-outline").
			GetContent()

		infobox4 := components.InfoBox().
			SetText("New Members").
			SetColor("yellow").SetNumber("2,000").
			SetIcon("ion-ios-people-outline").
			GetContent()

		var size = map[string]string{"md": "3", "sm": "6", "xs": "12"}
		infoboxCol1 := colComp.SetSize(size).SetContent(infobox).GetContent()
		infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
		infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
		infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()
		row1 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4).GetContent()

		/**************************
		 * Box
		/**************************/

		chartdata := `{"datasets":[{"data":[65,59,80,81,56,55,40],"fillColor":"rgb(210, 214, 222)","label":"Electronics","pointColor":"rgb(210, 214, 222)","pointHighlightFill":"#fff","pointHighlightStroke":"rgb(220,220,220)","pointStrokeColor":"#c1c7d1","strokeColor":"rgb(210, 214, 222)"},{"data":[28,48,40,19,86,27,90],"fillColor":"rgba(60,141,188,0.9)","label":"Digital Goods","pointColor":"#3b8bba","pointHighlightFill":"#fff","pointHighlightStroke":"rgba(60,141,188,1)","pointStrokeColor":"rgba(60,141,188,1)","strokeColor":"rgba(60,141,188,0.8)"}],"labels":["January","February","March","April","May","June","July"]}`

		AreaChart := components.AreaChart().SetID("salechart").
			SetData(chartdata).
			SetHeight(180).
			SetTitle("Sales: 1 Jan, 2014 - 30 Jul, 2014").GetContent()

		title := `<p class="text-center"><strong>Goal Completion</strong></p>`
		progressGroup := components.ProgressGroup().SetTitle("Add Products to Cart").
			SetColor("aqua").SetDenominator(200).SetMolecular(160).SetPercent(80).GetContent()
		progressGroup1 := components.ProgressGroup().SetTitle("Complete Purchase").
			SetColor("red").SetDenominator(400).SetMolecular(310).SetPercent(80).GetContent()
		progressGroup2 := components.ProgressGroup().SetTitle("Visit Premium Page").
			SetColor("green").SetDenominator(800).SetMolecular(490).SetPercent(80).GetContent()
		progressGroup3 := components.ProgressGroup().SetTitle("Send Inquiries").
			SetColor("yellow").SetDenominator(500).SetMolecular(250).SetPercent(50).GetContent()

		boxInternalCol1 := colComp.SetContent(AreaChart).SetSize(map[string]string{"md": "8"}).GetContent()
		boxInternalCol2 := colComp.
			SetContent(template.HTML(title) + progressGroup + progressGroup1 + progressGroup2 + progressGroup3).
			SetSize(map[string]string{"md": "4"}).
			GetContent()

		boxInternalRow := components.Row().SetContent(boxInternalCol1 + boxInternalCol2).GetContent()

		description1 := components.Description().SetPercent("17").
			SetNumber("¥100,000").SetTitle("TOTAL REVENUE").SetArrow("up").SetColor("green").
			SetBorder("right").GetContent()
		description2 := components.Description().SetPercent("17").
			SetNumber("¥100,000").SetTitle("TOTAL REVENUE").SetArrow("up").SetColor("green").
			SetBorder("right").GetContent()
		description3 := components.Description().SetPercent("17").
			SetNumber("¥100,000").SetTitle("TOTAL REVENUE").SetArrow("up").SetColor("green").
			SetBorder("right").GetContent()
		description4 := components.Description().SetPercent("17").
			SetNumber("¥100,000").SetTitle("TOTAL REVENUE").SetArrow("up").SetColor("green").GetContent()

		size2 := map[string]string{"sm": "3", "xs": "6"}
		boxInternalCol3 := colComp.SetContent(description1).SetSize(size2).GetContent()
		boxInternalCol4 := colComp.SetContent(description2).SetSize(size2).GetContent()
		boxInternalCol5 := colComp.SetContent(description3).SetSize(size2).GetContent()
		boxInternalCol6 := colComp.SetContent(description4).SetSize(size2).GetContent()

		boxInternalRow2 := components.Row().SetContent(boxInternalCol3 + boxInternalCol4 + boxInternalCol5 + boxInternalCol6).GetContent()

		box := components.Box().WithHeadBorder(true).SetHeader("Monthly Recap Report").
			SetBody(boxInternalRow).
			SetFooter(boxInternalRow2).
			GetContent()

		boxcol := colComp.SetContent(box).SetSize(map[string]string{"md": "12"}).GetContent()
		row2 := components.Row().SetContent(boxcol).GetContent()

		/**************************
		 * Small Box
		/**************************/

		smallbox := components.SmallBox().SetUrl("/").SetTitle("new users").SetValue("1000").GetContent()

		col1 := colComp.SetSize(size).SetContent(smallbox).GetContent()
		col2 := colComp.SetSize(size).SetContent(smallbox).GetContent()
		col3 := colComp.SetSize(size).SetContent(smallbox).GetContent()
		col4 := colComp.SetSize(size).SetContent(smallbox).GetContent()

		row3 := components.Row().SetContent(col1 + col2 + col3 + col4).GetContent()

		/**************************
		 * Pie Chart
		/**************************/

		table := components.Table().SetType("table").SetInfoList([]map[string]template.HTML{
			{
				"Order ID":   "OR9842",
				"Item":       "Call of Duty IV",
				"Status":     "shipped",
				"Popularity": "90%",
			}, {
				"Order ID":   "OR9842",
				"Item":       "Call of Duty IV",
				"Status":     "shipped",
				"Popularity": "90%",
			}, {
				"Order ID":   "OR9842",
				"Item":       "Call of Duty IV",
				"Status":     "shipped",
				"Popularity": "90%",
			},
		}).SetThead([]map[string]string{
			{
				"head":     "Order ID",
				"sortable": "0",
			}, {
				"head":     "Item",
				"sortable": "0",
			}, {
				"head":     "Status",
				"sortable": "0",
			}, {
				"head":     "Popularity",
				"sortable": "0",
			},
		}).GetContent()

		pieData := `[{"value":700,"color":"#f56954","highlight":"#f56954","label":"Chrome"},{"value":500,"color":"#00a65a","highlight":"#00a65a","label":"IE"},{"value":400,"color":"#f39c12","highlight":"#f39c12","label":"FireFox"},{"value":600,"color":"#00c0ef","highlight":"#00c0ef","label":"Safari"},{"value":300,"color":"#3c8dbc","highlight":"#3c8dbc","label":"Opera"},{"value":100,"color":"#d2d6de","highlight":"#d2d6de","label":"Navigator"}]`
		pie := components.PieChart().SetHeight(170).SetData(pieData).SetID("pieChart").GetContent()
		legend := components.ChartLegend().SetData([]map[string]string{
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

		boxInfo := components.Box().SetTheme("info").WithHeadBorder(true).SetHeader("Latest Orders").
			SetBody(table).
			SetFooter(`<div class="clearfix"><a href="javascript:void(0)" class="btn btn-sm btn-info btn-flat pull-left">Place New Order</a><a href="javascript:void(0)" class="btn btn-sm btn-default btn-flat pull-right">View All Orders</a> </div>`).
			GetContent()

		lineChart := components.LineChart().SetID("linechart").
			SetData(chartdata).
			SetHeight(200).
			SetTitle("Sales: 1 Jan, 2014 - 30 Jul, 2014").GetContent()

		boxLineChart := components.Box().WithHeadBorder(true).SetHeader("Monthly Recap Report").
			SetBody(lineChart).
			GetContent()

		barChart := components.BarChart().SetID("barchart").
			SetData(chartdata).
			SetWidth(180).
			SetTitle("Sales: 1 Jan, 2014 - 30 Jul, 2014").GetContent()

		boxBarChart := components.Box().WithHeadBorder(true).SetHeader("Monthly Recap Report").
			SetBody(barChart).
			GetContent()

		chartCol := components.Col().SetSize(map[string]string{"md": "6"})

		chartRow := components.Row().SetContent(chartCol.SetContent(boxBarChart).GetContent() +
			chartCol.SetContent(boxLineChart).GetContent()).GetContent()

		boxDanger := components.Box().SetTheme("danger").WithHeadBorder(true).SetHeader("Browser Usage").
			SetBody(components.Row().
				SetContent(colComp.SetSize(map[string]string{"md": "8"}).
					SetContent(pie).
					GetContent() + colComp.SetSize(map[string]string{"md": "4"}).
					SetContent(legend).
					GetContent()).GetContent()).
			SetFooter(`<p class="text-center"><a href="javascript:void(0)" class="uppercase">View All Users</a></p>`).
			GetContent()

		/**************************
		 * Product List
		/**************************/

		productList := components.ProductList().SetData([]map[string]string{
			{
				"img":         "http://adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
				"title":       "Samsung TV",
				"has_tabel":   "true",
				"labeltype":   "warning",
				"label":       "$1800",
				"description": `Samsung 32" 1080p 60Hz LED Smart HDTV.`,
			}, {
				"img":         "http://adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
				"title":       "Samsung TV",
				"has_tabel":   "true",
				"labeltype":   "warning",
				"label":       "$1800",
				"description": `Samsung 32" 1080p 60Hz LED Smart HDTV.`,
			}, {
				"img":         "http://adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
				"title":       "Samsung TV",
				"has_tabel":   "true",
				"labeltype":   "warning",
				"label":       "$1800",
				"description": `Samsung 32" 1080p 60Hz LED Smart HDTV.`,
			}, {
				"img":         "http://adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
				"title":       "Samsung TV",
				"has_tabel":   "true",
				"labeltype":   "warning",
				"label":       "$1800",
				"description": `Samsung 32" 1080p 60Hz LED Smart HDTV.`,
			},
		}).GetContent()

		boxWarning := components.Box().SetTheme("warning").WithHeadBorder(true).SetHeader("Recently Added Products").
			SetBody(productList).
			SetFooter(`<a href="javascript:void(0)" class="uppercase">View All Products</a>`).
			GetContent()

		col5 := colComp.SetSize(map[string]string{"md": "8"}).SetContent(boxInfo + chartRow).GetContent()
		col6 := colComp.SetSize(map[string]string{"md": "4"}).SetContent(boxDanger + boxWarning).GetContent()

		row4 := components.Row().SetContent(col5 + col6).GetContent()

		return types.Panel{
			Content:     template.HTML(row1) + template.HTML(row2) + template.HTML(row3) + template.HTML(row4),
			Title:       "仪表盘",
			Description: "仪表盘",
		}
	})
}

func ShowErrorPage(ctx *context.Context, errorMsg string) {
	page.SetPageContent(ctx, func() types.Panel {
		alert := template2.Get(Config.THEME).Alert().SetTitle(template.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template.HTML(errorMsg)).GetContent()

		return types.Panel{
			Content:     alert,
			Description: "Error",
			Title:       "Error",
		}
	})
}
