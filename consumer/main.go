package main

import (
	"context"
	"kafka-producer-consumer-example/common/pkg/gkafka"
	"kafka-producer-consumer-example/consumer/services"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

func init() {
	gkafka.SetupConsumerClient()
}

func main() {
	consumerClient := gkafka.GetConsumerClient()
	partitionConsumer, err := consumerClient.ConsumePartition(gkafka.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	consumerServvice := services.NewConsumerService()

	endConsumer := make(chan bool, 1)
	go consumerServvice.Consumer(partitionConsumer, endConsumer)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-quit

	endConsumer <- true
	defer func() {
		_ = partitionConsumer.Close()
	}()
	defer func() {
		_ = consumerClient.Close()
	}()
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
