package db

import (
	"context"

	"github.com/cifra-city/location-storage/internal/data/db/sqlcore"
	"github.com/google/uuid"
)

type Cities interface {
	Create(ctx context.Context, name string, location string) (sqlcore.City, error)

	Get(ctx context.Context, id uuid.UUID) (sqlcore.City, error)
	GetByName(ctx context.Context, name string) (sqlcore.City, error)

	UpdateName(ctx context.Context, id uuid.UUID, name string) (sqlcore.City, error)
	UpdateLocation(ctx context.Context, id uuid.UUID, location string) (sqlcore.City, error)

	Delete(ctx context.Context, id uuid.UUID) error
}

type cities struct {
	queries *sqlcore.Queries
}

func NewCities(queries *sqlcore.Queries) Cities {
	return &cities{queries: queries}
}

func (c *cities) Create(ctx context.Context, name string, location string) (sqlcore.City, error) {
	return c.queries.CreateCity(ctx, sqlcore.CreateCityParams{
		ID:       uuid.New(),
		Name:     name,
		Location: location,
	})
}

func (c *cities) Get(ctx context.Context, id uuid.UUID) (sqlcore.City, error) {
	return c.queries.GetCityByID(ctx, id)
}

func (c *cities) GetByName(ctx context.Context, name string) (sqlcore.City, error) {
	return c.queries.GetCityByName(ctx, name)
}

func (c *cities) UpdateName(ctx context.Context, id uuid.UUID, name string) (sqlcore.City, error) {
	return c.queries.UpdateCityName(ctx, sqlcore.UpdateCityNameParams{
		ID:   id,
		Name: name,
	})
}

func (c *cities) UpdateLocation(ctx context.Context, id uuid.UUID, location string) (sqlcore.City, error) {
	return c.queries.UpdateCityLocation(ctx, sqlcore.UpdateCityLocationParams{
		ID:       id,
		Location: location,
	})
}

func (c *cities) Delete(ctx context.Context, id uuid.UUID) error {
	return c.queries.DeleteCity(ctx, id)

}
