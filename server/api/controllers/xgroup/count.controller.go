package xgroup

import (
	"net/http"

	"github.com/mdiazp/xum/server/api"
	"github.com/mdiazp/xum/server/api/controllers"
	dbhandlers "github.com/mdiazp/xum/server/db/handlers"
)

// CountController ...
type CountController interface {
	controllers.BaseController
}

// NewCountController ...
func NewCountController(base api.Base) CountController {
	return &countController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type countController struct {
	api.Base
}

func (c *countController) GetRoute() string {
	return "/xgroupscount"
}

func (c *countController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *countController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrieveXGroup
}

// ServeHTTP ...
func (c *countController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := readXGroupFilter(c.Base, w, r)

	count, e := c.DB().CountXGroups(f)

	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, count)
}
