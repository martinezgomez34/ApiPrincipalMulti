package infrastructure

import (
	"database/sql"
	"api/src/sensor_bmp180/domain"
	"log"
)

type SensorBMP180RepositoryImpl struct {
	DB *sql.DB
}

func NewSensorBMP180Repository(db *sql.DB) *SensorBMP180RepositoryImpl {
	return &SensorBMP180RepositoryImpl{DB: db}
}

func (r *SensorBMP180RepositoryImpl) Save(data domain.SensorBMP180) error {
	query := `INSERT INTO sensorbmp180 (status, pressure) VALUES (?, ?)`
	_, err := r.DB.Exec(query, data.Status, data.Pressure)
	if err != nil {
		return err
	}
	log.Printf("Datos del sensor BMP180 guardados: %+v\n", data)
	return nil
}

func (r *SensorBMP180RepositoryImpl) GetSensorData() ([]domain.SensorBMP180, error) {
	query := `SELECT status, pressure FROM sensorbmp180 ORDER BY id_bmp180 DESC LIMIT 1`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensorData []domain.SensorBMP180
	for rows.Next() {
		var data domain.SensorBMP180
		err := rows.Scan(&data.Status, &data.Pressure)
		if err != nil {
			return nil, err
		}
		sensorData = append(sensorData, data)
	}
	return sensorData, nil
}
