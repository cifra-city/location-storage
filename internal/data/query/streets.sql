-- name: CreateStreet :one
INSERT INTO streets (
    name,
    district_id
) VALUES (
    $1, $2
)  RETURNING *;

-- name: GetStreetByID :one
SELECT * FROM streets WHERE id = $1;

-- name: GetStreetsByDistrict :many
SELECT * FROM streets WHERE district_id = $1;

-- name: GetStreetByName :one
SELECT * FROM streets WHERE name = $1;

-- name: DeleteStreet :exec
DELETE FROM streets WHERE id = $1;

-- name: UpdateStreetName :one
UPDATE streets SET
   name = $2
WHERE
   id = $1
RETURNING *;

-- name: UpdateStreetDistrict :one
UPDATE streets SET
   district_id = $2
WHERE
   id = $1
RETURNING *;
