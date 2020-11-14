package handlers

import (
	"net/http"

	"gitlab.com/norzion/temp0/command-handler/app"
	"gitlab.com/norzion/temp0/command-handler/commands"
	"gitlab.com/norzion/temp0/command-handler/errors"
	"gitlab.com/norzion/temp0/command-handler/queries"
)

// HTTPHandler is a http.Handler wrapper for the service
// and fullfills the "Handler" interface
type HTTPHandler struct {
	http.Handler
	app *app.App
}

// NewHTTPHandler sets up the full Event Horizon domain for the TodoMVC app and
// returns a handler exposing some of the components.
func NewHTTPHandler(app *app.App) (HTTPHandler, *errors.Error) {

	// Handle the API.
	h := http.NewServeMux()

	// example: bind some specific handler to an endpoint path
	h.Handle("/v1.0/devices/sensors/measurements/report", NewHTTPCommandHandler(app, commands.NewReportSensorMeasurements(app)))
	h.Handle("/v1.0/devices/alarm/set", NewHTTPCommandHandler(app, commands.NewSetAlarms(app)))
	h.Handle("/v1.0/devices/alarm/notification/set", NewHTTPCommandHandler(app, commands.NewSetAlarmNotification(app)))

	// Note: branch out into separate query-handler later
	h.Handle("/v1.0/devices/sensors/measurements/read", NewHTTPQueryHandler(app, queries.NewReadSensorMeasurements(app)))

	// Handle all static files, only allow what is needed.
	h.HandleFunc("/", staticHTTPHandler)

	return HTTPHandler{
		Handler: h,
		app:     app,
	}, nil
}

// NewHTTPHandlerWithPanicRecovery returns a new http handler with panic recovery
func NewHTTPHandlerWithPanicRecovery(app *app.App) (HTTPHandler, *errors.Error) {
	defer func() {
		r := recover()
		if r != nil {
			// probably fine to omit from unit tests - we will know if this happens xD
			app.Log.Panicln("Recovered from panic")
		}
	}()
	return NewHTTPHandler(app)
}

func staticHTTPHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	// case "/", "/index.html", "/styles.css", "/elm.js":
	// 	http.ServeFile(w, r, "ui"+r.URL.Path)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
