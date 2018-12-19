package xuser

import (
	"fmt"
	"net/http"

	"github.com/mdiazp/xum/server/api"
	"github.com/mdiazp/xum/server/api/controllers"
	"github.com/mdiazp/xum/server/db/models"
)

// UpdateController ...
type UpdateController interface {
	controllers.BaseController
}

// NewUpdateController ...
func NewUpdateController(base api.Base) UpdateController {
	return &updateController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type updateController struct {
	api.Base
}

func (c *updateController) GetRoute() string {
	return "/xgroup/{id}"
}

func (c *updateController) GetMethods() []string {
	return []string{"PATCH"}
}

// GetAccess ...
func (c *updateController) GetAccess() controllers.Permission {
	return controllers.PermissionUpdateXUser
}

// ServeHTTP ...
func (c *updateController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := (uint)(c.GetPInt(w, r, "id"))
	var xgroup models.User
	c.ReadJSON(w, r, &xgroup)
	xgroup.ID = id

	fmt.Println("xgroup = ", xgroup)

	e := c.DB().UpdateUser(id, &xgroup)
	c.WE(w, e, 500)
	c.WR(w, 200, xgroup)
}
