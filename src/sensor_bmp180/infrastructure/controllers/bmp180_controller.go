package controller

import (
	"api/src/sensor_bmp180/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SensorBMP180Controller struct {
	service *application.SensorBMP180Service
}

func NewSensorBMP180Controller(service *application.SensorBMP180Service) *SensorBMP180Controller {
	return &SensorBMP180Controller{service: service}
}

func (sc *SensorBMP180Controller) GetSensorData(c *gin.Context) {
	sensorData, err := sc.service.GetSensorData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensorData)
}
