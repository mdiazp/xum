package api

import (
	"github.com/mdiazp/xum/server/usersprovider"
	"github.com/mdiazp/xum/server/usersprovider/xxx"
)

// GetUsersProviderNames ...
func GetUsersProviderNames() []string {
	return []string{"XXX", "AD"}
}

// GetUsersProvider ...
func GetUsersProvider(b Base, provider string) usersprovider.Provider {
	switch provider {
	case "XXX":
		if b.GetEnv() == "dev" {
			return xxx.GetProvider()
		}
		return nil
	default:
		return nil
	}
}
