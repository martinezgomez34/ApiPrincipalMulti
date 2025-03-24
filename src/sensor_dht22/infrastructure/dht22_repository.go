package infrastructure

import (
	"database/sql"
	"api/src/sensor_dht22/domain"
	"log"
)

type SensorRepositoryImpl struct {
	DB *sql.DB
}

func NewSensorRepository(db *sql.DB) *SensorRepositoryImpl {
	return &SensorRepositoryImpl{DB: db}
}

func (r *SensorRepositoryImpl) ProcessData(data domain.SensorDHT22) error {
	query := `INSERT INTO sensor_data (status, temperature, humidity) VALUES (?, ?, ?)`
	_, err := r.DB.Exec(query, data.Status, data.Temperature, data.Humidity)
	if err != nil {
		return err
	}
	log.Printf("Datos del sensor guardados: %+v\n", data)
	return nil
}

func (r *SensorRepositoryImpl) GetSensorData() ([]domain.SensorDHT22, error) {
	query := `SELECT status, temperature, humidity FROM sensor_data ORDER By id_dht22 DESC LIMIT 1`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensorData []domain.SensorDHT22
	for rows.Next() {
		var data domain.SensorDHT22
		err := rows.Scan(&data.Status, &data.Temperature, &data.Humidity)
		if err != nil {
			return nil, err
		}
		sensorData = append(sensorData, data)
	}
	return sensorData, nil
}