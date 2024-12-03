package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
)

type Config struct {
	AppName                string `json:"app_name" toml:"AppName"`
	AppID                  string `json:"app_id" toml:"AppId"`
	Environment            string `json:"environment" toml:"Environment"`
	CertificateName        string `json:"-" toml:"CertificateName"`
	TelegramLoggerBotToken string `json:"-" toml:"-"`
	TelegramLoggerChatID   int64  `json:"-" toml:"-"`
	StaticPath             string `json:"-" toml:"StaticPath"`
	UsersDB                DB     `json:"users_db" toml:"UserDB"`
}

type DB struct {
	Server   string `json:"server" toml:"Server"`
	Port     string `json:"port" toml:"Port"`
	Database string `json:"database" toml:"Database"`
	Scheme   string `json:"scheme" toml:"Scheme"`
	SSLMode  bool   `json:"ssl_mode" toml:"SSLMode"`
	Username string `json:"-" toml:"-"`
	Password string `json:"-" toml:"-"`
}

func LoadConfig() Config {
	var conf Config
	if conf.Environment = os.Getenv("environment"); conf.Environment == "" {
		conf.Environment = local
	}
	if conf.AppID = os.Getenv("app_id"); conf.AppID == "" {
		conf.AppID = uuid.New().String()
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
	var err error
	if _, err = toml.DecodeFile(confFile, &conf); err != nil {
		log.Fatal("toml:", err)
	}
	conf.TelegramLoggerBotToken = os.Getenv("telegram_bot_token")
	if conf.TelegramLoggerChatID, err = strconv.ParseInt(os.Getenv("telegram_chat_id"), 10, 64); err != nil {
		log.Fatal("wrong telegram chat id")
	}
	conf.UsersDB.Username = os.Getenv("user_db_username")
	conf.UsersDB.Password = os.Getenv("user_db_password")
	return conf
}

const (
	local = "LOCAL"
	stage = "STAGE"
	prod  = "PRODUCTION"
)

func (conf *Config) LoadKey() (*rsa.PublicKey, error) {
	out, err := os.ReadFile(conf.CertificateName)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(out)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	cc, err := x509.ParseCertificates(block.Bytes)
	if err != nil {
		return nil, err
	}

	for _, c := range cc {
		if c.PublicKey == nil {
			continue
		}

		if pub, ok := c.PublicKey.(*rsa.PublicKey); ok {
			return pub, nil
		}
	}

	return nil, errors.New("could not find rsa256 public key")
}
