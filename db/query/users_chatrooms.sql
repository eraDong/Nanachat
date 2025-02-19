-- name: createUsersToChatrooms :one
INSERT INTO "users_chatrooms" (
    user_id,
    chatroom_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: getChatroomByUsername :one
SELECT UC.id AS chatroom_user_id, 
    U.id as user_id,
    U.username,
    U.nickname,
    U.email,
    U.created_at
FROM "users_chatrooms" UC
JOIN users U ON UC.user_id = U.id
WHERE
    U.username = $1;

-- name: getUserByChatroomName :one
SELECT UC.id AS chatroom_user_id,
    C.id as chatroom_id,
    C.chatroom_name,
    C.description,
    C.created_at
FROM "users_chatrooms" UC
JOIN chatrooms C ON UC.chatroom_id = C.id
WHERE
    C.chatroom_name = @chatroom_name;

-- name: listChatroomsByUsername :many
SELECT UC.id AS chatroom_user_id,
    C.id as chatroom_id,
    C.chatroom_name,
    C.description,
    C.created_at
FROM "users_chatrooms" UC
JOIN users U ON UC.user_id = U.id
JOIN chatrooms C ON UC.chatroom_id = C.id 
WHERE
    U.username = $1
ORDER BY UC.chatroom_id
LIMIT $2
OFFSET $3;

-- name: listUsersByChatroomName :many
SELECT UC.id AS chatroom_user_id, 
    U.id as user_id,
    U.username,
    U.nickname,
    U.email,
    U.created_at
FROM "users_chatrooms" UC
JOIN chatrooms C ON UC.chatroom_id = C.id
JOIN users U ON UC.user_id = U.id
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