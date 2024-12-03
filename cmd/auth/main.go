package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labmem0zero/go-logger"
	"github.com/labmem0zero/go-logger/impl"
	"github.com/labmem0zero/go-logger/impl/flogger"
	"github.com/labmem0zero/go-logger/impl/tlogger"

	"auth/config"
	"auth/provider"
	"auth/services"
	"auth/services/api"
	"auth/services/usecases"
	"auth/util"
)

func main() {
	conf := config.LoadConfig()

	tl, err := tlogger.NewLogger(conf.TelegramLoggerBotToken, conf.TelegramLoggerChatID, impl.LoggerSettings{
		AppName:     conf.AppName,
		AppID:       conf.AppID,
		Environment: conf.Environment,
		Levels: map[string]struct{}{
			logger.LevelErr:     {},
			logger.LevelFatal:   {},
			logger.LevelWarning: {},
			logger.LevelDebug:   {},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fl, err := flogger.NewFileLogger("log.log", impl.LoggerSettings{
		AppName:     conf.AppName,
		AppID:       conf.AppID,
		Environment: conf.Environment,
	})
	if err != nil {
		log.Fatal(err)
	}
	l := logger.New(fl, tl)

	var p provider.Provider
	if p, err = provider.New(conf, &l); err != nil {
		l.Fatal("launch", err)
		return
	}
	c := make(chan bool, 1)

	r := util.NewRecoverySystem(&l)
	u := usecases.New(p, &l, r)
	a := api.New(conf, &l, u)
	srv := services.StartServices("App start", u, a)
	l.Warn("App start", "Application has been started")
	AwaitClose(srv, l, c)
}

func AwaitClose(srv services.Services, l logger.Logger, c chan bool) {
	appExit := make(chan os.Signal, 1)
	signal.Notify(appExit, os.Interrupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer func() {
		srv.StopServices("App exit")
		l.Info("Application closed")
	}()
	select {
	case <-c:
		return
	case <-appExit:
		return
	}
}
