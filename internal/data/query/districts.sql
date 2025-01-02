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

-- name: UpdateDistrict :one
UPDATE districts
SET
    name = $2,
    city_id = $3
WHERE id = $1
RETURNING *;

