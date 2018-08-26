package template

var Adminlte = map[string]string{"admin_panel":`{{define "admin_panel"}}
<div class="navbar-custom-menu">
    <ul class="nav navbar-nav">
        <!-- User Account: style can be found in dropdown.less -->
        <li class="dropdown user user-menu">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                <img src="./{{.User.Avatar}}" class="user-image" alt="User Image">
                <span class="hidden-xs">{{.User.Name}}</span>
            </a>
            <ul class="dropdown-menu">
                <!-- User image -->
                <li class="user-header">
                    <img src="./{{.User.Avatar}}" class="img-circle" alt="User Image">
                    <p>
                        {{.User.Name}} -{{.User.Level}}
                        <small>{{.User.CreateAt}}</small>
                    </p>
                </li>
                <li class="user-footer">
                    <div class="pull-left">
                        <a href="/info/manager/edit?id={{.User.ID}}" class="btn btn-default btn-flat">Setting</a>
                    </div>
                    <div class="pull-right">
                        <a href="/logout" class="btn btn-default btn-flat">Sign out</a>
                    </div>
                </li>
            </ul>
        </li>
    </ul>
</div>
{{end}}`,"content":`{{define "content"}}
<section class="content-header">
    <h1>
        {{.Panel.Title}}
        <small>{{.Panel.Description}}</small>
    </h1>
</section>

<!-- Main content -->
<section class="content">
    {{.Panel.Content}}
</section>
{{end}}`,"footer":`{{define "footer"}}
<footer class="main-footer">
    <div class="pull-right hidden-xs">
        <b>Version</b> {{.System.Version}}
    </div>
    <strong>Copyright &copy; 2018- <a href="https://github.com/chenhg5/go-admin">GoAdmin</a>.</strong> All rights
    reserved.
</footer>
{{end}}`,"head":`{{define "head"}}
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>GoAdmin</title>
    <!-- Tell the browser to be responsive to screen width -->
    <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
    <!-- Bootstrap 3.3.7 -->
    <link rel="stylesheet" href="../../assets/bootstrap/dist/css/bootstrap.min.css">
    <!-- Font Awesome -->
    <link rel="stylesheet" href="../../assets/font-awesome/css/font-awesome.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="../../assets/Ionicons/css/ionicons.min.css">
    <!-- DataTables -->
    <link rel="stylesheet" href="../../assets/datatables.net-bs/css/dataTables.bootstrap.min.css">
    <!-- iCheck -->
    <link rel="stylesheet" href="../../assets/iCheck/minimal/_all.css">
    <link rel="stylesheet" href="../../assets/iCheck/futurico/futurico.css">
    <link rel="stylesheet" href="../../assets/iCheck/polaris/polaris.css">
    <link rel="stylesheet" href="../../assets/toastr/build/toastr.min.css">
    <link rel="stylesheet" href="../../assets/nprogress/nprogress.css">
    <link rel="stylesheet" href="../../assets/select2/select2.min.css">
    <link rel="stylesheet" href="../../assets/sweetalert/dist/sweetalert.css">
    <link rel="stylesheet" href="../../assets/fileinput/fileinput.min.css">
    <link rel="stylesheet" href="../../assets/nestable/nestable.css">
    <link rel="stylesheet" href="../../assets/duallistbox/bootstrap-duallistbox.min.css">
    <link rel="stylesheet" href="../../assets/fontawesome-iconpicker/dist/css/fontawesome-iconpicker.min.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="../../assets/dist/css/AdminLTE.min.css">
    <link rel="stylesheet" href="../../assets/dist/css/skins/skin-black.css">
    <!--[if lt IE 9]>
    <script src="../../assets/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="../../assets/respond/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- Google Font -->
    <link rel="stylesheet" href="../../assets/googleapis/font.css">

    <script src="../../assets/jQuery/jQuery-2.1.4.min.js"></script>
    <script src="../../assets/nestable/jquery.nestable.js"></script>
    <script src="../../assets/dist/js/adminlte.min.js"></script>
</head>
{{end}}`,"header":`{{define "header"}}
<header class="main-header">
    <!-- Logo -->
    <a href="/" class="logo">
        <!-- mini logo for sidebar mini 50x50 pixels -->
        <span class="logo-mini"><b>G</b>A</span>
        <!-- logo for regular state and mobile devices -->
        <span class="logo-lg"><b>Go</b>Admin</span>
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
{{end}}`,"js":`{{define "js"}}
<script src="../../assets/bootstrap/dist/js/bootstrap.min.js"></script>
<script src="../../assets/datatables.net/js/jquery.dataTables.min.js"></script>
<script src="../../assets/datatables.net-bs/js/dataTables.bootstrap.min.js"></script>
<script src="../../assets/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<script src="../../assets/fastclick/lib/fastclick.js"></script>
<script src="../../assets/select2/select2.full.min.js"></script>
<script src="../../assets/fileinput/fileinput.min.js"></script>
<script src="../../assets/iCheck/icheck.min.js"></script>
<script src="../../assets/nprogress/nprogress.js"></script>
<script src="../../assets/toastr/build/toastr.min.js"></script>
<script src="../../assets/bootstrap3-editable/js/bootstrap-editable.min.js"></script>
<script src="../../assets/jquery-pjax/jquery.pjax.js"></script>
<script src="../../assets/sweetalert/dist/sweetalert.min.js"></script>
<script src="../../assets/duallistbox/jquery.bootstrap-duallistbox.min.js"></script>
<script src="../../assets/fontawesome-iconpicker/dist/js/fontawesome-iconpicker.min.js"></script>
<script src="../../assets/dist/js/info.js"></script>
{{end}}`,"layout":`{{define "layout"}}

<!DOCTYPE html>
<html>

{{ template "head" }}

<body class="hold-transition skin-black sidebar-mini">
<div class="wrapper">

    {{ template "header" . }}

    {{ template "sidebar" . }}

    <div class="content-wrapper" id="pjax-container">

    {{ template "content" . }}

    </div>

    {{ template "footer" . }}

</div>

{{ template "js" }}

</body>
</html>

{{end}}
`,"menu":`{{define "menu"}}
<ul class="sidebar-menu" data-widget="tree">
{{range $key, $list := .}}
    {{if eq (len $list.ChildrenList) 0}}
        <li class='{{$list.Active}}'>
            <a href='{{$list.Url}}'>
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
                <li><a href="{{$item.Url}}"><i class="fa {{$item.Icon}}"></i> {{$item.Name}}</a></li>
            {{end}}
            </ul>
        </li>
    {{end}}
{{end}}
</ul>
{{end}}`,"sidebar":`{{define "sidebar"}}
<!-- Left side column. contains the logo and sidebar -->
<aside class="main-sidebar">
    <!-- sidebar: style can be found in sidebar.less -->
    <section class="sidebar">
        <!-- Sidebar user panel -->
        <div class="user-panel">
            <div class="pull-left image">
                <img src="../../assets/dist/img/avatar04.png" class="img-circle" alt="User Image">
            </div>
            <div class="pull-left info">
                admin
                <a href="#"><i class="fa fa-circle text-success"></i> Online</a>
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

        {{ template "menu" .Menu.GlobalMenuList }}

    </section>
    <!-- /.sidebar -->
</aside>
{{end}}`,"components/box":`{{define "box"}}
<div class="box">
    <div class="box-header {{.HeadBorder}}">
        {{.Header}}
    </div>
    <div class="box-body table-responsive no-padding">
        {{.Body}}
    </div>
    <div class="box-footer clearfix">
        {{.Footer}}
    </div>
</div>
{{end}}`,"components/col":`{{define "col"}}
<div class="col-{{.Type}}-{{.Width}}">{{.Content}}</div>
{{end}}`,"components/form":`{{define "form"}}
<script src="../../assets/select2/select2.full.min.js"></script>
<script src="../../assets/fileinput/fileinput.min.js"></script>
<script src="../../assets/duallistbox/jquery.bootstrap-duallistbox.min.js"></script>
<div class="box box-info">
    <div class="box-header with-border">
        <h3 class="box-title">Edit</h3>
        <div class="box-tools">
            <div class="btn-group pull-right" style="margin-right: 10px">
                <a href='{{.InfoUrl}}' class="btn btn-sm btn-default"><i class="fa fa-list"></i> List</a>
            </div>
            <div class="btn-group pull-right" style="margin-right: 10px">
                <a href='{{.InfoUrl}}' class="btn btn-sm btn-default form-history-back"><i class="fa fa-arrow-left"></i> Back</a>
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
                    {{else if eq $data.FormType "textarea"}}
                        {{ template "form_textarea" $data }}
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
                        Save
                    </button>
                </div>

                <div class="btn-group pull-left">
                    <button type="reset" class="btn btn-warning">Reset</button>
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
{{end}}`,"components/image":`{{define "image"}}
<img src="{{.Src}}" width="{{.Width}}" height="{{.Height}}">
{{end}}`,"components/infobox":`{{define "infobox"}}
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
{{end}}`,"components/label":`{{define "label"}}
<span class="label label-{{.Color}}">{{.Content}}</span>
{{end}}`,"components/paninator":`{{define "paninator"}}
Showing <b>0</b> to <b>10</b> of <b>1</b> entries
<ul class="pagination pagination-sm no-margin pull-right">
    <!-- Previous Page Link -->
    <li class="page-item disabled">
        <span class="page-link">«</span>
    </li>

    <!-- Array Of Links -->
    <li class="page-item active"><span class="page-link">1</span></li>


    <!-- Next Page Link -->
    <li class="page-item ">
        <a class="page-link" href="/info/manager?page=2&amp;pageSize=10&amp;sort=id&amp;sort_type=desc" rel="next">»</a>
    </li>
</ul>

<label class="control-label pull-right" style="margin-right: 10px; font-weight: 100;">

    <small>Show</small>&nbsp;
    <select class="input-sm grid-per-pager" name="per-page">
        <option value="/info/manager?page=1&amp;pageSize=10&amp;sort=id&amp;sort_type=desc&amp;pageSize=10" selected="''">
            10
        </option>
        <option value="/info/manager?page=1&amp;pageSize=10&amp;sort=id&amp;sort_type=desc&amp;pageSize=20">
            20
        </option>
        <option value="/info/manager?page=1&amp;pageSize=10&amp;sort=id&amp;sort_type=desc&amp;pageSize=30">
            30
        </option>
        <option value="/info/manager?page=1&amp;pageSize=10&amp;sort=id&amp;sort_type=desc&amp;pageSize=50">
            50
        </option>
        <option value="/info/manager?page=1&amp;pageSize=10&amp;sort=id&amp;sort_type=desc&amp;pageSize=100">
            100
        </option>s
    </select>
    &nbsp;<small>entries</small>
</label>
{{end}}`,"components/row":`{{define "row"}}
<div class="row">{{.Content}}</div>
{{end}}`,"components/table":`{{define "table"}}
    <table class="table table-hover">
        <tbody>
        <tr>
            {{if eq .Type "data-table"}}
                <th></th>
            {{end}}
            {{range $key, $head := .Thead}}
                <th>
                    {{index $head "head"}}
                    {{if eq (index $head "sortable") "1"}}
                        <a class="fa fa-fw fa-sort" href=""></a>
                    {{end}}
                </th>
            {{end}}
            {{if eq .Type "data-table"}}
                <th>操作</th>
            {{end}}
        </tr>

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
{{end}}`,"components/form/default":`{{define "form_default"}}
<label class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <div class="box box-solid box-default no-margin">
        <div class="box-body">
            {{.Value}}
        </div>
    </div>
</div>
{{end}}`,"components/form/file":`{{define "form_file"}}
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
{{end}}`,"components/form/password":`{{define "form_password"}}
<label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <div class="input-group">
        <span class="input-group-addon"><i class="fa fa-eye-slash"></i></span>
        <input type="password" id="{{.Field}}" name="{{.Field}}" value="{{.Value}}" class="form-control password" placeholder="Input {{.Head}}">
    </div>
</div>
{{end}}`,"components/form/select":`{{define "form_select"}}
<label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <select class="form-control {{.Field}} select2-hidden-accessible" style="width: 100%;" name="{{.Field}}[]" multiple="" data-placeholder="Input {{.Head}}" tabindex="-1" aria-hidden="true">
        {{range $key, $v := .Options }}
            <option value='{{index $v "value"}}' {{index $v "selected"}}>{{index $v "field"}}</option>
        {{end}}
    </select>
    <input type="hidden" name="{{.Field}}[]">
    <!--<span class="help-block">
        <i class="fa fa-info-circle"></i>&nbsp;All methods if empty
    </span>-->
</div>
<script>
    $(".{{.Field}}").select2({
        allowClear: true
    });
</script>
{{end}}`,"components/form/selectbox":`{{define "form_selectbox"}}
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
{{end}}`,"components/form/text":`{{define "form_text"}}
<label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <div class="input-group">
        <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
        <input type="text" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}' class="form-control json" placeholder="Input {{.Head}}">
    </div>
</div>
{{end}}`,"components/form/textarea":`{{define "form_textarea"}}
<label for="{{.Field}}" class="col-sm-2 control-label">{{.Head}}</label>
<div class="col-sm-8">
    <textarea name="{{.Field}}" class="form-control" rows="5" placeholder="Input {{.Head}}">{{.Value}}</textarea>
</div>
{{end}}`,"components/table/box-header":`{{define "box-header"}}
<div class="pull-right">
    <div class="btn-group pull-right" style="margin-right: 10px">
        <a href="" class="btn btn-sm btn-primary" data-toggle="modal" data-target="#filter-modal"><i class="fa fa-filter"></i>&nbsp;&nbsp;Filter</a>
        <a href="/story/word" class="btn btn-sm btn-facebook"><i class="fa fa-undo"></i>&nbsp;&nbsp;Reset</a>
    </div>
    <div class="modal fade" id="filter-modal" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">×</span>
                        <span class="sr-only">Close</span>
                    </button>
                    <h4 class="modal-title" id="myModalLabel">Filter</h4>
                </div>
                <form action="/story/word" method="get" pjax-container="">
                    <div class="modal-body">
                        <div class="form">
                            <div class="form-group">
                                <div class="form-group">
                                    <label>ID</label>
                                    <div class="input-group">
                                        <div class="input-group-addon">
                                            <i class="fa fa-pencil"></i>
                                        </div>
                                        <input type="text" class="form-control id" placeholder="ID" name="id" value="">
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="submit" class="btn btn-primary submit">Submit</button>
                        <button type="reset" class="btn btn-warning pull-left">Reset
                        </button>
                    </div>
                </form>
            </div>
        </div>
    </div>

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

    <div class="btn-group pull-right" style="margin-right: 10px">

        <a href="{{.NewUrl}}" class="btn btn-sm btn-success">

            <i class="fa fa-save"></i>&nbsp;&nbsp;New
        </a>
    </div>
</div>
<span>
    <input type="checkbox" class="grid-select-all" style="position: absolute; opacity: 0;">
    <div class="btn-group">
        <a class="btn btn-sm btn-default">Action</a>
        <button type="button" class="btn btn-sm btn-default dropdown-toggle" data-toggle="dropdown">
        <span class="caret"></span>
        <span class="sr-only">Toggle Dropdown</span>
        </button>
        <ul class="dropdown-menu" role="menu">
            <li><a href="#" class="grid-batch-0">Delete</a></li>
        </ul>
    </div>
    <a class="btn btn-sm btn-primary grid-refresh">
        <i class="fa fa-refresh"></i> Refresh
    </a>
</span>
{{end}}`,"login/theme1":`{{define "login_theme1"}}
<!DOCTYPE html>
<!--[if lt IE 7]>      <html class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
<!--[if IE 7]>         <html class="no-js lt-ie9 lt-ie8"> <![endif]-->
<!--[if IE 8]>         <html class="no-js lt-ie9"> <![endif]-->
<!--[if gt IE 8]><!-->
<html class="no-js"> <!--<![endif]-->
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>GoAdmin Login</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="Free HTML5 Template by FreeHTML5.co"/>
    <meta name="keywords" content="free html5, free template, free bootstrap, html5, css3, mobile first, responsive"/>

    <!-- Facebook and Twitter integration -->
    <meta property="og:title" content=""/>
    <meta property="og:image" content=""/>
    <meta property="og:url" content=""/>
    <meta property="og:site_name" content=""/>
    <meta property="og:description" content=""/>
    <meta name="twitter:title" content=""/>
    <meta name="twitter:image" content=""/>
    <meta name="twitter:url" content=""/>
    <meta name="twitter:card" content=""/>

    <!-- Place favicon.ico and apple-touch-icon.png in the root directory -->
    <link rel="shortcut icon" href="http://yanshi.sucaihuo.com/modals/26/2674/demo/favicon.ico">

    <link href='https://fonts.googleapis.com/css?family=Open+Sans:400,700,300' rel='stylesheet' type='text/css'>

    <link rel="stylesheet" href="../../assets/login/css/bootstrap.min.css">
    <link rel="stylesheet" href="../../assets/login/css/animate.css">
    <link rel="stylesheet" href="../../assets/login/css/style.css">

    <!-- Modernizr JS -->
    <script src="../../assets/login/js/modernizr-2.6.2.min.js"></script>
    <!-- FOR IE9 below -->
    <!--[if lt IE 9]>
    <script src="../../assets/login/js/respond.min.js"></script>
    <![endif]-->

</head>
<body>

<div class="container">
    <div class="copyrights">Collect from <a href="https://github.com/chenhg5/go-admin" title="GoAdmin">Go Admin</a>
    </div>
    <div class="row" style="margin-top: 80px;">
        <div class="col-md-4 col-md-offset-4">


            <!-- Start Sign In Form -->
            <form action="/signin" method="post" id="sign-up-form" class="fh5co-form animate-box"
                  data-animate-effect="fadeIn">
                <h2>GoAdmin</h2>
                <div class="form-group">
                    <label for="username" class="sr-only">Username</label>
                    <input type="text" class="form-control" id="username" placeholder="用户名" autocomplete="off">
                </div>
                <div class="form-group">
                    <label for="password" class="sr-only">Password</label>
                    <input type="password" class="form-control" id="password" placeholder="密码" autocomplete="off">
                </div>
                <input type="hidden" class="form-control" value="12313" autocomplete="off">
                <!--<div class="form-group">
                    <label for="remember"><input type="checkbox" id="remember"> Remember Me</label>
                </div>-->
                <!--<div class="form-group">
                    <p>Not registered? <a href="sign-up.html">Sign Up</a> | <a href="forgot.html">Forgot Password?</a>
                    </p>
                </div>-->
                <div class="form-group">
                    <button class="btn btn-primary">登录</button>
                </div>
            </form>
            <!-- END Sign In Form -->

        </div>
    </div>
    <div class="row" style="padding-top: 60px; clear: both;">
        <div class="col-md-12 text-center">
            <p>
                <small>&copy; All Rights Reserved. GoAdmin</small>
            </p>
        </div>
    </div>
</div>

<div id="particles-js">
    <canvas class="particles-js-canvas-el" width="1606" height="1862" style="width: 100%; height: 100%;"></canvas>
</div>

<!-- jQuery -->
<script src="../../assets/jQuery/jQuery-2.1.4.min.js"></script>
<!-- <script src="../../assets/login/js/jquery.min.js"></script> -->
<!-- Bootstrap -->
<script src="../../assets/login/js/bootstrap.min.js"></script>
<!-- Placeholder -->
<script src="../../assets/login/js/jquery.placeholder.min.js"></script>
<!-- Waypoints -->
<script src="../../assets/login/js/jquery.waypoints.min.js"></script>
<!-- Main JS -->
<script src="../../assets/login/js/main.js"></script>
<script src="../../assets/particles/particles.js"></script>

<style type="text/css">
    #particles-js {
        position: absolute;
        width: 100%;
        height: 100%;
        /*background: url(https://qcdn.vwvvwv.com/image/bg/m_bg.png);*/
        background-size: cover;
        top: 0;
        left: 0;
        z-index: -1;
    }
    body {
        color: #ffffff !important;
        background-color: #212121 !important
    }
</style>

<script>
    $("#sign-up-form").submit(function (e) {
        e.preventDefault();
        $.ajax({
            dataType: 'json',
            type: 'POST',
            url: '/signin',
            async: 'true',
            data: {
                'username': $("#username").val(),
                'password': $("#password").val()
            },
            success: function (data) {
                location.href = data.url
            },
            error: function (data) {
                alert("登录失败");
            }
        });
    });
</script>

<script>
    /* -----------------------------------------------
    /* How to use? : Check the GitHub README
    /* ----------------------------------------------- */
    /* To load a config file (particles.json) you need to host this demo (MAMP/WAMP/local)... */
    /*
    particlesJS.load('particles-js', 'particles.json', function() {
      console.log('particles.js loaded - callback');
    });
    */
    /* Otherwise just put the config content (json): */
    particlesJS('particles-js',
            {
                "particles": {
                    "number": {
                        "value": 80,
                        "density": {
                            "enable": true,
                            "value_area": 800
                        }
                    },
                    "color": {
                        "value": "#ffffff"
                    },
                    "shape": {
                        "type": "circle",
                        "stroke": {
                            "width": 0,
                            "color": "#000000"
                        },
                        "polygon": {
                            "nb_sides": 5
                        },
                        "image": {
                            "src": "img/github.svg",
                            "width": 100,
                            "height": 100
                        }
                    },
                    "opacity": {
                        "value": 0.5,
                        "random": false,
                        "anim": {
                            "enable": false,
                            "speed": 1,
                            "opacity_min": 0.1,
                            "sync": false
                        }
                    },
                    "size": {
                        "value": 5,
                        "random": true,
                        "anim": {
                            "enable": false,
                            "speed": 40,
                            "size_min": 0.1,
                            "sync": false
                        }
                    },
                    "line_linked": {
                        "enable": true,
                        "distance": 150,
                        "color": "#ffffff",
                        "opacity": 0.4,
                        "width": 1
                    },
                    "move": {
                        "enable": true,
                        "speed": 6,
                        "direction": "none",
                        "random": false,
                        "straight": false,
                        "out_mode": "out",
                        "attract": {
                            "enable": false,
                            "rotateX": 600,
                            "rotateY": 1200
                        }
                    }
                },
                "interactivity": {
                    "detect_on": "canvas",
                    "events": {
                        "onhover": {
                            "enable": true,
                            "mode": "repulse"
                        },
                        "onclick": {
                            "enable": true,
                            "mode": "push"
                        },
                        "resize": true
                    },
                    "modes": {
                        "grab": {
                            "distance": 400,
                            "line_linked": {
                                "opacity": 1
                            }
                        },
                        "bubble": {
                            "distance": 400,
                            "size": 40,
                            "duration": 2,
                            "opacity": 8,
                            "speed": 3
                        },
                        "repulse": {
                            "distance": 200
                        },
                        "push": {
                            "particles_nb": 4
                        },
                        "remove": {
                            "particles_nb": 2
                        }
                    }
                },
                "retina_detect": true,
                "config_demo": {
                    "hide_card": false,
                    "background_color": "#b61924",
                    "background_image": "",
                    "background_position": "50% 50%",
                    "background_repeat": "no-repeat",
                    "background_size": "cover"
                }
            }
    );
</script>

</body>
</html>
{{end}}`,}