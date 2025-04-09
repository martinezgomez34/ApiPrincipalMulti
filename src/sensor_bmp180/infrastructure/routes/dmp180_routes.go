package routes

import (
	"api/src/sensor_bmp180/infrastructure/controllers"
	"api/src/sensor_bmp180/application"
	"api/src/sensor_bmp180/infrastructure"
	"api/src/core"
	"github.com/gin-gonic/gin"
)

func BMP180Routes(router *gin.Engine, db *core.Database) {
	repo := infrastructure.NewBMP180Repository(db.DB)
	service := application.NewBMP180Service(repo)
	controller := controller.NewBMP180Controller(service)

	group := router.Group("/bmp180")
	{
		group.GET("/data", controller.GetSensorData)
	}
}