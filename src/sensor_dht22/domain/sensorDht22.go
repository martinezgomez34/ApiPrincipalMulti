package domain

import "time"

type SensorDHT22 struct {
	StationID string      `json:"station_id"`
	Status      string    `json:"status"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	CreatedAt   time.Time `json:"created_at"`
}