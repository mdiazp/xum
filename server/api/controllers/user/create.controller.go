package user

import (
	"fmt"
	"net/http"

	"github.com/mdiazp/xum/server/api"
	"github.com/mdiazp/xum/server/api/controllers"
	"github.com/mdiazp/xum/server/db/models"
)

// CreateController ...
type CreateController interface {
	controllers.BaseController
}

// NewCreateController ...
func NewCreateController(base api.Base) CreateController {
	return &createController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type createController struct {
	api.Base
}

func (c *createController) GetRoute() string {
	return "/user"
}

func (c *createController) GetMethods() []string {
	return []string{"POST"}
}

// GetAccess ...
func (c *createController) GetAccess() controllers.Permission {
	return controllers.PermissionCreateUser
}

// ServeHTTP ...
func (c *createController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	c.ReadJSON(w, r, &user)

	provider := api.GetUsersProvider(c, user.Provider)
	if provider == nil {
		c.WE(w, fmt.Errorf("Invalid Provider"), 400)
	}

	ur, e := provider.GetUserRecords(user.Username)
	c.WE(w, e, 400)

	user.Username = ur.Username
	user.Name = ur.Name

	e = c.DB().CreateUser(&user)
	c.WE(w, e, 500)

	c.WR(w, 201, user)
}
