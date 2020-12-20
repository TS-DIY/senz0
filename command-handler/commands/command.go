package commands

import (
	"context"

	"github.com/TS-DIY/senz0/command-handler/errors"
)

// Command provides an interface that all commands must implement
type Command interface {
	Handle(ctx context.Context, b []byte) *errors.Error
}
