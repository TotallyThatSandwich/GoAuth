-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password, user_token
) VALUES (
  $1, $2, uuid_generate_v4()
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;

-- name: CheckUserAuth :one
SELECT *
FROM users
WHERE username = $1
  AND hashed_password = $2;

-- name: GetUserToken :one
SELECT user_token
FROM users
WHERE user_id = $1;

-- name: UpdateUser :exec
UPDATE users
	SET username = $1,
    hashed_password = $2
WHERE user_id = $1;

-- name: ResetUserToken :one
UPDATE users
SET user_token = uuid_generate_v4()
WHERE user_id = $1
RETURNING *;


