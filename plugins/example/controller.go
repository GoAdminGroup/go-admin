package example

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/page"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/components/chart_legend"
	"github.com/GoAdminGroup/themes/adminlte/components/description"
	"github.com/GoAdminGroup/themes/adminlte/components/infobox"
	"github.com/GoAdminGroup/themes/adminlte/components/productlist"
	"github.com/GoAdminGroup/themes/adminlte/components/progress_group"
	"github.com/GoAdminGroup/themes/adminlte/components/smallbox"
)

func (e *Example) TestHandler(ctx *context.Context) {
	page.SetPageContent(ctx, auth.Auth(ctx), func(ctx interface{}) (types.Panel, error) {

		components := template2.Default()
		colComp := components.Col()

		/**************************
		 * Info Box
		/**************************/

		infobox1 := infobox.New().
			SetText("CPU TRAFFIC").
			SetColor("#3583af").
			SetNumber("100").
			SetIcon(`<svg t="1568904058859" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2216" width="48" height="48"><path d="M864 64l-704 0C142.336 64 128 78.336 128 96l0 832C128 945.664 142.336 960 160 960l704 0c17.664 0 32-14.336 32-32l0-832C896 78.336 881.664 64 864 64zM832 896 192 896 192 128l640 0L832 896z" fill="#e6e6e6" p-id="2217"></path><path d="M353.92 320c17.6 0 32-14.336 32-32S371.584 256 353.92 256L353.28 256C335.616 256 321.6 270.336 321.6 288S336.256 320 353.92 320z" fill="#e6e6e6" p-id="2218"></path><path d="M353.92 512c17.6 0 32-14.336 32-32S371.584 448 353.92 448L353.28 448C335.616 448 321.6 462.336 321.6 480S336.256 512 353.92 512z" fill="#e6e6e6" p-id="2219"></path><path d="M353.92 704c17.6 0 32-14.336 32-32S371.584 640 353.92 640L353.28 640c-17.6 0-31.616 14.336-31.616 32S336.256 704 353.92 704z" fill="#e6e6e6" p-id="2220"></path><path d="M480 320l192 0C689.664 320 704 305.664 704 288S689.664 256 672 256l-192 0C462.336 256 448 270.336 448 288S462.336 320 480 320z" fill="#e6e6e6" p-id="2221"></path><path d="M480 512l192 0C689.664 512 704 497.664 704 480S689.664 448 672 448l-192 0C462.336 448 448 462.336 448 480S462.336 512 480 512z" fill="#e6e6e6" p-id="2222"></path><path d="M480 704l192 0c17.664 0 32-14.336 32-32S689.664 640 672 640l-192 0C462.336 640 448 654.336 448 672S462.336 704 480 704z" fill="#e6e6e6" p-id="2223"></path></svg>`).
			GetContent()

		infobox2 := infobox.New().
			SetText("Likes").
			SetColor("#6a7c86").
			SetNumber("1030.00<small>$</small>").
			SetIcon(`<svg t="1570468923385" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1124" width="48" height="48"><path d="M508.416 104.96c-225.28 0-408.064 182.784-408.064 408.064s182.784 408.064 408.064 408.064 408.064-182.784 408.064-408.064c0-108.032-43.008-211.968-119.808-288.768-76.288-76.288-180.224-119.296-288.256-119.296z m120.32 460.8c16.384 0 30.208 13.312 30.208 30.208 0 16.384-13.312 30.208-30.208 30.208h-90.624V716.8c0 16.384-13.312 30.208-30.208 30.208-16.384 0-30.208-13.312-30.208-30.208v-91.136H387.584c-16.384 0-30.208-13.312-30.208-30.208 0-16.384 13.312-30.208 30.208-30.208h90.624V495.104H387.584c-16.384 0-30.208-13.312-30.208-30.208 0-16.384 13.312-30.208 30.208-30.208h77.312L387.584 356.864c-9.216-11.776-8.192-28.672 2.56-39.424 10.752-10.752 27.648-11.776 39.424-2.56l78.848 78.848 77.312-77.312c11.264-11.264 29.696-11.264 41.472 0 11.264 11.264 11.264 29.696 0 41.472L548.864 435.2h79.36c16.384 0 30.208 13.312 30.208 30.208 0 16.384-13.312 30.208-30.208 30.208h-90.112v70.144h90.624z m0 0" fill="#ffffff" p-id="1125"></path></svg>`).
			GetContent()

		infobox3 := infobox.New().
			SetText("Sales").
			SetColor("#d8cd68").
			SetNumber("760").
			SetIcon(`<svg t="1570469111431" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3801" width="48" height="48"><path d="M298.666667 128v768h426.666666V128H298.666667zM256 85.333333h512v853.333334H256V85.333333zM170.666667 128H85.333333V85.333333h128v853.333334H85.333333v-42.666667h85.333334V128z m768 768v42.666667h-128V85.333333h128v42.666667h-85.333334v768h85.333334z" p-id="3802" fill="#ffffff"></path></svg>`).
			GetContent()

		infobox4 := infobox.New().
			SetText("New Members").
			SetColor("#6cad6e").
			SetNumber("2,349").
			SetIcon(`<svg t="1570469079555" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2965" width="48" height="48"><path d="M702.9 293.4c26.6 48.9 41.8 105 41.8 164.6 0 190.7-155 345.3-346.2 345.3S52.3 648.7 52.3 458s155-345.3 346.2-345.3c127.4 0 238.7 68.7 298.8 170.9-0.5-1.8-0.7-3.6-0.7-5.5 0-12.1 9.8-21.8 21.8-21.8 4.6 0 9.1 0.1 13.6 0.3C663.8 144.1 539.9 69 398.5 69 183.2 69 8.6 243.1 8.6 458c0 188 133.7 344.8 311.3 381.1 25.2-5.8 51.6-8.9 78.6-8.9 27 0 53.4 3.1 78.6 8.9C654.8 802.8 788.4 646 788.4 458c0-55.1-11.5-107.5-32.2-155-12.3-2-24.9-3.1-37.7-3.1-6.1 0-11.6-2.5-15.6-6.5z" p-id="2966" fill="#ffffff"></path><path d="M319.9 839.1c-68.4 15.8-128.2 51.9-167.7 102.3-7.4 9.5-5.8 23.2 3.7 30.7 9.5 7.4 23.2 5.8 30.7-3.7 45.5-58.1 124.6-94.4 211.8-94.4 88.3 0 168.4 37.3 213.5 96.6 7.3 9.6 21 11.5 30.6 4.2 9.6-7.3 11.5-21 4.2-30.6-39.4-51.8-100-88.9-169.7-105-25.2-5.8-51.6-8.9-78.6-8.9-26.9-0.1-53.3 3-78.5 8.8z" p-id="2967" fill="#ffffff"></path><path d="M732.1 256.6c-4.5-0.2-9.1-0.3-13.6-0.3-12.1 0-21.8 9.8-21.8 21.8 0 1.9 0.2 3.7 0.7 5.5 1 3.8 2.9 7.1 5.6 9.8 4 4 9.5 6.5 15.6 6.5 12.8 0 25.4 1.1 37.7 3.1 132 21.6 229.6 153.8 215.3 290.1-15.7 149.4-146.3 258-291.4 242.7-12-1.3-22.8 7.4-24 19.4-0.1 0.5-0.1 1.1-0.1 1.6-0.1 0.5-0.2 1-0.2 1.6-1.3 12 7.4 22.8 19.4 24 66.7 7 124.1 42.3 153.3 91.9 6.1 10.4 19.5 13.9 29.9 7.7 10.4-6.1 13.9-19.5 7.7-29.9-19.5-33.1-48.6-60.8-83.8-80.7 122.3-31.5 218.3-138.3 232.6-273.8 17.8-169.3-112-332.7-282.9-341z" p-id="2968" fill="#ffffff"></path></svg>`).
			GetContent()

		var size = types.Size(6, 3, 0).XS(12)
		infoboxCol1 := colComp.SetSize(size).SetContent(infobox1).GetContent()
		infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
		infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
		infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()
		row1 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4).GetContent()

		/**************************
		 * Box
		/**************************/

		table := components.Table().SetInfoList([]map[string]types.InfoItem{
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
			SetFooter(`<div class="clearfix"><a href="javascript:void(0)" class="btn btn-sm btn-info btn-flat pull-left">处理订单</a><a href="javascript:void(0)" class="btn btn-sm btn-default btn-flat pull-right">查看所有新订单</a> </div>`).
			GetContent()

		tableCol := colComp.SetSize(types.SizeMD(8)).SetContent(row1 + boxInfo).GetContent()

		/**************************
		 * Product List
		/**************************/

		productList := productlist.New().SetData([]map[string]string{
			{
				"img":         "http://adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
				"title":       "GoAdmin",
				"has_tabel":   "true",
				"labeltype":   "warning",
				"label":       "free",
				"description": `a framework help you build the dataviz system`,
			}, {
				"img":         "http://adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
				"title":       "GoAdmin",
				"has_tabel":   "true",
				"labeltype":   "warning",
				"label":       "free",
				"description": `a framework help you build the dataviz system`,
			}, {
				"img":         "http://adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
				"title":       "GoAdmin",
				"has_tabel":   "true",
				"labeltype":   "warning",
				"label":       "free",
				"description": `a framework help you build the dataviz system`,
			},
		}).GetContent()

		boxWarning := components.Box().SetTheme("warning").WithHeadBorder().SetHeader("Recently Added Products").
			SetBody(productList).
			SetFooter(`<a href="javascript:void(0)" class="uppercase">View All Products</a>`).
			GetContent()

		newsCol := colComp.SetSize(types.SizeMD(4)).SetContent(boxWarning).GetContent()

		row5 := components.Row().SetContent(tableCol + newsCol).GetContent()

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

		title := `<p class="text-center"><strong>Goal Completion</strong></p>`
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

		size2 := types.SizeXS(6).SM(3)
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

		col5 := colComp.SetSize(types.SizeMD(8)).SetContent(tabs + template.HTML(buttonTest)).GetContent()
		col6 := colComp.SetSize(types.SizeMD(4)).SetContent(boxDanger + popup).GetContent()

		row4 := components.Row().SetContent(col5 + col6).GetContent()

		return types.Panel{
			Content:     row5 + row2 + row3 + row4,
			Title:       "Dashboard",
			Description: "dashboard example",
		}, nil
	}, e.Conn)
}
