package tmpl

var List = map[string]string{"admin_panel": `{{define "admin_panel"}}
<div class="navbar-custom-menu">
    <ul class="nav navbar-nav">
        <!-- User Account: style can be found in dropdown.less -->
        <li class="dropdown user user-menu">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                {{if eq .User.Avatar ""}}
                    <img src="{{.AssertRootUrl}}/assets/dist/img/avatar04.png" class="user-image" alt="User Image">
                {{else}}
                    <img src="{{.User.Avatar}}" class="user-image" alt="User Image">
                {{end}}
                <span class="hidden-xs">{{.User.Name}}</span>
            </a>
            <ul class="dropdown-menu">
                <!-- User image -->
                <li class="user-header">
                    {{if eq .User.Avatar ""}}
                        <img src="{{.AssertRootUrl}}/assets/dist/img/avatar04.png" class="img-circle" alt="User Image">
                    {{else}}
                        <img src="{{.User.Avatar}}" class="img-circle" alt="User Image">
                    {{end}}
                    <p>
                        {{.User.Name}} -{{.User.Level}}
                        <small>{{.User.CreateAt}}</small>
                    </p>
                </li>
                <li class="user-footer">
                    <div class="pull-left">
                        <a href="{{.AssertRootUrl}}/info/manager/edit?id={{.User.ID}}" class="btn btn-default btn-flat">{{lang "setting"}}</a>
                    </div>
                    <div class="pull-right">
                        <a href="{{.AssertRootUrl}}/logout" class="btn btn-default btn-flat">{{lang "sign out"}}</a>
                    </div>
                </li>
            </ul>
        </li>
    </ul>
</div>
{{end}}`, "components/alert": `{{define "alert"}}
<div class="alert alert-{{.Theme}} alert-dismissible">
    <button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
    <h4>{{.Title}}</h4>
    {{.Content}}
</div>
{{end}}`, "components/area-chart": `{{define "area-chart"}}
{{if ne .Title ""}}
<p class="text-center">
    <strong>{{.Title}}</strong>
</p>
{{end}}
<div class="chart">
    <canvas id="{{.ID}}" style="height: {{.Height}}px;"></canvas>
</div>
<script>
    // Get context with jQuery - using jQuery's .get() method.
    let salesChartCanvas = $('#{{.ID}}').get(0).getContext('2d');
    // This will get the first returned node in the jQuery collection.
    let salesChart       = new Chart(salesChartCanvas);

    let salesChartData = JSON.parse({{.Data}});

    let salesChartOptions = {
        // Boolean - If we should show the scale at all
        showScale               : true,
        // Boolean - Whether grid lines are shown across the chart
        scaleShowGridLines      : false,
        // String - Colour of the grid lines
        scaleGridLineColor      : 'rgba(0,0,0,.05)',
        // Number - Width of the grid lines
        scaleGridLineWidth      : 1,
        // Boolean - Whether to show horizontal lines (except X axis)
        scaleShowHorizontalLines: true,
        // Boolean - Whether to show vertical lines (except Y axis)
        scaleShowVerticalLines  : true,
        // Boolean - Whether the line is curved between points
        bezierCurve             : true,
        // Number - Tension of the bezier curve between points
        bezierCurveTension      : 0.3,
        // Boolean - Whether to show a dot for each point
        pointDot                : false,
        // Number - Radius of each point dot in pixels
        pointDotRadius          : 4,
        // Number - Pixel width of point dot stroke
        pointDotStrokeWidth     : 1,
        // Number - amount extra to add to the radius to cater for hit detection outside the drawn point
        pointHitDetectionRadius : 20,
        // Boolean - Whether to show a stroke for datasets
        datasetStroke           : true,
        // Number - Pixel width of dataset stroke
        datasetStrokeWidth      : 2,
        // Boolean - Whether to fill the dataset with a color
        datasetFill             : true,
        // String - A legend template
        legendTemplate          : '<ul class=\'<%=name.toLowerCase()%>-legend\'><% for (var i=0; i<datasets.length; i++){%><li><span style=\'background-color:<%=datasets[i].lineColor%>\'></span><%=datasets[i].label%></li><%}%></ul>',
        // Boolean - whether to maintain the starting aspect ratio or not when responsive, if set to false, will take up entire container
        maintainAspectRatio     : true,
        // Boolean - whether to make the chart responsive to window resizing
        responsive              : true
    };

    // Create the line chart
    salesChart.Line(salesChartData, salesChartOptions);
</script>
{{end}}`, "components/bar-chart": `{{define "bar-chart"}}
{{if ne .Title ""}}
<p class="text-center">
    <strong>{{.Title}}</strong>
</p>
{{end}}
<div class="chart">
    <canvas id="{{.ID}}" style="width: {{.Width}}px;"></canvas>
</div>
<script>
    let barChartCanvas                   = $('#{{.ID}}').get(0).getContext('2d');
    let barChart                         = new Chart(barChartCanvas);
    let barChartData                     = JSON.parse({{.Data}});
    barChartData.datasets[1].fillColor   = '#00a65a';
    barChartData.datasets[1].strokeColor = '#00a65a';
    barChartData.datasets[1].pointColor  = '#00a65a';
    let barChartOptions                  = {
        //Boolean - Whether the scale should start at zero, or an order of magnitude down from the lowest value
        scaleBeginAtZero        : true,
        //Boolean - Whether grid lines are shown across the chart
        scaleShowGridLines      : true,
        //String - Colour of the grid lines
        scaleGridLineColor      : 'rgba(0,0,0,.05)',
        //Number - Width of the grid lines
        scaleGridLineWidth      : 1,
        //Boolean - Whether to show horizontal lines (except X axis)
        scaleShowHorizontalLines: true,
        //Boolean - Whether to show vertical lines (except Y axis)
        scaleShowVerticalLines  : true,
        //Boolean - If there is a stroke on each bar
        barShowStroke           : true,
        //Number - Pixel width of the bar stroke
        barStrokeWidth          : 2,
        //Number - Spacing between each of the X value sets
        barValueSpacing         : 5,
        //Number - Spacing between data sets within X values
        barDatasetSpacing       : 1,
        //String - A legend template
        legendTemplate          : '<ul class="<%=name.toLowerCase()%>-legend"><% for (var i=0; i<datasets.length; i++){%><li><span style="background-color:<%=datasets[i].fillColor%>"></span><%if(datasets[i].label){%><%=datasets[i].label%><%}%></li><%}%></ul>',
        //Boolean - whether to make the chart responsive
        responsive              : true,
        maintainAspectRatio     : true
    };

    barChartOptions.datasetFill = false;
    barChart.Bar(barChartData, barChartOptions)
</script>
{{end}}`, "components/box": `{{define "box"}}
<div class="box box-{{.Theme}}">
    <div class="box-header {{.HeadBorder}}">
        {{.Header}}
    </div>
    <div class="box-body">
        {{.Body}}
    </div>
    {{if ne .Footer ""}}
    <div class="box-footer clearfix">
        {{.Footer}}
    </div>
    {{end}}
</div>
{{end}}`, "components/chart-legend": `{{define "chart-legend"}}
<ul class="chart-legend clearfix">
    {{range $key, $data := .Data}}
        <li><i class="fa fa-circle-o text-{{index $data "color"}}"></i>{{index $data "label"}}</li>
    {{end}}
</ul>
{{end}}`, "components/col": `{{define "col"}}
<div class="{{.Size}}">{{.Content}}</div>
{{end}}`, "components/description": `{{define "description"}}
<div class="description-block border-{{.Border}}">
    <span class="description-percentage text-{{.Color}}"><i class="fa fa-caret-{{.Arrow}}"></i>{{.Percent}}%</span>
    <h5 class="description-header">{{.Number}}</h5>
    <span class="description-text">{{.Title}}</span>
</div>
{{end}}`, "components/form/default": `{{define "form_default"}}
<label class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <div class="box box-solid box-default no-margin">
        <div class="box-body">
            {{.Value}}
        </div>
    </div>
</div>
{{end}}`, "components/form/file": `{{define "form_file"}}
<label for="{{.Field}}" class="col-sm-2  control-label">{{.Head}}</label>
<div class="col-sm-8">
    <input type="file" class="{{.Field}}" name="{{.Field}}" data-initial-preview="" data-initial-caption="{{.Value}}">
</div>
<script>
    $("input.{{.Field}}").fileinput({
        "overwriteInitial":true,
        "initialPreviewAsData":true,
        "browseLabel":"Browse",
        "showRemove":false,
        "showUpload":false,
        "deleteExtraData":{
            "avatar":"_file_del_",
            "_file_del_":"",
            "_token":"NVXb8lY1QT6oi7DMPFLwU9IlHiSON3C6isKIF1kD",
            "_method":"PUT"
        },
        "deleteUrl":"http:\/\/laravel-admin.org\/demo\/auth\/1",
        "allowedFileTypes":["image"]
    });
</script>
{{end}}`, "components/form/iconpicker": `{{define "form_iconpicker"}}
<label for="icon" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <div class="input-group iconpicker-container">
        <span class="input-group-addon"><i class="fa fa-bars"></i></span>
        <input style="width: 140px" type="text" id="icon" name="{{.Field}}" value="{{.Value}}"
               class="form-control icon iconpicker-element iconpicker-input"
               placeholder="{{lang "Input Icon"}}">
    </div>
    <span class="help-block">
        <i class="fa fa-info-circle"></i>&nbsp;For more icons please see <a href="http://fontawesome.io/icons/" target="_blank">http://fontawesome.io/icons/</a>
    </span>
    <div class="iconpicker-popover popover fade">
    <div class="arrow"></div>
    <div class="popover-title"><input type="search"
                                      class="form-control iconpicker-search"
                                      placeholder="Type to filter"></div>
    <div class="popover-content">
        <div class="iconpicker">
            <div class="iconpicker-items"><a role="button" href="#"
                                             class="iconpicker-item"
                                             title=".fa-500px"><i
                    class="fa fa-500px"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-adjust"><i
                    class="fa fa-adjust"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-adn"><i
                    class="fa fa-adn"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-align-center"><i
                    class="fa fa-align-center"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-align-justify"><i
                    class="fa fa-align-justify"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-align-left"><i
                    class="fa fa-align-left"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-align-right"><i
                    class="fa fa-align-right"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-amazon"><i
                    class="fa fa-amazon"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-ambulance"><i
                    class="fa fa-ambulance"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-american-sign-language-interpreting"><i
                    class="fa fa-american-sign-language-interpreting"></i></a><a
                    role="button" href="#" class="iconpicker-item"
                    title=".fa-anchor"><i class="fa fa-anchor"></i></a><a
                    role="button" href="#" class="iconpicker-item"
                    title=".fa-android"><i class="fa fa-android"></i></a><a
                    role="button" href="#" class="iconpicker-item"
                    title=".fa-angellist"><i class="fa fa-angellist"></i></a><a
                    role="button" href="#" class="iconpicker-item"
                    title=".fa-angle-double-down"><i
                    class="fa fa-angle-double-down"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-angle-double-left"><i
                    class="fa fa-angle-double-left"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-angle-double-right"><i
                    class="fa fa-angle-double-right"></i></a><a role="button"
                                                                href="#"
                                                                class="iconpicker-item"
                                                                title=".fa-angle-double-up"><i
                    class="fa fa-angle-double-up"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-angle-down"><i
                    class="fa fa-angle-down"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-angle-left"><i
                    class="fa fa-angle-left"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-angle-right"><i
                    class="fa fa-angle-right"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-angle-up"><i
                    class="fa fa-angle-up"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-apple"><i
                    class="fa fa-apple"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-archive"><i
                    class="fa fa-archive"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-area-chart"><i
                    class="fa fa-area-chart"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-arrow-circle-down"><i
                    class="fa fa-arrow-circle-down"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-arrow-circle-left"><i
                    class="fa fa-arrow-circle-left"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-arrow-circle-o-down"><i
                    class="fa fa-arrow-circle-o-down"></i></a><a role="button"
                                                                 href="#"
                                                                 class="iconpicker-item"
                                                                 title=".fa-arrow-circle-o-left"><i
                    class="fa fa-arrow-circle-o-left"></i></a><a role="button"
                                                                 href="#"
                                                                 class="iconpicker-item"
                                                                 title=".fa-arrow-circle-o-right"><i
                    class="fa fa-arrow-circle-o-right"></i></a><a role="button"
                                                                  href="#"
                                                                  class="iconpicker-item"
                                                                  title=".fa-arrow-circle-o-up"><i
                    class="fa fa-arrow-circle-o-up"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-arrow-circle-right"><i
                    class="fa fa-arrow-circle-right"></i></a><a role="button"
                                                                href="#"
                                                                class="iconpicker-item"
                                                                title=".fa-arrow-circle-up"><i
                    class="fa fa-arrow-circle-up"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-arrow-down"><i
                    class="fa fa-arrow-down"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-arrow-left"><i
                    class="fa fa-arrow-left"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-arrow-right"><i
                    class="fa fa-arrow-right"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-arrow-up"><i
                    class="fa fa-arrow-up"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-arrows"><i
                    class="fa fa-arrows"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-arrows-alt"><i
                    class="fa fa-arrows-alt"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-arrows-h"><i
                    class="fa fa-arrows-h"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-arrows-v"><i
                    class="fa fa-arrows-v"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-asl-interpreting"><i
                    class="fa fa-asl-interpreting"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-assistive-listening-systems"><i
                    class="fa fa-assistive-listening-systems"></i></a><a
                    role="button" href="#" class="iconpicker-item"
                    title=".fa-asterisk"><i class="fa fa-asterisk"></i></a><a
                    role="button" href="#" class="iconpicker-item" title=".fa-at"><i
                    class="fa fa-at"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-audio-description"><i
                    class="fa fa-audio-description"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-automobile"><i
                    class="fa fa-automobile"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-backward"><i
                    class="fa fa-backward"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-balance-scale"><i
                    class="fa fa-balance-scale"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-ban"><i
                    class="fa fa-ban"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-bank"><i
                    class="fa fa-bank"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-bar-chart"><i
                    class="fa fa-bar-chart"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-bar-chart-o"><i
                    class="fa fa-bar-chart-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-barcode"><i
                    class="fa fa-barcode"></i></a><a role="button" href="#"
                                                     class="iconpicker-item iconpicker-selected bg-primary"
                                                     title=".fa-bars"><i
                    class="fa fa-bars"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-battery-0"><i
                    class="fa fa-battery-0"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-battery-1"><i
                    class="fa fa-battery-1"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-battery-2"><i
                    class="fa fa-battery-2"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-battery-3"><i
                    class="fa fa-battery-3"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-battery-4"><i
                    class="fa fa-battery-4"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-battery-empty"><i
                    class="fa fa-battery-empty"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-battery-full"><i
                    class="fa fa-battery-full"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-battery-half"><i
                    class="fa fa-battery-half"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-battery-quarter"><i
                    class="fa fa-battery-quarter"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-battery-three-quarters"><i
                    class="fa fa-battery-three-quarters"></i></a><a role="button"
                                                                    href="#"
                                                                    class="iconpicker-item"
                                                                    title=".fa-bed"><i
                    class="fa fa-bed"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-beer"><i
                    class="fa fa-beer"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-behance"><i
                    class="fa fa-behance"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-behance-square"><i
                    class="fa fa-behance-square"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-bell"><i
                    class="fa fa-bell"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-bell-o"><i
                    class="fa fa-bell-o"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-bell-slash"><i
                    class="fa fa-bell-slash"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-bell-slash-o"><i
                    class="fa fa-bell-slash-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-bicycle"><i
                    class="fa fa-bicycle"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-binoculars"><i
                    class="fa fa-binoculars"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-birthday-cake"><i
                    class="fa fa-birthday-cake"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-bitbucket"><i
                    class="fa fa-bitbucket"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-bitbucket-square"><i
                    class="fa fa-bitbucket-square"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-bitcoin"><i
                    class="fa fa-bitcoin"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-black-tie"><i
                    class="fa fa-black-tie"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-blind"><i
                    class="fa fa-blind"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-bluetooth"><i
                    class="fa fa-bluetooth"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-bluetooth-b"><i
                    class="fa fa-bluetooth-b"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-bold"><i
                    class="fa fa-bold"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-bolt"><i
                    class="fa fa-bolt"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-bomb"><i
                    class="fa fa-bomb"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-book"><i
                    class="fa fa-book"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-bookmark"><i
                    class="fa fa-bookmark"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-bookmark-o"><i
                    class="fa fa-bookmark-o"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-braille"><i
                    class="fa fa-braille"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-briefcase"><i
                    class="fa fa-briefcase"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-btc"><i
                    class="fa fa-btc"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-bug"><i
                    class="fa fa-bug"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-building"><i
                    class="fa fa-building"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-building-o"><i
                    class="fa fa-building-o"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-bullhorn"><i
                    class="fa fa-bullhorn"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-bullseye"><i
                    class="fa fa-bullseye"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-bus"><i
                    class="fa fa-bus"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-buysellads"><i
                    class="fa fa-buysellads"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-cab"><i
                    class="fa fa-cab"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-calculator"><i
                    class="fa fa-calculator"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-calendar"><i
                    class="fa fa-calendar"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-calendar-check-o"><i
                    class="fa fa-calendar-check-o"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-calendar-minus-o"><i
                    class="fa fa-calendar-minus-o"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-calendar-o"><i
                    class="fa fa-calendar-o"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-calendar-plus-o"><i
                    class="fa fa-calendar-plus-o"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-calendar-times-o"><i
                    class="fa fa-calendar-times-o"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-camera"><i
                    class="fa fa-camera"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-camera-retro"><i
                    class="fa fa-camera-retro"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-car"><i
                    class="fa fa-car"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-caret-down"><i
                    class="fa fa-caret-down"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-caret-left"><i
                    class="fa fa-caret-left"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-caret-right"><i
                    class="fa fa-caret-right"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-caret-square-o-down"><i
                    class="fa fa-caret-square-o-down"></i></a><a role="button"
                                                                 href="#"
                                                                 class="iconpicker-item"
                                                                 title=".fa-caret-square-o-left"><i
                    class="fa fa-caret-square-o-left"></i></a><a role="button"
                                                                 href="#"
                                                                 class="iconpicker-item"
                                                                 title=".fa-caret-square-o-right"><i
                    class="fa fa-caret-square-o-right"></i></a><a role="button"
                                                                  href="#"
                                                                  class="iconpicker-item"
                                                                  title=".fa-caret-square-o-up"><i
                    class="fa fa-caret-square-o-up"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-caret-up"><i
                    class="fa fa-caret-up"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-cart-arrow-down"><i
                    class="fa fa-cart-arrow-down"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-cart-plus"><i
                    class="fa fa-cart-plus"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-cc"><i
                    class="fa fa-cc"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-cc-amex"><i
                    class="fa fa-cc-amex"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-cc-diners-club"><i
                    class="fa fa-cc-diners-club"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-cc-discover"><i
                    class="fa fa-cc-discover"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-cc-jcb"><i
                    class="fa fa-cc-jcb"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-cc-mastercard"><i
                    class="fa fa-cc-mastercard"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-cc-paypal"><i
                    class="fa fa-cc-paypal"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-cc-stripe"><i
                    class="fa fa-cc-stripe"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-cc-visa"><i
                    class="fa fa-cc-visa"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-certificate"><i
                    class="fa fa-certificate"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-chain"><i
                    class="fa fa-chain"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-chain-broken"><i
                    class="fa fa-chain-broken"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-check"><i
                    class="fa fa-check"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-check-circle"><i
                    class="fa fa-check-circle"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-check-circle-o"><i
                    class="fa fa-check-circle-o"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-check-square"><i
                    class="fa fa-check-square"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-check-square-o"><i
                    class="fa fa-check-square-o"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-chevron-circle-down"><i
                    class="fa fa-chevron-circle-down"></i></a><a role="button"
                                                                 href="#"
                                                                 class="iconpicker-item"
                                                                 title=".fa-chevron-circle-left"><i
                    class="fa fa-chevron-circle-left"></i></a><a role="button"
                                                                 href="#"
                                                                 class="iconpicker-item"
                                                                 title=".fa-chevron-circle-right"><i
                    class="fa fa-chevron-circle-right"></i></a><a role="button"
                                                                  href="#"
                                                                  class="iconpicker-item"
                                                                  title=".fa-chevron-circle-up"><i
                    class="fa fa-chevron-circle-up"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-chevron-down"><i
                    class="fa fa-chevron-down"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-chevron-left"><i
                    class="fa fa-chevron-left"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-chevron-right"><i
                    class="fa fa-chevron-right"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-chevron-up"><i
                    class="fa fa-chevron-up"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-child"><i
                    class="fa fa-child"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-chrome"><i
                    class="fa fa-chrome"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-circle"><i
                    class="fa fa-circle"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-circle-o"><i
                    class="fa fa-circle-o"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-circle-o-notch"><i
                    class="fa fa-circle-o-notch"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-circle-thin"><i
                    class="fa fa-circle-thin"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-clipboard"><i
                    class="fa fa-clipboard"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-clock-o"><i
                    class="fa fa-clock-o"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-clone"><i
                    class="fa fa-clone"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-close"><i
                    class="fa fa-close"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-cloud"><i
                    class="fa fa-cloud"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-cloud-download"><i
                    class="fa fa-cloud-download"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-cloud-upload"><i
                    class="fa fa-cloud-upload"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-cny"><i
                    class="fa fa-cny"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-code"><i
                    class="fa fa-code"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-code-fork"><i
                    class="fa fa-code-fork"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-codepen"><i
                    class="fa fa-codepen"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-codiepie"><i
                    class="fa fa-codiepie"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-coffee"><i
                    class="fa fa-coffee"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-cog"><i
                    class="fa fa-cog"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-cogs"><i
                    class="fa fa-cogs"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-columns"><i
                    class="fa fa-columns"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-comment"><i
                    class="fa fa-comment"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-comment-o"><i
                    class="fa fa-comment-o"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-commenting"><i
                    class="fa fa-commenting"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-commenting-o"><i
                    class="fa fa-commenting-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-comments"><i
                    class="fa fa-comments"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-comments-o"><i
                    class="fa fa-comments-o"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-compass"><i
                    class="fa fa-compass"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-compress"><i
                    class="fa fa-compress"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-connectdevelop"><i
                    class="fa fa-connectdevelop"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-contao"><i
                    class="fa fa-contao"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-copy"><i
                    class="fa fa-copy"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-copyright"><i
                    class="fa fa-copyright"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-creative-commons"><i
                    class="fa fa-creative-commons"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-credit-card"><i
                    class="fa fa-credit-card"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-credit-card-alt"><i
                    class="fa fa-credit-card-alt"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-crop"><i
                    class="fa fa-crop"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-crosshairs"><i
                    class="fa fa-crosshairs"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-css3"><i
                    class="fa fa-css3"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-cube"><i
                    class="fa fa-cube"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-cubes"><i
                    class="fa fa-cubes"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-cut"><i
                    class="fa fa-cut"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-cutlery"><i
                    class="fa fa-cutlery"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-dashboard"><i
                    class="fa fa-dashboard"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-dashcube"><i
                    class="fa fa-dashcube"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-database"><i
                    class="fa fa-database"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-deaf"><i
                    class="fa fa-deaf"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-deafness"><i
                    class="fa fa-deafness"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-dedent"><i
                    class="fa fa-dedent"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-delicious"><i
                    class="fa fa-delicious"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-desktop"><i
                    class="fa fa-desktop"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-deviantart"><i
                    class="fa fa-deviantart"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-diamond"><i
                    class="fa fa-diamond"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-digg"><i
                    class="fa fa-digg"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-dollar"><i
                    class="fa fa-dollar"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-dot-circle-o"><i
                    class="fa fa-dot-circle-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-download"><i
                    class="fa fa-download"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-dribbble"><i
                    class="fa fa-dribbble"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-dropbox"><i
                    class="fa fa-dropbox"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-drupal"><i
                    class="fa fa-drupal"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-edge"><i
                    class="fa fa-edge"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-edit"><i
                    class="fa fa-edit"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-eject"><i
                    class="fa fa-eject"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-ellipsis-h"><i
                    class="fa fa-ellipsis-h"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-ellipsis-v"><i
                    class="fa fa-ellipsis-v"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-empire"><i
                    class="fa fa-empire"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-envelope"><i
                    class="fa fa-envelope"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-envelope-o"><i
                    class="fa fa-envelope-o"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-envelope-square"><i
                    class="fa fa-envelope-square"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-envira"><i
                    class="fa fa-envira"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-eraser"><i
                    class="fa fa-eraser"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-eur"><i
                    class="fa fa-eur"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-euro"><i
                    class="fa fa-euro"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-exchange"><i
                    class="fa fa-exchange"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-exclamation"><i
                    class="fa fa-exclamation"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-exclamation-circle"><i
                    class="fa fa-exclamation-circle"></i></a><a role="button"
                                                                href="#"
                                                                class="iconpicker-item"
                                                                title=".fa-exclamation-triangle"><i
                    class="fa fa-exclamation-triangle"></i></a><a role="button"
                                                                  href="#"
                                                                  class="iconpicker-item"
                                                                  title=".fa-expand"><i
                    class="fa fa-expand"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-expeditedssl"><i
                    class="fa fa-expeditedssl"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-external-link"><i
                    class="fa fa-external-link"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-external-link-square"><i
                    class="fa fa-external-link-square"></i></a><a role="button"
                                                                  href="#"
                                                                  class="iconpicker-item"
                                                                  title=".fa-eye"><i
                    class="fa fa-eye"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-eye-slash"><i
                    class="fa fa-eye-slash"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-eyedropper"><i
                    class="fa fa-eyedropper"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-fa"><i
                    class="fa fa-fa"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-facebook"><i
                    class="fa fa-facebook"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-facebook-f"><i
                    class="fa fa-facebook-f"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-facebook-official"><i
                    class="fa fa-facebook-official"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-facebook-square"><i
                    class="fa fa-facebook-square"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-fast-backward"><i
                    class="fa fa-fast-backward"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-fast-forward"><i
                    class="fa fa-fast-forward"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-fax"><i
                    class="fa fa-fax"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-feed"><i
                    class="fa fa-feed"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-female"><i
                    class="fa fa-female"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-fighter-jet"><i
                    class="fa fa-fighter-jet"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-file"><i
                    class="fa fa-file"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-file-archive-o"><i
                    class="fa fa-file-archive-o"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-file-audio-o"><i
                    class="fa fa-file-audio-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-file-code-o"><i
                    class="fa fa-file-code-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-file-excel-o"><i
                    class="fa fa-file-excel-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-file-image-o"><i
                    class="fa fa-file-image-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-file-movie-o"><i
                    class="fa fa-file-movie-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-file-o"><i
                    class="fa fa-file-o"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-file-pdf-o"><i
                    class="fa fa-file-pdf-o"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-file-photo-o"><i
                    class="fa fa-file-photo-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-file-picture-o"><i
                    class="fa fa-file-picture-o"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-file-powerpoint-o"><i
                    class="fa fa-file-powerpoint-o"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-file-sound-o"><i
                    class="fa fa-file-sound-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-file-text"><i
                    class="fa fa-file-text"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-file-text-o"><i
                    class="fa fa-file-text-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-file-video-o"><i
                    class="fa fa-file-video-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-file-word-o"><i
                    class="fa fa-file-word-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-file-zip-o"><i
                    class="fa fa-file-zip-o"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-files-o"><i
                    class="fa fa-files-o"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-film"><i
                    class="fa fa-film"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-filter"><i
                    class="fa fa-filter"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-fire"><i
                    class="fa fa-fire"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-fire-extinguisher"><i
                    class="fa fa-fire-extinguisher"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-firefox"><i
                    class="fa fa-firefox"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-first-order"><i
                    class="fa fa-first-order"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-flag"><i
                    class="fa fa-flag"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-flag-checkered"><i
                    class="fa fa-flag-checkered"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-flag-o"><i
                    class="fa fa-flag-o"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-flash"><i
                    class="fa fa-flash"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-flask"><i
                    class="fa fa-flask"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-flickr"><i
                    class="fa fa-flickr"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-floppy-o"><i
                    class="fa fa-floppy-o"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-folder"><i
                    class="fa fa-folder"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-folder-o"><i
                    class="fa fa-folder-o"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-folder-open"><i
                    class="fa fa-folder-open"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-folder-open-o"><i
                    class="fa fa-folder-open-o"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-font"><i
                    class="fa fa-font"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-font-awesome"><i
                    class="fa fa-font-awesome"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-fonticons"><i
                    class="fa fa-fonticons"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-fort-awesome"><i
                    class="fa fa-fort-awesome"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-forumbee"><i
                    class="fa fa-forumbee"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-forward"><i
                    class="fa fa-forward"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-foursquare"><i
                    class="fa fa-foursquare"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-frown-o"><i
                    class="fa fa-frown-o"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-futbol-o"><i
                    class="fa fa-futbol-o"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-gamepad"><i
                    class="fa fa-gamepad"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-gavel"><i
                    class="fa fa-gavel"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-gbp"><i
                    class="fa fa-gbp"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-ge"><i
                    class="fa fa-ge"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-gear"><i
                    class="fa fa-gear"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-gears"><i
                    class="fa fa-gears"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-genderless"><i
                    class="fa fa-genderless"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-get-pocket"><i
                    class="fa fa-get-pocket"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-gg"><i
                    class="fa fa-gg"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-gg-circle"><i
                    class="fa fa-gg-circle"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-gift"><i
                    class="fa fa-gift"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-git"><i
                    class="fa fa-git"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-git-square"><i
                    class="fa fa-git-square"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-github"><i
                    class="fa fa-github"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-github-alt"><i
                    class="fa fa-github-alt"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-github-square"><i
                    class="fa fa-github-square"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-gitlab"><i
                    class="fa fa-gitlab"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-gittip"><i
                    class="fa fa-gittip"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-glass"><i
                    class="fa fa-glass"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-glide"><i
                    class="fa fa-glide"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-glide-g"><i
                    class="fa fa-glide-g"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-globe"><i
                    class="fa fa-globe"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-google"><i
                    class="fa fa-google"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-google-plus"><i
                    class="fa fa-google-plus"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-google-plus-circle"><i
                    class="fa fa-google-plus-circle"></i></a><a role="button"
                                                                href="#"
                                                                class="iconpicker-item"
                                                                title=".fa-google-plus-official"><i
                    class="fa fa-google-plus-official"></i></a><a role="button"
                                                                  href="#"
                                                                  class="iconpicker-item"
                                                                  title=".fa-google-plus-square"><i
                    class="fa fa-google-plus-square"></i></a><a role="button"
                                                                href="#"
                                                                class="iconpicker-item"
                                                                title=".fa-google-wallet"><i
                    class="fa fa-google-wallet"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-graduation-cap"><i
                    class="fa fa-graduation-cap"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-gratipay"><i
                    class="fa fa-gratipay"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-group"><i
                    class="fa fa-group"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-h-square"><i
                    class="fa fa-h-square"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-hacker-news"><i
                    class="fa fa-hacker-news"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hand-grab-o"><i
                    class="fa fa-hand-grab-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hand-lizard-o"><i
                    class="fa fa-hand-lizard-o"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-hand-o-down"><i
                    class="fa fa-hand-o-down"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hand-o-left"><i
                    class="fa fa-hand-o-left"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hand-o-right"><i
                    class="fa fa-hand-o-right"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-hand-o-up"><i
                    class="fa fa-hand-o-up"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-hand-paper-o"><i
                    class="fa fa-hand-paper-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-hand-peace-o"><i
                    class="fa fa-hand-peace-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-hand-pointer-o"><i
                    class="fa fa-hand-pointer-o"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-hand-rock-o"><i
                    class="fa fa-hand-rock-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hand-scissors-o"><i
                    class="fa fa-hand-scissors-o"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-hand-spock-o"><i
                    class="fa fa-hand-spock-o"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-hand-stop-o"><i
                    class="fa fa-hand-stop-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hard-of-hearing"><i
                    class="fa fa-hard-of-hearing"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-hashtag"><i
                    class="fa fa-hashtag"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-hdd-o"><i
                    class="fa fa-hdd-o"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-header"><i
                    class="fa fa-header"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-headphones"><i
                    class="fa fa-headphones"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-heart"><i
                    class="fa fa-heart"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-heart-o"><i
                    class="fa fa-heart-o"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-heartbeat"><i
                    class="fa fa-heartbeat"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-history"><i
                    class="fa fa-history"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-home"><i
                    class="fa fa-home"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-hospital-o"><i
                    class="fa fa-hospital-o"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-hotel"><i
                    class="fa fa-hotel"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-hourglass"><i
                    class="fa fa-hourglass"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-hourglass-1"><i
                    class="fa fa-hourglass-1"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hourglass-2"><i
                    class="fa fa-hourglass-2"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hourglass-3"><i
                    class="fa fa-hourglass-3"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hourglass-end"><i
                    class="fa fa-hourglass-end"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-hourglass-half"><i
                    class="fa fa-hourglass-half"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-hourglass-o"><i
                    class="fa fa-hourglass-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-hourglass-start"><i
                    class="fa fa-hourglass-start"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-houzz"><i
                    class="fa fa-houzz"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-html5"><i
                    class="fa fa-html5"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-i-cursor"><i
                    class="fa fa-i-cursor"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-ils"><i
                    class="fa fa-ils"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-image"><i
                    class="fa fa-image"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-inbox"><i
                    class="fa fa-inbox"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-indent"><i
                    class="fa fa-indent"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-industry"><i
                    class="fa fa-industry"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-info"><i
                    class="fa fa-info"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-info-circle"><i
                    class="fa fa-info-circle"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-inr"><i
                    class="fa fa-inr"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-instagram"><i
                    class="fa fa-instagram"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-institution"><i
                    class="fa fa-institution"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-internet-explorer"><i
                    class="fa fa-internet-explorer"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-intersex"><i
                    class="fa fa-intersex"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-ioxhost"><i
                    class="fa fa-ioxhost"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-italic"><i
                    class="fa fa-italic"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-joomla"><i
                    class="fa fa-joomla"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-jpy"><i
                    class="fa fa-jpy"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-jsfiddle"><i
                    class="fa fa-jsfiddle"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-key"><i
                    class="fa fa-key"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-keyboard-o"><i
                    class="fa fa-keyboard-o"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-krw"><i
                    class="fa fa-krw"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-language"><i
                    class="fa fa-language"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-laptop"><i
                    class="fa fa-laptop"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-lastfm"><i
                    class="fa fa-lastfm"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-lastfm-square"><i
                    class="fa fa-lastfm-square"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-leaf"><i
                    class="fa fa-leaf"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-leanpub"><i
                    class="fa fa-leanpub"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-legal"><i
                    class="fa fa-legal"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-lemon-o"><i
                    class="fa fa-lemon-o"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-level-down"><i
                    class="fa fa-level-down"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-level-up"><i
                    class="fa fa-level-up"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-life-bouy"><i
                    class="fa fa-life-bouy"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-life-buoy"><i
                    class="fa fa-life-buoy"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-life-ring"><i
                    class="fa fa-life-ring"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-life-saver"><i
                    class="fa fa-life-saver"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-lightbulb-o"><i
                    class="fa fa-lightbulb-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-line-chart"><i
                    class="fa fa-line-chart"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-link"><i
                    class="fa fa-link"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-linkedin"><i
                    class="fa fa-linkedin"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-linkedin-square"><i
                    class="fa fa-linkedin-square"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-linux"><i
                    class="fa fa-linux"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-list"><i
                    class="fa fa-list"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-list-alt"><i
                    class="fa fa-list-alt"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-list-ol"><i
                    class="fa fa-list-ol"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-list-ul"><i
                    class="fa fa-list-ul"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-location-arrow"><i
                    class="fa fa-location-arrow"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-lock"><i
                    class="fa fa-lock"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-long-arrow-down"><i
                    class="fa fa-long-arrow-down"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-long-arrow-left"><i
                    class="fa fa-long-arrow-left"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-long-arrow-right"><i
                    class="fa fa-long-arrow-right"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-long-arrow-up"><i
                    class="fa fa-long-arrow-up"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-low-vision"><i
                    class="fa fa-low-vision"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-magic"><i
                    class="fa fa-magic"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-magnet"><i
                    class="fa fa-magnet"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-mail-forward"><i
                    class="fa fa-mail-forward"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-mail-reply"><i
                    class="fa fa-mail-reply"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-mail-reply-all"><i
                    class="fa fa-mail-reply-all"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-male"><i
                    class="fa fa-male"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-map"><i
                    class="fa fa-map"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-map-marker"><i
                    class="fa fa-map-marker"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-map-o"><i
                    class="fa fa-map-o"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-map-pin"><i
                    class="fa fa-map-pin"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-map-signs"><i
                    class="fa fa-map-signs"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-mars"><i
                    class="fa fa-mars"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-mars-double"><i
                    class="fa fa-mars-double"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-mars-stroke"><i
                    class="fa fa-mars-stroke"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-mars-stroke-h"><i
                    class="fa fa-mars-stroke-h"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-mars-stroke-v"><i
                    class="fa fa-mars-stroke-v"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-maxcdn"><i
                    class="fa fa-maxcdn"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-meanpath"><i
                    class="fa fa-meanpath"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-medium"><i
                    class="fa fa-medium"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-medkit"><i
                    class="fa fa-medkit"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-meh-o"><i
                    class="fa fa-meh-o"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-mercury"><i
                    class="fa fa-mercury"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-microphone"><i
                    class="fa fa-microphone"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-microphone-slash"><i
                    class="fa fa-microphone-slash"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-minus"><i
                    class="fa fa-minus"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-minus-circle"><i
                    class="fa fa-minus-circle"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-minus-square"><i
                    class="fa fa-minus-square"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-minus-square-o"><i
                    class="fa fa-minus-square-o"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-mixcloud"><i
                    class="fa fa-mixcloud"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-mobile"><i
                    class="fa fa-mobile"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-mobile-phone"><i
                    class="fa fa-mobile-phone"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-modx"><i
                    class="fa fa-modx"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-money"><i
                    class="fa fa-money"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-moon-o"><i
                    class="fa fa-moon-o"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-mortar-board"><i
                    class="fa fa-mortar-board"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-motorcycle"><i
                    class="fa fa-motorcycle"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-mouse-pointer"><i
                    class="fa fa-mouse-pointer"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-music"><i
                    class="fa fa-music"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-navicon"><i
                    class="fa fa-navicon"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-neuter"><i
                    class="fa fa-neuter"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-newspaper-o"><i
                    class="fa fa-newspaper-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-object-group"><i
                    class="fa fa-object-group"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-object-ungroup"><i
                    class="fa fa-object-ungroup"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-odnoklassniki"><i
                    class="fa fa-odnoklassniki"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-odnoklassniki-square"><i
                    class="fa fa-odnoklassniki-square"></i></a><a role="button"
                                                                  href="#"
                                                                  class="iconpicker-item"
                                                                  title=".fa-opencart"><i
                    class="fa fa-opencart"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-openid"><i
                    class="fa fa-openid"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-opera"><i
                    class="fa fa-opera"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-optin-monster"><i
                    class="fa fa-optin-monster"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-outdent"><i
                    class="fa fa-outdent"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-pagelines"><i
                    class="fa fa-pagelines"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-paint-brush"><i
                    class="fa fa-paint-brush"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-paper-plane"><i
                    class="fa fa-paper-plane"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-paper-plane-o"><i
                    class="fa fa-paper-plane-o"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-paperclip"><i
                    class="fa fa-paperclip"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-paragraph"><i
                    class="fa fa-paragraph"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-paste"><i
                    class="fa fa-paste"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-pause"><i
                    class="fa fa-pause"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-pause-circle"><i
                    class="fa fa-pause-circle"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-pause-circle-o"><i
                    class="fa fa-pause-circle-o"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-paw"><i
                    class="fa fa-paw"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-paypal"><i
                    class="fa fa-paypal"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-pencil"><i
                    class="fa fa-pencil"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-pencil-square"><i
                    class="fa fa-pencil-square"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-pencil-square-o"><i
                    class="fa fa-pencil-square-o"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-percent"><i
                    class="fa fa-percent"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-phone"><i
                    class="fa fa-phone"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-phone-square"><i
                    class="fa fa-phone-square"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-photo"><i
                    class="fa fa-photo"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-picture-o"><i
                    class="fa fa-picture-o"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-pie-chart"><i
                    class="fa fa-pie-chart"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-pied-piper"><i
                    class="fa fa-pied-piper"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-pied-piper-alt"><i
                    class="fa fa-pied-piper-alt"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-pied-piper-pp"><i
                    class="fa fa-pied-piper-pp"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-pinterest"><i
                    class="fa fa-pinterest"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-pinterest-p"><i
                    class="fa fa-pinterest-p"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-pinterest-square"><i
                    class="fa fa-pinterest-square"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-plane"><i
                    class="fa fa-plane"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-play"><i
                    class="fa fa-play"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-play-circle"><i
                    class="fa fa-play-circle"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-play-circle-o"><i
                    class="fa fa-play-circle-o"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-plug"><i
                    class="fa fa-plug"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-plus"><i
                    class="fa fa-plus"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-plus-circle"><i
                    class="fa fa-plus-circle"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-plus-square"><i
                    class="fa fa-plus-square"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-plus-square-o"><i
                    class="fa fa-plus-square-o"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-power-off"><i
                    class="fa fa-power-off"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-print"><i
                    class="fa fa-print"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-product-hunt"><i
                    class="fa fa-product-hunt"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-puzzle-piece"><i
                    class="fa fa-puzzle-piece"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-qq"><i
                    class="fa fa-qq"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-qrcode"><i
                    class="fa fa-qrcode"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-question"><i
                    class="fa fa-question"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-question-circle"><i
                    class="fa fa-question-circle"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-question-circle-o"><i
                    class="fa fa-question-circle-o"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-quote-left"><i
                    class="fa fa-quote-left"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-quote-right"><i
                    class="fa fa-quote-right"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-ra"><i
                    class="fa fa-ra"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-random"><i
                    class="fa fa-random"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-rebel"><i
                    class="fa fa-rebel"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-recycle"><i
                    class="fa fa-recycle"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-reddit"><i
                    class="fa fa-reddit"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-reddit-alien"><i
                    class="fa fa-reddit-alien"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-reddit-square"><i
                    class="fa fa-reddit-square"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-refresh"><i
                    class="fa fa-refresh"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-registered"><i
                    class="fa fa-registered"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-remove"><i
                    class="fa fa-remove"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-renren"><i
                    class="fa fa-renren"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-reorder"><i
                    class="fa fa-reorder"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-repeat"><i
                    class="fa fa-repeat"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-reply"><i
                    class="fa fa-reply"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-reply-all"><i
                    class="fa fa-reply-all"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-resistance"><i
                    class="fa fa-resistance"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-retweet"><i
                    class="fa fa-retweet"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-rmb"><i
                    class="fa fa-rmb"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-road"><i
                    class="fa fa-road"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-rocket"><i
                    class="fa fa-rocket"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-rotate-left"><i
                    class="fa fa-rotate-left"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-rotate-right"><i
                    class="fa fa-rotate-right"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-rouble"><i
                    class="fa fa-rouble"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-rss"><i
                    class="fa fa-rss"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-rss-square"><i
                    class="fa fa-rss-square"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-rub"><i
                    class="fa fa-rub"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-ruble"><i
                    class="fa fa-ruble"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-rupee"><i
                    class="fa fa-rupee"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-safari"><i
                    class="fa fa-safari"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-save"><i
                    class="fa fa-save"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-scissors"><i
                    class="fa fa-scissors"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-scribd"><i
                    class="fa fa-scribd"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-search"><i
                    class="fa fa-search"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-search-minus"><i
                    class="fa fa-search-minus"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-search-plus"><i
                    class="fa fa-search-plus"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-sellsy"><i
                    class="fa fa-sellsy"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-send"><i
                    class="fa fa-send"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-send-o"><i
                    class="fa fa-send-o"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-server"><i
                    class="fa fa-server"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-share"><i
                    class="fa fa-share"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-share-alt"><i
                    class="fa fa-share-alt"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-share-alt-square"><i
                    class="fa fa-share-alt-square"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-share-square"><i
                    class="fa fa-share-square"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-share-square-o"><i
                    class="fa fa-share-square-o"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-shekel"><i
                    class="fa fa-shekel"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-sheqel"><i
                    class="fa fa-sheqel"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-shield"><i
                    class="fa fa-shield"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-ship"><i
                    class="fa fa-ship"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-shirtsinbulk"><i
                    class="fa fa-shirtsinbulk"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-shopping-bag"><i
                    class="fa fa-shopping-bag"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-shopping-basket"><i
                    class="fa fa-shopping-basket"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-shopping-cart"><i
                    class="fa fa-shopping-cart"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-sign-in"><i
                    class="fa fa-sign-in"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-sign-language"><i
                    class="fa fa-sign-language"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-sign-out"><i
                    class="fa fa-sign-out"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-signal"><i
                    class="fa fa-signal"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-signing"><i
                    class="fa fa-signing"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-simplybuilt"><i
                    class="fa fa-simplybuilt"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-sitemap"><i
                    class="fa fa-sitemap"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-skyatlas"><i
                    class="fa fa-skyatlas"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-skype"><i
                    class="fa fa-skype"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-slack"><i
                    class="fa fa-slack"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-sliders"><i
                    class="fa fa-sliders"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-slideshare"><i
                    class="fa fa-slideshare"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-smile-o"><i
                    class="fa fa-smile-o"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-snapchat"><i
                    class="fa fa-snapchat"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-snapchat-ghost"><i
                    class="fa fa-snapchat-ghost"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-snapchat-square"><i
                    class="fa fa-snapchat-square"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-soccer-ball-o"><i
                    class="fa fa-soccer-ball-o"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-sort"><i
                    class="fa fa-sort"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-sort-alpha-asc"><i
                    class="fa fa-sort-alpha-asc"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-sort-alpha-desc"><i
                    class="fa fa-sort-alpha-desc"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-sort-amount-asc"><i
                    class="fa fa-sort-amount-asc"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-sort-amount-desc"><i
                    class="fa fa-sort-amount-desc"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-sort-asc"><i
                    class="fa fa-sort-asc"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-sort-desc"><i
                    class="fa fa-sort-desc"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-sort-down"><i
                    class="fa fa-sort-down"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-sort-numeric-asc"><i
                    class="fa fa-sort-numeric-asc"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-sort-numeric-desc"><i
                    class="fa fa-sort-numeric-desc"></i></a><a role="button"
                                                               href="#"
                                                               class="iconpicker-item"
                                                               title=".fa-sort-up"><i
                    class="fa fa-sort-up"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-soundcloud"><i
                    class="fa fa-soundcloud"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-space-shuttle"><i
                    class="fa fa-space-shuttle"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-spinner"><i
                    class="fa fa-spinner"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-spoon"><i
                    class="fa fa-spoon"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-spotify"><i
                    class="fa fa-spotify"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-square"><i
                    class="fa fa-square"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-square-o"><i
                    class="fa fa-square-o"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-stack-exchange"><i
                    class="fa fa-stack-exchange"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-stack-overflow"><i
                    class="fa fa-stack-overflow"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-star"><i
                    class="fa fa-star"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-star-half"><i
                    class="fa fa-star-half"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-star-half-empty"><i
                    class="fa fa-star-half-empty"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-star-half-full"><i
                    class="fa fa-star-half-full"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-star-half-o"><i
                    class="fa fa-star-half-o"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-star-o"><i
                    class="fa fa-star-o"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-steam"><i
                    class="fa fa-steam"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-steam-square"><i
                    class="fa fa-steam-square"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-step-backward"><i
                    class="fa fa-step-backward"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-step-forward"><i
                    class="fa fa-step-forward"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-stethoscope"><i
                    class="fa fa-stethoscope"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-sticky-note"><i
                    class="fa fa-sticky-note"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-sticky-note-o"><i
                    class="fa fa-sticky-note-o"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-stop"><i
                    class="fa fa-stop"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-stop-circle"><i
                    class="fa fa-stop-circle"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-stop-circle-o"><i
                    class="fa fa-stop-circle-o"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-street-view"><i
                    class="fa fa-street-view"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-strikethrough"><i
                    class="fa fa-strikethrough"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-stumbleupon"><i
                    class="fa fa-stumbleupon"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-stumbleupon-circle"><i
                    class="fa fa-stumbleupon-circle"></i></a><a role="button"
                                                                href="#"
                                                                class="iconpicker-item"
                                                                title=".fa-subscript"><i
                    class="fa fa-subscript"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-subway"><i
                    class="fa fa-subway"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-suitcase"><i
                    class="fa fa-suitcase"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-sun-o"><i
                    class="fa fa-sun-o"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-superscript"><i
                    class="fa fa-superscript"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-support"><i
                    class="fa fa-support"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-table"><i
                    class="fa fa-table"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-tablet"><i
                    class="fa fa-tablet"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-tachometer"><i
                    class="fa fa-tachometer"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-tag"><i
                    class="fa fa-tag"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-tags"><i
                    class="fa fa-tags"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-tasks"><i
                    class="fa fa-tasks"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-taxi"><i
                    class="fa fa-taxi"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-television"><i
                    class="fa fa-television"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-tencent-weibo"><i
                    class="fa fa-tencent-weibo"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-terminal"><i
                    class="fa fa-terminal"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-text-height"><i
                    class="fa fa-text-height"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-text-width"><i
                    class="fa fa-text-width"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-th"><i
                    class="fa fa-th"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-th-large"><i
                    class="fa fa-th-large"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-th-list"><i
                    class="fa fa-th-list"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-themeisle"><i
                    class="fa fa-themeisle"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-thumb-tack"><i
                    class="fa fa-thumb-tack"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-thumbs-down"><i
                    class="fa fa-thumbs-down"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-thumbs-o-down"><i
                    class="fa fa-thumbs-o-down"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-thumbs-o-up"><i
                    class="fa fa-thumbs-o-up"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-thumbs-up"><i
                    class="fa fa-thumbs-up"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-ticket"><i
                    class="fa fa-ticket"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-times"><i
                    class="fa fa-times"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-times-circle"><i
                    class="fa fa-times-circle"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-times-circle-o"><i
                    class="fa fa-times-circle-o"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-tint"><i
                    class="fa fa-tint"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-toggle-down"><i
                    class="fa fa-toggle-down"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-toggle-left"><i
                    class="fa fa-toggle-left"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-toggle-off"><i
                    class="fa fa-toggle-off"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-toggle-on"><i
                    class="fa fa-toggle-on"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-toggle-right"><i
                    class="fa fa-toggle-right"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-toggle-up"><i
                    class="fa fa-toggle-up"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-trademark"><i
                    class="fa fa-trademark"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-train"><i
                    class="fa fa-train"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-transgender"><i
                    class="fa fa-transgender"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-transgender-alt"><i
                    class="fa fa-transgender-alt"></i></a><a role="button" href="#"
                                                             class="iconpicker-item"
                                                             title=".fa-trash"><i
                    class="fa fa-trash"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-trash-o"><i
                    class="fa fa-trash-o"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-tree"><i
                    class="fa fa-tree"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-trello"><i
                    class="fa fa-trello"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-tripadvisor"><i
                    class="fa fa-tripadvisor"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-trophy"><i
                    class="fa fa-trophy"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-truck"><i
                    class="fa fa-truck"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-try"><i
                    class="fa fa-try"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-tty"><i
                    class="fa fa-tty"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-tumblr"><i
                    class="fa fa-tumblr"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-tumblr-square"><i
                    class="fa fa-tumblr-square"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-turkish-lira"><i
                    class="fa fa-turkish-lira"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-tv"><i
                    class="fa fa-tv"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-twitch"><i
                    class="fa fa-twitch"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-twitter"><i
                    class="fa fa-twitter"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-twitter-square"><i
                    class="fa fa-twitter-square"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-umbrella"><i
                    class="fa fa-umbrella"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-underline"><i
                    class="fa fa-underline"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-undo"><i
                    class="fa fa-undo"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-universal-access"><i
                    class="fa fa-universal-access"></i></a><a role="button" href="#"
                                                              class="iconpicker-item"
                                                              title=".fa-university"><i
                    class="fa fa-university"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-unlink"><i
                    class="fa fa-unlink"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-unlock"><i
                    class="fa fa-unlock"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-unlock-alt"><i
                    class="fa fa-unlock-alt"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-unsorted"><i
                    class="fa fa-unsorted"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-upload"><i
                    class="fa fa-upload"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-usb"><i
                    class="fa fa-usb"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-usd"><i
                    class="fa fa-usd"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-user"><i
                    class="fa fa-user"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-user-md"><i
                    class="fa fa-user-md"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-user-plus"><i
                    class="fa fa-user-plus"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-user-secret"><i
                    class="fa fa-user-secret"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-user-times"><i
                    class="fa fa-user-times"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-users"><i
                    class="fa fa-users"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-venus"><i
                    class="fa fa-venus"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-venus-double"><i
                    class="fa fa-venus-double"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-venus-mars"><i
                    class="fa fa-venus-mars"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-viacoin"><i
                    class="fa fa-viacoin"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-viadeo"><i
                    class="fa fa-viadeo"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-viadeo-square"><i
                    class="fa fa-viadeo-square"></i></a><a role="button" href="#"
                                                           class="iconpicker-item"
                                                           title=".fa-video-camera"><i
                    class="fa fa-video-camera"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-vimeo"><i
                    class="fa fa-vimeo"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-vimeo-square"><i
                    class="fa fa-vimeo-square"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-vine"><i
                    class="fa fa-vine"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-vk"><i
                    class="fa fa-vk"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-volume-control-phone"><i
                    class="fa fa-volume-control-phone"></i></a><a role="button"
                                                                  href="#"
                                                                  class="iconpicker-item"
                                                                  title=".fa-volume-down"><i
                    class="fa fa-volume-down"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-volume-off"><i
                    class="fa fa-volume-off"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-volume-up"><i
                    class="fa fa-volume-up"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-warning"><i
                    class="fa fa-warning"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-wechat"><i
                    class="fa fa-wechat"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-weibo"><i
                    class="fa fa-weibo"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-weixin"><i
                    class="fa fa-weixin"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-whatsapp"><i
                    class="fa fa-whatsapp"></i></a><a role="button" href="#"
                                                      class="iconpicker-item"
                                                      title=".fa-wheelchair"><i
                    class="fa fa-wheelchair"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-wheelchair-alt"><i
                    class="fa fa-wheelchair-alt"></i></a><a role="button" href="#"
                                                            class="iconpicker-item"
                                                            title=".fa-wifi"><i
                    class="fa fa-wifi"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-wikipedia-w"><i
                    class="fa fa-wikipedia-w"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-windows"><i
                    class="fa fa-windows"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-won"><i
                    class="fa fa-won"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-wordpress"><i
                    class="fa fa-wordpress"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-wpbeginner"><i
                    class="fa fa-wpbeginner"></i></a><a role="button" href="#"
                                                        class="iconpicker-item"
                                                        title=".fa-wpforms"><i
                    class="fa fa-wpforms"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-wrench"><i
                    class="fa fa-wrench"></i></a><a role="button" href="#"
                                                    class="iconpicker-item"
                                                    title=".fa-xing"><i
                    class="fa fa-xing"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-xing-square"><i
                    class="fa fa-xing-square"></i></a><a role="button" href="#"
                                                         class="iconpicker-item"
                                                         title=".fa-y-combinator"><i
                    class="fa fa-y-combinator"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-y-combinator-square"><i
                    class="fa fa-y-combinator-square"></i></a><a role="button"
                                                                 href="#"
                                                                 class="iconpicker-item"
                                                                 title=".fa-yahoo"><i
                    class="fa fa-yahoo"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-yc"><i
                    class="fa fa-yc"></i></a><a role="button" href="#"
                                                class="iconpicker-item"
                                                title=".fa-yc-square"><i
                    class="fa fa-yc-square"></i></a><a role="button" href="#"
                                                       class="iconpicker-item"
                                                       title=".fa-yelp"><i
                    class="fa fa-yelp"></i></a><a role="button" href="#"
                                                  class="iconpicker-item"
                                                  title=".fa-yen"><i
                    class="fa fa-yen"></i></a><a role="button" href="#"
                                                 class="iconpicker-item"
                                                 title=".fa-yoast"><i
                    class="fa fa-yoast"></i></a><a role="button" href="#"
                                                   class="iconpicker-item"
                                                   title=".fa-youtube"><i
                    class="fa fa-youtube"></i></a><a role="button" href="#"
                                                     class="iconpicker-item"
                                                     title=".fa-youtube-play"><i
                    class="fa fa-youtube-play"></i></a><a role="button" href="#"
                                                          class="iconpicker-item"
                                                          title=".fa-youtube-square"><i
                    class="fa fa-youtube-square"></i></a></div>
        </div>
    </div>
</div>
</div>
{{end}}`, "components/form/password": `{{define "form_password"}}
<label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <div class="input-group">
        <span class="input-group-addon"><i class="fa fa-eye-slash"></i></span>
        <input type="password" id="{{.Field}}" name="{{.Field}}" value="{{.Value}}" class="form-control password" placeholder="{{lang "Input"}} {{.Head}}">
    </div>
</div>
{{end}}`, "components/form/select": `{{define "form_select"}}
<label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <select class="form-control {{.Field}} select2-hidden-accessible" style="width: 100%;" name="{{.Field}}[]" multiple="" data-placeholder="{{lang "Input"}} {{.Head}}" tabindex="-1" aria-hidden="true">
        {{range $key, $v := .Options }}
            <option value='{{index $v "value"}}' {{index $v "selected"}}>{{index $v "field"}}</option>
        {{end}}
    </select>
    <!--<span class="help-block">
        <i class="fa fa-info-circle"></i>&nbsp;All methods if empty
    </span>-->
</div>
<script>
    $(".{{.Field}}").select2({
        allowClear: true
    });
</script>
{{end}}`, "components/form/selectbox": `{{define "form_selectbox"}}
<label for="{{.Field}}" class="col-sm-2  control-label">{{.Head}}</label>
<div class="col-sm-8">
    <select class="form-control {{.Field}}" style="width: 100%;" name="{{.Field}}[]" multiple="multiple" data-placeholder="Input {{.Head}}"  >
        {{range  $key, $v := .Options }}
            <option value='{{index $v "value"}}' {{index $v "selected"}}>{{index $v "field"}}</option>
        {{end}}
    </select>
    <input type="hidden" name="{{.Field}}[]" />
</div>
<script>
    $(".{{.Field}}").bootstrapDualListbox({"infoText":"Showing all {0}","infoTextEmpty":"Empty list","infoTextFiltered":"{0} \/ {1}","filterTextClear":"Show all","filterPlaceHolder":"Filter"});
</script>
{{end}}`, "components/form/singleselect": `{{define "form_select_single"}}
<label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <select class="form-control {{.Field}} select2-hidden-accessible" style="width: 100%;" name="{{.Field}}" multiple="" data-placeholder="{{lang "Input"}} {{.Head}}" tabindex="-1" aria-hidden="true">
    {{range $key, $v := .Options }}
        <option value='{{index $v "value"}}' {{index $v "selected"}}>{{index $v "field"}}</option>
    {{end}}
    </select>
    <!--<span class="help-block">
        <i class="fa fa-info-circle"></i>&nbsp;All methods if empty
    </span>-->
</div>
<script>
    $(".{{.Field}}").select2({
        allowClear: true,
        maximumSelectionLength: 1
    });
</script>
{{end}}`, "components/form/text": `{{define "form_text"}}
<label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <div class="input-group">
        <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
        <input type="text" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}' class="form-control json" placeholder="{{lang "Input"}} {{.Head}}">
    </div>
</div>
{{end}}`, "components/form/textarea": `{{define "form_textarea"}}
<label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <textarea name="{{.Field}}" class="form-control" rows="5" placeholder="{{lang "Input"}} {{.Head}}">{{.Value}}</textarea>
</div>
{{end}}`, "components/form": `{{define "form"}}
<script src="{{.Prefix}}/assets/select2/select2.full.min.js"></script>
<script src="{{.Prefix}}/assets/fileinput/fileinput.min.js"></script>
<script src="{{.Prefix}}/assets/duallistbox/jquery.bootstrap-duallistbox.min.js"></script>
<div class="box box-info">
    <div class="box-header with-border">
        <h3 class="box-title">{{lang .Title}}</h3>
        <div class="box-tools">
            <!-- <div class="btn-group pull-right" style="margin-right: 10px">
                <a href='{{.InfoUrl}}' class="btn btn-sm btn-default"><i class="fa fa-list"></i> {{lang "List"}}</a>
            </div> -->
            <div class="btn-group pull-right" style="margin-right: 10px">
                <a href='{{.InfoUrl}}' class="btn btn-sm btn-default form-history-back"><i class="fa fa-arrow-left"></i> {{lang "Back"}}</a>
            </div>
        </div>
    </div>
    <form action='{{.Url}}' method="{{.Method}}" accept-charset="UTF-8" class="form-horizontal" pjax-container>
        <div class="box-body">
            <div class="fields-group">
                {{range $key, $data := .Content}}
                <div class="form-group">
                    {{if eq $data.FormType "default"}}
                        {{ template "form_default" $data }}
                    {{else if eq $data.FormType "text"}}
                        {{ template "form_text" $data }}
                    {{else if eq $data.FormType "file"}}
                        {{ template "form_file" $data }}
                    {{else if eq $data.FormType "password"}}
                        {{ template "form_password" $data }}
                    {{else if eq $data.FormType "selectbox"}}
                        {{ template "form_selectbox" $data }}
                    {{else if eq $data.FormType "select"}}
                        {{ template "form_select" $data }}
                    {{else if eq $data.FormType "select_single"}}
                        {{ template "form_select_single" $data }}
                    {{else if eq $data.FormType "textarea"}}
                        {{ template "form_textarea" $data }}
                    {{else if eq $data.FormType "iconpicker"}}
                        {{ template "form_iconpicker" $data }}
                    {{end}}
                </div>
                {{end}}
            </div>
        </div>
        <div class="box-footer">
            <div class="col-md-2">
            </div>
            <div class="col-md-8">

                <div class="btn-group pull-right">
                    <button type="submit" class="btn btn-info pull-right"
                            data-loading-text="&lt;i class='fa fa-spinner fa-spin '&gt;&lt;/i&gt; Save">
                        {{lang "Save"}}
                    </button>
                </div>

                <div class="btn-group pull-left">
                    <button type="reset" class="btn btn-warning">{{lang "Reset"}}</button>
                </div>

            </div>

        </div>

        <input type="hidden" name="_previous_" value='{{.InfoUrl}}' class="_previous_">
        {{range $key, $data := .Content}}
            {{if eq $data.Field "id"}}
                <input type="hidden" name="id" value='{{$data.Value}}' class="_previous_">
            {{end}}
        {{end}}
        <input type="hidden" name="_t" value='{{.CSRFToken}}' class="_previous_">
    </form>
</div>
{{end}}`, "components/image": `{{define "image"}}
<img src="{{.Src}}" width="{{.Width}}" height="{{.Height}}">
{{end}}`, "components/infobox": `{{define "infobox"}}
<div class="info-box">
    <span class="info-box-icon bg-{{.Color}}"><i class="fa {{.Icon}}"></i></span>
    <div class="info-box-content">
        <span class="info-box-text">{{.Text}}</span>
        <span class="info-box-number">{{.Number}}</span>
        {{.Content}}
    </div>
</div>
{{end}}`, "components/label": `{{define "label"}}
<span class="label label-{{.Color}}">{{.Content}}</span>
{{end}}`, "components/line-chart": `{{define "line-chart"}}
{{if ne .Title ""}}
<p class="text-center">
    <strong>{{.Title}}</strong>
</p>
{{end}}
<div class="chart">
    <canvas id="{{.ID}}" style="height: {{.Height}}px;"></canvas>
</div>
<script>
    let lineChartCanvas          = $('#{{.ID}}').get(0).getContext('2d');
    let lineChart                = new Chart(lineChartCanvas);
    let lineChartData = JSON.parse({{.Data}});

    let lineChartOptions = {
        // Boolean - If we should show the scale at all
        showScale               : true,
        // Boolean - Whether grid lines are shown across the chart
        scaleShowGridLines      : false,
        // String - Colour of the grid lines
        scaleGridLineColor      : 'rgba(0,0,0,.05)',
        // Number - Width of the grid lines
        scaleGridLineWidth      : 1,
        // Boolean - Whether to show horizontal lines (except X axis)
        scaleShowHorizontalLines: true,
        // Boolean - Whether to show vertical lines (except Y axis)
        scaleShowVerticalLines  : true,
        // Boolean - Whether the line is curved between points
        bezierCurve             : true,
        // Number - Tension of the bezier curve between points
        bezierCurveTension      : 0.3,
        // Boolean - Whether to show a dot for each point
        pointDot                : false,
        // Number - Radius of each point dot in pixels
        pointDotRadius          : 4,
        // Number - Pixel width of point dot stroke
        pointDotStrokeWidth     : 1,
        // Number - amount extra to add to the radius to cater for hit detection outside the drawn point
        pointHitDetectionRadius : 20,
        // Boolean - Whether to show a stroke for datasets
        datasetStroke           : true,
        // Number - Pixel width of dataset stroke
        datasetStrokeWidth      : 2,
        // Boolean - Whether to fill the dataset with a color
        datasetFill             : true,
        // String - A legend template
        legendTemplate          : '<ul class=\'<%=name.toLowerCase()%>-legend\'><% for (var i=0; i<datasets.length; i++){%><li><span style=\'background-color:<%=datasets[i].lineColor%>\'></span><%=datasets[i].label%></li><%}%></ul>',
        // Boolean - whether to maintain the starting aspect ratio or not when responsive, if set to false, will take up entire container
        maintainAspectRatio     : true,
        // Boolean - whether to make the chart responsive to window resizing
        responsive              : true
    };

    lineChartOptions.datasetFill = false;
    lineChart.Line(lineChartData, lineChartOptions)
</script>
{{end}}`, "components/paginator": `{{define "paginator"}}
Showing <b>{{.CurPageStartIndex}}</b> to <b>{{.CurPageEndIndex}}</b> of <b>{{.Total}}</b> entries
<ul class="pagination pagination-sm no-margin pull-right">
    <!-- Previous Page Link -->
    <li class="page-item {{.PreviousClass}}">
        {{if eq .PreviousClass "disabled"}}
            <span class="page-link">«</span>
        {{else}}
            <a class="page-link" href='{{.PreviousUrl}}' rel="next">«</a>
        {{end}}
    </li>

    <!-- Array Of Links -->
    {{range $key, $page := .Pages}}
        {{if eq (index $page "isSplit") "0"}}
            {{if eq (index $page "active") "active"}}
                <li class="page-item active"><span class="page-link">{{index $page "page"}}</span></li>
            {{else}}
                <li class="page-item"><a class="page-link" href='{{index $page "url"}}'>{{index $page "page"}}</a></li>
            {{end}}
        {{else}}
            <li class="page-item disabled"><span class="page-link">...</span></li>
        {{end}}
    {{end}}


    <!-- Next Page Link -->
    <li class='page-item {{.NextClass}}'>
        {{if eq .NextClass "disabled"}}
            <span class="page-link">»</span>
        {{else}}
            <a class="page-link" href='{{.NextUrl}}' rel="next">»</a>
        {{end}}
    </li>
</ul>

<label class="control-label pull-right" style="margin-right: 10px; font-weight: 100;">

    <small>Show</small>&nbsp;
    <select class="input-sm grid-per-pager" name="per-page">
        <option value="{{.Url}}&pageSize=10" {{index .Option "10"}}>
            10
        </option>
        <option value="{{.Url}}&pageSize=20" {{index .Option "20"}}>
            20
        </option>
        <option value="{{.Url}}&pageSize=30" {{index .Option "30"}}>
            30
        </option>
        <option value="{{.Url}}&pageSize=50" {{index .Option "50"}}>
            50
        </option>
        <option value="{{.Url}}&pageSize=100" {{index .Option "100"}}>
            100
        </option>
    </select>
    <small>entries</small>
</label>
{{end}}`, "components/pie-chart": `{{define "pie-chart"}}
{{if ne .Title ""}}
<p class="text-center">
    <strong>{{.Title}}</strong>
</p>
{{end}}
<div class="chart-responsive">
    <canvas id="{{.ID}}" style="height: {{.Height}}px"></canvas>
</div>
<script>
    let pieChartCanvas = $('#{{.ID}}').get(0).getContext('2d');
    let pieChart       = new Chart(pieChartCanvas);
    let PieData        = JSON.parse({{.Data}});
    let pieOptions     = {
        // Boolean - Whether we should show a stroke on each segment
        segmentShowStroke    : true,
        // String - The colour of each segment stroke
        segmentStrokeColor   : '#fff',
        // Number - The width of each segment stroke
        segmentStrokeWidth   : 1,
        // Number - The percentage of the chart that we cut out of the middle
        percentageInnerCutout: 50, // This is 0 for Pie charts
        // Number - Amount of animation steps
        animationSteps       : 100,
        // String - Animation easing effect
        animationEasing      : 'easeOutBounce',
        // Boolean - Whether we animate the rotation of the Doughnut
        animateRotate        : true,
        // Boolean - Whether we animate scaling the Doughnut from the centre
        animateScale         : false,
        // Boolean - whether to make the chart responsive to window resizing
        responsive           : true,
        // Boolean - whether to maintain the starting aspect ratio or not when responsive, if set to false, will take up entire container
        maintainAspectRatio  : false,
        // String - A legend template
        legendTemplate       : '<ul class=\'<%=name.toLowerCase()%>-legend\'><% for (var i=0; i<segments.length; i++){%><li><span style=\'background-color:<%=segments[i].fillColor%>\'></span><%if(segments[i].label){%><%=segments[i].label%><%}%></li><%}%></ul>',
        // String - A tooltip template
        tooltipTemplate      : '<%=value %> <%=label%> users'
    };
    // Create pie or douhnut chart
    // You can switch between pie and douhnut using the method below.
    pieChart.Doughnut(PieData, pieOptions);
</script>
{{end}}`, "components/popup": `{{define "popup"}}
<div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="{{.ID}}" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="{{.ID}}">{{.Title}}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <div class="modal-body">
                {{.Body}}
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-dismiss="modal">{{lang "Close"}}</button>
                <button type="button" class="btn btn-primary">{{.Footer}}</button>
            </div>
        </div>
    </div>
</div>
{{end}}`, "components/productlist": `{{define "productlist"}}
<ul class="products-list product-list-in-box">
    {{range $key, $data := .Data}}
    <li class="item">
        <div class="product-img">
            <img src="{{index $data "img"}}" alt="Product Image">
        </div>
        <div class="product-info">
            <a href="javascript:void(0)" class="product-title">{{index $data "title"}}
                {{if eq (index $data "has_tabel") "true"}}
                    <span class="label label-{{index $data "labeltype"}} pull-right">{{index $data "label"}}</span>
                {{end}}
            </a>
            <span class="product-description">
                {{index $data "description"}}
            </span>
        </div>
    </li>
    {{end}}
</ul>
{{end}}`, "components/progress-group": `{{define "progress-group"}}
<div class="progress-group">
    <span class="progress-text">{{.Title}}</span>
    <span class="progress-number"><b>{{.Molecular}}</b>/{{.Denominator}}</span>

    <div class="progress sm">
        <div class="progress-bar progress-bar-{{.Color}}" style="width: {{.Percent}}%"></div>
    </div>
</div>
{{end}}`, "components/row": `{{define "row"}}
<div class="row">{{.Content}}</div>
{{end}}`, "components/smallbox": `{{define "smallbox"}}
<div class="small-box bg-{{.Color}}">
    <div class="inner">
        <h3>{{.Value}}</h3>
        <p>{{.Title}}</p>
    </div>
    <div class="icon">
        <i class="fa fa-users"></i>
    </div>
    <a href="{{.Url}}" class="small-box-footer">
        More&nbsp;
        <i class="fa fa-arrow-circle-right"></i>
    </a>
</div>
{{end}}`, "components/table/box-header": `{{define "box-header"}}
<div class="pull-right">

    <div class="btn-group pull-right" style="margin-right: 10px">
        <a href="" class="btn btn-sm btn-primary" data-toggle="modal" data-target="#filter-modal"><i
                class="fa fa-filter"></i>&nbsp;&nbsp;{{lang "Filter"}}</a>
        <a href="{{.InfoUrl}}" class="btn btn-sm btn-facebook"><i class="fa fa-undo"></i>&nbsp;&nbsp;{{lang "Reset"}}
        </a>
    </div>

    <div class="modal fade" id="filter-modal" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">×</span>
                        <span class="sr-only">Close</span>
                    </button>
                    <h4 class="modal-title" id="myModalLabel">{{lang "Filter"}}</h4>
                </div>
                <form action="{{.InfoUrl}}" method="get" pjax-container="">
                    <div class="modal-body">
                        <div class="form">
                            {{range $key, $filter := .Filters}}
                                <div class="form-group">
                                    <label>{{index $filter "title"}}</label>
                                    <div class="input-group">
                                        <div class="input-group-addon">
                                            <i class="fa fa-pencil"></i>
                                        </div>
                                        <input type="text" class="form-control id" placeholder="{{index $filter "title"}}"
                                               name="{{index $filter "name"}}" value="">
                                    </div>
                                </div>
                            {{end}}
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="submit" id="filter-btn" class="btn btn-primary submit">{{lang "Submit"}}</button>
                        <button type="reset" class="btn btn-warning pull-left">{{lang "Reset"}}</button>
                    </div>
                </form>
            </div>
        </div>
    </div>

    <script>
        $("#filter-btn").click(function() {
            $('.modal-backdrop').hide();
        });
    </script>

    <!--
    <div class="btn-group pull-right" style="margin-right: 10px">
        <a class="btn btn-sm btn-twitter"><i class="fa fa-download"></i> Export</a>
        <button type="button" class="btn btn-sm btn-twitter dropdown-toggle" data-toggle="dropdown">
            <span class="caret"></span>
            <span class="sr-only">Toggle Dropdown</span>
        </button>
        <ul class="dropdown-menu" role="menu">
            <li><a href="/admin/story/word?_export_=all" target="_blank">All</a></li>
            <li><a href="/admin/story/word?_export_=page%3A1" target="_blank">Current
                page</a></li>
            <li><a href="/admin/story/word?_export_=selected%3A__rows__" target="_blank" class="export-selected">Selected rows</a></li>
        </ul>
    </div>
    -->

    <div class="btn-group pull-right" style="margin-right: 10px">

        <a href="{{.NewUrl}}" class="btn btn-sm btn-success">

            <i class="fa fa-save"></i>&nbsp;&nbsp;{{lang "New"}}
        </a>
    </div>
</div>
<span>
    <input type="checkbox" class="grid-select-all" style="position: absolute; opacity: 0;">
    <div class="btn-group">
        <a class="btn btn-sm btn-default">{{lang "Action"}}</a>
        <button type="button" class="btn btn-sm btn-default dropdown-toggle" data-toggle="dropdown">
        <span class="caret"></span>
        <span class="sr-only">{{lang "Toggle Dropdown"}}</span>
        </button>
        <ul class="dropdown-menu" role="menu">
            <li><a href="#" class="grid-batch-0">{{lang "Delete"}}</a></li>
        </ul>
    </div>
    <a class="btn btn-sm btn-primary grid-refresh">
        <i class="fa fa-refresh"></i> {{lang "Refresh"}}
    </a>
</span>
{{end}}`, "components/table": `{{define "table"}}
    <table class="table table-hover">
        {{if eq .Type "table"}}
            <thead>
                <tr>
                {{range $key, $head := .Thead}}
                    <th>
                    {{index $head "head"}}
                    {{if eq (index $head "sortable") "1"}}
                        <a class="fa fa-fw fa-sort" href=""></a>
                    {{end}}
                    </th>
                {{end}}
                </tr>
            </thead>
        {{end}}
        <tbody>
        {{if eq .Type "data-table"}}
            <tr>
                <th></th>
                {{range $key, $head := .Thead}}
                    <th>
                        {{index $head "head"}}
                        {{if eq (index $head "sortable") "1"}}
                            <a class="fa fa-fw fa-sort" href=""></a>
                        {{end}}
                    </th>
                {{end}}
                <th>{{lang "operation"}}</th>
            </tr>
        {{end}}

        {{$Thead := .Thead}}
        {{$Type := .Type}}
        {{$EditUrl := .EditUrl}}
        {{range $key1, $info := .InfoList}}
            <tr>
                {{if eq $Type "data-table"}}
                    <td>
                        <input type="checkbox" class="grid-row-checkbox" data-id="{{index $info "id"}}" style="position: absolute; opacity: 0;">
                    </td>
                {{end}}
                {{range $key2, $head2 := $Thead}}
                    <td>{{index $info (index $head2 "head")}}</td>
                {{end}}
                {{if eq $Type "data-table"}}
                    <td>
                        <a href='{{$EditUrl}}&id={{index $info "id"}}'><i class="fa fa-edit"></i></a>
                        <a href="javascript:void(0);" data-id='{{index $info "id"}}' class="grid-row-delete"><i class="fa fa-trash"></i></a>
                    </td>
                {{end}}
            </tr>
        {{end}}
        </tbody>
    </table>
    {{if eq $Type "data-table"}}
    <script>
        $('.grid-row-delete').unbind('click').click(function () {
            DeletePost($(this).data('id'))
        });

        let selectedRows = function () {
            let selected = [];
            $('.grid-row-checkbox:checked').each(function(){
                selected.push($(this).data('id'));
            });
            return selected;
        };

        $('.grid-select-all').on('ifChanged', function (event) {
            if (this.checked) {
                $('.grid-row-checkbox').iCheck('check');
            } else {
                $('.grid-row-checkbox').iCheck('uncheck');
            }
        });
        $('.grid-select-all').iCheck({checkboxClass: 'icheckbox_minimal-blue'});

        $(function () {
            $('.grid-row-checkbox').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function () {
                if (this.checked) {
                    $(this).closest('tr').css('background-color', "#ffffd5");
                } else {
                    $(this).closest('tr').css('background-color', '');
                }
            });
        });

        $('.grid-batch-0').on('click', function() {
            DeletePost(selectedRows().join())
        });

        function DeletePost(id) {
            swal({
                title: {{lang "are you sure to delete"}},
                type: "warning",
                showCancelButton: true,
                confirmButtonColor: "#DD6B55",
                confirmButtonText: {{lang "yes"}},
                closeOnConfirm: false,
                cancelButtonText: {{lang "cancel"}},
            },
            function(){
                $.ajax({
                    method: 'post',
                    url: {{.DeleteUrl}},
                    data: {
                        id: id
                    },
                    success: function (data) {
                        $.pjax.reload('#pjax-container');

                        data = JSON.parse(data);
                        if (data.code === 200) {
                            $('#_TOKEN').val(data.data);
                            swal(data.msg, '', 'success');
                        } else {
                            swal(data.msg, '', 'error');
                        }
                    }
                });
            });
        }
    </script>
    {{end}}
{{end}}`, "components/tabs": `{{define "tabs"}}
<div class="nav-tabs-custom">
    <ul class="nav nav-tabs">
        {{range $key, $data := .Data}}
            {{if eq $key 0}}
                <li class="active"><a href="#tab_{{$key}}" data-toggle="tab" aria-expanded="true">{{index $data "title"}}</a></li>
            {{else}}
                <li><a href="#tab_{{$key}}" data-toggle="tab" aria-expanded="true">{{index $data "title"}}</a></li>
            {{end}}
        {{end}}
    </ul>
    <div class="tab-content">
        {{range $key, $data := .Data}}
            {{if eq $key 0}}
                <div class="tab-pane active" id="tab_{{$key}}">
                {{index $data "content"}}
                </div>
            {{else}}
                <div class="tab-pane" id="tab_{{$key}}">
                {{index $data "content"}}
                </div>
            {{end}}
        {{end}}
    </div>
</div>
{{end}}`, "components/tree-header": `{{define "tree-header"}}
<div class="btn-group">
    <a class="btn btn-primary btn-sm tree-5b405b7481760-tree-tools" data-action="expand">
        <i class="fa fa-plus-square-o"></i>&nbsp;{{lang "expand"}}
    </a>
    <a class="btn btn-primary btn-sm tree-5b405b7481760-tree-tools" data-action="collapse">
        <i class="fa fa-minus-square-o"></i>&nbsp;{{lang "collapse"}}
    </a>
</div>

<div class="btn-group">
    <a class="btn btn-info btn-sm  tree-5b405b7481760-save"><i class="fa fa-save"></i>&nbsp;{{lang "save"}}</a>
</div>

<div class="btn-group">
    <a class="btn btn-warning btn-sm tree-5b405b7481760-refresh"><i class="fa fa-refresh"></i>&nbsp;{{lang "refresh"}}</a>
</div>
<div class="btn-group">
</div>
{{end}}`, "components/tree": `{{define "tree"}}
<div class="dd" id="tree-5b405b7481760">
    {{$EditUrl := .EditUrl}}
    <ol class="dd-list">
        {{range $key, $list := .Tree}}
        <li class="dd-item" data-id='{{$list.ID}}'>
            <div class="dd-handle">
                <i class="fa {{$list.Icon}}"></i>&nbsp;<strong>{{$list.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                    href="{{$list.Url}}" class="dd-nodrag">{{$list.Url}}</a>
                <span class="pull-right dd-nodrag">
                <a href="{{$EditUrl}}?id={{$list.ID}}"><i class="fa fa-edit"></i></a>
                <a href="javascript:void(0);" data-id="{{$list.ID}}" class="tree_branch_delete"><i class="fa fa-trash"></i></a>
            </span>
            </div>
            {{if gt (len $list.ChildrenList) 0}}
            <ol class="dd-list">
                {{range $key, $item := $list.ChildrenList}}
                    <li class="dd-item" data-id='{{$item.ID}}'>
                        <div class="dd-handle">
                            <i class="fa {{$item.Icon}}"></i>&nbsp;<strong>{{$item.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                                href="{{$item.Url}}" class="dd-nodrag">{{$item.Url}}</a>
                            <span class="pull-right dd-nodrag">
                                <a href="{{$EditUrl}}?id={{$item.ID}}"><i class="fa fa-edit"></i></a>
                                <a href="javascript:void(0);" data-id="{{$item.ID}}" class="tree_branch_delete"><i class="fa fa-trash"></i></a>
                            </span>
                        </div>
                    </li>
                {{end}}
            </ol>
            {{end}}
        </li>
        {{end}}
    </ol>
</div>
<script data-exec-on-popstate="">
    $(function () {
        $('#tree-5b405b7481760').nestable([]);
        $('.tree_branch_delete').click(function () {
            let id = $(this).data('id');
            swal({
                        title: "Are you sure to delete this item ?",
                        type: "warning",
                        showCancelButton: true,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "Confirm",
                        closeOnConfirm: false,
                        cancelButtonText: "Cancel"
                    },
                    function () {
                        $.ajax({
                            method: 'post',
                            url: {{.DeleteUrl}} + '?id=' + id,
                            data: {},
                            success: function (data) {
                                $.pjax.reload('#pjax-container');
                                if (data.code === 200) {
                                    swal(data.msg, '', 'success');
                                } else {
                                    swal(data.msg, '', 'error');
                                }
                            }
                        });
                    });
        });
        $('.tree-5b405b7481760-save').click(function () {
            let serialize = $('#tree-5b405b7481760').nestable('serialize');
            $.post({{.OrderUrl}}, {
                        _order: JSON.stringify(serialize)
                    },
                    function (data) {
                        $.pjax.reload('#pjax-container');
                        toastr.success('Save succeeded !');
                    });
        });
        $('.tree-5b405b7481760-refresh').click(function () {
            $.pjax.reload('#pjax-container');
            toastr.success('{{lang "Refresh succeeded"}} !');
        });
        $('.tree-5b405b7481760-tree-tools').on('click', function (e) {
            let target = $(e.target),
                    action = target.data('action');
            if (action === 'expand') {
                $('.dd').nestable('expandAll');
            }
            if (action === 'collapse') {
                $('.dd').nestable('collapseAll');
            }
        });
        $(".parent_id").select2({"allowClear": true, "placeholder": "Parent"});
        $('.icon').iconpicker({placement: 'bottomLeft'});
        $(".roles").select2({"allowClear": true, "placeholder": "Roles"});
    });
</script>
{{end}}`, "content": `{{define "content"}}
<script>
    $('.grid-per-pager').on("change", function (e) {
        console.log("changing...");
        $.pjax({url: this.value, container: '#pjax-container'});
    });
    $('.grid-refresh').on('click', function () {
        $.pjax.reload('#pjax-container');
        toastr.success('{{lang "Refresh succeeded"}} !');
    });
</script>
<script src="{{.AssertRootUrl}}/assets/chartjs/chart.js"></script>
<script src="{{.AssertRootUrl}}/assets/fontawesome-iconpicker/dist/js/fontawesome-iconpicker.min.js"></script>
<section class="content-header">
    <h1>
        {{lang .Panel.Title}}
        <small>{{lang .Panel.Description}}</small>
    </h1>
</section>

<!-- Main content -->
<section class="content">
    {{.Panel.Content}}
</section>
{{end}}`, "footer": `{{define "footer"}}
<footer class="main-footer">
    <div class="pull-right hidden-xs">
        <b>Version</b> {{.System.Version}}
    </div>
    <strong>Copyright &copy; 2018- <a href="https://github.com/chenhg5/go-admin">GoAdmin</a>.</strong> All rights
    reserved.
</footer>
{{end}}`, "head": `{{define "head"}}
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{.Title}}</title>
    <!-- Tell the browser to be responsive to screen width -->
    <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
    <!-- Bootstrap 3.3.7 -->
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/bootstrap/dist/css/bootstrap.min.css">
    <!-- Font Awesome -->
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/font-awesome/css/font-awesome.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/Ionicons/css/ionicons.min.css">
    <!-- DataTables -->
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/datatables.net-bs/css/dataTables.bootstrap.min.css">
    <!-- iCheck -->
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/iCheck/minimal/_all.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/iCheck/futurico/futurico.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/iCheck/polaris/polaris.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/toastr/build/toastr.min.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/nprogress/nprogress.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/select2/select2.min.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/sweetalert/dist/sweetalert.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/fileinput/fileinput.min.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/nestable/nestable.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/duallistbox/bootstrap-duallistbox.min.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/fontawesome-iconpicker/dist/css/fontawesome-iconpicker.min.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/dist/css/AdminLTE.min.css">
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/dist/css/skins/skin-black.css">
    <!--[if lt IE 9]>
    <script src="{{.AssertRootUrl}}/assets/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="{{.AssertRootUrl}}/assets/respond/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- Google Font -->
    <link rel="stylesheet" href="{{.AssertRootUrl}}/assets/googleapis/font.css">

    <script src="{{.AssertRootUrl}}/assets/jQuery/jQuery-2.1.4.min.js"></script>
    <script src="{{.AssertRootUrl}}/assets/nestable/jquery.nestable.js"></script>
    <script src="{{.AssertRootUrl}}/assets/dist/js/adminlte.min.js"></script>
    <script src="{{.AssertRootUrl}}/assets/fontawesome-iconpicker/dist/js/fontawesome-iconpicker.min.js"></script>
    <script src="{{.AssertRootUrl}}/assets/iCheck/icheck.min.js"></script>
</head>
{{end}}`, "header": `{{define "header"}}
<header class="main-header">
    <!-- Logo -->
    <a href="/" class="logo">
        <!-- mini logo for sidebar mini 50x50 pixels -->
        <span class="logo-mini">{{.MiniLogo}}</span>
        <!-- logo for regular state and mobile devices -->
        <span class="logo-lg">{{.Logo}}</span>
    </a>
    <!-- Header Navbar: style can be found in header.less -->
    <nav class="navbar navbar-static-top">
        <!-- Sidebar toggle button-->
        <a href="#" class="sidebar-toggle" data-toggle="offcanvas" role="button">
            <span class="sr-only">Toggle navigation</span>
        </a>

        {{ template "admin_panel" . }}
    </nav>
</header>
{{end}}`, "js": `{{define "js"}}
<script src="{{.AssertRootUrl}}/assets/bootstrap/dist/js/bootstrap.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/datatables.net/js/jquery.dataTables.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/datatables.net-bs/js/dataTables.bootstrap.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/fastclick/lib/fastclick.js"></script>
<script src="{{.AssertRootUrl}}/assets/select2/select2.full.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/fileinput/fileinput.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/nprogress/nprogress.js"></script>
<script src="{{.AssertRootUrl}}/assets/toastr/build/toastr.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/bootstrap3-editable/js/bootstrap-editable.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/jquery-pjax/jquery.pjax.js"></script>
<script src="{{.AssertRootUrl}}/assets/sweetalert/dist/sweetalert.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/duallistbox/jquery.bootstrap-duallistbox.min.js"></script>
<script src="{{.AssertRootUrl}}/assets/dist/js/info.js"></script>
{{end}}`, "layout": `{{define "layout"}}

<!DOCTYPE html>
<html>

{{ template "head" . }}

<body class="hold-transition skin-black sidebar-mini">
<div class="wrapper">

    {{ template "header" . }}

    {{ template "sidebar" . }}

    <div class="content-wrapper" id="pjax-container">

    {{ template "content" . }}

    </div>

    {{ template "footer" . }}

</div>

{{ template "js" . }}

</body>
</html>

{{end}}
`, "menu": `{{define "menu"}}
<ul class="sidebar-menu" data-widget="tree">
{{$AssertRootUrl := .AssertRootUrl}}
{{range $key, $list := .Menu.GlobalMenuList }}
    {{if eq (len $list.ChildrenList) 0}}
        <li class='{{$list.Active}}'>
            {{if eq $list.Url "/"}}
                <a href='{{$AssertRootUrl}}'>
            {{else}}
                <a href='{{$AssertRootUrl}}{{$list.Url}}'>
            {{end}}
                <i class="fa {{$list.Icon}}"></i><span>{{$list.Name}}</span>
                <span class="pull-right-container"><!-- <small class="label pull-right bg-green">new</small> --></span>
            </a>
        </li>
    {{else}}
        <li class="treeview {{$list.Active}}">
            <a href="#">
                <i class="fa {{$list.Icon}}"></i> <span>{{$list.Name}}</span>
                <span class="pull-right-container">
                        <i class="fa fa-angle-left pull-right"></i>
                    </span>
            </a>
            <ul class="treeview-menu">
            {{range $key2, $item := $list.ChildrenList}}
                <li><a href="{{$AssertRootUrl}}{{$item.Url}}"><i class="fa {{$item.Icon}}"></i> {{$item.Name}}</a></li>
            {{end}}
            </ul>
        </li>
    {{end}}
{{end}}
</ul>
{{end}}`, "sidebar": `{{define "sidebar"}}
<!-- Left side column. contains the logo and sidebar -->
<aside class="main-sidebar">
    <!-- sidebar: style can be found in sidebar.less -->
    <section class="sidebar">
        <!-- Sidebar user panel -->
        <div class="user-panel">
            <div class="pull-left image">
                {{if eq .User.Avatar ""}}
                    <img src="{{.AssertRootUrl}}/assets/dist/img/avatar04.png" class="img-circle" alt="User Image">
                {{else}}
                    <img src="{{.User.Avatar}}" class="img-circle" alt="User Image">
                {{end}}
            </div>
            <div class="pull-left info">
                {{.User.Name}}
                <a href="#"><i class="fa fa-circle text-success"></i> {{lang "online"}}</a>
            </div>
        </div>
        <!-- search form -->
        <form action="#" method="get" class="sidebar-form">
            <div class="input-group">
                <input type="text" name="q" class="form-control" placeholder="Search...">
                <span class="input-group-btn">
                <button type="submit" name="search" id="search-btn" class="btn btn-flat"><i class="fa fa-search"></i>
                </button>
              </span>
            </div>
        </form>
        <!-- /.search form -->

        {{ template "menu" . }}

    </section>
    <!-- /.sidebar -->
</aside>
{{end}}`}
