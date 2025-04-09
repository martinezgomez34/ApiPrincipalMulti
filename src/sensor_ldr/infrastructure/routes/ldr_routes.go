package routes

import (
	"api/src/sensor_ldr/application"
	"api/src/sensor_ldr/infrastructure/controllers"
	"api/src/sensor_ldr/infrastructure"
	"api/src/core"
	"github.com/gin-gonic/gin"
)

func LDRRoutes(router *gin.Engine, db *core.Database) {
	repo := infrastructure.NewLDRRepository(db.DB)
	service := application.NewLDRService(repo)
	controller := controller.NewLDRController(service)

	group := router.Group("/ldr")
	{
		group.GET("/data", controller.GetSensorData)
	}
}