-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
    $1, Now(), Now(), $2, Now() + INTERVAL '60 Day', null
)
RETURNING *;

-- name: GetUserFromRefreshToken :one
SELECT user_id FROM refresh_tokens WHERE token=$1 AND revoked_at IS NULL and expires_at > Now();

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens SET updated_at = Now(), revoked_at = Now()
where token=$1;