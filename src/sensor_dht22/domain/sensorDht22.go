package domain

type SensorDHT22 struct {
	Status      string  `json:"status"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}