package domain

type LDRRepository interface {
	ProcessData(data SensorLDR) error
	GetSensorData() ([]SensorLDR, error)
}