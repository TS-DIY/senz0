package handlers

import (
	"context"
	"io/ioutil"
	"net/http"

	"gitlab.com/norzion/temp0/command-handler/app"
	"gitlab.com/norzion/temp0/command-handler/commands"
	"gitlab.com/norzion/temp0/command-handler/errors"
)

// HTTPCommandHandler ...
type HTTPCommandHandler struct {
	app *app.App
	cmd commands.Command
}

// NewHTTPCommandHandler is a HTTP handler for Commands. It expects a POST with a JSON
// body that will be unmarshalled into the command.
func NewHTTPCommandHandler(app *app.App, cmd commands.Command) *HTTPCommandHandler {
	return &HTTPCommandHandler{cmd: cmd, app: app}
}

// ServeHTTP makes HTTPCommandHandler fulfill the http.Handler interface
func (ch *HTTPCommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// reject anything but POST requests
	if r.Method != "POST" {
		http.Error(w, "Unsupported method: "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	// read the binary data and complain if it's bad
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// No point in unit testing this dude (says the pros)
		http.Error(w, "Could not read command", http.StatusBadRequest)
		return
	}

	// NOTE: Use a new context when handling, else it will be cancelled with
	// the HTTP request which will cause stuff to fail, if it runs
	// async in goroutines past the request.
	if err := ch.cmd.Handle(context.Background(), b); err != nil {
		switch err.Code() {
		case errors.ErrInvalidRequest, errors.ErrIllegalAction:
			http.Error(w, "Could not handle command: "+err.Error(), http.StatusBadRequest)
			return
		case errors.ErrServerInternal:
			http.Error(w, "Server failed to handle command: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)

}

// HTTPCommandHandleFunc provides a convenient wrapper for binding several commands and handlers with a router
func HTTPCommandHandleFunc(app *app.App, cmd commands.Command) http.Handler {
	ch := NewHTTPCommandHandler(app, cmd)
	return http.HandlerFunc(ch.ServeHTTP)
}
