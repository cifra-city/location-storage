// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: cities.sql

package dbcore

import (
	"context"

	"github.com/google/uuid"
)

const createCity = `-- name: CreateCity :one
INSERT INTO cities (
    name,
    country_id
) VALUES (
    $1, $2
) RETURNING id, name, country_id
`

type CreateCityParams struct {
	Name      string
	CountryID uuid.UUID
}

func (q *Queries) CreateCity(ctx context.Context, arg CreateCityParams) (City, error) {
	row := q.db.QueryRowContext(ctx, createCity, arg.Name, arg.CountryID)
	var i City
	err := row.Scan(&i.ID, &i.Name, &i.CountryID)
	return i, err
}

const deleteCity = `-- name: DeleteCity :exec
DELETE FROM cities WHERE id = $1
`

func (q *Queries) DeleteCity(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCity, id)
	return err
}

const getCitiesByCountry = `-- name: GetCitiesByCountry :many
SELECT id, name, country_id FROM cities WHERE country_id = $1
`

func (q *Queries) GetCitiesByCountry(ctx context.Context, countryID uuid.UUID) ([]City, error) {
	rows, err := q.db.QueryContext(ctx, getCitiesByCountry, countryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []City
	for rows.Next() {
		var i City
		if err := rows.Scan(&i.ID, &i.Name, &i.CountryID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCityByID = `-- name: GetCityByID :one
SELECT id, name, country_id FROM cities WHERE id = $1
`

func (q *Queries) GetCityByID(ctx context.Context, id uuid.UUID) (City, error) {
	row := q.db.QueryRowContext(ctx, getCityByID, id)
	var i City
	err := row.Scan(&i.ID, &i.Name, &i.CountryID)
	return i, err
}

const getCityByName = `-- name: GetCityByName :one
SELECT id, name, country_id FROM cities WHERE name = $1
`

func (q *Queries) GetCityByName(ctx context.Context, name string) (City, error) {
	row := q.db.QueryRowContext(ctx, getCityByName, name)
	var i City
	err := row.Scan(&i.ID, &i.Name, &i.CountryID)
	return i, err
}

const updateCityCountry = `-- name: UpdateCityCountry :one
UPDATE cities SET
    country_id = $2
WHERE id = $1
RETURNING id, name, country_id
`

type UpdateCityCountryParams struct {
	ID        uuid.UUID
	CountryID uuid.UUID
}

func (q *Queries) UpdateCityCountry(ctx context.Context, arg UpdateCityCountryParams) (City, error) {
	row := q.db.QueryRowContext(ctx, updateCityCountry, arg.ID, arg.CountryID)
	var i City
	err := row.Scan(&i.ID, &i.Name, &i.CountryID)
	return i, err
}

const updateCityName = `-- name: UpdateCityName :one
UPDATE cities SET
    name = $2
WHERE id = $1
RETURNING id, name, country_id
`

type UpdateCityNameParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) UpdateCityName(ctx context.Context, arg UpdateCityNameParams) (City, error) {
	row := q.db.QueryRowContext(ctx, updateCityName, arg.ID, arg.Name)
	var i City
	err := row.Scan(&i.ID, &i.Name, &i.CountryID)
	return i, err
}
