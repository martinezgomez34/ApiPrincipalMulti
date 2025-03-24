package application

import "api/src/sensor_dht22/domain"

type SensorService struct {
	repo domain.SensorRepository
}

func NewSensorService(repo domain.SensorRepository) *SensorService {
	return &SensorService{repo: repo}
}

func (s *SensorService) ProcessSensorData(data domain.SensorDHT22) error {
	return s.repo.ProcessData(data)
}

func (s *SensorService) GetSensorData() ([]domain.SensorDHT22, error) {
	return s.repo.GetSensorData()
}