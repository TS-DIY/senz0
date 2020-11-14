package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gitlab.com/norzion/temp0/command-handler/app"
	"gitlab.com/norzion/temp0/command-handler/queries"
)

func TestNewHTTPQueryHandler_ServeHTTP(
	t *testing.T) {
	type args struct {
		query queries.Query
	}

	app := &app.App{
		Log: log.New(os.Stdout, "testing", log.Ldate|log.Ltime|log.Lshortfile),
	}

	h := HTTPQueryHandleFunc(app, NewTestQuery(app))

	tests := []struct {
		name   string
		req    *http.Request
		status int
	}{
		{
			name:   "GET Message",
			req:    newTestRequest(http.NewRequest("GET", "/example", nil)),
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "PUT Message",
			req:    newTestRequest(http.NewRequest("PUT", "/example", nil)),
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "POST Message - non-valid json",
			req:    newTestRequest(http.NewRequest("POST", "/example", bytes.NewReader(nil))),
			status: http.StatusBadRequest,
		},
		{
			name:   "POST Message - OK",
			req:    newTestRequest(http.NewRequest("POST", "/example", bytes.NewReader([]byte(`{"flaf":"giraf"}`)))),
			status: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, tt.req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}

		})
	}
}
