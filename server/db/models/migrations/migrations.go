package migrations

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mdiazp/xum/server/db/models"
)

// InitDB ...
func InitDB(db *gorm.DB) (e error) {
	xumModel := []models.XUMModel{
		&models.XGroupAdmin{},
		&models.User{},
		&models.XUser{},
		&models.XGroup{},
	}

	model := make([]interface{}, 0)
	for _, x := range xumModel {
		model = append(model, x)
	}

	// db.DropTableIfExists(model...)
	// return
	db.SingularTable(true)
	e = db.AutoMigrate(model...).Error
	if e != nil {
		return fmt.Errorf("db.Automigrate: %s", e.Error())
	}

	for _, x := range xumModel {
		if e = x.AddSQLConstrainsts(db); e != nil {
			return e
		}
	}

	return
}
