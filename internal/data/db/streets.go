package db

import (
	"context"

	"github.com/cifra-city/location-storage/internal/data/db/dbcore"
	"github.com/google/uuid"
)

type Streets interface {
	Create(ctx context.Context, name string, districtId uuid.UUID, location string) (dbcore.Street, error)

	Get(ctx context.Context, id uuid.UUID) (dbcore.Street, error)
	GetByName(ctx context.Context, name string) (dbcore.Street, error)
	ListStreetsByCity(ctx context.Context, cityId uuid.UUID) ([]dbcore.Street, error)

	UpdateName(ctx context.Context, id uuid.UUID, name string) (dbcore.Street, error)
	UpdateCity(ctx context.Context, id uuid.UUID, cityId uuid.UUID) (dbcore.Street, error)
	UpdateLocation(ctx context.Context, id uuid.UUID, location string) (dbcore.Street, error)

	Delete(ctx context.Context, id uuid.UUID) error
}

type streets struct {
	queries *dbcore.Queries
}

func NewStreets(queries *dbcore.Queries) Streets {
	return &streets{queries: queries}
}

func (s *streets) Create(ctx context.Context, name string, cityId uuid.UUID, location string) (dbcore.Street, error) {
	return s.queries.CreateStreet(ctx, dbcore.CreateStreetParams{
		Name:     name,
		CityID:   cityId,
		Location: location,
	})
}

func (s *streets) Get(ctx context.Context, id uuid.UUID) (dbcore.Street, error) {
	return s.queries.GetStreetByID(ctx, id)
}

func (s *streets) GetByName(ctx context.Context, name string) (dbcore.Street, error) {
	return s.queries.GetStreetByName(ctx, name)
}

func (s *streets) ListStreetsByCity(ctx context.Context, cityId uuid.UUID) ([]dbcore.Street, error) {
	return s.queries.ListStreetsByCity(ctx, cityId)
}

func (s *streets) UpdateName(ctx context.Context, id uuid.UUID, name string) (dbcore.Street, error) {
	return s.queries.UpdateStreetName(ctx, dbcore.UpdateStreetNameParams{
		ID:   id,
		Name: name,
	})
}

func (s *streets) UpdateCity(ctx context.Context, id uuid.UUID, cityId uuid.UUID) (dbcore.Street, error) {
	return s.queries.UpdateStreetCity(ctx, dbcore.UpdateStreetCityParams{
		ID:     id,
		CityID: cityId,
	})
}

func (s *streets) UpdateLocation(ctx context.Context, id uuid.UUID, location string) (dbcore.Street, error) {
	return s.queries.UpdateStreetLocation(ctx, dbcore.UpdateStreetLocationParams{
		ID:       id,
		Location: location,
	})
}

func (s *streets) Delete(ctx context.Context, id uuid.UUID) error {
	return s.queries.DeleteStreet(ctx, id)
}
