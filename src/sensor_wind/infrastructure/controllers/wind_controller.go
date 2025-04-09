package controllers

import (
	"api/consumer"
	"api/src/core"
	"api/src/sensor_wind/application"
	"api/src/sensor_wind/domain"
	"api/src/sensor_wind/infrastructure"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WindController struct {
	service *application.WindService
}

func NewWindController(service *application.WindService) *WindController {
	return &WindController{service: service}
}

func (wc *WindController) GetSensorData(c *gin.Context) {
	sensorData, err := wc.service.GetSensorData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensorData)
}

func StartWindConsumer(rabbitMQ *consumer.RabbitMQ, queueName string, db *core.Database, wsClient *infrastructure.WindClient) {
	_, err := rabbitMQ.DeclareQueue(queueName)
	if err != nil {
		log.Fatal("Error declarando la cola Wind:", err)
	}

	msgs, err := rabbitMQ.ConsumeMessages(queueName)
	if err != nil {
		log.Fatal("Error consumiendo mensajes Wind:", err)
	}

	repo := infrastructure.NewWindRepository(db.DB)
	sensorService := application.NewWindService(repo)

	for msg := range msgs {
		var sensorData domain.SensorWind
		if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
			log.Println("Error decodificando mensaje Wind:", err)
			continue
		}

		if err := sensorService.ProcessSensorData(sensorData); err != nil {
			log.Println("Error procesando datos Wind:", err)
		}

		if err := wsClient.SendSensorData(sensorData); err != nil {
			log.Println("Error enviando datos Wind al WebSocket:", err)
		}
	}
}