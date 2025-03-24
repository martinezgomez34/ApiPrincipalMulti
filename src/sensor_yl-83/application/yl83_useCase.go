package application

import "api/src/sensor_yl-83/domain"

type SensorService struct {
	repo domain.SensorRepository
}

func NewSensorService(repo domain.SensorRepository) *SensorService {
	return &SensorService{repo: repo}
}

func (s *SensorService) ProcessSensorData(data domain.SensorYL83) error {
	return s.repo.ProcessData(data)
}