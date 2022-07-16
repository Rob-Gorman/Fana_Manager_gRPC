package models

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Flag struct {
	gorm.Model
	DisplayName string
	Key         string
	Audiences   []Audience `gorm:"many2many:flag_audiences"`
}

type Audience struct {
	gorm.Model
	DisplayName string
	Key         string
	Flags       []Flag `gorm:"many2many:flag_audiences"`
	Conditions  []Condition
}

type Attribute struct {
	ID          uint `gorm:"primaryKey"`
	Key         string
	DisplayName string
}

type Condition struct {
	ID        uint      `gorm:"primaryKey"`
	Attribute Attribute `gorm:"foreignKey:UserName"`
	Operator  string
	Vals      pgtype.JSONB `gorm:"type:jsonb;default:'[]';not null"`
}
