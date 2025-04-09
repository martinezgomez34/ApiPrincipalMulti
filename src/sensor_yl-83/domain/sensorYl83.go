package domain

type SensorYL83 struct {
	StationID string    `json:"station_id"`
	Status    string `json:"status"`
	IsRaining bool   `json:"is_raining"`
}
