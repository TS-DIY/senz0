package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gitlab.com/norzion/temp0/command-handler/app"
	"gitlab.com/norzion/temp0/command-handler/commands"
)

func TestNewHTTPCommandHandler_ServeHTTP(
	t *testing.T) {
	type args struct {
		cmd commands.Command
	}

	app := &app.App{
		Log: log.New(os.Stdout, "testing", log.Ldate|log.Ltime|log.Lshortfile),
	}

	h := HTTPCommandHandleFunc(app, NewTestCommand(app))

	tests := []struct {
		name   string
		args   args
		req    *http.Request
		status int
	}{
		{
			name: "GET Message",
			args: args{
				cmd: NewTestCommand(app),
			},
			req:    newTestRequest(http.NewRequest("GET", "/test", nil)),
			status: http.StatusMethodNotAllowed,
		},
		{
			name: "POST Message - non-valid json",
			args: args{
				cmd: NewTestCommand(app),
			},
			req:    newTestRequest(http.NewRequest("POST", "/test", bytes.NewReader(nil))),
			status: http.StatusBadRequest,
		},
		{
			name: "POST Message - OK",
			args: args{
				cmd: NewTestCommand(app),
			},
			req:    newTestRequest(http.NewRequest("POST", "/test", bytes.NewReader([]byte(`{"test":"giraf"}`)))),
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
