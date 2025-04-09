package controller

import (
	"api/consumer"
	"api/src/core"
	"api/src/sensor_ldr/application"
	"api/src/sensor_ldr/domain"
	"api/src/sensor_ldr/infrastructure"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LDRController struct {
	service *application.LDRService
}

func NewLDRController(service *application.LDRService) *LDRController {
	return &LDRController{service: service}
}

func (lc *LDRController) GetSensorData(c *gin.Context) {
	sensorData, err := lc.service.GetSensorData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensorData)
}

func StartLDRConsumer(rabbitMQ *consumer.RabbitMQ, queueName string, db *core.Database, wsClient *infrastructure.LDRClient) {
	_, err := rabbitMQ.DeclareQueue(queueName)
	if err != nil {
		log.Fatal("Error declarando la cola LDR:", err)
	}

	msgs, err := rabbitMQ.ConsumeMessages(queueName)
	if err != nil {
		log.Fatal("Error consumiendo mensajes LDR:", err)
	}

	repo := infrastructure.NewLDRRepository(db.DB)
	sensorService := application.NewLDRService(repo)

	for msg := range msgs {
		var sensorData domain.SensorLDR
		if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
			log.Println("Error decodificando mensaje LDR:", err)
			continue
		}

		if err := sensorService.ProcessSensorData(sensorData); err != nil {
			log.Println("Error procesando datos LDR:", err)
		}

		if err := wsClient.SendSensorData(sensorData); err != nil {
			log.Println("Error enviando datos LDR al WebSocket:", err)
		}
	}
}