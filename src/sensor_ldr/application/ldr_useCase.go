package application

import "api/src/sensor_ldr/domain"

type LDRService struct {
	repo domain.LDRRepository
}

func NewLDRService(repo domain.LDRRepository) *LDRService {
	return &LDRService{repo: repo}
}

func (s *LDRService) ProcessSensorData(data domain.SensorLDR) error {
	return s.repo.ProcessData(data)
}

func (s *LDRService) GetSensorData() ([]domain.SensorLDR, error) {
	return s.repo.GetSensorData()
}