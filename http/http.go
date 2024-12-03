package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"auth/errs"
)

func CheckErrWriteResp(w http.ResponseWriter, code int, out any, err error) {
	if err != nil {
		var appErr *errs.AppError
		var ok bool
		appErr, ok = err.(*errs.AppError)
		if !ok {
			appErr = errs.New(err, "", code)
		}
		w.WriteHeader(appErr.Code)
		w.Write([]byte(appErr.Err))
		return
	}
	w.WriteHeader(code)
	if out == nil {
		return
	}
	if err = json.NewEncoder(w).Encode(out); err == nil {
		return
	}
	w.Write([]byte(fmt.Sprintln(out)))
}
