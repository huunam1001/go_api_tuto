
-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 AND hashed_password = $2;