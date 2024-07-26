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

func (r *Repository) GetInfo(id uint) (*models.Regions, error) {
	var info models.Regions
	err := r.DB.Where("id = ?", id).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
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

func (r *Repository) GetRegionTable1Archive() ([]models.RegionTable_1Archive, error) {
	var archives []models.RegionTable_1Archive
	result := r.DB.Find(&archives)
	if result.Error != nil {
		return nil, result.Error
	}
	return archives, nil
}

func (r *Repository) GetRegionTable2Archive() ([]models.RegionTable_2Archive, error) {
	var archives []models.RegionTable_2Archive
	result := r.DB.Find(&archives)
	if result.Error != nil {
		return nil, result.Error
	}
	return archives, nil
}

func (r *Repository) GetRegionTable3Archive() ([]models.RegionTable_3Archive, error) {
	var archives []models.RegionTable_3Archive
	result := r.DB.Find(&archives)
	if result.Error != nil {
		return nil, result.Error
	}
	return archives, nil
}

func (r *Repository) GetFoundationArchive() ([]models.FoundationsArchive, error) {
	var archives []models.FoundationsArchive
	result := r.DB.Find(&archives)
	if result.Error != nil {
		return nil, result.Error
	}
	return archives, nil
}

func (r *Repository) GetInfoArchives() ([]models.RegionsArchive, error) {
	var archives []models.RegionsArchive
	result := r.DB.Find(&archives)
	if result.Error != nil {
		return nil, result.Error
	}
	return archives, nil
}

func (r *Repository) GetArchivedTotalComputersCounts() ([]models.TotalComputersCountArchive, error) {
	var archives []models.TotalComputersCountArchive
	if err := r.DB.Find(&archives).Error; err != nil {
		return nil, err
	}
	return archives, nil
}

func (r *Repository) GetArchivedComputerModelTypes(totalPCId uint) ([]models.ComputerModelTypeArchive, error) {
	var archives []models.ComputerModelTypeArchive
	if err := r.DB.Where("total_pc_id = ?", totalPCId).Find(&archives).Error; err != nil {
		return nil, err
	}
	return archives, nil
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

func (r *Repository) ArchiveInfo(info *models.Regions) error {
	archive := models.RegionsArchive{
		School:         info.School,
		NameSurname:    info.NameSurname,
		Specialization: info.Specialization,
		WhereGraduate:  info.WhereGraduate,
		WhenGraduate:   info.WhenGraduate,
		WorkExperience: info.WorkExperience,
		BirthYear:      info.BirthYear,
		Education:      info.Education,
		PhoneNumber:    info.PhoneNumber,
		RegionID:       info.RegionID,
		ArchivedAt:     time.Now(),
	}
	return r.DB.Create(&archive).Error
}

func (r *Repository) ArchiveTotalComputersCount(info *models.TotalComputersCount) error {
	archive := models.TotalComputersCountArchive{
		School:                info.School,
		TotalCompCount:        info.TotalCompCount,
		TotalWorkingCompCount: info.TotalWorkingCompCount,
		TotalRepairCompCount:  info.TotalRepairCompCount,
		TotalBrokenCompCount:  info.TotalBrokenCompCount,
		ComputersNeed:         info.ComputersNeed,
		Trash:                 info.Trash,
		RegionId:              info.RegionId,
		ArchivedAt:            time.Now(),
	}
	if err := r.DB.Create(&archive).Error; err != nil {
		return err
	}

	// Архивирование связанных данных ComputerModelType
	var modelsToArchive []models.ComputerModelType
	if err := r.DB.Where("total_pc_id = ?", info.Id).Find(&modelsToArchive).Error; err != nil {
		return err
	}

	// Создаем архивные записи для каждой ComputerModelType
	for _, model := range modelsToArchive {
		archiveModel := models.ComputerModelTypeArchive{
			PintiumDDR: model.PintiumDDR,
			Monoblock:  model.Monoblock,
			Laptop:     model.Laptop,
			TotalPCId:  archive.Id, // Используем ID архивной записи TotalComputersCount
			ArchivedAt: time.Now(),
		}
		if err := r.DB.Create(&archiveModel).Error; err != nil {
			return err
		}
	}

	return nil
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

type RegionWithTeacherInfo struct {
	models.Regions
	TeacherRegion string `json:"teacher_region"`
}

func (r *Repository) GetRegionsByRegionID(regionID uint) ([]RegionWithTeacherInfo, error) {
	var regionsWithInfo []RegionWithTeacherInfo
	result := r.DB.Table("regions").
		Select("regions.*, information_about_it_teachers_2023.region as teacher_region").
		Joins("left join information_about_it_teachers_2023 on regions.region_id = information_about_it_teachers_2023.id").
		Where("regions.region_id = ?", regionID).
		Scan(&regionsWithInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return regionsWithInfo, nil
}

type TableSums struct {
	Table1      models.RegionTable_1
	Table2      models.RegionTable_2
	Table3      models.RegionTable_3
	Foundations []models.Foundations
}

type FoundationSums struct {
	Foundation      string `json:"foundation"`
	ComputersCount  int    `json:"computers_count"`
	BoardsCount     int    `json:"boards_count"`
	ProjectorsCount int    `json:"projectors_count"`
	PrinterCount    int    `json:"printer_count"`
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
        SUM(computers_buy_plan6_month) as computers_buy_plan_6_month,
        AVG(computers_buy_percent) as computers_buy_percent,
        SUM(boards_buy_plan6_month) as boards_buy_plan_6_month,
        SUM(boards_buy6_month) as boards_buy_6_month,
        AVG(boards_buy_percent) as boards_buy_percent,
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
        foundation,
        SUM(computers_count) as computers_count,
        SUM(boards_count) as boards_count,
        SUM(projectors_count) as projectors_count,
        SUM(printer_count) as printer_count
    FROM foundations
    GROUP BY foundation`
	if err := r.DB.Raw(query4).Scan(&sums.Foundations).Error; err != nil {
		return sums, err
	}

	return sums, nil
}

func (r *Repository) GetAllRegions() ([]models.InformationAboutItTeachers_2023, error) {
	var regions []models.InformationAboutItTeachers_2023
	if err := r.DB.Find(&regions).Error; err != nil {
		return nil, err
	}
	return regions, nil
}

func (r *Repository) CreateInfo(info *models.Regions) error {
	return r.DB.Create(info).Error
}

func (r *Repository) UpdateInfo(info *models.Regions) error {
	return r.DB.Save(info).Error
}

func (r *Repository) DeleteInfo(id uint) error {
	info := &models.Regions{}
	return r.DB.Where("id = ?", id).Delete(info).Error
}

func (r *Repository) GetPcRegions() ([]models.TotalComputersCountRegions, error) {
	var regions []models.TotalComputersCountRegions
	if err := r.DB.Find(&regions).Error; err != nil {
		return nil, err
	}
	return regions, nil
}

func (r *Repository) CreateTotalComputersCount(totalComputersCount *models.TotalComputersCount) error {
	return r.DB.Create(totalComputersCount).Error
}

func (r *Repository) CreateComputersModelType(computerModelsType *models.ComputerModelType) error {
	return r.DB.Save(computerModelsType).Error
}

func (r *Repository) UpdateTotalComputersCount(totalComputersCount *models.TotalComputersCount) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var existing models.TotalComputersCount
		if err := tx.Where("id = ?", totalComputersCount.Id).First(&existing).Error; err != nil {
			return err
		}

		if err := r.ArchiveTotalComputersCount(&existing); err != nil {
			return err
		}

		if err := tx.Model(&existing).Updates(totalComputersCount).Error; err != nil {
			return err
		}

		for _, compModel := range totalComputersCount.CompModel {
			if err := tx.Where("id = ?", compModel.Id).Updates(&compModel).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *Repository) GetCompCountByRegionID(regionID uint) ([]models.TotalComputersCount, error) {
	var totalComputersCounts []models.TotalComputersCount
	err := r.DB.Preload("CompModel").Where("region_id = ?", regionID).Find(&totalComputersCounts).Error
	if err != nil {
		return nil, err
	}
	return totalComputersCounts, nil
}

func (r *Repository) DeleteTotalComputersCount(id uint) error {
	// Сначала удалите связанные записи
	if err := r.DB.Where("total_pc_id = ?", id).Delete(&models.ComputerModelType{}).Error; err != nil {
		return err
	}

	// Теперь удалите основную запись
	result := r.DB.Delete(&models.TotalComputersCount{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
