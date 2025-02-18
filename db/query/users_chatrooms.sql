-- name: createUsersToChatrooms :one
INSERT INTO "users_chatrooms" (
    user_id,
    chatroom_id,
    role
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: getChatroomByUsername :one
SELECT * FROM "users_chatrooms" UC
JOIN users U ON UC.user_id = U.id
WHERE
    U.username = $1;

-- name: getUserByChatroomName :one
SELECT * FROM "users_chatrooms" UC
JOIN chatrooms C ON UC.chatroom_id = C.chatroom_id
WHERE
    C.chatroom_name = @chatroom_name;

-- name: listChatroomsByUsername :many
SELECT * FROM "users_chatrooms" UC
JOIN users U ON UC.user_id = U.id
WHERE
    U.username = $1
ORDER BY UC.chatroom_id
LIMIT $2
OFFSET $3;

-- name: listUsersByChatroomName :many
SELECT * FROM "users_chatrooms" UC
JOIN chatrooms C ON UC.chatroom_id = C.chatroom_id
WHERE
    C.chatroom_name = $1
ORDER BY UC.user_id
LIMIT $2
OFFSET $3;

-- name: deleteUsersFromChatrooms :exec
DELETE FROM "users_chatrooms"
WHERE id=$1;

-- name: updateUsersFromChatrooms :one
UPDATE "users_chatrooms" UC
SET 
    role = COALESCE(sqlc.narg(role),role)
FROM "users" U
WHERE 
    UC.user_id = U.id
    AND U.username = @username
RETURNING *;