package consumer

import (
	"api/src/core"
	"api/src/sensor_bmp180/application"
	"api/src/sensor_bmp180/domain"
	infraestructure "api/src/sensor_bmp180/infrastructure"
	"encoding/json"
	"log"
)

func StartBMP180Consumer(rabbitMQ *RabbitMQ, queueName string, db *core.Database) {
	_, err := rabbitMQ.DeclareQueue(queueName)
	if err != nil {
		log.Fatal("Error declarando la cola:", err)
	}

	msgs, err := rabbitMQ.ConsumeMessages(queueName)
	if err != nil {
		log.Fatal("Error consumiendo mensajes:", err)
	}

	log.Println("Consumidor de RabbitMQ para BMP180 iniciado correctamente")

	repo := infraestructure.NewSensorBMP180Repository(db.DB)
	sensorService := application.NewSensorBMP180Service(repo)

	for msg := range msgs {
		var sensorData domain.SensorBMP180
		if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
			log.Println("Error decodificando mensaje:", err)
			continue
		}

	if err := sensorService.ProcessSensorData(sensorData); err != nil {
		log.Println("Error procesando datos del sensor BMP180:", err)
	}

	}
}
