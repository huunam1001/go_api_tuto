
-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 AND hashed_password = $2;

-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;