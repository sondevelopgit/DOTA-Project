package config

import (
	"fmt"

	"github.com/caarlos0/env/v7"
)

type Configuration struct {
	Debug        bool     `env:"DEBUG" envDefault:"local"`
	SecretKey    string   `env:"SECRET_KEY"`
	ServerPort   string   `env:"SERVER_PORT" envDefault:"4000"`
	AllowOrigins []string `env:"ALLOW_ORIGINS" envSeparator:":"`
}

var configuration Configuration

func init() {
	configuration = Configuration{}
	if err := env.Parse(&configuration); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func ConfigInstance() *Configuration {
	return &configuration
}
