package file

import (
	"mime/multipart"
	"github.com/valyala/fasthttp"
)

type LocalFileUploader struct {
	BasePath string
}

func GetLocalFileUploader() *LocalFileUploader {
	return &LocalFileUploader{
		"./resources/adminlte/uploads",
	}
}

func (local *LocalFileUploader) Upload(form *multipart.Form) (*multipart.Form, error) {
	return Upload(func(fileObj *multipart.FileHeader, filename string) (string, error) {
		if err := fasthttp.SaveMultipartFile(fileObj, (*local).BasePath+filename); err != nil {
			return "", err
		}
		return "/uploads/" + filename, nil
	}, form)
}

