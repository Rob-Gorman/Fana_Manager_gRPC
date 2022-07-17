package db

import (
	"sovereign/configs"
	"sovereign/models"
	"sovereign/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dbUri := configs.DBConnStr()
	DB, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	utils.HandleErr(err, "Could not connect to DB at URI")

	DB.AutoMigrate(&models.Flag{})
	DB.AutoMigrate(&models.Audience{})
	DB.AutoMigrate(&models.Attribute{})
	DB.AutoMigrate(&models.Condition{})
	return DB
}
