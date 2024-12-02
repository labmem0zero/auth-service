package util

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"runtime/debug"

	"github.com/labmem0zero/go-logger"

	"auth/errs"
)

type Recovery struct {
	l *logger.Logger
}

func (rc *Recovery) Recover(reqID string) {
	r := recover()
	if r != nil {
		var err error
		stacktrace := string(debug.Stack())
		switch t := r.(type) {
		case string:
			err = errs.New(fmt.Errorf(`panic: %s, stacktrace: %s`, t, stacktrace), "", 500)
		case error:
			err = errs.New(fmt.Errorf(`panic: %v, stacktrace: %s`, t, stacktrace), "", 500)
		default:
			err = errs.New(fmt.Errorf(`unknown panic`), "", 500)
		}
		rc.l.Error(reqID, fmt.Sprintf("panic: %v", err))
	}
}

func NewRecoverySystem(l *logger.Logger) *Recovery {
	return &Recovery{l: l}
}

func Convert[T any](in any) (out T, err error) {
	var ok bool
	out, ok = in.(T)
	if !ok {
		err = errs.New(errors.New("wrong input"), fmt.Sprintf("in should be %T", out), 500)
	}
	return
}

func ToReadCloserWithName(in any) (out io.ReadCloser, name string, err error) {
	switch v := in.(type) {
	case multipart.FileHeader:
		name = v.Filename
		out, err = v.Open()
	case os.File:
		name = v.Name()
		out = &v
	default:
		err = errors.New("wrong input")
	}
	return
}
