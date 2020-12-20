package queries

import (
	"context"

	"github.com/TS-DIY/senz0/command-handler/errors"
)

// Query provides an interface that all queries must implement
type Query interface {
	Handle(ctx context.Context, b []byte) ([]byte, *errors.Error)
}
