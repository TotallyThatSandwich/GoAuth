-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password, user_token
) VALUES ( $1, $2, uuid_generate_v4() )
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1
	AND user_token = $2;

-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
  AND hashed_password = $2;

-- name: GetUserFromID :one
SELECT *
FROM users
WHERE user_id = $1;

-- name: GetUserToken :one
SELECT user_token
FROM users
WHERE user_id = $1;

-- name: CheckUserToken :one
SELECT *
FROM users
WHERE user_token = $1;

-- name: UpdateUser :one
UPDATE users
	SET username = $1,
    hashed_password = $2
WHERE user_id = $3
RETURNING *;

-- name: ResetUserToken :one
UPDATE users
SET user_token = uuid_generate_v4()
WHERE user_id = $1
RETURNING *;


