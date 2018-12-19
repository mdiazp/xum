package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// XUser ...
type XUser struct {
	ID             uint
	Name           string `gorm:"type:varchar(255); not null"`
	LastName       string `gorm:"type:varchar(255); not null"`
	Identification string `gorm:"type:varchar(255); unique_index; not null"`
	PhoneNumber    string `gorm:"type:varchar(100);not null"`
	Description    string `gorm:"type:varchar(500); not null"`
	Actived        bool   `gorm:"not null"`

	XGroupName string `gorm:"-"`
	XGroupID   uint   `gorm:"column:xgroup_id;not null"`
}

// TableName ...
func (XUser) TableName() string {
	return "xuser"
}

// AddSQLConstrainsts ...
func (o *XUser) AddSQLConstrainsts(db *gorm.DB) error {
	e := db.Model(o).AddForeignKey("xgroup_id", "xgroup(id)", "CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s.AddSQLConstrainsts(): %s", o.TableName(), e.Error())
	}
	return e
}
