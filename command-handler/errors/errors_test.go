package errors

import (
	"testing"
)

func TestErr_Error(t *testing.T) {
	type fields struct {
		msg        string
		code       string
		stacktrace []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Error",
			fields: fields{msg: "This is an error", code: "", stacktrace: nil},
			want:   "This is an error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				msg:        tt.fields.msg,
				code:       tt.fields.code,
				stacktrace: tt.fields.stacktrace,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Err.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErr_SetStackTrace(t *testing.T) {
	tests := []struct {
		name           string
		wantStacktrace bool
	}{
		{name: "has stacktrace", wantStacktrace: true},
		{name: "has no stacktrace", wantStacktrace: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{}
			if tt.wantStacktrace {
				e.SetStackTrace()
			}

			if tt.wantStacktrace && e.GetStackTrace() == "" {
				t.Errorf("Has no stacktrace")
			}

			if !tt.wantStacktrace && e.GetStackTrace() != "" {
				t.Errorf("Has stacktrace")
			}

		})
	}
}
