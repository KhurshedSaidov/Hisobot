package configs

import (
	"Hisobot/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase() error {
	config, err := InitConfigs()
	if err != nil {
		return err
	}

	dbUrl := fmt.Sprintf(
		"host = %s port = %s user = %s password = %s dbname = %s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBname)

	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.ReportSixMonth{})
	DB.AutoMigrate(&models.RegionTable_1{})
	DB.AutoMigrate(&models.RegionTable_1Archive{})
	DB.AutoMigrate(&models.RegionTable_2{})
	DB.AutoMigrate(&models.RegionTable_3{})
	DB.AutoMigrate(&models.Foundations{})
	return nil
}