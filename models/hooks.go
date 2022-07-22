package models

import (
	"gorm.io/gorm"
)

func (fl *Flag) AfterUpdate(db *gorm.DB) (err error) {
	return nil
}
