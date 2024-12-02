package db

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"auth/config"
)

func Connect(conf config.DB) (conn *sqlx.DB, err error) {
	query := url.Values{}

	switch conf.Scheme {
	case "postgres":
		query.Add("dbname", conf.Database)
		if !conf.SSLMode {
			query.Add("sslmode", "disable")
		}
	default:
		return nil, fmt.Errorf("unknown db scheme")
	}

	host := conf.Server
	if conf.Port != "" {
		host += ":" + conf.Port
	}

	u := &url.URL{
		Scheme:   conf.Scheme,
		User:     url.UserPassword(conf.Username, conf.Password),
		Host:     host,
		RawQuery: query.Encode(),
	}

	conn, err = sqlx.Open(conf.Scheme, u.String())
	return
}
