package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labmem0zero/go-logger"

	"auth/config"
	"auth/http/middlewares"
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
	cert, err := conf.LoadKey()
	if err != nil {
		l.Fatal("App start", err)
	}
	mw := middlewares.New(l, cert)
	h := handlers.New(l, u)
	r := mux.NewRouter().StrictSlash(true)
	r.Use(mw.MiddlewareRequestLogging)
	api := r.PathPrefix("/api").Subrouter()
	ApiV1NoAuth(l, h, api)
	return Api{
		l:      l,
		router: r,
	}
}
