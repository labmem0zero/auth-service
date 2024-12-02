package handlers

import (
	"encoding/json"
	"net/http"

	"auth/models"
)

func (h Handlers) UserCreate(w http.ResponseWriter, r *http.Request) {
	reqID := "reqID"
	var u models.UserCreate
	var err error
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.Write([]byte(err.Error()))
		h.l.Error(reqID, err)
		return
	}
	var id int64
	if id, err = h.u.UserCreate(reqID, u); err != nil {
		w.Write([]byte(err.Error()))
		h.l.Error(reqID, err)
		return
	}
	res := models.UserView{
		UserID:   id,
		Username: u.Username,
	}
	if err = json.NewEncoder(w).Encode(res); err != nil {
		w.Write([]byte(err.Error()))
		h.l.Error(reqID, err)
	}
	return
}
