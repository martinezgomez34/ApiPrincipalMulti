package domain

type WindRepository interface {
	ProcessData(data SensorWind) error
	GetSensorData() ([]SensorWind, error)
}