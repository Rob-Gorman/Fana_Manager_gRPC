package models

import (
	"gorm.io/gorm"
)

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
	ID          uint `gorm:"primaryKey"`
	AudienceID  uint
	AttributeID uint `gorm:"foreignKey:ID; references:Key"`
	Operator    string
	Vals        string `gorm:"default:'[]';not null"`
}
