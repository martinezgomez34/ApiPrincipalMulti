package domain

type SensorRepository interface {
	ProcessData(data SensorDHT22) error
	GetSensorData() ([]SensorDHT22, error)
}