package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mdiazp/xum/server/api"
	dbhandlers "github.com/mdiazp/xum/server/db/handlers"
	"github.com/mdiazp/xum/server/db/models"
)

// AuthHeader ...
const AuthHeader = "AuthToken"

// MustAuth ...
func MustAuth(base api.Base, next http.Handler) http.Handler {
	return &Auth{
		next: next,
		Base: base,
	}
}

// Auth ...
type Auth struct {
	next http.Handler
	api.Base
}

func (c *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get(AuthHeader)
	claims, e := c.JWTHandler().GetClaims(tokenString)
	c.WE(w, e, 401)

	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		c.WE(w, fmt.Errorf("Token expired"), 403)
	}

	var user models.User
	e = c.DB().RetrieveUserByUsername(claims.Username, &user)
	if e != nil {
		if e == dbhandlers.ErrRecordNotFound {
			c.WE(w, fmt.Errorf("User Not Found"), 403)
		}
		c.WE(w, e, 500)
	}

	//Check Enabled property
	if !user.Enabled {
		c.WE(w, fmt.Errorf("User is not enabled"), 403)
	}

	c.ContextWriteAuthor(r, &user)

	c.next.ServeHTTP(w, r)
}
