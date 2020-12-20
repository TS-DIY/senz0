package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/TS-DIY/senz0/command-handler/app"
)

func TestNewHTTPHandler(t *testing.T) {
	app := &app.App{
		Log: log.New(os.Stdout, "testing", log.Ldate|log.Ltime|log.Lshortfile),
	}
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "valid http handler",
			// want:    h,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewHTTPHandler(app)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTPHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewHTTPHandlerWithPanicRecovery(t *testing.T) {
	app := &app.App{
		Log: log.New(os.Stdout, "testing", log.Ldate|log.Ltime|log.Lshortfile),
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "valid http handler",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewHTTPHandlerWithPanicRecovery(app)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTPHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_staticHTTPHandler(t *testing.T) {
	type args struct {
		r *http.Request
	}

	tests := []struct {
		name   string
		req    *http.Request
		status int
	}{
		{
			name:   "Default URL behavior",
			req:    newTestRequest(http.NewRequest("POST", "/", bytes.NewReader(nil))),
			status: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			staticHTTPHandler(rr, tt.req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}
		})
	}
}
