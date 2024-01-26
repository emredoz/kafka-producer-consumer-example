package service

import (
	"fmt"
	"kafka-producer-consumer-example/common/model"
	"kafka-producer-consumer-example/common/pkg/gkafka"
	"log"

	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
)

type ProducerService interface {
	SendNotification(n model.Notification) error
}
type producerService struct {
	client sarama.SyncProducer
}

func NewProducer(client sarama.SyncProducer) ProducerService {
	return &producerService{client: client}
}

func (p producerService) SendNotification(notification model.Notification) error {
	data, err := proto.Marshal(&notification)
	if err != nil {
		log.Print("proto marshal err: ", err)
		return err
	}
	message := &sarama.ProducerMessage{
		Topic: gkafka.Topic,
		Key:   nil,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := p.client.SendMessage(message)
	if err != nil {
		log.Print("p.client.SendMessage err: ", err)
		return err
	}
	log.Println(fmt.Sprintf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", gkafka.Topic, partition, offset))
	return nil
}
