package fileserver

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/labmem0zero/go-logger"

	"auth/errs"
	"auth/util"
)

const staticFilesPrefix = "static"

type FileServerMock struct {
	*logger.Logger
	staticDir string
}

func New(staticDir string, l *logger.Logger) FileServerMock {
	return FileServerMock{
		Logger:    l,
		staticDir: staticDir,
	}
}

func (fsm *FileServerMock) Upload(reqID string, file any, p string) (u string, err error) {
	r, name, err := util.ToReadCloserWithName(file)
	if err != nil {
		return "", err
	}
	u = "/" + path.Join(staticFilesPrefix, p, name)
	p = path.Join(fsm.staticDir, p, name)
	if _, err = os.Open(p); !os.IsNotExist(err) {
		err = errs.New(errors.New("file already exists"), "change path", 400)
		return
	}
	var b []byte
	var f *os.File
	if f, err = os.Create(p); err != nil {
		err = errs.New(err, "create file error", 500)
		return
	}
	var n int64
	n, err = io.Copy(f, r)
	if _, err = f.Write(b); err != nil {
		err = errs.New(err, "write file error", 500)
		return
	}
	fsm.Info(reqID, fmt.Sprintf("wrote %v bytes to file %v", n, p))
	return
}

func (fsm *FileServerMock) Download(reqID string, u string) (f io.ReadCloser, err error) {
	split := strings.Split(u, staticFilesPrefix)
	if len(split) != 2 {
		err = errs.New(errors.New("wrong file name"), "", 400)
	}
	p := path.Join(fsm.staticDir, split[1])
	if f, err = os.Open(p); err != nil {
		err = errs.New(err, "read file error", 500)
		if os.IsNotExist(err) {
			err = errs.New(err, "missing file", 404)
		}
		return
	}
	return
}
