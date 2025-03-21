package infraestructure

import (
	"api/src/sensor_dht22/domain"
	"log"
)

type SensorRepositoryImpl struct {
	// Aquí podrías agregar dependencias como una base de datos, etc.
}

func NewSensorRepository() *SensorRepositoryImpl {
	return &SensorRepositoryImpl{}
}

func (r *SensorRepositoryImpl) ProcessData(data domain.SensorDHT22) error {
	// Aquí procesarías los datos del sensor, por ejemplo, guardarlos en una base de datos.
	log.Printf("Procesando datos del sensor DHT22: %+v\n", data)
	return nil
}