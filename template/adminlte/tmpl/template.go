package tmpl

var List = map[string]string{"admin_panel": `{{define "admin_panel"}}
    <div class="navbar-custom-menu">
        <ul class="nav navbar-nav">

            <li>
                <a href="javascript:void(0);" class="fixed-btn" data-click="false">
                    <i class="fa fa-anchor"></i>
                </a>
            </li>

            <li>
                <a href="javascript:void(0);" class="fullpage-btn">
                    <i class="fa fa-arrows-alt"></i>
                </a>
            </li>
            <li>
                <a href="javascript:void(0);" class="exit-fullpage-btn" style="display: none;">
                    <i class="fa fa-compress"></i>
                </a>
            </li>
            <li>
                <a href="javascript:void(0);" class="container-refresh">
                    <i class="fa fa-refresh"></i>
                </a>
            </li>
            <!-- User Account: style can be found in dropdown.less -->
            <li class="dropdown user user-menu">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                    {{if eq .User.Avatar ""}}
                        <img src="{{.UrlPrefix}}/assets/dist/img/avatar04.png" class="user-image" alt="User Image">
                    {{else}}
                        <img src="{{.User.Avatar}}" class="user-image" alt="User Image">
                    {{end}}
                    <span class="hidden-xs">{{.User.Name}}</span>
                </a>
                <ul class="dropdown-menu">
                    <!-- User image -->
                    <li class="user-header">
                        {{if eq .User.Avatar ""}}
                            <img src="{{.UrlPrefix}}/assets/dist/img/avatar04.png" class="img-circle" alt="User Image">
                        {{else}}
                            <img src="{{.User.Avatar}}" class="img-circle" alt="User Image">
                        {{end}}
                        <p>
                            {{.User.Name}} -{{.User.Level}}
                            <small>{{.User.CreatedAt}}</small>
                        </p>
                    </li>
                    <li class="user-footer">
                        <div class="pull-left">
                            <a href="{{.UrlPrefix}}/info/manager/edit?id={{.User.Id}}"
                               class="btn btn-default btn-flat">{{lang "setting"}}</a>
                        </div>
                        <div class="pull-right">
                            <a href="{{.UrlPrefix}}/logout" class="btn btn-default btn-flat">{{lang "sign out"}}</a>
                        </div>
                    </li>
                </ul>
            </li>
            <!-- <li class="hidden-xs">
                 <a href="javascript:;" data-toggle="control-sidebar"><i class="fa fa-gears"></i></a>
             </li> -->
        </ul>
    </div>
{{end}}`, "components/alert": `{{define "alert"}}
<div class="alert alert-{{.Theme}} alert-dismissible">
    <button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
    <h4>{{langHtml .Title}}</h4>
    {{langHtml .Content}}
</div>
{{end}}`, "components/area-chart": `{{define "area-chart"}}
{{if ne .Title ""}}
<p class="text-center">
    <strong>{{langHtml .Title}}</strong>
</p>
{{end}}
<div class="chart">
    <canvas id="{{.ID}}" style="height: {{.Height}}px;"></canvas>
</div>
<script>
    (new Chart($('#{{.ID}}').get(0).getContext('2d'))).Line(JSON.parse({{.Data}}), {
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
    });
</script>
{{end}}`, "components/bar-chart": `{{define "bar-chart"}}
{{if ne .Title ""}}
<p class="text-center">
    <strong>{{langHtml .Title}}</strong>
</p>
{{end}}
<div class="chart">
    <canvas id="{{.ID}}" style="width: {{.Width}}px;"></canvas>
</div>
<script>
    barChartData = JSON.parse({{.Data}});
    barChartData.datasets[1].fillColor = '#00a65a';
    barChartData.datasets[1].strokeColor = '#00a65a';
    barChartData.datasets[1].pointColor = '#00a65a';
    (new Chart($('#{{.ID}}').get(0).getContext('2d'))).Bar(barChartData, {
        //Boolean - Whether the scale should start at zero, or an order of magnitude down from the lowest value
        scaleBeginAtZero: true,
        //Boolean - Whether grid lines are shown across the chart
        scaleShowGridLines: true,
        //String - Colour of the grid lines
        scaleGridLineColor: 'rgba(0,0,0,.05)',
        //Number - Width of the grid lines
        scaleGridLineWidth: 1,
        //Boolean - Whether to show horizontal lines (except X axis)
        scaleShowHorizontalLines: true,
        //Boolean - Whether to show vertical lines (except Y axis)
        scaleShowVerticalLines: true,
        //Boolean - If there is a stroke on each bar
        barShowStroke: true,
        //Number - Pixel width of the bar stroke
        barStrokeWidth: 2,
        //Number - Spacing between each of the X value sets
        barValueSpacing: 5,
        //Number - Spacing between data sets within X values
        barDatasetSpacing: 1,
        //String - A legend template
        legendTemplate: '<ul class="<%=name.toLowerCase()%>-legend"><% for (var i=0; i<datasets.length; i++){%><li><span style="background-color:<%=datasets[i].fillColor%>"></span><%if(datasets[i].label){%><%=datasets[i].label%><%}%></li><%}%></ul>',
        //Boolean - whether to make the chart responsive
        responsive: true,
        maintainAspectRatio: true,
        datasetFill: false
    })
</script>
{{end}}`, "components/box": `{{define "box"}}
<div class="box box-{{.Theme}}">
    {{if eq .HeadColor ""}}
        <div class="box-header {{.HeadBorder}}">
    {{else}}
        <div class="box-header {{.HeadBorder}}" style="background-color: {{.HeadColor}};">
    {{end}}
        {{langHtml .Header}}
    </div>
    <div class="box-body" style="overflow: scroll;">
        {{langHtml .Body}}
    </div>
    {{if ne .Footer ""}}
    <div class="box-footer clearfix">
        {{langHtml .Footer}}
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
<div class="{{.Size}}">{{langHtml .Content}}</div>
{{end}}`, "components/description": `{{define "description"}}
<div class="description-block border-{{.Border}}">
    <span class="description-percentage text-{{.Color}}"><i class="fa fa-caret-{{.Arrow}}"></i>{{langHtml .Percent}}%</span>
    <h5 class="description-header">{{langHtml .Number}}</h5>
    <span class="description-text">{{langHtml .Title}}</span>
</div>
{{end}}`, "components/form/color": `{{define "form_color"}}
    {{if eq .Must true}}
    <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
    <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <div class="input-group colorpicker-element">
            <span class="input-group-addon"><i style="background-color: rgb(0, 0, 0);"></i></span>
            <input style="width: 140px" type="text" id="{{.Field}}" name="{{.Field}}" value="" class="form-control {{.Field}}" placeholder="{{.Value}}">
        </div>
    </div>
    <script>
        $('.{{.Field}}').parent().colorpicker([]);
    </script>
{{end}}`, "components/form/currency": `{{define "form_currency"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <div class="input-group">
            <span class="input-group-addon">$</span>
            <input style="width: 120px; text-align: right;" type="text" id="{{.Field}}" name="{{.Field}}"
                   value="{{.Value}}" class="form-control {{.Field}}" placeholder="{{.Head}}">
        </div>
    </div>
    <script>
        $(function () {
            $('.{{.Field}}').inputmask({"alias": "currency", "radixPoint": ".", "prefix": "", "removeMaskOnSubmit": true});
        });
    </script>
{{end}}`, "components/form/datetime": `{{define "form_datetime"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-calendar fa-fw"></i></span>
            <input style="width: 160px" type="text" id="{{.Field}}" name="{{.Field}}" value="{{.Value}}"
                   class="form-control {{.Field}}" placeholder="{{lang "Input"}} {{.Head}}">
        </div>
    </div>
    <script>
        $(function () {
            $('.{{.Field}}').parent().datetimepicker({
                "format": "YYYY-MM-DD HH:mm:ss",
                "locale": "zh-CN",
                "allowInputToggle": true
            });
        });
    </script>
{{end}}`, "components/form/default": `{{define "form_default"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <div class="box box-solid box-default no-margin">
            <div class="box-body">
                {{.Value}}
            </div>
        </div>
    </div>
{{end}}`, "components/form/email": `{{define "form_email"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-envelope fa-fw"></i></span>
            <input type="text" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}' class="form-control json"
                   placeholder="{{lang "Input"}} {{.Head}}">
        </div>
    </div>
{{end}}`, "components/form/file": `{{define "form_file"}}
{{if eq .Must true}}
    <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
{{else}}
    <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
{{end}}
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
        "deleteUrl":"",
        "allowedFileTypes":["image"]
    });
</script>
{{end}}`, "components/form/iconpicker": `{{define "form_iconpicker"}}
{{if eq .Must true}}
    <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
{{else}}
    <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
{{end}}
<div class="col-sm-8">
    <div class="input-group">
        <span class="input-group-addon"><i class="fa"></i></span>
        {{if eq .Value ""}}
            <input style="width: 140px" type="text" id="icon" name="{{.Field}}" value="fa-bars"
                   class="form-control icon iconpicker-element iconpicker-input"
                   placeholder="{{lang "Input Icon"}}">
        {{else}}
            <input style="width: 140px" type="text" id="icon" name="{{.Field}}" value="{{.Value}}"
                   class="form-control icon iconpicker-element iconpicker-input"
                   placeholder="{{lang "Input Icon"}}">
        {{end}}
    </div>
</div>
<script>
    $('.iconpicker-input').iconpicker({placement:'bottomLeft'});
</script>
{{end}}`, "components/form/ip": `{{define "form_ip"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-laptop fa-fw"></i></span>
            <input style="width: 130px" type="text" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}' class="form-control json"
                   placeholder="{{lang "Input"}} {{.Head}}">
        </div>
    </div>
{{end}}`, "components/form/number": `{{define "form_number"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <div class="input-group">
            <input style="width: 100px; text-align: center;" type="text" id="{{.Field}}" name="{{.Field}}"
                   value="{{.Value}}" class="form-control {{.Field}}"
                   placeholder="{{.Head}}">
        </div>
    </div>
    <script>
        $(function () {
            $('.{{.Field}}:not(.initialized)')
                .addClass('initialized')
                .bootstrapNumber({
                    upClass: 'success',
                    downClass: 'primary',
                    center: true
                });
        })
    </script>
{{end}}`, "components/form/password": `{{define "form_password"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
<div class="col-sm-8">
    {{if .Editable}}
        <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-eye-slash"></i></span>
            <input type="password" id="{{.Field}}" name="{{.Field}}" value="{{.Value}}" class="form-control password" placeholder="{{lang "Input"}} {{.Head}}">
        </div>
    {{else}}
        <div class="box box-solid box-default no-margin">
            <div class="box-body">********</div>
         </div>
    {{end}}
</div>
{{end}}`, "components/form/radio": `{{define "form_radio"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        {{$field := .Field}}
        {{range $key, $v := .Options }}
            <input type="radio" name="{{index $v "field"}}" value="{{index $v "value"}}"
                   class="minimal {{$field}}" {{index $v "selected"}}
                   style="position: absolute; opacity: 0;">&nbsp;{{index $v "label"}}&nbsp;&nbsp;
        {{end}}
    </div>
    <script>
        $(function () {
            $('.{{.Field}}').iCheck({radioClass: 'iradio_minimal-blue'});
        });
    </script>
{{end}}`, "components/form/richtext": `{{define "form_rich_text"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <div id="{{.Field}}-editor">
            <p>欢迎使用 <b>wangEditor</b> 富文本编辑器</p>
        </div>
        <input type="hidden" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}'
               placeholder="{{lang "Input"}} {{.Head}}">
    </div>
    <script type="text/javascript">
        editor = new window.wangEditor('#{{.Field}}-editor');
        editor.customConfig.onchange = function (html) {
            $('#{{.Field}}').val(html)
        };
        editor.create();
        editor.txt.html('{{.Value}}');
        {{if not .Editable}}
        editor.$textElem.attr('contenteditable', false);
        {{end}}
    </script>
{{end}}`, "components/form/select": `{{define "form_select"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <select class="form-control {{.Field}} select2-hidden-accessible" style="width: 100%;" name="{{.Field}}[]"
                multiple="" data-placeholder="{{lang "Input"}} {{.Head}}" tabindex="-1" aria-hidden="true"
                {{if not .Editable}}disabled="disabled"{{end}}>
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
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
<div class="col-sm-8">
    <select class="form-control {{.Field}}" style="width: 100%;" name="{{.Field}}[]" multiple="multiple" data-placeholder="Input {{.Head}}"  {{if not .Editable}}disabled="disabled"{{end}}>
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
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
<div class="col-sm-8">
    <select class="form-control {{.Field}} select2-hidden-accessible" style="width: 100%;" name="{{.Field}}"
            multiple="" data-placeholder="{{lang "Input"}} {{.Head}}" tabindex="-1" aria-hidden="true" {{if not .Editable}}disabled="disabled"{{end}}>
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
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
<div class="col-sm-8">
    {{if .Editable}}
        <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
            <input type="text" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}' class="form-control json" placeholder="{{lang "Input"}} {{.Head}}">
        </div>
    {{else}}
        <div class="box box-solid box-default no-margin">
            <div class="box-body">{{.Value}}</div>
        </div>
    {{end}}
</div>
{{end}}`, "components/form/textarea": `{{define "form_textarea"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
<div class="col-sm-8">
    <textarea name="{{.Field}}" class="form-control" rows="5" placeholder="{{lang "Input"}} {{.Head}}" {{if not .Editable}}disabled="disabled"{{end}}>{{.Value}}</textarea>
</div>
{{end}}`, "components/form/url": `{{define "form_url"}}
    {{if eq .Must true}}
        <label for="{{.Field}}" class="col-sm-2 asterisk control-label">{{.Head}}</label>
    {{else}}
        <label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
    {{end}}
    <div class="col-sm-8">
        <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-internet-explorer fa-fw"></i></span>
            <input type="text" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}' class="form-control json"
                   placeholder="{{lang "Input"}} {{.Head}}">
        </div>
    </div>
{{end}}`, "components/form": `{{define "form"}}
    <script src="{{.Prefix}}/assets/dist/js/form.min.js"></script>
    <div class="box box-info">
        <div class="box-header with-border">
            <h3 class="box-title">{{langHtml .Title}}</h3>
            <div class="box-tools">
                <!-- <div class="btn-group pull-right" style="margin-right: 10px">
                <a href='{{.InfoUrl}}' class="btn btn-sm btn-default"><i class="fa fa-list"></i> {{lang "List"}}</a>
                </div> -->
                <div class="btn-group pull-right" style="margin-right: 10px">
                    <a href='{{.InfoUrl}}' class="btn btn-sm btn-default form-history-back"><i
                                class="fa fa-arrow-left"></i> {{lang "Back"}}</a>
                </div>
            </div>
        </div>
        {{.Header}}
        {{$PrimaryKey := .PrimaryKey}}
        <form action='{{.Url}}' method="{{.Method}}" accept-charset="UTF-8" class="form-horizontal" pjax-container>
            <div class="box-body">
                {{if eq (len .TabHeaders) 0}}
                    <div class="fields-group">
                        {{ template "form_components" .Content }}
                    </div>

                    {{range $key, $data := .Content}}
                        {{if eq $data.Field $PrimaryKey}}
                            <input type="hidden" name="{{$PrimaryKey}}" value='{{$data.Value}}'>
                        {{end}}
                    {{end}}
                {{else}}
                    <div class="nav-tabs-custom">
                        <ul class="nav nav-tabs">
                            {{range $key, $data := .TabHeaders}}
                                {{if eq $key 0}}
                                    <li class="active">
                                {{else}}
                                    <li class="">
                                {{end}}
                                <a href="#tab-form-{{$key}}" data-toggle="tab" aria-expanded="true">
                                    {{$data}} <i class="fa fa-exclamation-circle text-red hide"></i>
                                </a>
                            </li>
                            {{end}}
                        </ul>
                        <div class="tab-content fields-group">

                            {{range $key, $data := .TabContents}}

                            {{if eq $key 0}}
                                <div class="tab-pane active" id="tab-form-{{$key}}">
                            {{else}}
                                <div class="tab-pane" id="tab-form-{{$key}}">
                            {{end}}
                                {{ template "form_components" $data}}
                                {{range $key, $d := $data}}
                                    {{if eq $d.Field $PrimaryKey}}
                                        <input type="hidden" name="{{$PrimaryKey}}" value='{{$d.Value}}'>
                                    {{end}}
                                {{end}}
                            </div>

                            {{end}}

                        </div>
                    </div>
                {{end}}
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

            <input type="hidden" name="_previous_" value='{{.InfoUrl}}'>
            <input type="hidden" name="_t" value='{{.CSRFToken}}'>
        </form>
        {{.Footer}}
    </div>
{{end}}`, "components/form_components": `{{define "form_components"}}
    {{range $key, $data := .}}
        <div class="form-group">
            {{if eq $data.FormType.String "default"}}
                {{ template "form_default" $data }}
            {{else if eq $data.FormType.String "text"}}
                {{ template "form_text" $data }}
            {{else if eq $data.FormType.String "file"}}
                {{ template "form_file" $data }}
            {{else if eq $data.FormType.String "password"}}
                {{ template "form_password" $data }}
            {{else if eq $data.FormType.String "selectbox"}}
                {{ template "form_selectbox" $data }}
            {{else if eq $data.FormType.String "select"}}
                {{ template "form_select" $data }}
            {{else if eq $data.FormType.String "select_single"}}
                {{ template "form_select_single" $data }}
            {{else if eq $data.FormType.String "textarea"}}
                {{ template "form_textarea" $data }}
            {{else if eq $data.FormType.String "iconpicker"}}
                {{ template "form_iconpicker" $data }}
            {{else if eq $data.FormType.String "richtext"}}
                {{ template "form_rich_text" $data }}
            {{else if eq $data.FormType.String "datetime"}}
                {{ template "form_datetime" $data }}
            {{else if eq $data.FormType.String "radio"}}
                {{ template "form_radio" $data }}
            {{else if eq $data.FormType.String "email"}}
                {{ template "form_email" $data }}
            {{else if eq $data.FormType.String "url"}}
                {{ template "form_url" $data }}
            {{else if eq $data.FormType.String "ip"}}
                {{ template "form_ip" $data }}
            {{else if eq $data.FormType.String "color"}}
                {{ template "form_color" $data }}
            {{else if eq $data.FormType.String "currency"}}
                {{ template "form_currency" $data }}
            {{else if eq $data.FormType.String "number"}}
                {{ template "form_number" $data }}
            {{end}}
        </div>
    {{end}}
{{end}}`, "components/image": `{{define "image"}}
<img src="{{.Src}}" width="{{.Width}}" height="{{.Height}}">
{{end}}`, "components/infobox": `{{define "infobox"}}
<div class="info-box">
    {{if .IsHexColor}}
        <span class="info-box-icon" style="background-color: {{.Color}} !important;">
    {{else}}
        <span class="info-box-icon bg-{{.Color}}">
    {{end}}
    {{if .IsSvg}}
        {{.Icon}}
    {{else}}
        <i class="fa {{.Icon}}"></i>
    {{end}}
    </span>
    <div class="info-box-content">
        <span class="info-box-text">{{langHtml .Text}}</span>
        <span class="info-box-number">{{langHtml .Number}}</span>
        {{langHtml .Content}}
    </div>
</div>
{{end}}`, "components/label": `{{define "label"}}
<span class="label label-{{.Color}}">{{langHtml .Content}}</span>
{{end}}`, "components/line-chart": `{{define "line-chart"}}
{{if ne .Title ""}}
<p class="text-center">
    <strong>{{langHtml .Title}}</strong>
</p>
{{end}}
<div class="chart">
    <canvas id="{{.ID}}" style="height: {{.Height}}px;"></canvas>
</div>
<script>
    (new Chart($('#{{.ID}}').get(0).getContext('2d'))).Line(JSON.parse({{.Data}}), {
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
        datasetFill             : false,
        // String - A legend template
        legendTemplate          : '<ul class=\'<%=name.toLowerCase()%>-legend\'><% for (var i=0; i<datasets.length; i++){%><li><span style=\'background-color:<%=datasets[i].lineColor%>\'></span><%=datasets[i].label%></li><%}%></ul>',
        // Boolean - whether to maintain the starting aspect ratio or not when responsive, if set to false, will take up entire container
        maintainAspectRatio     : true,
        // Boolean - whether to make the chart responsive to window resizing
        responsive              : true
    })
</script>
{{end}}`, "components/paginator": `{{define "paginator"}}
{{lang "showing"}} <b>{{.CurPageStartIndex}}</b> {{lang "to"}} <b>{{.CurPageEndIndex}}</b> {{lang "of"}} <b>{{.Total}}</b> {{lang "entries"}}
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

    <small>{{lang "show"}}</small>&nbsp;
    <select class="input-sm grid-per-pager" name="per-page">
        <option value="{{.Url}}&__pageSize=10" {{index .Option "10"}}>
            10
        </option>
        <option value="{{.Url}}&__pageSize=20" {{index .Option "20"}}>
            20
        </option>
        <option value="{{.Url}}&__pageSize=30" {{index .Option "30"}}>
            30
        </option>
        <option value="{{.Url}}&__pageSize=50" {{index .Option "50"}}>
            50
        </option>
        <option value="{{.Url}}&__pageSize=100" {{index .Option "100"}}>
            100
        </option>
    </select>
    <small>{{lang "entries"}}</small>
</label>
{{end}}`, "components/pie-chart": `{{define "pie-chart"}}
{{if ne .Title ""}}
<p class="text-center">
    <strong>{{langHtml .Title}}</strong>
</p>
{{end}}
<div class="chart-responsive">
    <canvas id="{{.ID}}" style="height: {{.Height}}px"></canvas>
</div>
<script>
    // You can switch between pie and douhnut using the method below.
    (new Chart($('#{{.ID}}').get(0).getContext('2d'))).Doughnut(JSON.parse({{.Data}}), {
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
    });
</script>
{{end}}`, "components/popup": `{{define "popup"}}
<div class="modal fade" id="{{.ID}}" tabindex="-1" role="dialog" aria-labelledby="{{.ID}}" aria-hidden="true">
    <div class="modal-dialog modal-{{.Size}}" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="{{.ID}}Title">{{langHtml .Title}}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <div class="modal-body">
                {{langHtml .Body}}
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-dismiss="modal">{{lang "Close"}}</button>
                {{if .Footer}}
                <button type="button" class="btn btn-primary">{{langHtml .Footer}}</button>
                {{end}}
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
        <span class="progress-text">{{langHtml .Title}}</span>
        <span class="progress-number"><b>{{.Molecular}}</b>/{{.Denominator}}</span>

        <div class="progress sm">
            {{if .IsHexColor}}
                <div class="progress-bar" style="width: {{.Percent}}%;background-color: {{.Color}} !important;"></div>
            {{else}}
                <div class="progress-bar progress-bar-{{.Color}}" style="width: {{.Percent}}%"></div>
            {{end}}
        </div>
    </div>
{{end}}`, "components/row": `{{define "row"}}
<div class="row">{{langHtml .Content}}</div>
{{end}}`, "components/smallbox": `{{define "smallbox"}}
    <div class="small-box bg-{{.Color}}">
        <div class="inner">
            <h3>{{langHtml .Value}}</h3>
            <p>{{langHtml .Title}}</p>
        </div>
        <div class="icon">
            <i class="fa {{.Icon}}"></i>
        </div>
        <a href="{{.Url}}" class="small-box-footer">
            {{lang "more"}}
            <i class="fa fa-arrow-circle-right"></i>
        </a>
    </div>
{{end}}`, "components/table/box-header": `{{define "box-header"}}
    <div class="pull-right">

        <div class="dropdown pull-right column-selector" style="margin-right: 10px">
            <button type="button" class="btn btn-sm btn-instagram dropdown-toggle" data-toggle="dropdown">
                <i class="fa fa-table"></i>
                &nbsp;
                <span class="caret"></span>
            </button>
            <ul class="dropdown-menu" role="menu" style="padding: 10px;max-height: 400px;overflow: scroll;">
                <li>
                    <ul style="padding: 0;">
                        {{range $key, $head := .Thead}}
                            <li class="checkbox icheck" style="margin: 0;">
                                <label style="width: 100%;padding: 3px;">
                                    <input type="checkbox" class="column-select-item" data-id="{{index $head "field"}}"
                                           style="position: absolute; opacity: 0;">&nbsp;&nbsp;&nbsp;{{index $head "head"}}
                                </label>
                            </li>
                        {{end}}
                    </ul>
                </li>
                <li class="divider">
                </li>
                <li class="text-right">
                    <button class="btn btn-sm btn-default column-select-all">{{lang "all"}}</button>&nbsp;&nbsp;
                    <button class="btn btn-sm btn-primary column-select-submit">{{lang "submit"}}</button>
                </li>
            </ul>
        </div>

        <div class="btn-group pull-right" style="margin-right: 10px">
            <a href="" class="btn btn-sm btn-primary" data-toggle="modal" data-target="#filter-modal"><i
                        class="fa fa-filter"></i>&nbsp;&nbsp;{{lang "filter"}}</a>
            <a href="{{.InfoUrl}}" class="btn btn-sm btn-facebook"><i
                        class="fa fa-undo"></i>&nbsp;&nbsp;{{lang "Reset"}}
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
                        <h4 class="modal-title" id="myModalLabel">{{lang "filter"}}</h4>
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
                                            <input type="text" class="form-control id"
                                                   placeholder="{{index $filter "title"}}"
                                                   name="{{index $filter "name"}}" value="">
                                        </div>
                                    </div>
                                {{end}}
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button type="submit" id="filter-btn"
                                    class="btn btn-primary submit">{{lang "Submit"}}</button>
                            <button type="reset" class="btn btn-warning pull-left">{{lang "Reset"}}</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>

        <script>
            $("#filter-btn").click(function () {
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
            {{if .NewUrl}}
                <a href="{{.NewUrl}}" class="btn btn-sm btn-success">
                    <i class="fa fa-save"></i>&nbsp;&nbsp;{{lang "New"}}
                </a>
            {{end}}
            {{if .ExportUrl}}
                <a href="javascript:;" class="btn btn-sm btn-success" id="export-btn">
                    <i class="fa fa-save"></i>&nbsp;&nbsp;{{lang "Export"}}
                </a>

                <script>
                    $("#export-btn").click(function () {
                        let form = $("<form>");
                        form.attr("style", "display:none");
                        form.attr("target", "");
                        form.attr("method", "post");
                        form.attr("action",{{.ExportUrl}});
                        let input1 = $("<input>");
                        input1.attr("type", "hidden");
                        input1.attr("name", "time");
                        input1.attr("value", (new Date()).getMilliseconds());
                        $("body").append(form);
                        form.append(input1);
                        form.submit();
                        form.remove()
                    });
                </script>
            {{end}}
        </div>
    </div>
    <span>
    {{if .DeleteUrl}}
        <input type="checkbox" class="grid-select-all" style="position: absolute; opacity: 0;">
        <div class="btn-group">
            <a class="btn btn-sm btn-default">{{lang "Action"}}</a>
            <button type="button" class="btn btn-sm btn-default dropdown-toggle" data-toggle="dropdown">
            <span class="caret"></span>
            <span class="sr-only">{{lang "Toggle Dropdown"}}</span>
            </button>
            <ul class="dropdown-menu" role="menu">
                <li><a href="#" class="grid-batch-0">{{lang "Delete"}}</a></li>
                {{if .ExportUrl}}
                    <li><a href="#" class="grid-batch-1">{{lang "Export"}}</a></li>
                {{end}}
            </ul>
        </div>
    {{end}}
    <a class="btn btn-sm btn-primary grid-refresh">
        <i class="fa fa-refresh"></i> {{lang "Refresh"}}
    </a>
</span>
{{end}}`, "components/table": `{{define "table"}}
    <table class="table table-hover" style="min-width: {{.MinWidth}}px;">
        {{if eq .Type "table"}}
            <thead>
            <tr>
                {{range $key, $head := .Thead}}
                    <th>
                        {{index $head "head"}}
                    </th>
                {{end}}
            </tr>
            </thead>
        {{end}}
        <tbody>
        {{if eq .Type "data-table"}}
            <tr>
                {{if eq .IsTab false}}
                    <th></th>
                {{end}}
                {{range $key, $head := .Thead}}
                    {{if eq (index $head "hide") "0"}}
                        {{if eq (index $head "width") "0"}}
                            <th>
                        {{else}}
                            <th style="width: {{index $head "width"}}px">
                        {{end}}
                        {{index $head "head"}}
                        {{if eq (index $head "sortable") "1"}}
                            <a class="fa fa-fw fa-sort" id="sort-{{index $head "field"}}"
                               href="?__sort={{index $head "field"}}&__sort_type=desc"></a>
                        {{end}}
                        </th>
                    {{end}}
                {{end}}
                {{if eq .NoAction false}}
                    <th>{{lang "operation"}}</th>
                {{end}}
            </tr>
        {{end}}


        {{$NoAction := .NoAction}}
        {{$Action := .Action}}
        {{$Thead := .Thead}}
        {{$Type := .Type}}
        {{$EditUrl := .EditUrl}}
        {{$IsTab := .IsTab}}
        {{$DeleteUrl := .DeleteUrl}}
        {{$PrimaryKey := .PrimaryKey}}
        {{range $key1, $info := .InfoList}}
            <tr>
                {{if eq $Type "data-table"}}
                    {{if eq $IsTab false}}
                        <td>
                            {{if $DeleteUrl}}
                                <input type="checkbox" class="grid-row-checkbox" data-id="{{index $info $PrimaryKey}}"
                                       style="position: absolute; opacity: 0;">
                            {{end}}
                        </td>
                    {{end}}
                {{end}}
                {{if eq $Type "data-table"}}
                    {{range $key2, $head2 := $Thead}}
                        {{if eq (index $head2 "hide") "0"}}
                            <td>{{index $info (index $head2 "field")}}</td>
                        {{end}}
                    {{end}}
                    {{if eq $NoAction false}}
                        <td>
                            {{$Action}}
                            {{if $EditUrl}}
                                <a href='{{$EditUrl}}&id={{index $info $PrimaryKey}}'><i class="fa fa-edit"></i></a>
                            {{end}}
                            {{if $DeleteUrl}}
                                <a href="javascript:void(0);" data-id='{{index $info $PrimaryKey}}'
                                   class="grid-row-delete"><i
                                            class="fa fa-trash"></i></a>
                            {{end}}
                        </td>
                    {{end}}
                {{else}}
                    {{range $key2, $head2 := $Thead}}
                        <td>{{index $info (index $head2 "head")}}</td>
                    {{end}}
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

            selectedRows = function () {
                let selected = [];
                $('.grid-row-checkbox:checked').each(function () {
                    selected.push($(this).data('id'));
                });
                return selected;
            };

            $(function () {
                $('.grid-select-all').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function (event) {
                    console.log("event", event);
                    if (this.checked) {
                        $('.grid-row-checkbox').iCheck('check');
                    } else {
                        $('.grid-row-checkbox').iCheck('uncheck');
                    }
                });
                let items = $('.column-select-item');
                iCheck(items);
                iCheck($('.grid-row-checkbox'));
                let columns = getQueryVariable("__columns");
                if (columns === "") {
                    items.iCheck('check');
                } else {
                    let columnsArr = columns.split(",");
                    for (let i = 0; i < columnsArr.length; i++) {
                        for (let j = 0; j < items.length; j++) {
                            if (columnsArr[i] === $(items[j]).attr("data-id")) {
                                $(items[j]).iCheck('check');
                            }
                        }
                    }
                }
            });

            selectedAllFieldsRows = function () {
                let selected = [];
                $('.column-select-item:checked').each(function () {
                    selected.push($(this).data('id'));
                });
                return selected;
            };

            $('.column-select-all').on('click', function (event) {
                $('.column-select-item').iCheck('check');
            });

            $('.column-select-submit').on('click', function () {

                let newUrl = "";
                if (getQueryVariable('__columns') !== '') {
                    newUrl = replaceParamVal('__columns', selectedAllFieldsRows().join(","));
                } else {
                    if (location.href.indexOf("?") > 0) {
                        newUrl = location.href + "&__columns=" + selectedAllFieldsRows().join(",");
                    } else {
                        newUrl = location.href + "?__columns=" + selectedAllFieldsRows().join(",");
                    }
                }

                $.pjax({url: newUrl, container: '#pjax-container'})

                toastr.success('{{lang "reload succeeded"}} !');
            });

            function replaceParamVal(paramName, replaceWith) {
                let oUrl = this.location.href.toString();
                let re = eval('/(' + paramName + '=)([^&]*)/gi');
                return oUrl.replace(re, paramName + '=' + replaceWith);
            }

            $('.grid-batch-0').on('click', function () {
                DeletePost(selectedRows().join())
            });

            {{if .ExportUrl}}
            $('.grid-batch-1').on('click', function () {
                ExportAll(selectedRows().join())
            });

            function ExportAll(id) {
                let form = $("<form>");
                form.attr("style", "display:none");
                form.attr("target", "");
                form.attr("method", "post");
                form.attr("action",{{.ExportUrl}});
                let input1 = $("<input>");
                input1.attr("type", "hidden");
                input1.attr("name",{{.PrimaryKey}});
                input1.attr("value", id);
                $("body").append(form);
                form.append(input1);
                form.submit();
                form.remove()
            }
            {{end}}

            function iCheck(el) {
                el.iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function () {
                    if (this.checked) {
                        $(this).closest('tr').css('background-color', "#ffffd5");
                    } else {
                        $(this).closest('tr').css('background-color', '');
                    }
                });
            }

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
                    function () {
                        $.ajax({
                            method: 'post',
                            url: {{.DeleteUrl}},
                            data: {
                                id: id
                            },
                            success: function (data) {
                                $.pjax.reload('#pjax-container');
                                if (typeof (data) === "string") {
                                    data = JSON.parse(data);
                                }
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

            $(function () {
                let sort = getQueryVariable("__sort");
                let sort_type = getQueryVariable("__sort_type");

                if (sort !== "" && sort_type !== "") {
                    let sortFa = $('#sort-' + sort);
                    if (sort_type === 'asc') {
                        sortFa.attr('href', '?__sort=' + sort + "&__sort_type=desc")
                    } else {
                        sortFa.attr('href', '?__sort=' + sort + "&__sort_type=asc")
                    }
                    sortFa.removeClass('fa-sort');
                    sortFa.addClass('fa-sort-amount-' + sort_type);
                }
            });

            function getQueryVariable(variable) {
                let query = window.location.search.substring(1);
                let vars = query.split("&");
                for (let i = 0; i < vars.length; i++) {
                    let pair = vars[i].split("=");
                    if (pair[0] === variable) {
                        return pair[1];
                    }
                }
                return "";
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
    <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/chart.js"}}"></script>
    <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/fontawesome-iconpicker.min.js"}}"></script>
    {{if lang .Panel.Title}}
        <section class="content-header">
            <h1>
                {{lang .Panel.Title}}
                <small>{{lang .Panel.Description}}</small>
            </h1>
            <ol class="breadcrumb" style="margin-right: 30px;">
                <li><a href="{{.IndexUrl}}"><i class="fa fa-dashboard"></i> {{lang "home"}}</a></li>
                {{.Menu.FormatPath}}
            </ol>
        </section>
    {{end}}

    <!-- Main content -->
    <section class="content">
        {{.Panel.Content}}
    </section>
{{end}}`, "control_panel": `{{define "control_panel"}}
    <div class="control-sidebar-bg" style="position: fixed; height: auto;"></div>
    <aside class="control-sidebar control-sidebar-dark control-sidebar-open"
           style="position: fixed; max-height: 100%; overflow: auto;">
        <ul class="nav nav-tabs nav-justified control-sidebar-tabs">
            <li class="active"><a href="#control-sidebar-setting-tab" data-toggle="tab" aria-expanded="true"><i
                            class="fa fa-wrench"></i></a></li>
            <li><a href="#control-sidebar-home-tab" data-toggle="tab"><i class="fa fa-home"></i></a></li>
            <li><a href="#control-sidebar-settings-tab" data-toggle="tab"><i class="fa fa-gears"></i></a></li>
        </ul>
        <div class="tab-content">
            <div class="tab-pane active" id="control-sidebar-setting-tab">
                <h4 class="control-sidebar-heading">{{lang "layout"}}</h4>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-layout="fixed"
                                                                                         class="pull-right"
                                                                                         checked="checked"> 固定布局</label>
                    <p>盒子模型和固定布局不能同时启作用</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-layout="layout-boxed"
                                                                                         class="pull-right">
                        盒子布局</label>
                    <p>盒子布局最大宽度将被限定为1250px</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-layout="sidebar-collapse"
                                                                                         class="pull-right">
                        切换菜单栏</label>
                    <p>切换菜单栏的展示或收起</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-enable="expandOnHover"
                                                                                         class="pull-right">
                        菜单栏自动展开</label>
                    <p>鼠标移到菜单栏自动展开</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-menu="show-submenu"
                                                                                         class="pull-right">
                        显示菜单栏子菜单</label>
                    <p>菜单栏子菜单将始终显示</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-menu="disable-top-badge"
                                                                                         class="pull-right"
                                                                                         checked="checked">
                        禁用顶部彩色小角标</label>
                    <p>左边菜单栏的彩色小角标不受影响</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-controlsidebar="control-sidebar-open"
                                                                                         class="pull-right">
                        切换右侧操作栏</label>
                    <p>切换右侧操作栏覆盖或独占</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-sidebarskin="toggle"
                                                                                         class="pull-right">
                        切换右侧操作栏背景</label>
                    <p>将右侧操作栏背景亮色或深色切换</p></div>
                <h4 class="control-sidebar-heading">{{lang "skin"}}</h4>
                <ul class="list-unstyled clearfix skin-list">
                    <li><a href="javascript:;" data-skin="skin-blue" style="" class="clearfix full-opacity-hover">
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 7px; background: #367fa9;"></span><span
                                        class="bg-light-blue"
                                        style="display:block; width: 80%; float: left; height: 7px;"></span></div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Blue</p></li>
                    <li><a href="javascript:;" data-skin="skin-white" class="clearfix full-opacity-hover">
                            <div style="box-shadow: 0 0 2px rgba(0,0,0,0.1)" class="clearfix"><span
                                        style="display:block; width: 20%; float: left; height: 7px; background: #fefefe;"></span><span
                                        style="display:block; width: 80%; float: left; height: 7px; background: #fefefe;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">White</p></li>
                    <li><a href="javascript:;" data-skin="skin-purple" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-purple-active"></span><span class="bg-purple"
                                                                             style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Purple</p></li>
                    <li><a href="javascript:;" data-skin="skin-green" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-green-active"></span><span class="bg-green"
                                                                            style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Green</p></li>
                    <li><a href="javascript:;" data-skin="skin-red" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-red-active"></span><span class="bg-red"
                                                                          style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Red</p></li>
                    <li><a href="javascript:;" data-skin="skin-yellow" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-yellow-active"></span><span class="bg-yellow"
                                                                             style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Yellow</p></li>
                    <li><a href="javascript:;" data-skin="skin-blue-light" class="clearfix full-opacity-hover">
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 7px; background: #367fa9;"></span><span
                                        class="bg-light-blue"
                                        style="display:block; width: 80%; float: left; height: 7px;"></span></div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">Blue Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-white-light" class="clearfix full-opacity-hover">
                            <div style="box-shadow: 0 0 2px rgba(0,0,0,0.1)" class="clearfix"><span
                                        style="display:block; width: 20%; float: left; height: 7px; background: #fefefe;"></span><span
                                        style="display:block; width: 80%; float: left; height: 7px; background: #fefefe;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">White Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-purple-light" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-purple-active"></span><span class="bg-purple"
                                                                             style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">Purple Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-green-light" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-green-active"></span><span class="bg-green"
                                                                            style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">Green Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-red-light" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-red-active"></span><span class="bg-red"
                                                                          style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">Red Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-yellow-light" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-yellow-active"></span><span class="bg-yellow"
                                                                             style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px;">Yellow Light</p></li>
                </ul>
            </div>
            <div class="tab-pane" id="control-sidebar-home-tab">
                <h4 class="control-sidebar-heading">{{lang "home"}}</h4>
            </div>
            <div class="tab-pane" id="control-sidebar-settings-tab">
                <h4 class="control-sidebar-heading">{{lang "config"}}</h4>
            </div>
        </div>
    </aside>
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

        <link rel="stylesheet" href="{{link .CdnUrl .UrlPrefix "/assets/dist/css/all.min.css"}}">

        <!--[if lt IE 9]>
        <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/html5shiv.min.js"}}"></script>
        <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/respond.min.js"}}"></script>
        <![endif]-->

        <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/all.min.js"}}"></script>

        {{.CustomHeadHtml}}
    </head>
{{end}}`, "header": `{{define "header"}}
    <header class="main-header">
        <!-- Logo -->
        <a href="{{.UrlPrefix}}" class="logo">
            <!-- mini logo for sidebar mini 50x50 pixels -->
            <span class="logo-mini">{{.MiniLogo}}</span>
            <!-- logo for regular state and mobile devices -->
            <span class="logo-lg">{{.Logo}}</span>
        </a>
        <!-- Header Navbar: style can be found in header.less -->
        <nav class="navbar navbar-static-top">
            <div id="firstnav">
                <a href="#" class="sidebar-toggle" data-toggle="offcanvas" role="button">
                    <span class="sr-only">Toggle navigation</span>
                </a>

                <ul class="nav nav-tabs nav-addtabs" style="max-width: 60%">
                </ul>

                {{ template "admin_panel" . }}
            </div>
        </nav>
    </header>
{{end}}`, "js": `{{define "js"}}
    <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/all_2.min.js"}}"></script>
{{end}}`, "layout": `{{define "layout"}}

    <!DOCTYPE html>
    <html>

    {{ template "head" . }}

    <body class="hold-transition {{.ColorScheme}} sidebar-mini">
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
    {{.CustomFootHtml}}
    </html>

{{end}}
`, "menu": `{{define "menu"}}
    <ul class="sidebar-menu" data-widget="tree">
        {{$UrlPrefix := .UrlPrefix}}
        {{range $key, $list := .Menu.List }}
            {{if eq (len $list.ChildrenList) 0}}
                {{if $list.Header}}
                    <li class="header" data-rel="external">{{$list.Header}}</li>
                {{end}}
                <li class='{{$list.Active}}'>
                    {{if eq $list.Url "/"}}
                    <a href='{{$UrlPrefix}}'>
                    {{else if isLinkUrl $list.Url}}
                    <a href='{{$list.Url}}'>
                    {{else}}
                    <a href='{{$UrlPrefix}}{{$list.Url}}'>
                    {{end}}
                        <i class="fa {{$list.Icon}}"></i><span> {{$list.Name}}</span>
                        <span class="pull-right-container"><!-- <small class="label pull-right bg-green">new</small> --></span>
                    </a>
                </li>
            {{else}}
                <li class="treeview {{$list.Active}}">
                    <a href="#">
                        <i class="fa {{$list.Icon}}"></i><span> {{$list.Name}}</span>
                        <span class="pull-right-container">
                        <i class="fa fa-angle-left pull-right"></i>
                    </span>
                    </a>
                    <ul class="treeview-menu">
                        {{range $key2, $item := $list.ChildrenList}}
                            <li><a href="{{$UrlPrefix}}{{$item.Url}}"><i class="fa {{$item.Icon}}"></i> {{$item.Name}}
                                </a></li>
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
                        <img src="{{.UrlPrefix}}/assets/dist/img/avatar04.png" class="img-circle" alt="User Image">
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
