package models

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Flag struct {
	gorm.Model
	DisplayName string
	Key         string     `gorm:"UNIQUE; NOT NULL"`
	Audiences   []Audience `gorm:"many2many:flag_audiences"`
}

type Audience struct {
	gorm.Model
	DisplayName string
	Key         string `gorm:"UNIQUE; NOT NULL"`
	Flags       []Flag `gorm:"many2many:flag_audiences"`
	Conditions  []Condition
}

type Attribute struct {
	ID          uint   `gorm:"primaryKey"`
	Key         string `gorm:"UNIQUE; NOT NULL"`
	DisplayName string
}

type Condition struct {
	ID          uint `gorm:"primaryKey"`
	AudienceID  uint
	AttributeID Attribute `gorm:"foreignKey:ID; references:Key"`
	Operator    string
	Vals        pgtype.JSONB `gorm:"type:jsonb;default:'[]';not null"`
}
