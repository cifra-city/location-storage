-- name: CreateCity :one
INSERT INTO cities (id, name, location)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateCity :one
UPDATE cities
SET name = $2, location = $3
WHERE id = $1
RETURNING *;

-- name: UpdateCityLocation :one
UPDATE cities
SET location = $2
WHERE id = $1
RETURNING *;

-- name: UpdateCityName :one
UPDATE cities
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCity :exec
DELETE FROM cities
WHERE id = $1;

-- name: GetCityByID :one
SELECT * FROM cities
WHERE id = $1;

-- name: GetCityByName :one
SELECT * FROM cities
WHERE name = $1;

-- name: ListCities :many
SELECT * FROM cities
ORDER BY name;
