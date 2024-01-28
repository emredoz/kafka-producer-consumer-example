package main

import (
	"context"
	"fmt"
	"kafka-producer-consumer-example/common/pkg/gkafka"
	"kafka-producer-consumer-example/producer/controller"
	"kafka-producer-consumer-example/producer/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	gkafka.SetupProducerClient()
}

const (
	defaultPort  = 8081
	readTimeout  = 10 * time.Second
	writeTimeout = 10 * time.Second
)

func main() {
	producerClient := gkafka.GetProducerClient()
	defer func() {
		_ = producerClient.Close()
	}()
	gin.SetMode(gin.ReleaseMode)

	producerService := service.NewProducer(producerClient)
	router := newRouter()
	controller.NewProducerController(producerService, router)
	server := newServer(router, defaultPort)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Server failed to start: ", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Print("Server Shutdown err: ", err)
	}
}

func newServer(router *gin.Engine, port int) *http.Server {
	endPoint := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}

func newRouter() *gin.Engine {
	return gin.New()
}
