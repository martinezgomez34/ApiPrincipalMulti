package routes

import (
	"api/src/sensor_wind/application"
	"api/src/sensor_wind/infrastructure/controllers"
	"api/src/sensor_wind/infrastructure"
	"api/src/core"
	"github.com/gin-gonic/gin"
)

func WindRoutes(router *gin.Engine, db *core.Database) {
	repo := infrastructure.NewWindRepository(db.DB)
	service := application.NewWindService(repo)
	controller := controllers.NewWindController(service)

	group := router.Group("/wind")
	{
		group.GET("/data", controller.GetSensorData)
	}
}