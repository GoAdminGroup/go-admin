package file

import (
	"mime/multipart"
	"path"
	"goAdmin/modules"
	"github.com/valyala/fasthttp"
	"fmt"
)

type FileUploader interface {
	Upload(*multipart.Form) (*multipart.Form, error)
}

func GetFileEngine(name string) FileUploader {
	if name == "local" {
		return GetLocalFileUploader()
	}
	return nil
}

type UploadFun func(*multipart.FileHeader, string) (string, error)

func Upload(c UploadFun, form *multipart.Form) (*multipart.Form, error) {
	var (
		suffix   string
		filename string
	)

	for k, _ := range (*form).File {
		fileObj := form.File[k][0]

		suffix = path.Ext(fileObj.Filename)
		filename = modules.Uuid(50) + suffix
		if err := fasthttp.SaveMultipartFile(fileObj, "./resources/uploads/"+filename); err != nil {
			fmt.Println("save upload file error:", err)
		}

		pathStr, err := c(fileObj, filename)

		if err != nil {
			return nil, err
		}

		(*form).Value[k] = []string{pathStr}
	}

	return form, nil
}