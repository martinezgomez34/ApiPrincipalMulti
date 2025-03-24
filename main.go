package main

import (
	"api/consumer"
	"api/src/core"
	controller "api/src/sensor_dht22/infrastructure/controllers"
	controllerYL83 "api/src/sensor_yl-83/infrastructure/controllers"
	route "api/src/sensor_dht22/infrastructure/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	db := core.NewDatabase()
	rabbitMQ := consumer.NewRabbitMQ()
	defer rabbitMQ.Close()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	route.SensorRoutes(router, db)

	queueName := "dht22_queue"
	yl83QueueName := "rain_queue"

	go controller.StartDHT22Consumer(rabbitMQ, queueName, db)
	go controllerYL83.StartYL83Consumer(rabbitMQ, yl83QueueName)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error iniciando el servidor HTTP:", err)
	}
}