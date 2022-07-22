package dev

import (
	"manager/models"

	"gorm.io/gorm"
)

func RefreshSchema(db *gorm.DB) {
	// this drops all of this projects tables
	// re-creates them with the defined schema
	// and seeds some data
	var tables []interface{}
	tables = append(tables,
		&models.Flag{},
		&models.Audience{},
		&models.Attribute{},
		&models.Condition{},
		&models.FlagLog{},
		&models.AudienceLog{},
		&models.Sdkkey{},
	)

	// drop all relevant tables
	db.Migrator().DropTable(tables...)
	db.Migrator().DropTable("flag_audiences")

	// create all relevant tables
	db.AutoMigrate(tables...)
	// seed some data
	SeedDB(db)
}
