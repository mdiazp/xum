package user

import (
	"database/sql"
	"net/http"

	"github.com/mdiazp/xum/server/api"
	"github.com/mdiazp/xum/server/api/controllers"
	"github.com/mdiazp/xum/server/db/models"
)

// RetrieveController ...
type RetrieveController interface {
	controllers.BaseController
}

// NewRetrieveController ...
func NewRetrieveController(base api.Base) RetrieveController {
	return &retrieveController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type retrieveController struct {
	api.Base
}

func (c *retrieveController) GetRoute() string {
	return "/user/{id}"
}

func (c *retrieveController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrieveUser
}

// ServeHTTP ...
func (c *retrieveController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := (uint)(c.GetPInt(w, r, "id"))

	var user models.User
	e := c.DB().RetrieveUserByID(id, &user)
	if e == sql.ErrNoRows {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, user)
}
