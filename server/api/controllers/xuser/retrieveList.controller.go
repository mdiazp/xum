package xuser

import (
	"net/http"

	"github.com/mdiazp/xum/server/api"
	"github.com/mdiazp/xum/server/api/controllers"
	dbhandlers "github.com/mdiazp/xum/server/db/handlers"
	"github.com/mdiazp/xum/server/db/models"
)

// RetrieveListController ...
type RetrieveListController interface {
	controllers.BaseController
}

// NewRetrieveListController ...
func NewRetrieveListController(base api.Base) RetrieveListController {
	return &retrieveListController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type retrieveListController struct {
	api.Base
}

func (c *retrieveListController) GetRoute() string {
	return "/xusers"
}

func (c *retrieveListController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveListController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrieveXUser
}

// ServeHTTP ...
func (c *retrieveListController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := readXUserFilter(c.Base, w, r)
	ob := c.GetQOrderBy(w, r)
	p := c.GetQPaginator(w, r)

	var xusers []models.XUser
	e := c.DB().RetrieveXUserList(&xusers, f, ob, p)

	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, xusers)
}
