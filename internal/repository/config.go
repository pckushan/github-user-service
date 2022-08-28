package repository

import (
	"github.com/caarlos0/env/v6"
	"log"
)

var Config RepoConfig

type RepoConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	UserName string `env:"DB_USERNAME" envDefault:"user"`
}

func (r RepoConfig) Register() error {
	err := env.Parse(&Config)
	if err != nil {
		log.Fatal("register failed , error parsing repository config")
	}

	return nil
}

func (r RepoConfig) Validate() error {
	if Config.Host == "" {
		log.Fatal("application repo host cannot be empty")
	}
	if Config.Port == "" {
		log.Fatal("application repo port cannot be empty")
	}
	return nil
}

func (r RepoConfig) Print() interface{} {
	defer log.Println("---loading repository configs---")
	return &Config
}
