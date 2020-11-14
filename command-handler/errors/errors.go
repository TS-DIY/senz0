package errors

import (
	"fmt"
	"runtime/debug"
)

const (
	// ErrInvalidRequest represents client requests that are not valid, due to e.g., error in the request
	ErrInvalidRequest string = "S401"

	// ErrIllegalAction represents client requests that are not valid, according to domain logic
	ErrIllegalAction string = "S402"

	// ErrServerInternal represents serverside errors due to either faulty code or e.g., interrupted data streams/internal requests
	ErrServerInternal string = "S501"
)

// Error is a customized implementation of the error interface, with more information
type Error struct {
	msg        string
	code       string
	stacktrace []byte
}

// NewError ...
func NewError(msg string) *Error {
	return &Error{msg: msg}
}

// WithCode assigns an internal service code to the service
func (e *Error) WithCode(code string) *Error {
	e.code = code
	return e
}

// Implement the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.msg)

}

// Code returns the Code which represents the internal server state for this specific error
func (e *Error) Code() string {
	return e.code
}

// SetStackTrace sets the stacktrace on an error, typically used for situations where a panic was recovered and we want to know where it occured
func (e *Error) SetStackTrace() {
	e.stacktrace = debug.Stack()
}

// GetStackTrace returns a formatted stacktrace, if any has been set
func (e *Error) GetStackTrace() string {
	if e.stacktrace != nil {
		return fmt.Sprintf("%s", e.stacktrace)
	}

	return ""
}
