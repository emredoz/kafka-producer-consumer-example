package controller

import "C"
import (
	"kafka-producer-consumer-example/common/model"
	"kafka-producer-consumer-example/producer/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type producerController struct {
	service service.ProducerService
}

func NewProducerController(service service.ProducerService, R *gin.Engine) {
	controller := &producerController{service: service}
	api := R.Group("/api/v1")
	api.POST("/notification", controller.SendNotification)
}

func (s *producerController) SendNotification(c *gin.Context) {
	notification := &model.Notification{}
	err := c.ShouldBindJSON(notification)
	if err != nil {
		log.Print("c.ShouldBindJSON err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}
	err = s.service.SendNotification(*notification)
	if err != nil {
		log.Print("s.p.SendNotification err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
