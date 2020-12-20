package app

import (
	"log"

	"github.com/TS-DIY/senz0/command-handler/datastore"
)

// App holds all dependencies for our service, generally to be used in handlers
type App struct {
	Log       *log.Logger
	Datastore datastore.Datastore
}
