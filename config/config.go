package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Environment            string `json:"environment"`
	TelegramLoggerBotToken string `toml:"TelegramBotToken"`
	TelegramLoggerChatID   int64  `toml:"TelegramChatID"`
}

func LoadConfig() Config {
	var conf Config
	if conf.Environment = os.Getenv("environment"); conf.Environment == "" {
		conf.Environment = local
	}
	var confFile = ""
	switch conf.Environment {
	case local:
		confFile = "./config/conf_local.toml"
	case stage:
		confFile = "./config/conf_stage.toml"
	case prod:
		confFile = "./config/conf_prod.toml"
	}
	if _, err := toml.DecodeFile(confFile, &conf); err != nil {
		log.Fatal("toml:", err)
	}
	return conf
}

const (
	local = "LOCAL"
	stage = "STAGE"
	prod  = "PRODUCTION"
)
