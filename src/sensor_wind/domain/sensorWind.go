package domain

import "time"

type SensorWind struct {
	StationID string    `json:"station_id"`
	Status    string    `json:"status"`
	WindSpeed float64   `json:"wind_speed"`
	CreatedAt time.Time `json:"created_at"`
}