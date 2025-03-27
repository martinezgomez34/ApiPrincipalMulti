package application

import "api/src/sensor_bmp180/domain"

type SensorBMP180Service struct {
	repo domain.SensorBMP180Repository
}

func NewSensorBMP180Service(repo domain.SensorBMP180Repository) *SensorBMP180Service {
	return &SensorBMP180Service{repo: repo}
}

func (s *SensorBMP180Service) ProcessSensorData(data domain.SensorBMP180) error {
	return s.repo.Save(data)
}

func (s *SensorBMP180Service) GetSensorData() ([]domain.SensorBMP180, error) {
	return s.repo.GetSensorData()
}
