// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package file

import (
	"mime/multipart"

	"github.com/GoAdminGroup/go-admin/modules/config"
)

// LocalFileUploader is an Uploader of local file engine.
type LocalFileUploader struct {
	BasePath string
}

// GetLocalFileUploader return the default Uploader.
func GetLocalFileUploader() Uploader {
	return &LocalFileUploader{
		config.GetStore().Path,
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
