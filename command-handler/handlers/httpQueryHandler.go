package handlers

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/TS-DIY/senz0/command-handler/app"
	"github.com/TS-DIY/senz0/command-handler/errors"
	"github.com/TS-DIY/senz0/command-handler/queries"
)

// HTTPQueryHandler ...
type HTTPQueryHandler struct {
	app   *app.App
	query queries.Query
}

// NewHTTPQueryHandler is a HTTP handler for Queries. It expects a GET/POST with a JSON
// body that will be unmarshalled into the query.
func NewHTTPQueryHandler(app *app.App, query queries.Query) *HTTPQueryHandler {
	return &HTTPQueryHandler{query: query, app: app}
}

// ServeHTTP makes HTTPCommandHandler fulfill the http.Handler interface
func (qh *HTTPQueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	body := make([]byte, 0)

	// assuming queries are always POST requests, using json bodies (empty, if there are no parameters)
	switch r.Method {
	case "GET":
		http.Error(w, "Not implemented", http.StatusMethodNotAllowed)
		break
	case "POST":
		// read the binary data and complain if it's bad
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// No point in unit testing this dude (says the pros)
			http.Error(w, "Could not read command", http.StatusBadRequest)
			return
		}
		body = b
		break
	default:
		http.Error(w, "unsuported method: "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	res, err := qh.query.Handle(context.Background(), body)
	if err != nil {
		switch err.Code() {
		case errors.ErrInvalidRequest, errors.ErrIllegalAction:
			http.Error(w, "Could not handle command: "+err.Error(), http.StatusBadRequest)
			return
		case errors.ErrServerInternal:
			// Only happens on external panics - or if forced by 'hard'checks - no need for testing
			http.Error(w, "Server failed to handle command: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Write(res)

}

// HTTPQueryHandleFunc provides a convenient wrapper for binding several commands and handlers with a router
func HTTPQueryHandleFunc(app *app.App, query queries.Query) http.Handler {
	ch := NewHTTPQueryHandler(app, query)
	return http.HandlerFunc(ch.ServeHTTP)
}
