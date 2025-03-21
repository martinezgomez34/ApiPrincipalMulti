package main

import (
	"api/consumer"
	controller "api/src/sensor_dht22/infrastructure/controllers"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	rabbitMQ := consumer.NewRabbitMQ()
	defer rabbitMQ.Close()

	// Nombre de la cola para el sensor DHT22
	queueName := "dht22_queue"

	// Iniciar consumidor del sensor DHT22
	go controller.StartDHT22Consumer(rabbitMQ, queueName)

	// Aquí podrías agregar más consumidores para otros sensores
	// Ejemplo:
	// go controller.StartYL83Consumer(rabbitMQ, "sensor_yl83_queue")
	// go controller.StartBMP180Consumer(rabbitMQ, "sensor_bmp180_queue")
	// go controller.StartLDRConsumer(rabbitMQ, "sensor_ldr_queue")

	// Mantener el programa en ejecución
	select {}
}