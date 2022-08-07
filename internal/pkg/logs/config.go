package logs

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"log"
)

var Config LoggerConfig

type LoggerConfig struct {
	Level string `env:"LOG_LEVEL" envDefault:"TRACE"`
}

func (l *LoggerConfig) Register() error {
	err := env.Parse(&Config)
	if err != nil {
		return errors.Wrap(err, "register failed, error parsing logger config")
	}
	return nil
}

func (l *LoggerConfig) Validate() error {
	return nil
}

func (l *LoggerConfig) Print() interface{} {
	defer log.Println("loading logger configurations... ")
	return Config
}
