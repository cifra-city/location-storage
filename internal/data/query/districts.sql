-- name: CreateDistrict :one
INSERT INTO districts (
   name,
   city_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetDistrictByID :one
SELECT id, name, city_id FROM districts WHERE id = $1;

-- name: GetDistrictsByCity :many
SELECT id, name FROM districts WHERE city_id = $1;

-- name: GetDistrictByName :one
SELECT * FROM districts WHERE name = $1;

-- name: DeleteDistrict :exec
DELETE FROM districts WHERE id = $1;

-- name: UpdateDistrictName :one
UPDATE districts
SET
    name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateDistrictCity :one
UPDATE districts
SET
    city_id = $2
WHERE id = $1
RETURNING *;


