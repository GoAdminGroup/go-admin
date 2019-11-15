package file

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"mime/multipart"
)

// LocalFileUploader is an Uploader of local file engine.
type LocalFileUploader struct {
	BasePath string
}

// GetLocalFileUploader return the default Uploader.
func GetLocalFileUploader() Uploader {
	return &LocalFileUploader{
		config.Get().Store.Path,
	}
}

// Upload implements the Uploader.Upload.
func (local *LocalFileUploader) Upload(form *multipart.Form) error {
	return Upload(func(fileObj *multipart.FileHeader, filename string) (string, error) {
		if err := SaveMultipartFile(fileObj, (*local).BasePath+"/"+filename); err != nil {
			return "", err
		}
		return filename, nil
	}, form)
}
