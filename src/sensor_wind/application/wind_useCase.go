package application

import "api/src/sensor_wind/domain"

type WindService struct {
	repo domain.WindRepository
}

func NewWindService(repo domain.WindRepository) *WindService {
	return &WindService{repo: repo}
}

func (s *WindService) ProcessSensorData(data domain.SensorWind) error {
	return s.repo.ProcessData(data)
}

func (s *WindService) GetSensorData() ([]domain.SensorWind, error) {
	return s.repo.GetSensorData()
}