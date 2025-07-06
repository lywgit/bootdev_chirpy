-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(), Now(), Now(), $1, $2
)
RETURNING *;

-- name: GetChirps :many
SELECT id, created_at, updated_at, body, user_Id FROM chirps c  
ORDER BY created_at ASC;

-- name: GetChirpByID :one
SELECT id, created_at, updated_at, body, user_Id FROM chirps c where id = $1;