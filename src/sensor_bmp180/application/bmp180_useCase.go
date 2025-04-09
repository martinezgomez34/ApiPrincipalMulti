package application

import "api/src/sensor_bmp180/domain"

type BMP180Service struct {
	repo domain.BMP180Repository
}

func NewBMP180Service(repo domain.BMP180Repository) *BMP180Service {
	return &BMP180Service{repo: repo}
}

func (s *BMP180Service) ProcessSensorData(data domain.SensorBMP180) error {
	return s.repo.ProcessData(data)
}

func (s *BMP180Service) GetSensorData() ([]domain.SensorBMP180, error) {
	return s.repo.GetSensorData()
}