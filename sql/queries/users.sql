-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password, is_chirpy_red)
VALUES (
    gen_random_uuid(), Now(), Now(), $1, $2, false
)
RETURNING *;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, email, hashed_password, is_chirpy_red FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT id, created_at, updated_at, email, hashed_password, is_chirpy_red FROM users WHERE id = $1;

-- name: UpdateUsersByID :one
UPDATE users SET email=$2, hashed_password=$3, updated_at = Now() WHERE id = $1 RETURNING *;

-- name: UpdateUsersSetChirpyRed :one
UPDATE users SET is_chirpy_red=true WHERE id = $1 RETURNING *; 