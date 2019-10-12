package file

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"mime/multipart"
)

type LocalFileUploader struct {
	BasePath string
}

func GetLocalFileUploader() Uploader {
	return &LocalFileUploader{
		config.Get().Store.Path,
	}
}

func (local *LocalFileUploader) Upload(form *multipart.Form) error {
	return Upload(func(fileObj *multipart.FileHeader, filename string) (string, error) {
		if err := SaveMultipartFile(fileObj, (*local).BasePath+"/"+filename); err != nil {
			return "", err
		}
		return filename, nil
	}, form)
}
