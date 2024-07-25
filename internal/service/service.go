package service

import (
	"Hisobot/internal/repository"
	"Hisobot/models"
)

type Service struct {
	Repository *repository.Repository
}

func (s *Service) UpdateFirstTable(studentsCount, itRoomsCount, furnitureCount int, localNet bool, electroLibCount, materialCount,
	itTeachersCount, withHighEduCount, withSecondEduCount, unqualifiedCount, completeCourseCount, libraryCount int, region_id uint) error {
	report, err := s.Repository.GetReportFromFirstTableById(region_id)
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

func (s *Service) UpdateSecondTable(schoolsCount, computersCount, learningComputersCount, repairComputerCount,
	brokeComputersCount, decommissionedComputersCount, purchasedComputersCount, licenseComputersCount, printersCount,
	brokePrintersCount, scannersCount, connectedToInternetCount int, connectionType string, electronicBoardsCount,
	projectorsCount, websiteCount int, regionID uint) error {
	report, err := s.Repository.GetReportFromSecondTableById(regionID)
	if err != nil {
		return err
	}

	if err := s.Repository.ArchiveSecondTable(report); err != nil {
		return err
	}

	report.SchoolsCount = schoolsCount
	report.ComputersCount = computersCount
	report.LearningComputersCount = learningComputersCount
	report.RepairComputerCount = repairComputerCount
	report.BrokeComputerCount = brokeComputersCount
	report.DecommissionedComputersCount = decommissionedComputersCount
	report.PurchasedComputersCount = purchasedComputersCount
	report.LicenseComputersCount = licenseComputersCount
	report.PrintersCount = printersCount
	report.BrokePrintersCount = brokePrintersCount
	report.ScannersCount = scannersCount
	report.ConnectedToInternetCount = connectedToInternetCount
	report.ConnectionType = connectionType
	report.ElectronicBoardsCount = electronicBoardsCount
	report.ProjectorsCount = projectorsCount
	report.WebsiteCount = websiteCount

	return s.Repository.UpdateSecondTable(report)
}

func (s *Service) UpdateThirdTable(computersBuyPlan, computersBuyPlan6Month int,
	boardsBuyPlan6Month, boardsBuy6Month int, boardsBuyPercent float64, videoProjectorCount, printersBuy, ictFinancing int,
	regionID uint) error {
	report, err := s.Repository.GetReportFromThirdTableById(regionID)
	if err != nil {
		return err
	}

	if err := s.Repository.ArchiveThirdTable(report); err != nil {
		return err
	}

	report.ComputersBuyPlan = computersBuyPlan
	report.ComputersBuyPlan6Month = computersBuyPlan6Month
	report.BoardsBuyPlan6Month = boardsBuyPlan6Month
	report.BoardsBuy6Month = boardsBuy6Month
	report.BoardsBuyPercent = boardsBuyPercent
	report.VideoProjectorCount = videoProjectorCount
	report.PrintersBuy = printersBuy
	report.IctFinancing = ictFinancing

	return s.Repository.UpdateThirdTable(report)
}

func (s *Service) UpdateFoundations(id uint, foundation string, computersCount, boardsCount,
	projectorsCount, printerCount int, regionID uint) error {
	report, err := s.Repository.GetReportFromFoundationsById(id, regionID)
	if err != nil {
		return err
	}

	if err := s.Repository.ArchiveFoundations(report); err != nil {
		return err
	}

	report.Foundation = foundation
	report.ComputersCount = computersCount
	report.BoardsCount = boardsCount
	report.ProjectorsCount = projectorsCount
	report.PrinterCount = printerCount

	return s.Repository.UpdateFoundations(report)
}

func (s *Service) UpdateInfo(id uint, school, nameSurname, specialization, whereGraduate string,
	whenGraduate, workExperience, birthYear int, education, phoneNumber string) error {
	info, err := s.Repository.GetInfo(id)
	if err != nil {
		return err
	}

	if err := s.Repository.ArchiveInfo(info); err != nil {
		return err
	}
	info.School = school
	info.NameSurname = nameSurname
	info.Specialization = specialization
	info.WhereGraduate = whereGraduate
	info.WhenGraduate = whenGraduate
	info.WorkExperience = workExperience
	info.BirthYear = birthYear
	info.Education = education
	info.PhoneNumber = phoneNumber

	return s.Repository.UpdateInfo(info)

}

func (s *Service) GetReportByRegionId(regionID uint) (*models.ReportSixMonth, error) {
	return s.Repository.GetReportByRegionId(regionID)
}

func (s *Service) GetAllRegions() ([]models.ReportSixMonth, error) {
	return s.Repository.GetRegions()
}

func (s *Service) GetAllSums() (repository.TableSums, error) {
	return s.Repository.GetAllSums()
}

func (s *Service) CreateInfo(school, nameSurname, specialization, whereGraduate string,
	whenGraduate, workExperience, birthYear int, education, phoneNumber string, regionID uint) error {
	info := &models.Regions{
		School:         school,
		NameSurname:    nameSurname,
		Specialization: specialization,
		WhereGraduate:  whereGraduate,
		WhenGraduate:   whenGraduate,
		WorkExperience: workExperience,
		BirthYear:      birthYear,
		Education:      education,
		PhoneNumber:    phoneNumber,
		RegionID:       regionID,
	}

	return s.Repository.CreateInfo(info)
}

func (s *Service) DeleteInfo(id uint) error {
	return s.Repository.DeleteInfo(id)
}

func (s *Service) GetRegionsByRegionID(regionID uint) ([]repository.RegionWithTeacherInfo, error) {
	return s.Repository.GetRegionsByRegionID(regionID)
}
