package ecommerce

import (
	"github.com/chenhg5/go-admin/modules/config"
	template2 "github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

func GetContent() types.Panel {

	components := template2.Get(config.Get().THEME)
	colComp := components.Col()

	/**************************
	 * Info Box
	/**************************/

	infobox := components.InfoBox().
		SetText("支付订单数").
		SetColor("blue").
		SetNumber("100").
		SetIcon(`<svg t="1568904058859" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2216" width="48" height="48"><path d="M864 64l-704 0C142.336 64 128 78.336 128 96l0 832C128 945.664 142.336 960 160 960l704 0c17.664 0 32-14.336 32-32l0-832C896 78.336 881.664 64 864 64zM832 896 192 896 192 128l640 0L832 896z" fill="#e6e6e6" p-id="2217"></path><path d="M353.92 320c17.6 0 32-14.336 32-32S371.584 256 353.92 256L353.28 256C335.616 256 321.6 270.336 321.6 288S336.256 320 353.92 320z" fill="#e6e6e6" p-id="2218"></path><path d="M353.92 512c17.6 0 32-14.336 32-32S371.584 448 353.92 448L353.28 448C335.616 448 321.6 462.336 321.6 480S336.256 512 353.92 512z" fill="#e6e6e6" p-id="2219"></path><path d="M353.92 704c17.6 0 32-14.336 32-32S371.584 640 353.92 640L353.28 640c-17.6 0-31.616 14.336-31.616 32S336.256 704 353.92 704z" fill="#e6e6e6" p-id="2220"></path><path d="M480 320l192 0C689.664 320 704 305.664 704 288S689.664 256 672 256l-192 0C462.336 256 448 270.336 448 288S462.336 320 480 320z" fill="#e6e6e6" p-id="2221"></path><path d="M480 512l192 0C689.664 512 704 497.664 704 480S689.664 448 672 448l-192 0C462.336 448 448 462.336 448 480S462.336 512 480 512z" fill="#e6e6e6" p-id="2222"></path><path d="M480 704l192 0c17.664 0 32-14.336 32-32S689.664 640 672 640l-192 0C462.336 640 448 654.336 448 672S462.336 704 480 704z" fill="#e6e6e6" p-id="2223"></path></svg>`).
		GetContent()

	infobox2 := components.InfoBox().
		SetText("支付金额").
		SetColor("red").
		SetNumber("1030.00<small>$</small>").
		SetIcon("fa-google-plus").
		GetContent()

	infobox3 := components.InfoBox().
		SetText("浏览量").
		SetColor("green").
		SetNumber("760").
		SetIcon("ion-ios-cart-outline").
		GetContent()

	infobox4 := components.InfoBox().
		SetText("累积客户数").
		SetColor("yellow").
		SetNumber("2,349").
		SetIcon("ion-ios-people-outline").
		GetContent()

	var size = map[string]string{"md": "3", "sm": "6", "xs": "12"}
	infoboxCol1 := colComp.SetSize(size).SetContent(infobox).GetContent()
	infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
	infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
	infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()
	row1 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4).GetContent()

	/**************************
	 * Info Box
	/**************************/

	infobox5 := components.InfoBox().
		SetText("本月目标(元)").
		SetColor("blue").
		SetNumber("--").
		SetIcon("ion-ios-gear-outline").
		GetContent()

	infobox6 := components.InfoBox().
		SetText("完成进度(%)").
		SetColor("red").
		SetNumber("--").
		SetIcon("fa-google-plus").
		GetContent()

	infobox7 := components.InfoBox().
		SetText("可用店铺余额(元)").
		SetColor("green").
		SetNumber("0.00").
		SetIcon("ion-ios-cart-outline").
		GetContent()

	infobox8 := components.InfoBox().
		SetText("待结算(元)").
		SetColor("yellow").
		SetNumber("0.00").
		SetIcon("ion-ios-people-outline").
		GetContent()

	infoboxCol5 := colComp.SetSize(size).SetContent(infobox5).GetContent()
	infoboxCol6 := colComp.SetSize(size).SetContent(infobox6).GetContent()
	infoboxCol7 := colComp.SetSize(size).SetContent(infobox7).GetContent()
	infoboxCol8 := colComp.SetSize(size).SetContent(infobox8).GetContent()
	row5 := components.Row().SetContent(infoboxCol5 + infoboxCol6 + infoboxCol7 + infoboxCol8).GetContent()

	/**************************
	 * Box
	/**************************/

	chartdata := `{"datasets":[{"data":[65,59,80,81,56,55,40],"fillColor":"rgb(210, 214, 222)","label":"Electronics","pointColor":"rgb(210, 214, 222)","pointHighlightFill":"#fff","pointHighlightStroke":"rgb(220,220,220)","pointStrokeColor":"#c1c7d1","strokeColor":"rgb(210, 214, 222)"},{"data":[28,48,40,19,86,27,90],"fillColor":"rgba(60,141,188,0.9)","label":"Digital Goods","pointColor":"#3b8bba","pointHighlightFill":"#fff","pointHighlightStroke":"rgba(60,141,188,1)","pointStrokeColor":"rgba(60,141,188,1)","strokeColor":"rgba(60,141,188,0.8)"}],"labels":["January","February","March","April","May","June","July"]}`

	lineChart := components.AreaChart().SetID("salechart").
		SetData(chartdata).
		SetHeight(180).
		SetTitle("Sales: 1 Jan, 2014 - 30 Jul, 2014").GetContent()

	title := `<p class="text-center"><strong>Goal Completion</strong></p>`
	progressGroup := components.ProgressGroup().
		SetTitle("Add Products to Cart").
		SetColor("aqua").
		SetDenominator(200).
		SetMolecular(160).
		SetPercent(80).
		GetContent()

	progressGroup1 := components.ProgressGroup().
		SetTitle("Complete Purchase").
		SetColor("red").
		SetDenominator(400).
		SetMolecular(310).
		SetPercent(80).
		GetContent()

	progressGroup2 := components.ProgressGroup().
		SetTitle("Visit Premium Page").
		SetColor("green").
		SetDenominator(800).
		SetMolecular(490).
		SetPercent(80).
		GetContent()

	progressGroup3 := components.ProgressGroup().
		SetTitle("Send Inquiries").
		SetColor("yellow").
		SetDenominator(500).
		SetMolecular(250).
		SetPercent(50).
		GetContent()

	boxInternalCol1 := colComp.SetContent(lineChart).SetSize(map[string]string{"md": "8"}).GetContent()
	boxInternalCol2 := colComp.
		SetContent(template.HTML(title) + progressGroup + progressGroup1 + progressGroup2 + progressGroup3).
		SetSize(map[string]string{"md": "4"}).
		GetContent()

	boxInternalRow := components.Row().SetContent(boxInternalCol1 + boxInternalCol2).GetContent()

	description1 := components.Description().SetPercent("17").
		SetNumber("¥100,000").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	description2 := components.Description().
		SetPercent("17").
		SetNumber("¥100,000").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	description3 := components.Description().
		SetPercent("17").
		SetNumber("¥100,000").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	description4 := components.Description().
		SetPercent("17").
		SetNumber("¥100,000").
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

	boxDanger := components.Box().SetTheme("danger").WithHeadBorder(true).SetHeader("Browser Usage").
		SetBody(components.Row().
			SetContent(colComp.SetSize(map[string]string{"md": "8"}).
				SetContent(pie).
				GetContent() + colComp.SetSize(map[string]string{"md": "4"}).
				SetContent(legend).
				GetContent()).GetContent()).
		SetFooter(`<p class="text-center"><a href="javascript:void(0)" class="uppercase">View All Users</a></p>`).
		GetContent()

	tabs := components.Tabs().SetData([]map[string]template.HTML{
		{
			"title": "tabs1",
			"content": template.HTML(`<b>How to use:</b>

                <p>Exactly like the original bootstrap tabs except you should use
                  the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
                A wonderful serenity has taken possession of my entire soul,
                like these sweet mornings of spring which I enjoy with my whole heart.
                I am alone, and feel the charm of existence in this spot,
                which was created for the bliss of souls like mine. I am so happy,
                my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
                that I neglect my talents. I should be incapable of drawing a single stroke
                at the present moment; and yet I feel that I never was a greater artist than now.`),
		}, {
			"title": "tabs2",
			"content": template.HTML(`
                The European languages are members of the same family. Their separate existence is a myth.
                For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
                in their grammar, their pronunciation and their most common words. Everyone realizes why a
                new common language would be desirable: one could refuse to pay expensive translators. To
                achieve this, it would be necessary to have uniform grammar, pronunciation and more common
                words. If several languages coalesce, the grammar of the resulting language is more simple
                and regular than that of the individual languages.
              `),
		}, {
			"title": "tabs3",
			"content": template.HTML(`
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                It has survived not only five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
                like Aldus PageMaker including versions of Lorem Ipsum.
              `),
		},
	}).GetContent()

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

	buttonTest := `<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#exampleModal" data-whatever="@mdo">Open modal for @mdo</button>`
	popupForm := `<form>
          <div class="form-group">
            <label for="recipient-name" class="col-form-label">Recipient:</label>
            <input type="text" class="form-control" id="recipient-name">
          </div>
          <div class="form-group">
            <label for="message-text" class="col-form-label">Message:</label>
            <textarea class="form-control" id="message-text"></textarea>
          </div>
        </form>`
	popup := components.Popup().SetID("exampleModal").
		SetFooter("Save Change").
		SetTitle("this is a popup").
		SetBody(template.HTML(popupForm)).
		GetContent()

	col5 := colComp.SetSize(map[string]string{"md": "8"}).SetContent(boxInfo + tabs + template.HTML(buttonTest)).GetContent()
	col6 := colComp.SetSize(map[string]string{"md": "4"}).SetContent(boxDanger + boxWarning + popup).GetContent()

	row4 := components.Row().SetContent(col5 + col6).GetContent()

	return types.Panel{
		Content:     row1 + row5 + row2 + row3 + row4,
		Title:       "概况",
		Description: "今日概况",
	}
}
