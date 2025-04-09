package controller

import (
	"api/consumer"
	"api/src/core"
	"api/src/sensor_bmp180/application"
	"api/src/sensor_bmp180/domain"
	"api/src/sensor_bmp180/infrastructure"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BMP180Controller struct {
	service *application.BMP180Service
}

func NewBMP180Controller(service *application.BMP180Service) *BMP180Controller {
	return &BMP180Controller{service: service}
}

func (bc *BMP180Controller) GetSensorData(c *gin.Context) {
	sensorData, err := bc.service.GetSensorData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensorData)
}

func StartBMP180Consumer(rabbitMQ *consumer.RabbitMQ, queueName string, db *core.Database, wsClient *infrastructure.BmpClient) {
	_, err := rabbitMQ.DeclareQueue(queueName)
	if err != nil {
		log.Fatal("Error declarando la cola BMP180:", err)
	}

	msgs, err := rabbitMQ.ConsumeMessages(queueName)
	if err != nil {
		log.Fatal("Error consumiendo mensajes BMP180:", err)
	}

	repo := infrastructure.NewBMP180Repository(db.DB)
	sensorService := application.NewBMP180Service(repo)

	for msg := range msgs {
		var sensorData domain.SensorBMP180
		if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
			log.Println("Error decodificando mensaje BMP180:", err)
			continue
		}

		if err := sensorService.ProcessSensorData(sensorData); err != nil {
			log.Println("Error procesando datos BMP180:", err)
		}

        if err := wsClient.SendSensorData(sensorData); err != nil {
            log.Println("Error enviando datos BMP180 al WebSocket:", err)
        }
	}
}