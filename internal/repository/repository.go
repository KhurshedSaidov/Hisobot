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

func (r *Repository) GetReportFromFirstTableById(regionId uint) (*models.RegionTable_1, error) {
	var report models.RegionTable_1
	err := r.DB.Where("region_id = ?", regionId).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *Repository) GetReportFromSecondTableById(regionId uint) (*models.RegionTable_2, error) {
	var report models.RegionTable_2
	err := r.DB.Where("region_id = ?", regionId).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *Repository) GetReportFromThirdTableById(regionId uint) (*models.RegionTable_3, error) {
	var report models.RegionTable_3
	err := r.DB.Where("region_id = ?", regionId).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *Repository) GetReportFromFoundationsById(id, regionId uint) (*models.Foundations, error) {
	var report models.Foundations
	err := r.DB.Where("id = ? AND region_id = ?", id, regionId).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *Repository) UpdateFirstTable(table1 *models.RegionTable_1) error {
	return r.DB.Save(table1).Error
}

func (r *Repository) UpdateSecondTable(table1 *models.RegionTable_2) error {
	return r.DB.Save(table1).Error
}

func (r *Repository) UpdateThirdTable(table1 *models.RegionTable_3) error {
	return r.DB.Save(table1).Error
}

func (r *Repository) UpdateFoundations(table1 *models.Foundations) error {
	return r.DB.Save(table1).Error
}

func (r *Repository) ArchiveFirstTable(report *models.RegionTable_1) error {
	archive := models.RegionTable_1Archive{
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

func (r *Repository) ArchiveSecondTable(report *models.RegionTable_2) error {
	archive := models.RegionTable_2Archive{
		SchoolsCount:                 report.SchoolsCount,
		ComputersCount:               report.ComputersCount,
		LearningComputersCount:       report.LearningComputersCount,
		RepairComputerCount:          report.RepairComputerCount,
		BrokeComputerCount:           report.BrokeComputerCount,
		DecommissionedComputersCount: report.DecommissionedComputersCount,
		PurchasedComputersCount:      report.PurchasedComputersCount,
		LicenseComputersCount:        report.LicenseComputersCount,
		PrintersCount:                report.PrintersCount,
		BrokePrintersCount:           report.BrokePrintersCount,
		ScannersCount:                report.ScannersCount,
		ConnectedToInternetCount:     report.ConnectedToInternetCount,
		ConnectionType:               report.ConnectionType,
		ElectronicBoardsCount:        report.ElectronicBoardsCount,
		ProjectorsCount:              report.ProjectorsCount,
		WebsiteCount:                 report.WebsiteCount,
		RegionID:                     report.RegionID,
		ArchivedAt:                   time.Now(),
	}
	return r.DB.Create(&archive).Error
}

func (r *Repository) ArchiveThirdTable(report *models.RegionTable_3) error {
	archive := models.RegionTable_3Archive{
		ComputersBuyPlan:       report.ComputersBuyPlan,
		ComputersBuyPlan6Month: report.ComputersBuyPlan6Month,
		BoardsBuy6Month:        report.BoardsBuy6Month,
		BoardsBuyPlan6Month:    report.BoardsBuyPlan6Month,
		BoardsBuyPercent:       report.BoardsBuyPercent,
		VideoProjectorCount:    report.VideoProjectorCount,
		PrintersBuy:            report.PrintersBuy,
		IctFinancing:           report.IctFinancing,
		RegionID:               report.RegionID,
		ArchivedAt:             time.Now(),
	}
	return r.DB.Create(&archive).Error
}

func (r *Repository) ArchiveFoundations(report *models.Foundations) error {
	archive := models.FoundationsArchive{
		Foundation:      report.Foundation,
		ComputersCount:  report.ComputersCount,
		BoardsCount:     report.BoardsCount,
		ProjectorsCount: report.ProjectorsCount,
		PrinterCount:    report.PrinterCount,
		RegionID:        report.RegionID,
		ArchivedAt:      time.Now(),
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
