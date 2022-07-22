package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

func (fl *Flag) AfterUpdate(db *gorm.DB) (err error) {
	publishThis := BuildFlagset(db)
	res, _ := json.Marshal(publishThis)
	fmt.Println(string(res))
	return nil
}
