package models

import (
	"github.com/jinzhu/gorm"
)

// XUMModel ...
type XUMModel interface {
	AddSQLConstrainsts(db *gorm.DB) error
}
