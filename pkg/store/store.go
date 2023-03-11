package store

import (
	"context"

	"github.com/smallcase/go-be-template/pkg/binotto"
)

type BinottoStore interface {
	GetById(ctx context.Context, id string) (*binotto.Binotto, error)
	Create(ctx context.Context, venue string) (*binotto.Binotto, error)
	UpdateVenue(ctx context.Context, id, venue string) error
}
