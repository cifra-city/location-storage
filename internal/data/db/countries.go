package db

import (
	"context"

	"github.com/cifra-city/location-storage/internal/data/db/dbcore"
	"github.com/google/uuid"
)

type Countries interface {
	Create(ctx context.Context, name string) (dbcore.Country, error)

	Get(ctx context.Context, id uuid.UUID) (dbcore.Country, error)
	GetByName(ctx context.Context, name string) (dbcore.Country, error)

	Update(ctx context.Context, id uuid.UUID, name string) (dbcore.Country, error)

	Delete(ctx context.Context, id uuid.UUID) error
}

type countries struct {
	queries *dbcore.Queries
}

func NewCountries(queries *dbcore.Queries) Countries {
	return &countries{queries: queries}
}

func (c *countries) Create(ctx context.Context, name string) (dbcore.Country, error) {
	return c.queries.CreateCountry(ctx, name)
}

func (c *countries) Get(ctx context.Context, id uuid.UUID) (dbcore.Country, error) {
	return c.queries.GetCountryByID(ctx, id)
}

func (c *countries) GetByName(ctx context.Context, name string) (dbcore.Country, error) {
	return c.queries.GetCountryByName(ctx, name)
}

func (c *countries) Update(ctx context.Context, id uuid.UUID, name string) (dbcore.Country, error) {
	return c.queries.UpdateCountry(ctx, dbcore.UpdateCountryParams{
		ID:   id,
		Name: name,
	})
}

func (c *countries) Delete(ctx context.Context, id uuid.UUID) error {
	return c.queries.DeleteCountry(ctx, id)
}
