package models

import (
	"gorm.io/gorm"
)

func (fl *Flag) AfterCreate(db *gorm.DB) error {
	log := BuildFlagLog(*fl, "created")
	err := db.Create(&log).Error
	return err
}

func (fl *Flag) AfterUpdate(db *gorm.DB) error {
	log := BuildFlagLog(*fl, "updated")
	err := db.Create(&log).Error
	return err
}

func (fl *Flag) AfterSave(db *gorm.DB) error {
	log := BuildFlagLog(*fl, "updated")
	err := db.Create(&log).Error
	return err
}

func (fl *Flag) AfterDelete(db *gorm.DB) error {
	log := BuildFlagLog(*fl, "deleted")
	err := db.Create(&log).Error
	return err
}

func (a *Audience) AfterCreate(db *gorm.DB) error {
	log := BuildAudLog(*a, "created")
	err := db.Create(&log).Error
	return err
}

func (a *Audience) AfterUpdate(db *gorm.DB) error {
	log := BuildAudLog(*a, "updated")
	err := db.Create(&log).Error
	return err
}

func (a *Audience) AfterSave(db *gorm.DB) error {
	log := BuildAudLog(*a, "updated")
	err := db.Create(&log).Error
	return err
}

func (a *Audience) AfterDelete(db *gorm.DB) error {
	log := BuildAudLog(*a, "deleted")
	err := db.Create(&log).Error
	return err
}

func (a *Attribute) AfterCreate(db *gorm.DB) error {
	log := BuildAttrLog(*a, "created")
	err := db.Create(&log).Error
	return err
}

func (a *Attribute) AfterUpdate(db *gorm.DB) error {
	log := BuildAttrLog(*a, "updated")
	err := db.Create(&log).Error
	return err
}

func (a *Attribute) AfterSave(db *gorm.DB) error {
	log := BuildAttrLog(*a, "updated")
	err := db.Create(&log).Error
	return err
}

func (a *Attribute) AfterDelete(db *gorm.DB) error {
	log := BuildAttrLog(*a, "deleted")
	err := db.Create(&log).Error
	return err
}
