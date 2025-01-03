-- name: CreateStreet :one
INSERT INTO streets (id, name, city_id, location)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateStreet :one
UPDATE streets
SET name = $2, location = $3
WHERE id = $1
RETURNING *;

-- name: UpdateStreetLocation :one
UPDATE streets
SET location = $2
WHERE id = $1
RETURNING *;

-- name: UpdateStreetName :one
UPDATE streets
SET name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateStreetCity :one
UPDATE streets
SET city_id = $2
WHERE id = $1
RETURNING *;

-- name: DeleteStreet :exec
DELETE FROM streets
WHERE id = $1;

-- name: GetStreetByID :one
SELECT * FROM streets
WHERE id = $1;

-- name: GetStreetByName :one
SELECT * FROM streets
WHERE name = $1;

-- name: ListStreetsByCity :many
SELECT *
FROM streets
WHERE city_id = $1
ORDER BY name;
