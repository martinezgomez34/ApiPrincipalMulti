package domain

type SensorRepository interface {
	ProcessData(data SensorYL83) error
}