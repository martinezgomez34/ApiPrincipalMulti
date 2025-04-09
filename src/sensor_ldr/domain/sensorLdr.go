package domain

import "time"

type SensorLDR struct {
	StationID string    `json:"station_id"`
	Status    string    `json:"status"`
	LDRPercent int      `json:"LDR_percent"`
	CreatedAt time.Time `json:"created_at"`
}