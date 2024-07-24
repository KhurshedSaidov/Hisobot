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

type TableSums struct {
	Table1 models.RegionTable_1
	Table2 models.RegionTable_2
	Table3 models.RegionTable_3
	Table4 models.Foundations
}

func (r *Repository) GetAllSums() (TableSums, error) {
	var sums TableSums

	// Table 1
	query1 := `
    SELECT
        SUM(students_count) as students_count,
        SUM(it_rooms_count) as it_rooms_count,
        SUM(furniture_count) as furniture_count,
        SUM(electro_lib_count) as electro_lib_count,
        SUM(materials_count) as materials_count,
        SUM(it_teachers_count) as it_teachers_count,
        SUM(with_high_edu_count) as with_high_edu_count,
        SUM(with_second_edu_count) as with_second_edu_count,
        SUM(unqualified_count) as unqualified_count,
        SUM(complete_course_count) as complete_course_count,
        SUM(library_count) as library_count
    FROM region_table_1`
	if err := r.DB.Raw(query1).Scan(&sums.Table1).Error; err != nil {
		return sums, err
	}

	// Table 2
	query2 := `
    SELECT
        SUM(schools_count) as schools_count,
        SUM(computers_count) as computers_count,
        SUM(learning_computers_count) as learning_computers_count,
        SUM(repair_computer_count) as repair_computer_count,
        SUM(broke_computer_count) as broke_computer_count,
        SUM(decommissioned_computers_count) as decommissioned_computers_count,
        SUM(purchased_computers_count) as purchased_computers_count,
        SUM(license_computers_count) as license_computers_count,
        SUM(printers_count) as printers_count,
        SUM(broke_printers_count) as broke_printers_count,
        SUM(scanners_count) as scanners_count,
        SUM(connected_to_internet_count) as connected_to_internet_count,
        SUM(electronic_boards_count) as electronic_boards_count,
        SUM(projectors_count) as projectors_count,
        SUM(website_count) as website_count
    FROM region_table_2`
	if err := r.DB.Raw(query2).Scan(&sums.Table2).Error; err != nil {
		return sums, err
	}

	// Table 3
	query3 := `
    SELECT
        SUM(computers_buy_plan) as computers_buy_plan,
        SUM(computers_buy_plan_6_month) as computers_buy_plan_6_month,
        SUM(boards_buy_plan_6_month) as boards_buy_plan_6_month,
        SUM(boards_buy_6_month) as boards_buy_6_month,
        SUM(video_projector_count) as video_projector_count,
        SUM(printers_buy) as printers_buy,
        SUM(ict_financing) as ict_financing
    FROM region_table_3`
	if err := r.DB.Raw(query3).Scan(&sums.Table3).Error; err != nil {
		return sums, err
	}

	// Table 4
	query4 := `
    SELECT
        SUM(computers_count) as computers_count,
        SUM(boards_count) as boards_count,
        SUM(projectors_count) as projectors_count,
        SUM(printer_count) as printer_count
    FROM foundations`
	if err := r.DB.Raw(query4).Scan(&sums.Table4).Error; err != nil {
		return sums, err
	}

	return sums, nil
}
