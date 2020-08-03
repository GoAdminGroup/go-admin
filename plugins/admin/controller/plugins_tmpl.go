package controller

import (
	"bytes"
	"html/template"

	"github.com/GoAdminGroup/go-admin/modules/remote_server"

	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
)

func GetPluginsPageJS(data PluginsPageJSData) template.JS {
	t := template.New("plugins_page_js").Funcs(map[string]interface{}{
		"lang":     language.Get,
		"plugWord": plugWord,
	})
	t, err := t.Parse(pluginsPageJS)
	if err != nil {
		logger.Error(err)
		return ""
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return template.JS(buf.String())

}

type PluginsPageJSData struct {
	Prefix string
}

var pluginsPageJS = `

function pluginInstall(name){
	location.href="{{.Prefix}}/info/plugin_"+name+"/new"
}

var downloadLock = false;
const apiTokenKey = "GOADMIN_SERVER_API_TOKEN";
const apiTokenExpireKey = "GOADMIN_SERVER_API_TOKEN_EXPIRE";
const serverHost = "` + remote_server.ServerHost + `"

function login() {
	$.ajax({
		type: 'POST',
		url:  '{{.Prefix}}/server/login',
		dataType: "json",
		contentType: "application/json",
		data: JSON.stringify({
			"account": $("#login-popup-modal form input.name").val(),
			"password": $("#login-popup-modal form input.password").val()
		}),
		success: function (data) {
			if (data.code == 0) {
				localStorage.setItem(apiTokenKey, data.data.token);
				localStorage.setItem(apiTokenExpireKey, (new Date()).getTime() + data.data.expire);
				$('#login-popup-modal').modal('hide');
				if (typeof loginSuccessCallback != "undefined") {
					loginSuccessCallback(data);
				}
				$.pjax.reload('#pjax-container');
			} else {
				swal({
					type: "error",
					title: data.msg,
					showCancelButton: false,
					html: true,
					confirmButtonColor: "#3c8dbc",
					confirmButtonText: '{{lang "got it"}}',
        		})
			}
		},
		error: function (data) {
			swal({
				type: "error",
				title: data.responseJSON.msg ? data.responseJSON.msg : "error",
				showCancelButton: false,
				html: true,
				confirmButtonColor: "#3c8dbc",
				confirmButtonText: '{{lang "got it"}}',
			})	
		}
	});	
}

function pluginBuy(name, uuid) {
	window.open(serverHost + "/plugins/detail/" + uuid); 
}

function pluginDetail(name, uuid) {
	$.ajax({
		dataType: 'json',
		type: 'POST',
		url: '{{.Prefix}}/plugin/detail?name=' + name,
		async: 'true',
		success: function (data) {
			if (data.code == 0) {
				let head_ele = "#detail-popup-modal .plugin-detail .plugin-detail-head "
				$(head_ele + ".plugin-detail-head-logo img").attr('src', data.data.mini_cover);
				$(head_ele + ".plugin-detail-head-title .plugin-detail-title").html(data.data.title);
				$(head_ele + ".plugin-detail-head-title .plugin-detail-provider").html(data.data.author);
				let item_ele = "#detail-popup-modal .plugin-detail .plugin-detail-info .plugin-detail-info-item";
				$(item_ele).eq(0).find(".plugin-detail-info-item-content").html(data.data.introduction);
				$(item_ele).eq(1).find(".plugin-detail-info-item-content").html(data.data.website);
				$(item_ele).eq(2).find(".plugin-detail-info-item-content").html(data.data.version);
				$(item_ele).eq(3).find(".plugin-detail-info-item-content").html(data.data.created_at);
				$(item_ele).eq(4).find(".plugin-detail-info-item-content").html(data.data.updated_at);
				$(item_ele).eq(5).find("a").attr("href", "https://www.go-admin.cn/plugins/detail/" + data.data.uuid);

				let footer_ele = "#detail-popup-modal .modal-footer .btn.btn-primary";

				if (data.data.install) {
					if (data.data.upgrade) {
						$(footer_ele).html('{{plugWord "upgrade"}}').attr("onclick", "pluginDownload('"+ name +"', this)");
						$(footer_ele).show();
					} else {
						$(footer_ele).hide();
					}
				} else {
					if (data.data.downloaded) {
						if (data.data.download_reboot) {
							if (data.data.skip) {
								$(footer_ele).html('{{plugWord "install"}}')
							} else {
								$(footer_ele).html('{{plugWord "install"}}').attr("onclick", "pluginInstall('"+ name +"')")
							}
						} else {
							$(footer_ele).html('{{plugWord "install"}}').attr("onclick", 'pluginRebootInstall()')
						}
					} else {
						if (data.data.free) {
							$(footer_ele).html('{{plugWord "download"}}').attr("onclick", "pluginInstall('"+ name +"', this)")
						} else {
							$(footer_ele).html('{{plugWord "buy"}}').attr("onclick", "pluginBuy('"+ name +"', '"+ uuid +"')")
						}
					}
					$(footer_ele).show();
				}

				$("#detail-popup-modal").modal('show');
			} else {
				swal({
					type: "error",
					title: data.msg,
					showCancelButton: false,
					html: true,
					confirmButtonColor: "#3c8dbc",
					confirmButtonText: '{{lang "got it"}}',
        		})
			}
		},
		error: function (data) {
			swal({
				type: "error",
				title: data.responseJSON.msg ? data.responseJSON.msg : "error",
				showCancelButton: false,
				html: true,
				confirmButtonColor: "#3c8dbc",
				confirmButtonText: '{{lang "got it"}}',
			})	
		}
	});
}

function pluginDownload(name, ele) {
	if (downloadLock) {
		return
	}
	NProgress.start();
	downloadLock = true;
	$(ele).html('<i class="fa fa-refresh fa-spin text-primary" style="color: white;"></i> {{plugWord "downloading"}}')
	$.ajax({
		dataType: 'json',
		type: 'POST',
		url: '{{.Prefix}}/plugin/download',
		async: 'true',
		data: {
			'name': name
		},
		success: function (data) {
			NProgress.done();
			downloadLock = false;
			if (data.code == 0) {
				$(ele).attr('onclick', 'pluginRebootInstall()')
				$(ele).html('{{plugWord "install"}}')
				swal({
					type: "success",
					title: data.msg,
					showCancelButton: false,
					confirmButtonColor: "#3c8dbc",
					confirmButtonText: '{{lang "got it"}}',
        		})
			} else {
				$(ele).html('{{plugWord "download"}}')
				swal({
					type: "error",
					title: data.msg,
					showCancelButton: false,
					text: "<a target='_blank' href='//baidu.com'>点这里手动下载</a>",
					html: true,
					confirmButtonColor: "#3c8dbc",
					confirmButtonText: '{{lang "got it"}}',
        		})
			}
		},
		error: function (data) {
			downloadLock = false;
			$(ele).html('{{plugWord "download"}}')
			alert('download fail');	
		}
	});
}

function pluginRebootInstall() {
	swal({
		type: "success",
		title: "{{plugWord "restart to install"}}",
		showCancelButton: false,
		confirmButtonColor: "#3c8dbc",
		confirmButtonText: '{{lang "got it"}}',
	})
}
`

var pluginsPageCSS = template.CSS(`
	.plugin-item-img {
		text-align: center;
	}
	.plugin-item-img img {
		width: 100%;
		height: auto;
		border: 1px solid #f5f4f4;
	}
	.plugin-item-title {
		margin-top: 10px;
		color: #717171;
		overflow: hidden;
        text-overflow: ellipsis;
        word-break: break-all;
        display: -webkit-box;
        -webkit-line-clamp: 1;
        -webkit-box-orient: vertical;
	}
	@media screen and (min-width: 1200px) {
		.pjax-container-content .row .col-lg-2 {
			width: 14%;
			max-width: 200px;
		}
	}
	.plugin-item-label {
		color: white;
		background-color: red;
		padding: 4px;
		border-radius: 3px;
		font-size: 12px;
		position: absolute;
		top: 5px;
		right: 5px;
	}
`)

var pluginsStorePageCSS = template.CSS(`
	.plugin-item-content {
		margin-left: 15px;
	}
	.plugin-store-item-img img, .plugin-detail-head-logo img {
		border: 1px solid #f5f4f4;
	}
	.plugin-item-content-title {
		font-size: 15px;
		margin-bottom: 6px;
		font-weight: bold;
		padding-top: 2px;
	}
	.plugin-item-content {
		position: absolute;
		margin-left: 121px;
		padding-right: 10px;
		top: 7px;
    	width: 100%;
    	padding-right: 139px;
	}
	.plugin-item-content-description {
		overflow: hidden;
		text-overflow: ellipsis;
		word-break: break-all;
		display: -webkit-box;
		font-size: 15px;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		height: 42px;
	}
	.installation {
		float: right;
		margin-top: 10px;
	}
	.plugin-info {
		float: right;
		margin-top: 10px;
		margin-left: 10px;
	}
	.plugin-detail {
		padding: 10px;
	}
	.plugin-detail-head {
		clear: both;
		height: 112px;
    	margin-bottom: 33px;
	}
	.plugin-detail-title {
		font-size: 30px;
	}
	.plugin-detail-provider {
    	font-size: 15px;
    	margin-top: 4px;
	}
	.plugin-detail-head-logo {
		width: 120px;
		float: left;
	}
	.plugin-detail-head-title {
		float: left;
		margin-left: 10px;
	}
	.plugin-detail-info-item {
		clear: both;
		height: 15px;
		margin-bottom: 17px;
	}
	.plugin-detail-info-item-head {
		width: 120px;
		float: left;
		font-weight: bold;
	}
	.plugin-detail-info-item-content {
		float: left;
		margin-left: 10px;
	}
`)

var pluginStore404 = func() template.HTML {
	return template.HTML(`
<div class="plugin-store-404-content"><p>` + plugWord("can not connect to the goadmin remote server") + `</p></div>
`)
}

var pluginsPageDetailPopupBody = func() template.HTML {
	return template.HTML(`
<div class="plugin-detail">
	<div class="plugin-detail-head">
		<div class="plugin-detail-head-logo">
			<img src="" width="110px" height="110px">
		</div>
		<div class="plugin-detail-head-title">
			<div class="plugin-detail-title"></div>
			<div class="plugin-detail-provider"></div>
		</div>
	</div>
<div class="plugin-detail-info">
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("introduction") + `</div>
		<div class="plugin-detail-info-item-content"></div>
	</div>
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("website") + `</div>
		<div class="plugin-detail-info-item-content"></div>
	</div>
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("version") + `</div>
		<div class="plugin-detail-info-item-content"></div>
	</div>
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("created at") + `</div>
		<div class="plugin-detail-info-item-content"></div>
	</div>
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("updated at") + `</div>
		<div class="plugin-detail-info-item-content"></div>
	</div>
	<div class="plugin-detail-info-item">
		<a target="_blank" href=""><div class="plugin-detail-info-item-head">` + plugWord("learn more") + `</div></a>
	</div>
</div>
</div>`)
}
