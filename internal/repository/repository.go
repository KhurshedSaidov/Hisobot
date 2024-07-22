package repository

import (
	"Hisobot/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetReportByRegionId(regionId uint) (*models.ReportSixMonth, error) {
	var report models.ReportSixMonth
	err := r.DB.Preload("RegionTables1").Preload("RegionTables2").Preload("RegionTables3").First(&report, regionId).Error
	return &report, err
}
