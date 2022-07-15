package configs

import (
	"fmt"
	"sovereign/utils"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() *gorm.DB {
	dbUri := dbConnStr()
	DB, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	utils.HandleErr(err, "Could not connect to the string")
	return DB
}

func dbConnStr() string {
	variables := getEnvVars("DB_HOST", "DB_USER", "DB_NAME", "DB_PW", "DB_PORT")
	dbUri := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		variables...,
	)

	return dbUri
}
