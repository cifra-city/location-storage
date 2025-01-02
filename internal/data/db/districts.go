package db

import (
	"context"

	"github.com/cifra-city/location-storage/internal/data/db/dbcore"
	"github.com/google/uuid"
)

type Districts interface {
	Create(ctx context.Context, name string, cityId uuid.UUID) (dbcore.District, error)

	Get(ctx context.Context, id uuid.UUID) (dbcore.District, error)
	GetByName(ctx context.Context, name string) (dbcore.District, error)
	GetByCity(ctx context.Context, cityID uuid.UUID) ([]dbcore.GetDistrictsByCityRow, error)

	Update(ctx context.Context, id uuid.UUID, name string) (dbcore.District, error)

	Delete(ctx context.Context, id uuid.UUID) error
}

type district struct {
	queries *dbcore.Queries
}

func NewDistricts(queries *dbcore.Queries) Districts {
	return &district{queries: queries}
}

func (d *district) Create(ctx context.Context, name string, cityId uuid.UUID) (dbcore.District, error) {
	return d.queries.CreateDistrict(ctx, dbcore.CreateDistrictParams{
		Name:   name,
		CityID: cityId,
	})
}

func (d *district) Get(ctx context.Context, id uuid.UUID) (dbcore.District, error) {
	return d.queries.GetDistrictByID(ctx, id)
}

func (d *district) GetByName(ctx context.Context, name string) (dbcore.District, error) {
	return d.queries.GetDistrictByName(ctx, name)
}

func (d *district) GetByCity(ctx context.Context, cityID uuid.UUID) ([]dbcore.GetDistrictsByCityRow, error) {
	return d.queries.GetDistrictsByCity(ctx, cityID)
}

func (d *district) Update(ctx context.Context, id uuid.UUID, name string) (dbcore.District, error) {
	return d.queries.UpdateDistrict(ctx, dbcore.UpdateDistrictParams{
		ID:   id,
		Name: name,
	})
}

func (d *district) Delete(ctx context.Context, id uuid.UUID) error {
	return d.queries.DeleteDistrict(ctx, id)
}
