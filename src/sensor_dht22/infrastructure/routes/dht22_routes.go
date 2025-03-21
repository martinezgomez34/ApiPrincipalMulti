package routes

import (
	"api/src/sensor_dht22/infrastructure/controllers"
	"api/src/sensor_dht22/application"
	"api/src/sensor_dht22/infrastructure"
	"api/src/core"
	"github.com/gin-gonic/gin"
)

func SensorRoutes(router *gin.Engine, db *core.Database) {
	repo := infrastructure.NewSensorRepository(db.DB)
	service := application.NewSensorService(repo)
	controller := controller.NewSensorController(service)

	group := router.Group("/sensor")
	{
		group.GET("/data", controller.GetSensorData)
	}
}