-- name: CreateCountry :one
INSERT INTO countries (name)
VALUES ($1) RETURNING *;

-- name: GetCountryByID :one
SELECT * FROM countries WHERE id = $1;

-- name: GetCountryByName :one
SELECT * FROM countries WHERE name = $1;

-- name: GetAllCountries :many
SELECT * FROM countries;

-- name: DeleteCountry :exec
DELETE FROM countries WHERE id = $1;

-- name: UpdateCountry :one
UPDATE countries
SET
    name = $2
WHERE id = $1
RETURNING *;

