package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.com/norzion/temp0/command-handler/app"
	"gitlab.com/norzion/temp0/command-handler/errors"
)

func newTestRequest(req *http.Request, err error) *http.Request {
	return req
}

type TestBody struct {
	TestVariable string `json:"test"`
}

type TestCommand struct {
}

func NewTestCommand(app *app.App) *TestCommand {
	return &TestCommand{}
}

func (tc *TestCommand) Handle(ctx context.Context, b []byte) *errors.Error {
	body := &TestBody{}
	if err := json.Unmarshal(b, &body); err != nil {
		return errors.NewError("could not decode payload").WithCode(errors.ErrInvalidRequest)
	}
	return nil
}

// TestQuery ..
type TestQuery struct {
}

// NewTestQuery ..
func NewTestQuery(app *app.App) *TestQuery {
	return &TestQuery{}
}

// Handle is a demo Handle for a command..
func (query *TestQuery) Handle(ctx context.Context, b []byte) ([]byte, *errors.Error) {

	// populate our command body using appropriate unmarshalling
	body := &TestBody{}

	/* if using JSON */
	if err := json.Unmarshal(b, &body); err != nil {
		return nil, errors.NewError("could not decode payload").WithCode(errors.ErrInvalidRequest)
	}

	return nil, nil
}
