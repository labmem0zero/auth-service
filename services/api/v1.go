package api

import (
	"github.com/gorilla/mux"
	"github.com/labmem0zero/go-logger"

	"auth/services/api/handlers"
)

func ApiV1(l *logger.Logger, h handlers.Handlers, r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/user", h.UserCreate).Methods("POST")
}
