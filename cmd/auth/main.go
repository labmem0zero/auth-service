package main

import (
	"log"

	"github.com/labmem0zero/go-logger"
	"github.com/labmem0zero/go-logger/impl/filelogger"
	"github.com/labmem0zero/go-logger/impl/telegram"

	"auth/config"
)

func main() {
	cfg := config.LoadConfig()
	tl, err := telegram.NewLogger(cfg.TelegramLoggerBotToken, cfg.TelegramLoggerChatID)
	log.Println(cfg.TelegramLoggerBotToken, cfg.TelegramLoggerChatID)
	if err != nil {
		log.Fatal(err)
	}
	fl, err := filelogger.NewFileLogger("log.log")
	if err != nil {
		log.Fatal(err)
	}
	lg := logger.New(fl, tl)
	lg.Debug("Debug", "test")
	lg.Info("Info", "test")
}
