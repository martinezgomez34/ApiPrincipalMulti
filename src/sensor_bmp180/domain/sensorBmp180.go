package domain

import "time"

type SensorBMP180 struct {
	StationID string    `json:"station_id"`
	Status    string    `json:"status"`
	Pressure  float64   `json:"pressure"`
	CreatedAt time.Time `json:"created_at"`
}
