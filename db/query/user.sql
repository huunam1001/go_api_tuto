
-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 AND hashed_password = $2;

-- name: GetLoginUser :one
SELECT * FROM users
WHERE (username = $1 OR email = $2) AND hashed_password = $3;

-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetListUserWithAccountOrEmail :many
SELECT * FROM users
WHERE username = $1 OR  email = $2;
