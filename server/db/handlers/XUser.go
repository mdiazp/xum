package handlers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mdiazp/xum/server/db/models"
)

// XUserHandler ...
type XUserHandler interface {
	CreateXUser(o *models.XUser) error
	RetrieveXUserByID(id uint, o *models.XUser) error
	RetrieveXUserByIdentification(identification string, o *models.XUser) error
	UpdateXUser(id uint, o *models.XUser) error
	DeleteXUser(id uint) error

	RetrieveXUserList(l *[]models.XUser, filter *XUserFilter,
		orderBy *OrderBy, pag *Paginator) error
	CountXUsers(filter *XUserFilter) (count int, e error)
}

// XUserFilter ...
type XUserFilter struct {
	NameSubstr           *string
	LastNameSubstr       *string
	PhoneNumberSubstr    *string
	IdentificationPrefix *string
	Actived              *bool
	XGroupID             *int
}

/////////////////////////////////////////////////////////////////////////////////////

func (h *handler) CreateXUser(o *models.XUser) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveXUserByID(id uint, o *models.XUser) error {
	return h.Where("id = ?", id).First(o).Error
}

func (h *handler) RetrieveXUserByIdentification(identification string, o *models.XUser) error {
	return h.Where("identification = ?", identification).First(o).Error
}

func (h *handler) UpdateXUser(id uint, o *models.XUser) error {
	return h.Save(o).Error
	// return h.Model(models.XUser{}).Where("id=?", id).Update(o).Error
}

func (h *handler) DeleteXUser(id uint) error {
	return h.Delete(models.User{ID: id}).Error
}

func (h *handler) RetrieveXUserList(l *[]models.XUser, filter *XUserFilter,
	orderBy *OrderBy, pag *Paginator) error {
	db := h.DB.Table("xuser")
	if orderBy == nil {
		orderBy = &OrderBy{By: "id", DESC: false}
	}
	db = makeXUserFilter(db, filter)
	db = orderByAndPaginator(db, orderBy, pag, "xuser")

	db = db.
		Select(
			"xuser.id, xuser.name, xuser.last_name, xuser.identification, " +
				"xuser.phone_number, xuser.description, xuser.actived, " +
				"xgroup.name, xgroup.id").Joins("left join xgroup on xuser.xgroup_id = xgroup.id")

	fmt.Println("SQL Query = ", db.QueryExpr())

	rows, e := db.Rows()
	if e != nil {
		return e
	}
	defer rows.Close()

	for rows.Next() {
		u := models.XUser{}
		e = rows.Scan(&u.ID, &u.Name, &u.LastName, &u.Identification, &u.PhoneNumber,
			&u.Description, &u.Actived, &u.XGroupName, &u.XGroupID)
		if e != nil {
			return e
		}
		*l = append(*l, u)
	}

	return e
}

func (h *handler) CountXUsers(filter *XUserFilter) (count int, e error) {
	e = makeXUserFilter(h.DB.Model(&models.XUser{}), filter).Count(&count).Error
	return
}

func makeXUserFilter(db *gorm.DB, filter *XUserFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	if filter.NameSubstr != nil {
		db = db.Where("xuser.name ILIKE ?", "%"+*(filter.NameSubstr)+"%")
	}
	if filter.LastNameSubstr != nil {
		db = db.Where("xuser.last_name ILIKE ?", "%"+*(filter.LastNameSubstr)+"%")
	}
	if filter.PhoneNumberSubstr != nil {
		db = db.Where("xuser.phone_number ILIKE ?", "%"+*(filter.PhoneNumberSubstr)+"%")
	}
	if filter.Actived != nil {
		db = db.Where("xuser.actived = ?", *(filter.Actived))
	}
	if filter.IdentificationPrefix != nil {
		db = db.Where("xuser.identification ILIKE ?", *(filter.IdentificationPrefix)+"%")
	}
	if filter.XGroupID != nil {
		db = db.Where("xuser.xgroup_id = ?", *(filter.XGroupID))
	}
	return db
}
