package routes

import (
	"api/src/sensor_bmp180/application"
	"api/src/sensor_bmp180/infrastructure"
	"api/src/sensor_bmp180/infrastructure/controllers"
	"api/src/core"
	"github.com/gin-gonic/gin"
)

func SensorBMP180Routes(router *gin.Engine, db *core.Database) {
	repo := infrastructure.NewSensorBMP180Repository(db.DB)
	service := application.NewSensorBMP180Service(repo)
	controller := controller.NewSensorBMP180Controller(service)

	group := router.Group("/sensorbmp180")
	{
		group.GET("/data", controller.GetSensorData)
	}
}
