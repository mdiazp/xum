package xgroup

import (
	"net/http"

	dbhandlers "github.com/mdiazp/xum/server/db/handlers"

	"github.com/mdiazp/xum/server/api"
)

func readXGroupFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.XGroupFilter {
	f := dbhandlers.XGroupFilter{}
	f.NameSubstr = c.GetQString(w, r, "nameSubstr", false)
	f.Actived = c.GetQBool(w, r, "actived", false)
	f.AdminID = c.GetQInt(w, r, "adminID", false)
	return &f
}
