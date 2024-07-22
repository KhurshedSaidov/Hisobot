package service

import (
	"Hisobot/internal/repository"
	"Hisobot/models"
)

type Service struct {
	Repository *repository.Repository
}

func (s *Service) GetReportByRegionId(regionID uint) (*models.ReportSixMonth, error) {
	return s.Repository.GetReportByRegionId(regionID)
}
