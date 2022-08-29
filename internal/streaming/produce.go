package streaming

import (
	"encoding/json"
	"fmt"
	"github-user-service/internal/domain/adaptors/streaming"
	"github-user-service/internal/domain/events"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

const (
	githubUsersTopic = `github.users`
	userIDHeader     = `user-id`
)

type Producer struct {
	syncProducer sarama.SyncProducer
}

func NewSyncProducer() streaming.Producer {
	brokers := Config.Brokers
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	syncProd, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		log.Fatal("error initiating streaming", err)
	}

	return &Producer{
		syncProducer: syncProd,
	}
}

func (p Producer) Produce(message interface{}) error {
	userEvent, ok := message.(events.UserChanged)
	if !ok {
		return fmt.Errorf("event type mismatch, expected: UserChanged")
	}

	b, err := json.Marshal(userEvent)
	if err != nil {
		return fmt.Errorf("event marshall error")
	}
	msg := &sarama.ProducerMessage{
		Topic: githubUsersTopic,
		Key:   sarama.StringEncoder(userEvent.Payload.UniqueID.String()),
		Value: sarama.StringEncoder(b),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte(userIDHeader),
				Value: []byte(userEvent.Payload.UniqueID.String()),
			},
		},
		Timestamp: time.Now(),
	}

	partition, offset, err := p.syncProducer.SendMessage(msg)
	log.Println(fmt.Sprintf("message produced to partition [%d] with offset [%d]", partition, offset))

	return nil
}
