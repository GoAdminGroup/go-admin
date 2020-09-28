package action

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type FileUploadAction struct {
	BaseAction
	Url      string
	Method   string
	FileName string
	Handlers []context.Handler
}

func FileUpload(id string, handler types.Handler) *FileUploadAction {
	if id == "" {
		panic("wrong file upload action parameter, empty id")
	}
	return &FileUploadAction{
		Url:      URL(id),
		Method:   "post",
		FileName: "file",
		Handlers: context.Handlers{handler.Wrap()},
	}
}

func (file *FileUploadAction) SetUrl(url string) *FileUploadAction {
	file.Url = url
	return file
}

func (file *FileUploadAction) SetMethod(method string) *FileUploadAction {
	file.Method = method
	return file
}

func (file *FileUploadAction) GetCallbacks() context.Node {
	return context.Node{
		Path:     file.Url,
		Method:   file.Method,
		Handlers: file.Handlers,
		Value:    map[string]interface{}{constant.ContextNodeNeedAuth: 1},
	}
}

func (file *FileUploadAction) Js() template.JS {

	return template.JS(`$('` + file.BtnId + `').on('click', function(){
  $('` + file.BtnId + `_input').click();
});

$("` + file.BtnId + `_input").on("change", function () {
  var files = $('` + file.BtnId + `_input').prop('files');

  var data = new FormData();
  data.append('` + file.FileName + `', files[0]);
  NProgress.start();

  $.ajax({
		url: '` + file.Url + `',
		type: '` + file.Method + `',
		data: data,
		cache: false,
  		processData: false,
  		contentType: false,
		success: function (data) { 
		  	NProgress.done();
			if (data.code === 0) {
				swal(data.msg, '', 'success');
				$.pjax.reload('#pjax-container');
			} else {
				swal(data.msg, '', 'error');
			}
		},
		error: function (data) {
			NProgress.done();
			if (data.responseText !== "") {
				swal(data.responseJSON.msg, '', 'error');								
			} else {
				swal('error', '', 'error');
			}
		},
  });
});`)
}

func (file *FileUploadAction) BtnAttribute() template.HTML {
	return template.HTML(`href="javascript:;"`)
}

func (file *FileUploadAction) FooterContent() template.HTML {
	return template.HTML(`<input class="` + file.BtnId[1:] + `_input" type="file" multiple="multiple" style="display:none" />`)

}
