package models

import (
	"github.com/jinzhu/gorm"
)

// XGroup ...
type XGroup struct {
	ID          uint
	Name        string `gorm:"type:varchar(255); unique_index; not null"`
	Description string `gorm:"type:varchar(500); not null"`
	Actived     bool   `gorm:"not null"`
}

// TableName ...
func (XGroup) TableName() string {
	return "xgroup"
}

// AddSQLConstrainsts ...
func (o *XGroup) AddSQLConstrainsts(db *gorm.DB) error {
	return nil
}
