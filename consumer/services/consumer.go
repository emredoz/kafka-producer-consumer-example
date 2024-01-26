package services

import (
	"kafka-producer-consumer-example/common/model"
	"log"

	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
)

type ConsumerService interface {
	Consumer(partitionConsumer sarama.PartitionConsumer, endConsumerChan chan bool)
}

type consumerService struct {
}

func NewConsumerService() ConsumerService {
	return &consumerService{}
}

func (c *consumerService) Consumer(partitionConsumer sarama.PartitionConsumer, endConsumerChan chan bool) {
	var endConsumer = false
	for !endConsumer {
		select {
		case message := <-partitionConsumer.Messages():
			n, err := convertToNotification(message)
			if err != nil {
				log.Print("convertToNotification err: ", err)
				continue
			}
			handleNotification(n)
			break
		case endConsumer = <-endConsumerChan:
			break
		}
	}
}

func convertToNotification(message *sarama.ConsumerMessage) (*model.Notification, error) {
	data := message.Value
	n := &model.Notification{}
	err := proto.Unmarshal(data, n)
	if err != nil {
		log.Print("proto unmarshal err: ", err)
		return nil, err
	}
	return n, nil
}

func handleNotification(n *model.Notification) {
	// do something
	log.Println("notification received, notification: ", n)
}
