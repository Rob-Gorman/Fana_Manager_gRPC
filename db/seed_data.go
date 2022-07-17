package db

import (
	"fmt"
	"sovereign/models"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	seedFlags(db)
	seedAttributes(db)
	seedAudiences(db)
	seedFlagAuds(db)
}

func seedFlags(db *gorm.DB) {
	var flags = []models.Flag{
		{Key: "fake-flag-1"},
		{Key: "experimental-flag-1"},
		{Key: "development-flag-1"},
	}
	db.Create(&flags)
}

func seedAttributes(db *gorm.DB) {
	var attrs = []models.Attribute{
		{Key: "state", Type: "STR"},
		{Key: "student", Type: "BOOL"},
		{Key: "beta", Type: "BOOL"},
	}
	db.Create(&attrs)
}

func seedAudiences(db *gorm.DB) {
	ca_stu := models.Audience{
		Key: "california_students",
		Conditions: []models.Condition{
			{
				AttributeID: 1,
				Operator:    "EQ",
				Vals:        "california",
			},
			{
				AttributeID: 2,
				Operator:    "EQ",
				Vals:        "true",
			},
		},
	}

	db.Create(&ca_stu)
}

func seedFlagAuds(db *gorm.DB) {
	var caliStuAud models.Audience
	db.Where("key = ?", "california_students").First(&caliStuAud)

	var singleFlag models.Flag
	db.First(&singleFlag)
	fmt.Println(singleFlag.Key)
	// type flagAud struct {
	// 	flagID      uint
	// 	audienceKey string
	// }
	// var insVal flagAud
	// insVal := flagAud{
	// 	flagID:      singleFlag.ID,
	// 	audienceKey: "california_students",
	// }
	// db.Table("flag_audiences").Create(&insVal)
	// db.Model(&singleFlag).Association("Audiences").Append("audiences", db.Model(&models.Audience{}).Select("key").Where("audiences.key = ?", "california_students"))
	// db.Model(&singleFlag).Association("Audiences").Append(&caliStuAud)
	// UPDATE "users" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);

	// fmt.Println("aud Key?", caliStuAud.Key)

	db.Model(&singleFlag).Association("Audiences")
	singleFlag.Audiences = []models.Audience{caliStuAud}
	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&singleFlag)
	db.First(&singleFlag)
	fmt.Println(singleFlag.Audiences[0].Key) // HOLY FUCKING SHIT YES GOD YES

	// db.Save(&singleFlag)
}
