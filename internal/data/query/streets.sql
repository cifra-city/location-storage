-- name: CreateStreet :exec
INSERT INTO streets (name, district_id) VALUES ($1, $2);

-- name: GetStreetByID :one
SELECT id, name, district_id FROM streets WHERE id = $1;

-- name: GetStreetsByDistrict :many
SELECT id, name FROM streets WHERE district_id = $1;

-- name: DeleteStreet :exec
DELETE FROM streets WHERE id = $1;

-- name: UpdateStreet :exec
UPDATE streets SET name = $2, district_id = $3 WHERE id = $1;
