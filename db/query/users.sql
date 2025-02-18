-- name: createUser :one
INSERT INTO "users" (
    username,
    nickname,
    hashed_password,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: getUser :one
SELECT * FROM "users"
WHERE username=$1;

-- name: deleteUser :exec
DELETE FROM "users"
WHERE id=$1;

-- name: updateUser :one
UPDATE "users"
SET 
    nickname = COALESCE(sqlc.narg(nickname),nickname),
    hashed_password = COALESCE(sqlc.narg(hashed_password),hashed_password),
    email = COALESCE(sqlc.narg(email),email)
WHERE 
    username = @username
RETURNING *;