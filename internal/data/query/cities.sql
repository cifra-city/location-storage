-- name: CreateCity :exec
INSERT INTO cities (name, country_id) VALUES ($1, $2);

-- name: GetCityByID :one
SELECT id, name, country_id FROM cities WHERE id = $1;

-- name: GetCitiesByCountry :many
SELECT id, name FROM cities WHERE country_id = $1;

-- name: DeleteCity :exec
DELETE FROM cities WHERE id = $1;

-- name: UpdateCity :exec
UPDATE cities SET name = $2, country_id = $3 WHERE id = $1;
