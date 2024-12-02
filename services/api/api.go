package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labmem0zero/go-logger"

	"auth/config"
	"auth/services/api/handlers"
	"auth/services/usecases"
)

type Api struct {
	l      *logger.Logger
	router *mux.Router
	c      chan bool
	u      usecases.Usecases
}

func (a Api) Start(reqID string) {
	a.l.Info(reqID, "Api service has started")
	err := http.ListenAndServe(":8080", a.router)
	if err != nil {
		a.l.Fatal(reqID, err)
	}
}

func (a Api) Stop(reqID string) {
	a.l.Info(reqID, "Api service has stopped")
}

func New(conf config.Config, l *logger.Logger, u usecases.Usecases) Api {
	h := handlers.New(l, u)
	r := mux.NewRouter().StrictSlash(true)
	api := r.PathPrefix("/api").Subrouter()
	ApiV1(l, h, api)
	return Api{
		l:      l,
		router: r,
	}
}
