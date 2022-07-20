package models

import (
	"time"

	"gorm.io/gorm"
)

// getting flag_audiences right with just GORM tags was brutal
// settled on flag_id and audience_id
// looking into flag_key and audience_key instead
// not sure if that matters in real flow of our system though

type Flag struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Key         string         `json:"key" gorm:"type:varchar(30); UNIQUE; NOT NULL"`
	DisplayName string         `json:"displayName" gorm:"type:varchar(30)"`
	Sdkkey      string         `json:"sdkKey" gorm:"type:varchar(30)"`
	Status      bool           `json:"status" gorm:"default:false; NOT NULL"`
	Audiences   []Audience     `json:"audiences" gorm:"many2many:flag_audiences; joinForeignKey:FlagID;joinReferences:AudienceID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type Audience struct {
	gorm.Model
	DisplayName string `gorm:"type:varchar(30)"`
	Key         string `gorm:"type:varchar(30); UNIQUE; NOT NULL"`
	Flags       []Flag `gorm:"many2many:flag_audiences; foreignKey:ID"`
	Conditions  []Condition
}

type Attribute struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Key         string    `json:"key" gorm:"type:varchar(30); UNIQUE; NOT NULL"`
	Type        string    `json:"type" gorm:"type:varchar(10)"`
	DisplayName string    `json:"displayName" gorm:"type:varchar(30)"`
	CreatedAt   time.Time `json:"created_at"`
}

type Condition struct {
	ID           uint `gorm:"primaryKey"`
	AudienceID   uint
	AttributeKey string
	Attribute    Attribute `gorm:"foreignKey:AttributeKey; references:Key"`
	Operator     string
	Vals         string `gorm:"default:'[]';not null"`
}
