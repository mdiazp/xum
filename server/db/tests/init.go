package tests

import (
	"fmt"
	"testing"

	"github.com/mdiazp/xum/server/conf"
	dbhandlers "github.com/mdiazp/xum/server/db/handlers"
)

// CONFIG ...
var CONFIG *conf.Configuration

func init() {
	var e error
	configPath := "/home/kino/my_configs/xum"
	CONFIG, e = conf.LoadConfiguration(configPath, "dev")
	if e != nil {
		panic(fmt.Errorf("Fail loading configuration: %s", e.Error()))
	}
}

// GetDBHandler ...
func GetDBHandler(t *testing.T) dbhandlers.Handler {
	db, e := dbhandlers.NewHandler(CONFIG)
	if e != nil {
		t.Fatalf("Fail newModelHandler: %s", e.Error())
	}

	return db
}
