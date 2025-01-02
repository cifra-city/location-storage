-- name: CreateDistrict :exec
INSERT INTO districts (name, city_id) VALUES ($1, $2);

-- name: GetDistrictByID :one
SELECT id, name, city_id FROM districts WHERE id = $1;

-- name: GetDistrictsByCity :many
SELECT id, name FROM districts WHERE city_id = $1;

-- name: DeleteDistrict :exec
DELETE FROM districts WHERE id = $1;

-- name: UpdateDistrict :exec
UPDATE districts SET name = $2, city_id = $3 WHERE id = $1;
