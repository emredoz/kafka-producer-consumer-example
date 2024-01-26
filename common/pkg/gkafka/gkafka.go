package gkafka

import (
	"github.com/IBM/sarama"
	"log"
)

const (
	KafkaHost = "localhost:9092"
	Topic     = "notification"
)

var consumerClient sarama.Consumer
var producerClient sarama.SyncProducer

func SetupConsumerClient() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{KafkaHost}, config)

	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}
	consumerClient = consumer
}

func GetConsumerClient() sarama.Consumer {
	return consumerClient
}

func SetupProducerClient() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{KafkaHost}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err: ", err)
	}
	producerClient = producer
}

func GetProducerClient() sarama.SyncProducer {
	return producerClient
}
