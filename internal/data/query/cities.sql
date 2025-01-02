-- name: CreateCity :one
INSERT INTO cities (
    name,
    country_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetCityByID :one
SELECT * FROM cities WHERE id = $1;

-- name: GetCitiesByCountry :many
SELECT * FROM cities WHERE country_id = $1;

-- name: GetCityByName :one
SELECT * FROM cities WHERE name = $1;

-- name: DeleteCity :exec
DELETE FROM cities WHERE id = $1;

-- name: UpdateCity :one
UPDATE cities SET
    name = $2
WHERE id = $1
RETURNING *;
