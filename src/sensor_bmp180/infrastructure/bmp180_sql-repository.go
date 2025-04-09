package infrastructure

import (
	"database/sql"
	"api/src/sensor_bmp180/domain"
	"log"
	"time"
)

type BMP180RepositoryImpl struct {
	DB *sql.DB
}

func NewBMP180Repository(db *sql.DB) *BMP180RepositoryImpl {
	return &BMP180RepositoryImpl{DB: db}
}

func (r *BMP180RepositoryImpl) ProcessData(data domain.SensorBMP180) error {
	lastTime, err := r.GetLastInsertTime()
	if err != nil {
		return err
	}
	if time.Since(lastTime) < time.Hour {
		return nil
	}

	query := `INSERT INTO bmp180_data (station_id, status, pressure, created_at) VALUES (?, ?, ?, ?)`
	_, err = r.DB.Exec(query, data.StationID, data.Status, data.Pressure, time.Now())
	if err != nil {
		return err
	}
	log.Printf("Datos BMP180 guardados: %+v\n", data)
	return nil
}


func (r *BMP180RepositoryImpl) GetSensorData() ([]domain.SensorBMP180, error) {
    query := `SELECT station_id, status, pressure, created_at FROM bmp180_data ORDER BY created_at ASC`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var sensorData []domain.SensorBMP180
    for rows.Next() {
        var data domain.SensorBMP180
        err := rows.Scan(&data.StationID, &data.Status, &data.Pressure, &data.CreatedAt)
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

func (r *BMP180RepositoryImpl) GetLastInsertTime() (time.Time, error) {
	var lastTime time.Time
	query := `SELECT created_at FROM bmp180_data ORDER BY created_at DESC LIMIT 1`
	err := r.DB.QueryRow(query).Scan(&lastTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, nil
		}
		return time.Time{}, err
	}
	return lastTime, nil
}
