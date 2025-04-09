package infrastructure

import (
	"database/sql"
	"api/src/sensor_dht22/domain"
	"log"
	"time"
)

type SensorRepositoryImpl struct {
	DB *sql.DB
}

func NewSensorRepository(db *sql.DB) *SensorRepositoryImpl {
	return &SensorRepositoryImpl{DB: db}
}

func (r *SensorRepositoryImpl) ProcessData(data domain.SensorDHT22) error {
    if data.CreatedAt.IsZero() {
        data.CreatedAt = time.Now()
    }

    lastTime, err := r.GetLastInsertTime()
    if err != nil {
        return err
    }

    if time.Since(lastTime) < time.Hour {
        return nil
    }

    query := `INSERT INTO sensor_data (station_id, status, temperature, humidity, created_at) VALUES (?, ?, ?, ?, ?)`
    _, err = r.DB.Exec(query, data.StationID, data.Status, data.Temperature, data.Humidity, data.CreatedAt)
    if err != nil {
        return err
    }
    log.Printf("Datos del sensor guardados: %+v\n", data)
    return nil
}

func (r *SensorRepositoryImpl) GetSensorData() ([]domain.SensorDHT22, error) {
    query := `SELECT station_id, status, temperature, humidity, created_at FROM sensor_data ORDER BY created_at ASC` 
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var sensorData []domain.SensorDHT22
    for rows.Next() {
        var data domain.SensorDHT22
        err := rows.Scan(&data.StationID, &data.Status, &data.Temperature, &data.Humidity, &data.CreatedAt)
        if err != nil {
            return nil, err
        }
        sensorData = append(sensorData, data)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return sensorData, nil
}

func (r *SensorRepositoryImpl) GetLastInsertTime() (time.Time, error) {
    query := `SELECT created_at FROM sensor_data ORDER BY created_at DESC LIMIT 1`
    var lastTime time.Time
    err := r.DB.QueryRow(query).Scan(&lastTime)
    if err != nil {
        return time.Time{}, err
    }
    return lastTime, nil
}
