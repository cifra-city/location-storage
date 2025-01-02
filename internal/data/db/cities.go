package db

import (
	"context"

	"github.com/cifra-city/location-storage/internal/data/db/dbcore"
	"github.com/google/uuid"
)

type Cities interface {
	Create(ctx context.Context, name string, countryId uuid.UUID) (dbcore.City, error)

	Get(ctx context.Context, id uuid.UUID) (dbcore.City, error)
	GetByName(ctx context.Context, name string) (dbcore.City, error)
	GetByCountry(ctx context.Context, countryID uuid.UUID) ([]dbcore.City, error)

	UpdateName(ctx context.Context, id uuid.UUID, name string) (dbcore.City, error)
	UpdateCountry(ctx context.Context, CityId uuid.UUID, countryID uuid.UUID) (dbcore.City, error)

	Delete(ctx context.Context, id uuid.UUID) error
}

type cities struct {
	queries *dbcore.Queries
}

func NewCities(queries *dbcore.Queries) Cities {
	return &cities{queries: queries}
}

func (c *cities) Create(ctx context.Context, name string, countryId uuid.UUID) (dbcore.City, error) {
	return c.queries.CreateCity(ctx, dbcore.CreateCityParams{
		Name:      name,
		CountryID: countryId,
	})
}

func (c *cities) Get(ctx context.Context, id uuid.UUID) (dbcore.City, error) {
	return c.queries.GetCityByID(ctx, id)
}

func (c *cities) GetByName(ctx context.Context, name string) (dbcore.City, error) {
	return c.queries.GetCityByName(ctx, name)
}

func (c *cities) GetByCountry(ctx context.Context, countryID uuid.UUID) ([]dbcore.City, error) {
	return c.queries.GetCitiesByCountry(ctx, countryID)
}

func (c *cities) UpdateName(ctx context.Context, id uuid.UUID, name string) (dbcore.City, error) {
	return c.queries.UpdateCityName(ctx, dbcore.UpdateCityNameParams{
		ID:   id,
		Name: name,
	})
}

func (c *cities) UpdateCountry(ctx context.Context, CityId uuid.UUID, countryID uuid.UUID) (dbcore.City, error) {
	return c.queries.UpdateCityCountry(ctx, dbcore.UpdateCityCountryParams{
		ID:        CityId,
		CountryID: countryID,
	})
}

func (c *cities) Delete(ctx context.Context, id uuid.UUID) error {
	return c.queries.DeleteCity(ctx, id)

}
