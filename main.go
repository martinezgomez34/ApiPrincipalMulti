package main

import (
	"api/consumer"
	"api/src/core"
	controller "api/src/sensor_dht22/infrastructure/controllers"
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

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Puedes especificar dominios en lugar de "*" si es necesario
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Duración máxima de la caché de preflight request
	}))
	
	route.SensorRoutes(router, db)

	// Nombre de la cola para el sensor DHT22
	queueName := "dht22_queue"

	// Iniciar consumidor del sensor DHT22
	go controller.StartDHT22Consumer(rabbitMQ, queueName, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error iniciando el servidor HTTP:", err)
	}
}