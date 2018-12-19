package user

import (
	"net/http"

	dbhandlers "github.com/mdiazp/xum/server/db/handlers"

	"github.com/mdiazp/xum/server/api"
)

func readUserFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.UserFilter {
	f := dbhandlers.UserFilter{}
	f.NameSubstr = c.GetQString(w, r, "nameSubstr", false)
	f.UsernamePrefix = c.GetQString(w, r, "usernamePrefix", false)
	f.Provider = c.GetQString(w, r, "provider", false)
	f.Enabled = c.GetQBool(w, r, "enabled", false)
	f.Rol = c.GetQString(w, r, "rol", false)
	f.XGroupWhichAdmin = c.GetQInt(w, r, "xgroupWichAdmin", false)
	return &f
}
