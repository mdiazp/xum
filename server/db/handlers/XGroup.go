package handlers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mdiazp/xum/server/db/models"
)

// XGroupHandler ...
type XGroupHandler interface {
	CreateXGroup(o *models.XGroup) error
	RetrieveXGroupByID(id uint, o *models.XGroup) error
	RetrieveXGroupByIdentification(identification string, o *models.XGroup) error
	UpdateXGroup(id uint, o *models.XGroup) error
	DeleteXGroup(id uint) error

	RetrieveXGroupList(l *[]models.XGroup, filter *XGroupFilter,
		orderBy *OrderBy, pag *Paginator) error
	CountXGroups(filter *XGroupFilter) (count int, e error)
}

// XGroupFilter ...
type XGroupFilter struct {
	NameSubstr *string
	Actived    *bool
	AdminID    *int
}

/////////////////////////////////////////////////////////////////////////////////////

func (h *handler) CreateXGroup(o *models.XGroup) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveXGroupByID(id uint, o *models.XGroup) error {
	return h.Where("id = ?", id).First(o).Error
}

func (h *handler) RetrieveXGroupByIdentification(identification string, o *models.XGroup) error {
	return h.Where("identification = ?", identification).First(o).Error
}

func (h *handler) UpdateXGroup(id uint, o *models.XGroup) error {
	return h.Save(o).Error
	// return h.Model(models.XGroup{}).Where("id=?", id).Update(o).Error
}

func (h *handler) DeleteXGroup(id uint) error {
	return h.Delete(models.User{ID: id}).Error
}

func (h *handler) RetrieveXGroupList(l *[]models.XGroup, filter *XGroupFilter,
	orderBy *OrderBy, pag *Paginator) error {
	db := makeXGroupFilter(h.DB, filter)
	db = orderByAndPaginator(db, orderBy, pag, "xgroup")
	return db.Find(l).Error
}

func (h *handler) CountXGroups(filter *XGroupFilter) (count int, e error) {
	e = makeXGroupFilter(h.DB.Model(&models.XGroup{}), filter).Count(&count).Error
	return
}

func makeXGroupFilter(db *gorm.DB, filter *XGroupFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	if filter.NameSubstr != nil {
		db = db.Where("name ILIKE ?", "%"+*(filter.NameSubstr)+"%")
	}
	if filter.Actived != nil {
		db = db.Where("actived = ?", *(filter.Actived))
	}

	fmt.Println("AdminID filtering has to be doed by join to table xgroup_admin")
	if filter.AdminID != nil {
		db = db.Where("id IN (?)",
			db.Model(&models.XGroupAdmin{}).
				Select("xgroup_id").Where("user_id = ?", *(filter.AdminID)).QueryExpr())
	}

	return db
}
