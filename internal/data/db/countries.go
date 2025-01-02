package db

import (
	"context"

	"github.com/cifra-city/location-storage/internal/data/db/dbcore"
)

type Countries interface {
	Create(ctx context.Context, name string) (dbcore.Country, error)
	Get(ctx context.Context, id uuid.UUID) (dbcore.Country, error)
	GetByName(ctx context.Context, name string) (dbcore.Country, error)
}
