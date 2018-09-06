package controller

import (
	"html/template"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/page"
	"github.com/chenhg5/go-admin/template/types"
	template2 "github.com/chenhg5/go-admin/template"
)

func ShowDashboard(ctx *context.Context) {
	page.SetPageContent(Config.THEME, Config.ADMIN_PREFIX, ctx, func() types.Panel {

		/**************************
		 * Info Box
		/**************************/

		infobox := template2.Get(Config.THEME).InfoBox().
			SetText("CPU TRAFFIC").
			SetColor("blue").SetNumber("41,410").
			SetIcon("ion-ios-gear-outline").
			GetContent()

		infobox2 := template2.Get(Config.THEME).InfoBox().
			SetText("Likes").
			SetColor("red").SetNumber("90<small>%</small>").
			SetIcon("fa-google-plus").
			GetContent()

		infobox3 := template2.Get(Config.THEME).InfoBox().
			SetText("Sales").
			SetColor("green").SetNumber("760").
			SetIcon("ion-ios-cart-outline").
			GetContent()

		infobox4 := template2.Get(Config.THEME).InfoBox().
			SetText("New Members").
			SetColor("yellow").SetNumber("2,000").
			SetIcon("ion-ios-people-outline").
			GetContent()

		var size = map[string]string{"md": "3","sm": "6","xs": "12"}
		infoboxCol1 := template2.Get(Config.THEME).Col().SetSize(size).SetContent(infobox).GetContent()
		infoboxCol2 := template2.Get(Config.THEME).Col().SetSize(size).SetContent(infobox2).GetContent()
		infoboxCol3 := template2.Get(Config.THEME).Col().SetSize(size).SetContent(infobox3).GetContent()
		infoboxCol4 := template2.Get(Config.THEME).Col().SetSize(size).SetContent(infobox4).GetContent()
		row1 := template2.Get(Config.THEME).Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4).GetContent()

		/**************************
		 * Box
		/**************************/

		chartdata := `{"datasets":[{"data":[65,59,80,81,56,55,40],"fillColor":"rgb(210, 214, 222)","label":"Electronics","pointColor":"rgb(210, 214, 222)","pointHighlightFill":"#fff","pointHighlightStroke":"rgb(220,220,220)","pointStrokeColor":"#c1c7d1","strokeColor":"rgb(210, 214, 222)"},{"data":[28,48,40,19,86,27,90],"fillColor":"rgba(60,141,188,0.9)","label":"Digital Goods","pointColor":"#3b8bba","pointHighlightFill":"#fff","pointHighlightStroke":"rgba(60,141,188,1)","pointStrokeColor":"rgba(60,141,188,1)","strokeColor":"rgba(60,141,188,0.8)"}],"labels":["January","February","March","April","May","June","July"]}`

		lineChart := template2.Get(Config.THEME).LineChart().SetID("salechart").
			SetPrefix(Config.ADMIN_PREFIX).SetData(chartdata).
			SetHeight(180).
			SetTitle("Sales: 1 Jan, 2014 - 30 Jul, 2014").GetContent()

		title := `<p class="text-center"><strong>Goal Completion</strong></p>`
		progressGroup := template2.Get(Config.THEME).ProgressGroup().SetTitle("Add Products to Cart").
			SetColor("aqua").SetDenominator(200).SetMolecular(160).SetPercent(80).GetContent()
		progressGroup1 := template2.Get(Config.THEME).ProgressGroup().SetTitle("Complete Purchase").
			SetColor("red").SetDenominator(400).SetMolecular(310).SetPercent(80).GetContent()
		progressGroup2 := template2.Get(Config.THEME).ProgressGroup().SetTitle("Visit Premium Page").
			SetColor("green").SetDenominator(800).SetMolecular(490).SetPercent(80).GetContent()
		progressGroup3 := template2.Get(Config.THEME).ProgressGroup().SetTitle("Send Inquiries").
			SetColor("yellow").SetDenominator(500).SetMolecular(250).SetPercent(50).GetContent()

		boxInternalCol1 := template2.Get(Config.THEME).Col().SetContent(lineChart).SetSize(map[string]string{"md": "8"}).GetContent()
		boxInternalCol2 := template2.Get(Config.THEME).Col().
			SetContent(template.HTML(title) + progressGroup + progressGroup1 + progressGroup2 + progressGroup3).
			SetSize(map[string]string{"md": "4"}).
			GetContent()

		boxInternalRow := template2.Get(Config.THEME).Row().SetContent(boxInternalCol1 + boxInternalCol2).GetContent()

		description1 := template2.Get(Config.THEME).Description().SetPercent("17").
			SetNumber("¥100,000").SetTitle("TOTAL REVENUE").SetArrow("up").SetColor("green").
			SetBorder("border-right").GetContent()
		description2 := template2.Get(Config.THEME).Description().SetPercent("17").
			SetNumber("¥100,000").SetTitle("TOTAL REVENUE").SetArrow("up").SetColor("green").
			SetBorder("border-right").GetContent()
		description3 := template2.Get(Config.THEME).Description().SetPercent("17").
			SetNumber("¥100,000").SetTitle("TOTAL REVENUE").SetArrow("up").SetColor("green").
			SetBorder("border-right").GetContent()
		description4 := template2.Get(Config.THEME).Description().SetPercent("17").
			SetNumber("¥100,000").SetTitle("TOTAL REVENUE").SetArrow("up").SetColor("green").GetContent()

		size2 := map[string]string{"sm": "3", "xs":"6"}
		boxInternalCol3 := template2.Get(Config.THEME).Col().SetContent(description1).SetSize(size2).GetContent()
		boxInternalCol4 := template2.Get(Config.THEME).Col().SetContent(description2).SetSize(size2).GetContent()
		boxInternalCol5 := template2.Get(Config.THEME).Col().SetContent(description3).SetSize(size2).GetContent()
		boxInternalCol6 := template2.Get(Config.THEME).Col().SetContent(description4).SetSize(size2).GetContent()

		boxInternalRow2 := template2.Get(Config.THEME).Row().SetContent(boxInternalCol3 + boxInternalCol4 + boxInternalCol5 + boxInternalCol6).GetContent()

		box := template2.Get(Config.THEME).Box().WithHeadBorder(true).SetHeader("Monthly Recap Report").
			SetBody(boxInternalRow).
			SetFooter(boxInternalRow2).
			GetContent()

		boxcol := template2.Get(Config.THEME).Col().SetContent(box).SetSize(map[string]string{"md": "12"}).GetContent()
		row2 := template2.Get(Config.THEME).Row().SetContent(boxcol).GetContent()

		/**************************
		 * Small Box
		/**************************/

		smallbox := template2.Get(Config.THEME).SmallBox().SetUrl("/").SetTitle("new users").SetValue("1000").GetContent()

		col1 := template2.Get(Config.THEME).Col().SetSize(size).SetContent(smallbox).GetContent()
		col2 := template2.Get(Config.THEME).Col().SetSize(size).SetContent(smallbox).GetContent()
		col3 := template2.Get(Config.THEME).Col().SetSize(size).SetContent(smallbox).GetContent()
		col4 := template2.Get(Config.THEME).Col().SetSize(size).SetContent(smallbox).GetContent()

		row3 := template2.Get(Config.THEME).Row().SetContent(col1 + col2 + col3 + col4).GetContent()

		return types.Panel{
			Content:     template.HTML(row1) + template.HTML(row2) + template.HTML(row3),
			Title:       "仪表盘",
			Description: "仪表盘",
		}
	})
}
