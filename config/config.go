package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type Config struct {
	Server   *Server
	Postgres *Postgres
}

type Server struct {
	Port  int  `envconfig:"SERVER_PORT" default:"80"`
	Debug bool `envconifg:"SERVER_DEBUG" default:"false"`
}

type Postgres struct {
	Addr     string `envconfig:"PG_ADDR" required:"false"`
	Port     int    `envconfig:"PG_PORT" required:"false"`
	Database string `envconfig:"PG_DB" required:"false"`
	User     string `envconfig:"PG_USER" required:"false"`
	Password string `envconfig:"PG_PASSWORD" required:"false"`
}

func GetConfig() *Config {
	var cfg Config
	_ = gotenv.Load(".env")

	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
