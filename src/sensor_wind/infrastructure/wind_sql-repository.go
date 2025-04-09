package infrastructure

import (
	"database/sql"
	"api/src/sensor_wind/domain"
	"log"
	"time"
)

type WindRepositoryImpl struct {
	DB *sql.DB
}

func NewWindRepository(db *sql.DB) *WindRepositoryImpl {
	return &WindRepositoryImpl{DB: db}
}

func (r *WindRepositoryImpl) ProcessData(data domain.SensorWind) error {
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

	query := `INSERT INTO wind_data (station_id, status, wind_speed, created_at) VALUES (?, ?, ?, ?)`
	_, err = r.DB.Exec(query, data.StationID, data.Status, data.WindSpeed, data.CreatedAt)
	if err != nil {
		return err
	}
	log.Printf("Datos Wind guardados: %+v\n", data)
	return nil
}

func (r *WindRepositoryImpl) GetSensorData() ([]domain.SensorWind, error) {
    query := `SELECT station_id, status, wind_speed, created_at FROM wind_data ORDER BY created_at ASC`  // Orden ascendente
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var sensorData []domain.SensorWind
    for rows.Next() {
        var data domain.SensorWind
        err := rows.Scan(&data.StationID, &data.Status, &data.WindSpeed, &data.CreatedAt)
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

func (r *WindRepositoryImpl) GetLastInsertTime() (time.Time, error) {
	query := `SELECT created_at FROM wind_data ORDER BY created_at DESC LIMIT 1`
	var lastTime time.Time
	err := r.DB.QueryRow(query).Scan(&lastTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, nil
		}
		return time.Time{}, err
	}
	return lastTime, nil
}
