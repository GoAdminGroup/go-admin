package file

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"io"
	"mime/multipart"
	"os"
	"path"
	"sync"
)

type Uploader interface {
	Upload(*multipart.Form) error
}

type UploaderGenerator func() Uploader

var uploaderList = map[string]UploaderGenerator{
	"local": GetLocalFileUploader,
}

var upMu sync.Mutex

// AddUploader makes a uploader generator available by the provided theme name.
// If Add is called twice with the same name or if uploader is nil,
// it panics.
func AddUploader(name string, up UploaderGenerator) {
	upMu.Lock()
	defer upMu.Unlock()
	if up == nil {
		panic("uploader generator is nil")
	}
	if _, dup := uploaderList[name]; dup {
		panic("add uploader generator twice " + name)
	}
	uploaderList[name] = up
}

func GetFileEngine(name string) Uploader {
	if up, ok := uploaderList[name]; ok {
		return up()
	}
	panic("wrong uploader name")
}

type UploadFun func(*multipart.FileHeader, string) (string, error)

func Upload(c UploadFun, form *multipart.Form) error {
	var (
		suffix   string
		filename string
	)

	for k := range form.File {
		fileObj := form.File[k][0]

		suffix = path.Ext(fileObj.Filename)
		filename = modules.Uuid() + suffix

		pathStr, err := c(fileObj, filename)

		if err != nil {
			return err
		}

		form.Value[k] = []string{pathStr}
	}

	return nil
}

func SaveMultipartFile(fh *multipart.FileHeader, path string) (err error) {
	var f multipart.File
	f, err = fh.Open()
	if err != nil {
		return err
	}
	defer func() {
		if err2 := f.Close(); err2 != nil {
			err = err2
		}
	}()

	if ff, ok := f.(*os.File); ok {
		return os.Rename(ff.Name(), path)
	}

	ff, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err2 := ff.Close(); err2 != nil {
			err = err2
		}
	}()
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
