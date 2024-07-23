package repository

import (
	"Hisobot/models"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetReportFromFirstTableById(regionId, reportId uint) (*models.RegionTable_1, error) {
	var report models.RegionTable_1
	err := r.DB.Where("region_id = ? AND id = ?", regionId, reportId).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *Repository) UpdateFirstTable(table1 *models.RegionTable_1) error {
	return r.DB.Save(table1).Error
}

func (r *Repository) ArchiveFirstTable(report *models.RegionTable_1) error {
	archive := models.RegionTable_1Archive{
		Id:                  report.Id,
		StudentsCount:       report.StudentsCount,
		ItRoomsCount:        report.ItRoomsCount,
		FurnitureCount:      report.FurnitureCount,
		LocalNet:            report.LocalNet,
		ElectroLibCount:     report.ElectroLibCount,
		LibraryCount:        report.LibraryCount,
		MaterialsCount:      report.MaterialsCount,
		ItTeachersCount:     report.ItTeachersCount,
		WithHighEduCount:    report.WithHighEduCount,
		WithSecondEduCount:  report.WithSecondEduCount,
		UnqualifiedCount:    report.UnqualifiedCount,
		CompleteCourseCount: report.CompleteCourseCount,
		RegionID:            report.RegionID,
		ArchivedAt:          time.Now(),
	}
	return r.DB.Create(&archive).Error
}

func (r *Repository) GetReportByRegionId(regionId uint) (*models.ReportSixMonth, error) {
	var report models.ReportSixMonth
	err := r.DB.Preload("RegionTables1").Preload("RegionTables2").Preload("RegionTables3").Preload("Foundations").First(&report, regionId).Error
	return &report, err
}

func (r *Repository) GetRegions() ([]models.ReportSixMonth, error) {
	var regions []models.ReportSixMonth
	if err := r.DB.Find(&regions).Error; err != nil {
		return nil, err
	}
	return regions, nil
}
