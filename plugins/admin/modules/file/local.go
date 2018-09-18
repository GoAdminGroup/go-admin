package file

import (
	"github.com/chenhg5/go-admin/modules/config"
	"mime/multipart"
)

type LocalFileUploader struct {
	BasePath string
}

func GetLocalFileUploader() *LocalFileUploader {
	return &LocalFileUploader{
		config.Get().STORE.PATH,
	}
}

func (local *LocalFileUploader) Upload(form *multipart.Form) (*multipart.Form, error) {
	return Upload(func(fileObj *multipart.FileHeader, filename string) (string, error) {
		if err := SaveMultipartFile(fileObj, (*local).BasePath+"/"+filename); err != nil {
			return "", err
		}
		return filename, nil
	}, form)
}
