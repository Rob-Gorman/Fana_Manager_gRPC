package models

import (
	"time"

	_ "gorm.io/gorm"
)

type FlagLog struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	FlagID    uint      `json:"flagID"`
	EventDesc string    `json:"eventDesc"`
	CreatedAt time.Time `json:"created_at"`
}

type AudienceLog struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	AudienceID uint      `json:"audienceID"`
	EventDesc  string    `json:"eventDesc"`
	CreatedAt  time.Time `json:"created_at"`
}
