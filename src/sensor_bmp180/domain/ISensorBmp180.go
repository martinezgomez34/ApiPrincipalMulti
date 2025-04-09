package domain

type BMP180Repository interface {
	ProcessData(data SensorBMP180) error
	GetSensorData() ([]SensorBMP180, error)
}