package controller

import (
	"api/consumer"
	"api/src/core"
	"api/src/sensor_dht22/application"
	"api/src/sensor_dht22/domain"
	"api/src/sensor_dht22/infrastructure"
	infraestructure "api/src/sensor_dht22/infrastructure"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


type SensorController struct {
	service *application.SensorService
}

func NewSensorController(service *application.SensorService) *SensorController {
	return &SensorController{service: service}
}

func (sc *SensorController) GetSensorData(c *gin.Context) {
	sensorData, err := sc.service.GetSensorData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensorData)
}

func StartDHT22Consumer(rabbitMQ *consumer.RabbitMQ, queueName string, db *core.Database, wsClient *infrastructure.DhtClient) {
	// Declarar la cola
	_, err := rabbitMQ.DeclareQueue(queueName)
	if err != nil {
		log.Fatal("Error declarando la cola:", err)
	}

	// Consumir mensajes de la cola
	msgs, err := rabbitMQ.ConsumeMessages(queueName)
	if err != nil {
		log.Fatal("Error consumiendo mensajes:", err)
	}

	log.Println("Consumidor de RabbitMQ iniciado correctamente")

	// Inicializar el repositorio y el servicio
	repo := infraestructure.NewSensorRepository(db.DB)
	sensorService := application.NewSensorService(repo)

	// Procesar mensajes
	for msg := range msgs {
		var sensorData domain.SensorDHT22
		if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
			log.Println("Error decodificando mensaje:", err)
			continue
		}

		// Procesar los datos del sensor
		if err := sensorService.ProcessSensorData(sensorData); err != nil {
			log.Println("Error procesando datos del sensor dht22:", err)
		}

		// Enviar datos al WebSocket API
        if err := wsClient.SendSensorData(sensorData); err != nil {
            log.Println("Error enviando datos al WebSocket:", err)
        }

		// Crear y publicar notificaciones de eventos especiales
		if sensorData.Temperature > 30 {
			emptyImage := "" 
			notification := domain.Message{
				Header:      "Alerta de Temperatura",
				Description: "La temperatura es demasiado alta.",
				Image:       &emptyImage, 
				Status:      "Alerta",
			}
			publishNotification(rabbitMQ, notification)
		} else if sensorData.Temperature < 25 {
			emptyImage := "" 
			notification := domain.Message{
				Header:      "Alerta de Temperatura",
				Description: "La temperatura es demasiado baja.",
				Image:       &emptyImage, 
				Status:      "Alerta",
			}
			publishNotification(rabbitMQ, notification)
		}
	}
}

// Función auxiliar para publicar notificaciones
func publishNotification(rabbitMQ *consumer.RabbitMQ, notification domain.Message) {
	// Serializar notificación a JSON
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		log.Println("Error serializando notificación:", err)
		return
	}

	// Publicar notificación en la cola
	if err := rabbitMQ.PublishMessage("temperature_messages", notificationJSON); err != nil {
		log.Println("Error publicando notificación:", err)
	}
}