package db

import (
	"context"

	"github.com/cifra-city/location-storage/internal/data/db/sqlcore"
	"github.com/google/uuid"
)

type Streets interface {
	Create(ctx context.Context, name string, districtId uuid.UUID, location string) (sqlcore.Street, error)

	Get(ctx context.Context, id uuid.UUID) (sqlcore.Street, error)
	GetByName(ctx context.Context, name string) (sqlcore.Street, error)
	ListStreetsByCity(ctx context.Context, cityId uuid.UUID) ([]sqlcore.Street, error)

	UpdateName(ctx context.Context, id uuid.UUID, name string) (sqlcore.Street, error)
	UpdateCity(ctx context.Context, id uuid.UUID, cityId uuid.UUID) (sqlcore.Street, error)
	UpdateLocation(ctx context.Context, id uuid.UUID, location string) (sqlcore.Street, error)

	Delete(ctx context.Context, id uuid.UUID) error
}

type streets struct {
	queries *sqlcore.Queries
}

func NewStreets(queries *sqlcore.Queries) Streets {
	return &streets{queries: queries}
}

func (s *streets) Create(ctx context.Context, name string, cityId uuid.UUID, location string) (sqlcore.Street, error) {
	return s.queries.CreateStreet(ctx, sqlcore.CreateStreetParams{
		Name:     name,
		CityID:   cityId,
		Location: location,
	})
}

func (s *streets) Get(ctx context.Context, id uuid.UUID) (sqlcore.Street, error) {
	return s.queries.GetStreetByID(ctx, id)
}

func (s *streets) GetByName(ctx context.Context, name string) (sqlcore.Street, error) {
	return s.queries.GetStreetByName(ctx, name)
}

func (s *streets) ListStreetsByCity(ctx context.Context, cityId uuid.UUID) ([]sqlcore.Street, error) {
	return s.queries.ListStreetsByCity(ctx, cityId)
}

func (s *streets) UpdateName(ctx context.Context, id uuid.UUID, name string) (sqlcore.Street, error) {
	return s.queries.UpdateStreetName(ctx, sqlcore.UpdateStreetNameParams{
		ID:   id,
		Name: name,
	})
}

func (s *streets) UpdateCity(ctx context.Context, id uuid.UUID, cityId uuid.UUID) (sqlcore.Street, error) {
	return s.queries.UpdateStreetCity(ctx, sqlcore.UpdateStreetCityParams{
		ID:     id,
		CityID: cityId,
	})
}

func (s *streets) UpdateLocation(ctx context.Context, id uuid.UUID, location string) (sqlcore.Street, error) {
	return s.queries.UpdateStreetLocation(ctx, sqlcore.UpdateStreetLocationParams{
		ID:       id,
		Location: location,
	})
}

func (s *streets) Delete(ctx context.Context, id uuid.UUID) error {
	return s.queries.DeleteStreet(ctx, id)
}
