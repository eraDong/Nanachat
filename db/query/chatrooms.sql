-- name: createChatrooms :one
INSERT INTO "chatrooms" (
    chatroom_name,
    description
) VALUES (
    $1, $2
) RETURNING *;

-- name: getChatrooms :one
SELECT * FROM "chatrooms"
WHERE chatroom_name=@chatroom_name;

-- name: deleteChatrooms :exec
DELETE FROM "chatrooms"
WHERE id=$1;

-- name: updateChatrooms :one
UPDATE "chatrooms"
SET 
    description = COALESCE(sqlc.narg(description),description)
WHERE 
    chatroom_name = @chatroom_name
RETURNING *;