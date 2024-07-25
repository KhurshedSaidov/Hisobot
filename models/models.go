package models

import (
	"gorm.io/gorm"
	"time"
)

type ReportSixMonth struct {
	gorm.Model
	Region        string          `json:"region"`
	RegionTables1 []RegionTable_1 `gorm:"foreignKey:RegionID"`
	RegionTables2 []RegionTable_2 `gorm:"foreignKey:RegionID"`
	RegionTables3 []RegionTable_3 `gorm:"foreignKey:RegionID"`
	Foundations   []Foundations   `gorm:"foreignKey:RegionID""`
}

type RegionTable_1 struct {
	Id                  uint `gorm:"primaryKey"`
	StudentsCount       int  `json:"students_Count"`
	ItRoomsCount        int  `json:"it_rooms_count"`
	FurnitureCount      int  `json:"furniture_count"`
	LocalNet            bool `json:"local_net"`
	ElectroLibCount     int  `json:"electro_lib_count"`
	LibraryCount        int  `json:"library_count"`
	MaterialsCount      int  `json:"materials_count"`
	ItTeachersCount     int  `json:"it_teachers_count"`
	WithHighEduCount    int  `json:"with_high_edu_count"`
	WithSecondEduCount  int  `json:"with_second_edu_count"`
	UnqualifiedCount    int  `json:"unqualified_count"`
	CompleteCourseCount int  `json:"complete_course_count"`
	RegionID            uint `json:"region_id"`
}

type RegionTable_1Archive struct {
	Id                  uint      `gorm:"primaryKey"`
	StudentsCount       int       `json:"students_Count"`
	ItRoomsCount        int       `json:"it_rooms_count"`
	FurnitureCount      int       `json:"furniture_count"`
	LocalNet            bool      `json:"local_net"`
	ElectroLibCount     int       `json:"electro_lib_count"`
	LibraryCount        int       `json:"library_count"`
	MaterialsCount      int       `json:"materials_count"`
	ItTeachersCount     int       `json:"it_teachers_count"`
	WithHighEduCount    int       `json:"with_high_edu_count"`
	WithSecondEduCount  int       `json:"with_second_edu_count"`
	UnqualifiedCount    int       `json:"unqualified_count"`
	CompleteCourseCount int       `json:"complete_course_count"`
	RegionID            uint      `json:"region_id"`
	ArchivedAt          time.Time `json:"archived_at"`
}

type RegionTable_2 struct {
	Id                           uint   `gorm:"primarykey"`
	SchoolsCount                 int    `json:"schools_count"`
	ComputersCount               int    `json:"computers_count"`
	LearningComputersCount       int    `json:"learning_computers_count"`
	RepairComputerCount          int    `json:"repair_computer_count"`
	BrokeComputerCount           int    `json:"broke_computer_count"`
	DecommissionedComputersCount int    `json:"decommissioned_computers_count"`
	PurchasedComputersCount      int    `json:"purchased_computers_count"`
	LicenseComputersCount        int    `json:"license_computers_count"`
	PrintersCount                int    `json:"printers_count"`
	BrokePrintersCount           int    `json:"broke_printers_count"`
	ScannersCount                int    `json:"scanners_count"`
	ConnectedToInternetCount     int    `json:"connected_to_internet_count"`
	ConnectionType               string `json:"connection_type"`
	ElectronicBoardsCount        int    `json:"electronic_boards_count"`
	ProjectorsCount              int    `json:"projectors_count"`
	WebsiteCount                 int    `json:"website_count"`
	RegionID                     uint   `json:"region_id"`
}

type RegionTable_2Archive struct {
	Id                           uint      `gorm:"primarykey"`
	SchoolsCount                 int       `json:"schools_count"`
	ComputersCount               int       `json:"computers_count"`
	LearningComputersCount       int       `json:"learning_computers_count"`
	RepairComputerCount          int       `json:"repair_computer_count"`
	BrokeComputerCount           int       `json:"broke_computer_count"`
	DecommissionedComputersCount int       `json:"decommissioned_computers_count"`
	PurchasedComputersCount      int       `json:"purchased_computers_count"`
	LicenseComputersCount        int       `json:"license_computers_count"`
	PrintersCount                int       `json:"printers_count"`
	BrokePrintersCount           int       `json:"broke_printers_count"`
	ScannersCount                int       `json:"scanners_count"`
	ConnectedToInternetCount     int       `json:"connected_to_internet_count"`
	ConnectionType               string    `json:"connection_type"`
	ElectronicBoardsCount        int       `json:"electronic_boards_count"`
	ProjectorsCount              int       `json:"projectors_count"`
	WebsiteCount                 int       `json:"website_count"`
	RegionID                     uint      `json:"region_id"`
	ArchivedAt                   time.Time `json:"archived_at"`
}

type RegionTable_3 struct {
	Id                     uint    `gorm:"primaryKey"`
	ComputersBuyPlan       int     `json:"computers_buy_plan"`
	ComputersBuyPlan6Month int     `json:"computers_buy_plan_6_month"`
	BoardsBuyPlan6Month    int     `json:"boards_buy_plan_6_month"`
	BoardsBuy6Month        int     `json:"boards_buy_6_month"`
	BoardsBuyPercent       float64 `json:"boards_buy_percent"`
	VideoProjectorCount    int     `json:"video_projector_count"`
	PrintersBuy            int     `json:"printers_buy"`
	IctFinancing           int     `json:"ict_financing"`
	RegionID               uint    `json:"region_id"`
}

type RegionTable_3Archive struct {
	Id                     uint      `gorm:"primaryKey"`
	ComputersBuyPlan       int       `json:"computers_buy_plan"`
	ComputersBuyPlan6Month int       `json:"computers_buy_plan_6_month"`
	BoardsBuyPlan6Month    int       `json:"boards_buy_plan_6_month"`
	BoardsBuy6Month        int       `json:"boards_buy_6_month"`
	BoardsBuyPercent       float64   `json:"boards_buy_percent"`
	VideoProjectorCount    int       `json:"video_projector_count"`
	PrintersBuy            int       `json:"printers_buy"`
	IctFinancing           int       `json:"ict_financing"`
	RegionID               uint      `json:"region_id"`
	ArchivedAt             time.Time `json:"archived_at"`
}

type Foundations struct {
	Id              uint   `gorm:"primaryKey"`
	Foundation      string `json:"foundation"`
	ComputersCount  int    `json:"computers_count"`
	BoardsCount     int    `json:"boards_count"`
	ProjectorsCount int    `json:"projectors_count"`
	PrinterCount    int    `json:"printer_count"`
	RegionID        uint   `json:"region_id"`
}

type FoundationsArchive struct {
	Id              uint      `gorm:"primaryKey"`
	Foundation      string    `json:"foundation"`
	ComputersCount  int       `json:"computers_count"`
	BoardsCount     int       `json:"boards_count"`
	ProjectorsCount int       `json:"projectors_count"`
	PrinterCount    int       `json:"printer_count"`
	RegionID        uint      `json:"region_id"`
	ArchivedAt      time.Time `json:"archived_at"`
}

type InformationAboutItTeachers_2023 struct {
	Id      uint      `gorm:"primaryKey"`
	Region  string    `json:"region"`
	Regions []Regions `gorm:"foreignKey:RegionID"`
}

type Regions struct {
	Id             uint   `gorm:"primaryKey"`
	School         string `json:"school"`
	NameSurname    string `json:"name_surname"`
	Specialization string `json:"specialization"`
	WhereGraduate  string `json:"where_graduate"`
	WhenGraduate   int    `json:"when_graduate"`
	WorkExperience int    `json:"work_experience"`
	BirthYear      int    `json:"birth_year"`
	Education      string `json:"education"`
	PhoneNumber    string `json:"phone_number"`
	RegionID       uint   `json:"region_id"`
}

type RegionsArchive struct {
	Id             uint      `gorm:"primaryKey"`
	School         string    `json:"school"`
	NameSurname    string    `json:"name_surname"`
	Specialization string    `json:"specialization"`
	WhereGraduate  string    `json:"where_graduate"`
	WhenGraduate   int       `json:"when_graduate"`
	WorkExperience int       `json:"work_experience"`
	BirthYear      int       `json:"birth_year"`
	Education      string    `json:"education"`
	PhoneNumber    string    `json:"phone_number"`
	RegionID       uint      `json:"region_id"`
	ArchivedAt     time.Time `json:"archived_at"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"DBname"`
}
