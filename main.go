package main

import (
	"api/consumer"
	"api/src/core"
	bmpClient "api/src/sensor_bmp180/infrastructure"
	dhtClient "api/src/sensor_dht22/infrastructure"
	ldrClient "api/src/sensor_ldr/infrastructure"
	windClient "api/src/sensor_wind/infrastructure"
	controllerBMP180 "api/src/sensor_bmp180/infrastructure/controllers"
	bmp180routes "api/src/sensor_bmp180/infrastructure/routes"
	controllerDHT22 "api/src/sensor_dht22/infrastructure/controllers"
	dht22routes "api/src/sensor_dht22/infrastructure/routes"
	controllerYL83 "api/src/sensor_yl-83/infrastructure/controllers"
	ldrroutes "api/src/sensor_ldr/infrastructure/routes"
	controllerLDR "api/src/sensor_ldr/infrastructure/controllers"
	windroutes "api/src/sensor_wind/infrastructure/routes"
	controllerWIND "api/src/sensor_wind/infrastructure/controllers"
	"log"
	"os"
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
	
	dht22routes.SensorRoutes(router, db)
	bmp180routes.BMP180Routes(router, db)
	ldrroutes.LDRRoutes(router, db)
	windroutes.WindRoutes(router, db)

	dht22QueueName := "dht22_queue"
	yl83QueueName := "rain_queue"
	bmp180QueueName:= "bmp180_queue"
	ldrQueueName := "ldr_queue"
	windQueueName := "wind_queue"

	wsAPIURL := os.Getenv("WS_API_URL") 
    dhtClient := dhtClient.NewDhtClient(wsAPIURL)
	bmpClient := bmpClient.NewBmpClient(wsAPIURL)
	ldrClient := ldrClient.NewLDRClient(wsAPIURL)
	windClient := windClient.NewWindClient(wsAPIURL)

	go controllerDHT22.StartDHT22Consumer(rabbitMQ, dht22QueueName, db , dhtClient)
	go controllerYL83.StartYL83Consumer(rabbitMQ, yl83QueueName)
	go controllerBMP180.StartBMP180Consumer(rabbitMQ, bmp180QueueName, db, bmpClient)
	go controllerLDR.StartLDRConsumer(rabbitMQ, ldrQueueName, db, ldrClient)
	go controllerWIND.StartWindConsumer(rabbitMQ, windQueueName, db, windClient)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal("Error iniciando el servidor HTTP:", err)
	}
}