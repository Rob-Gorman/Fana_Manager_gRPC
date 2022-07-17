package models

import (
	"gorm.io/gorm"
)

// getting flag_audiences right with just GORM tags was brutal
// settled on flag_id and audience_id
// looking into flag_key and audience_key instead
// not sure if that matters in real flow of our system though

type Flag struct {
	gorm.Model
	Key         string     `gorm:"type:varchar(30); UNIQUE; NOT NULL"`
	DisplayName string     `gorm:"type:varchar(30)"`
	Audiences   []Audience `gorm:"many2many:flag_audiences; joinForeignKey:FlagID;joinReferences:AudienceID"`
}

type Audience struct {
	gorm.Model
	DisplayName string `gorm:"type:varchar(30)"`
	Key         string `gorm:"type:varchar(30); UNIQUE; NOT NULL"`
	Flags       []Flag `gorm:"many2many:flag_audiences; foreignKey:ID"`
	Conditions  []Condition
}

type Attribute struct {
	ID          uint   `gorm:"primaryKey"`
	Key         string `gorm:"type:varchar(30); UNIQUE; NOT NULL"`
	Type        string `gorm:"type:varchar(10)"`
	DisplayName string `gorm:"type:varchar(30)"`
}

type Condition struct {
	ID           uint `gorm:"primaryKey"`
	AudienceID   uint
	AttributeKey string
	Attribute    Attribute `gorm:"foreignKey:AttributeKey; references:Key"`
	Operator     string
	Vals         string `gorm:"default:'[]';not null"`
}
