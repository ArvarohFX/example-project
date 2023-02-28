package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type Config struct {
	Server *Server
}

type Server struct {
	Port  int  `envconfig:"SERVER_PORT" default:"80"`
	Debug bool `envconifg:"SERVER_DEBUG" default:"false"`
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
