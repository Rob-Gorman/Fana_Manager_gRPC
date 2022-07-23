package models

import (
	"time"

	_ "gorm.io/gorm"
)

type FlagLog struct {
	ID        uint      `json:"logID" gorm:"primarykey"`
	FlagID    uint      `json:"id"`
	FlagKey   string    `json:"key"`
	EventDesc string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
}

type AudienceLog struct {
	ID          uint      `json:"logID" gorm:"primarykey"`
	AudienceID  uint      `json:"id"`
	AudienceKey string    `json:"key"`
	EventDesc   string    `json:"action"`
	CreatedAt   time.Time `json:"created_at"`
}

type AttributeLog struct {
	ID           uint      `json:"logID" gorm:"primarykey"`
	AttributeID  uint      `json:"id"`
	AttributeKey string    `json:"attributeKey"`
	EventDesc    string    `json:"action"`
	CreatedAt    time.Time `json:"created_at"`
}
