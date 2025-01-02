package db

import (
	"context"

	"github.com/cifra-city/location-storage/internal/data/db/dbcore"
	"github.com/google/uuid"
)

type Streets interface {
	Create(ctx context.Context, name string, districtId uuid.UUID) (dbcore.Street, error)

	Get(ctx context.Context, id uuid.UUID) (dbcore.Street, error)
	GetByName(ctx context.Context, name string) (dbcore.Street, error)
	GetByDistrict(ctx context.Context, districtID uuid.UUID) ([]dbcore.Street, error)

	UpdateName(ctx context.Context, id uuid.UUID, name string) (dbcore.Street, error)
	UpdateDistrict(ctx context.Context, streetId uuid.UUID, districtId uuid.UUID) (dbcore.Street, error)

	Delete(ctx context.Context, id uuid.UUID) error
}

type streets struct {
	queries *dbcore.Queries
}

func NewStreets(queries *dbcore.Queries) Streets {
	return &streets{queries: queries}
}

func (s *streets) Create(ctx context.Context, name string, districtId uuid.UUID) (dbcore.Street, error) {
	return s.queries.CreateStreet(ctx, dbcore.CreateStreetParams{
		Name:       name,
		DistrictID: districtId,
	})
}

func (s *streets) Get(ctx context.Context, id uuid.UUID) (dbcore.Street, error) {
	return s.queries.GetStreetByID(ctx, id)
}

func (s *streets) GetByName(ctx context.Context, name string) (dbcore.Street, error) {
	return s.queries.GetStreetByName(ctx, name)
}

func (s *streets) GetByDistrict(ctx context.Context, districtID uuid.UUID) ([]dbcore.Street, error) {
	return s.queries.GetStreetsByDistrict(ctx, districtID)
}

func (s *streets) UpdateName(ctx context.Context, id uuid.UUID, name string) (dbcore.Street, error) {
	return s.queries.UpdateStreetName(ctx, dbcore.UpdateStreetNameParams{
		ID:   id,
		Name: name,
	})
}

func (s *streets) UpdateDistrict(ctx context.Context, streetId uuid.UUID, districtId uuid.UUID) (dbcore.Street, error) {
	return s.queries.UpdateStreetDistrict(ctx, dbcore.UpdateStreetDistrictParams{
		ID:         streetId,
		DistrictID: districtId,
	})
}

func (s *streets) Delete(ctx context.Context, id uuid.UUID) error {
	return s.queries.DeleteStreet(ctx, id)
}
