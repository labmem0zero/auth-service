package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	http2 "auth/http"
	"auth/http/middlewares"
	"auth/models"
)

func (h Handlers) UserCreate(w http.ResponseWriter, r *http.Request) {
	reqID := middlewares.GetReqID(r)
	var u models.UserSet
	var err error
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.Write([]byte(err.Error()))
		h.l.Error(reqID, err)
		return
	}
	if len(u.Username) < 5 || len(u.Password) == 0 {
		err = errors.New("Wrong username or password")
		http2.CheckErrWriteResp(w, 400, nil, err)
		return
	}
	var res models.UserView
	if res, err = h.u.UserCreate(reqID, u); err != nil {
		http2.CheckErrWriteResp(w, 500, res, err)
		return
	}
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http2.CheckErrWriteResp(w, 500, res, err)
	}
	return
}
