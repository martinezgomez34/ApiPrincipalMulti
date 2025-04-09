package infrastructure

import (
	"database/sql"
	"api/src/sensor_ldr/domain"
	"log"
	"time"
)

type LDRRepositoryImpl struct {
	DB *sql.DB
}

func NewLDRRepository(db *sql.DB) *LDRRepositoryImpl {
	return &LDRRepositoryImpl{DB: db}
}

func (r *LDRRepositoryImpl) ProcessData(data domain.SensorLDR) error {
	lastTime, err := r.GetLastInsertTime()
	if err != nil {
		return err
	}

	if time.Since(lastTime) < time.Hour {
		return nil
	}

	query := `INSERT INTO ldr_data (station_id, status, ldr_percent, created_at) VALUES (?, ?, ?, ?)`
	_, err = r.DB.Exec(query, data.StationID, data.Status, data.LDRPercent, time.Now())
	if err != nil {
		return err
	}
	log.Printf("Datos LDR guardados: %+v\n", data)
	return nil
}


func (r *LDRRepositoryImpl) GetSensorData() ([]domain.SensorLDR, error) {
	query := `SELECT station_id, status, ldr_percent, created_at FROM ldr_data ORDER BY created_at DESC LIMIT 1`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var sensorData []domain.SensorLDR
    for rows.Next() {
        var data domain.SensorLDR
        err := rows.Scan(&data.StationID, &data.Status, &data.LDRPercent, &data.CreatedAt)
        if err != nil {
            return nil, err
        }
        sensorData = append(sensorData, data)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return sensorData, nil
}

func (r *LDRRepositoryImpl) GetLastInsertTime() (time.Time, error) {
	var lastTime time.Time
	query := `SELECT created_at FROM ldr_data ORDER BY created_at DESC LIMIT 1`
	err := r.DB.QueryRow(query).Scan(&lastTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, nil 
		}
		return time.Time{}, err
	}
	return lastTime, nil
}
