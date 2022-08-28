package streaming

import (
	"github.com/caarlos0/env/v6"
	"log"
)

var Config KafkaConfig

type KafkaConfig struct {
	Brokers []string `env:"KAFKA_BROKERS" envDefault:"[localhost:9090]"`
}

func (k *KafkaConfig) Register() error {
	err := env.Parse(&Config)
	if err != nil {
		log.Fatal("register failed , error parsing kafka config")
	}

	return nil
}

func (k *KafkaConfig) Validate() error {
	if len(Config.Brokers) == 0 {
		log.Fatal("application kafka brokers cannot be empty")
	}
	return nil
}

func (k *KafkaConfig) Print() interface{} {
	defer log.Println("---loading kafka configs---")
	return &Config
}
