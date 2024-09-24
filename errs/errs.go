package errs

import "encoding/json"

type AppError struct {
	Err  string `json:"err"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (e *AppError) Error() string {
	b, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		return ""
	}
	return string(b)
}

func New(err error, msg string, code int) *AppError {
	if val, ok := err.(*AppError); ok {
		return val
	}
	var e string
	if err != nil {
		e = err.Error()
	}
	return &AppError{
		Err:  e,
		Msg:  msg,
		Code: code,
	}
}
