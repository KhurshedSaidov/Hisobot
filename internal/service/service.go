package service

import (
	"Hisobot/internal/repository"
	"Hisobot/models"
)

type Service struct {
	Repository *repository.Repository
}

func (s *Service) UpdateFirstTable(id uint, studentsCount, itRoomsCount, furnitureCount int, localNet bool, electroLibCount, materialCount,
	itTeachersCount, withHighEduCount, withSecondEduCount, unqualifiedCount, completeCourseCount, libraryCount int, region_id uint) error {
	report, err := s.Repository.GetReportFromFirstTableById(region_id, id)
	if err != nil {
		return err
	}

	if err := s.Repository.ArchiveFirstTable(report); err != nil {
		return err
	}

	report.StudentsCount = studentsCount
	report.ItRoomsCount = itRoomsCount
	report.FurnitureCount = furnitureCount
	report.LocalNet = localNet
	report.ElectroLibCount = electroLibCount
	report.MaterialsCount = materialCount
	report.ItTeachersCount = itTeachersCount
	report.WithHighEduCount = withHighEduCount
	report.WithSecondEduCount = withSecondEduCount
	report.UnqualifiedCount = unqualifiedCount
	report.CompleteCourseCount = completeCourseCount
	report.LibraryCount = libraryCount

	return s.Repository.UpdateFirstTable(report)
}

func (s *Service) GetReportByRegionId(regionID uint) (*models.ReportSixMonth, error) {
	return s.Repository.GetReportByRegionId(regionID)
}

func (s *Service) GetAllRegions() ([]models.ReportSixMonth, error) {
	return s.Repository.GetRegions()
}
