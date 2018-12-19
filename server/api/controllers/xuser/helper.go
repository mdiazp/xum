package xuser

import (
	"net/http"

	dbhandlers "github.com/mdiazp/xum/server/db/handlers"

	"github.com/mdiazp/xum/server/api"
)

func readXUserFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.XUserFilter {
	f := dbhandlers.XUserFilter{}
	f.NameSubstr = c.GetQString(w, r, "nameSubstr", false)
	f.LastNameSubstr = c.GetQString(w, r, "lastnameSubstr", false)
	f.PhoneNumberSubstr = c.GetQString(w, r, "phoneSubstr", false)
	f.IdentificationPrefix = c.GetQString(w, r, "identificationPrefix", false)
	f.Actived = c.GetQBool(w, r, "actived", false)
	f.XGroupID = c.GetQInt(w, r, "xgroupid", false)
	return &f
}
