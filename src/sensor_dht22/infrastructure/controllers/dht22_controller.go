package controller

import (
	"api/consumer"
	"api/src/sensor_dht22/application"
	"api/src/sensor_dht22/domain"
	infraestructure "api/src/sensor_dht22/infrastructure"
	"encoding/json"
	"log"
)

func StartDHT22Consumer(rabbitMQ *consumer.RabbitMQ, queueName string) {
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

	// Inicializar el repositorio y el servicio
	repo := infraestructure.NewSensorRepository()
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
			log.Println("Error procesando datos del sensor:", err)
		}

		// Crear y publicar notificaciones de eventos especiales
		if sensorData.Temperature > 30 {
			emptyImage := "" // Cadena vacía
			notification := domain.Message{
				Header:      "Alerta de Temperatura",
				Description: "La temperatura es demasiado alta.",
				Image:       &emptyImage, // Puntero a cadena vacía
				Status:      "Alerta",
			}
			publishNotification(rabbitMQ, notification)
		} else if sensorData.Temperature < 10 {
			emptyImage := "" // Cadena vacía
			notification := domain.Message{
				Header:      "Alerta de Temperatura",
				Description: "La temperatura es demasiado baja.",
				Image:       &emptyImage, // Puntero a cadena vacía
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
	if err := rabbitMQ.PublishMessage("temperature_notifications", notificationJSON); err != nil {
		log.Println("Error publicando notificación:", err)
	}
}