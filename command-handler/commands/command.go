package commands

import (
	"context"

	"gitlab.com/norzion/temp0/command-handler/errors"
)

// Command provides an interface that all commands must implement
type Command interface {
	Handle(ctx context.Context, b []byte) *errors.Error
}
