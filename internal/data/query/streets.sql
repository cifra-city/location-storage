-- name: CreateStreet :one
INSERT INTO streets (
    name,
    district_id
) VALUES (
    $1, $2
)  RETURNING *;

-- name: GetStreetByID :one
SELECT id, name, district_id FROM streets WHERE id = $1;

-- name: GetStreetsByDistrict :many
SELECT id, name FROM streets WHERE district_id = $1;

-- name: GetStreetByName :one
SELECT * FROM streets WHERE name = $1;

-- name: DeleteStreet :exec
DELETE FROM streets WHERE id = $1;

-- name: UpdateStreet :one
UPDATE streets SET
   name = $2,
   district_id = $3
WHERE
   id = $1
RETURNING *;
