-- name: createMessage :one
INSERT INTO "messages" (
    user_chatroom_id,
    text
) VALUES (
    $1, $2
) RETURNING *;

-- name: deleteMessage :exec
DELETE FROM "messages"
WHERE id=$1;

-- name: listMessages :many
SELECT 
    M.id AS message_id,
    M.text,
    M.created_at,
    U.username AS sender_name,
    C.chatroom_name
FROM "messages" M
JOIN "users_chatrooms" UC ON M.user_chatroom_id = UC.id
JOIN "users" U ON UC.user_id = U.id
JOIN "chatrooms" C ON UC.chatroom_id = C.id
WHERE C.id = $1
ORDER BY M.created_at DESC
LIMIT $2
OFFSET $3;