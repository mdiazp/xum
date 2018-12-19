package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// XGroupAdmin ...
type XGroupAdmin struct {
	ID       uint
	UserID   uint `gorm:"column:user_id; not null"`
	XGroupID uint `gorm:"column:xgroup_id; not null"`
}

// TableName ...
func (XGroupAdmin) TableName() string {
	return "xgroup_admin"
}

//AddSQLConstrainsts ...
func (o *XGroupAdmin) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Model(o).AddForeignKey("user_id", "system_user(id)", "CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(user_id): %s", o.TableName(), e.Error())
		return
	}
	e = db.Model(o).AddForeignKey("xgroup_id", "xgroup(id)", "CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(xgroup_id): %s", o.TableName(), e.Error())
	}

	e = db.Model(o).AddUniqueIndex("idx_system_user_xgroup", "user_id", "xgroup_id").Error
	if e != nil {
		e = fmt.Errorf("%s - AddIndex(): %s", o.TableName(), e.Error())
	}
	return
}
