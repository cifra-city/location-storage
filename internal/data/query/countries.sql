-- name: CreateCountry :exec
INSERT INTO countries (name) VALUES ($1);

-- name: GetCountryByID :one
SELECT id, name FROM countries WHERE id = $1;

-- name: GetAllCountries :many
SELECT id, name FROM countries;

-- name: DeleteCountry :exec
DELETE FROM countries WHERE id = $1;

-- name: UpdateCountry :exec
UPDATE countries SET name = $2 WHERE id = $1;
