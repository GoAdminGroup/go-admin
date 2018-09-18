package file

import (
	"github.com/chenhg5/go-admin/plugins/admin/modules"
	"io"
	"mime/multipart"
	"os"
	"path"
	"sync"
)

type Uploader interface {
	Upload(*multipart.Form) (*multipart.Form, error)
}

func GetFileEngine(name string) Uploader {
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

		pathStr, err := c(fileObj, filename)

		if err != nil {
			return nil, err
		}

		(*form).Value[k] = []string{pathStr}
	}

	return form, nil
}

func SaveMultipartFile(fh *multipart.FileHeader, path string) error {
	f, err := fh.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	if ff, ok := f.(*os.File); ok {
		return os.Rename(ff.Name(), path)
	}

	ff, err := os.Create(path)
	if err != nil {
		return err
	}
	defer ff.Close()
	_, err = copyZeroAlloc(ff, f)
	return err
}

func copyZeroAlloc(w io.Writer, r io.Reader) (int64, error) {
	vbuf := copyBufPool.Get()
	buf := vbuf.([]byte)
	n, err := io.CopyBuffer(w, r, buf)
	copyBufPool.Put(vbuf)
	return n, err
}

var copyBufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 4096)
	},
}
