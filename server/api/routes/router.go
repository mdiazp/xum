package routes

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"

	"github.com/mdiazp/xum/server/api"
	"github.com/mdiazp/xum/server/api/controllers"
	"github.com/mdiazp/xum/server/api/controllers/session"
	"github.com/mdiazp/xum/server/api/controllers/user"
	"github.com/mdiazp/xum/server/api/controllers/xgroup"
	"github.com/mdiazp/xum/server/api/controllers/xuser"
	"github.com/mdiazp/xum/server/api/middlewares"
)

// Router ...
func Router(base api.Base) http.Handler {
	// Middlewares
	logger := middlewares.Logger(base)

	ctrs := []controllers.BaseController{
		xuser.NewRetrieveController(base),
		xuser.NewRetrieveListController(base),
		xuser.NewCountController(base),
		//xuser.NewCountController(base),
		//xuser.NewUpdateController(base),

		xgroup.NewRetrieveController(base),
		xgroup.NewRetrieveListController(base),
		//xgroup.NewCountController(base),
		//xgroup.NewUpdateController(base),

		session.NewLoginController(base),
		session.NewLogoutController(base),
		session.NewProvidersController(base),

		user.NewCreateController(base),
		user.NewRetrieveController(base),
		user.NewRetrieveListController(base),
	}

	router := mux.NewRouter()

	router.
		PathPrefix("/swagger/").
		Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir(base.PublicFolderPath()))))

	router.Use(logger)

	for _, ctr := range ctrs {
		var h http.Handler = ctr
		if ctr.GetAccess() != "" {
			h = middlewares.CheckAccessControl(base, ctr)
			h = middlewares.MustAuth(base, h)
		}
		router.Handle(ctr.GetRoute(), h).Methods(ctr.GetMethods()...)
	}

	h := cors.AllowAll().Handler(router)

	return h
}
