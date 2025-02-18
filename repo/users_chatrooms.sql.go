// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users_chatrooms.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUsersToChatrooms = `-- name: createUsersToChatrooms :one
INSERT INTO "users_chatrooms" (
    user_id,
    chatroom_id,
    role
) VALUES (
    $1, $2, $3
) RETURNING id, user_id, chatroom_id, created_at, role
`

type createUsersToChatroomsParams struct {
	UserID     int32  `json:"user_id"`
	ChatroomID int32  `json:"chatroom_id"`
	Role       string `json:"role"`
}

func (q *Queries) createUsersToChatrooms(ctx context.Context, arg createUsersToChatroomsParams) (UsersChatrooms, error) {
	row := q.db.QueryRow(ctx, createUsersToChatrooms, arg.UserID, arg.ChatroomID, arg.Role)
	var i UsersChatrooms
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ChatroomID,
		&i.CreatedAt,
		&i.Role,
	)
	return i, err
}

const deleteUsersFromChatrooms = `-- name: deleteUsersFromChatrooms :exec
DELETE FROM "users_chatrooms"
WHERE id=$1
`

func (q *Queries) deleteUsersFromChatrooms(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUsersFromChatrooms, id)
	return err
}

const getChatroomByUsername = `-- name: getChatroomByUsername :one
SELECT uc.id, user_id, chatroom_id, uc.created_at, role, u.id, username, nickname, hashed_password, email, u.created_at FROM "users_chatrooms" UC
JOIN users U ON UC.user_id = U.id
WHERE
    U.username = $1
`

type getChatroomByUsernameRow struct {
	ID             int32              `json:"id"`
	UserID         int32              `json:"user_id"`
	ChatroomID     int32              `json:"chatroom_id"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	Role           string             `json:"role"`
	ID_2           int32              `json:"id_2"`
	Username       string             `json:"username"`
	Nickname       string             `json:"nickname"`
	HashedPassword string             `json:"hashed_password"`
	Email          string             `json:"email"`
	CreatedAt_2    pgtype.Timestamptz `json:"created_at_2"`
}

func (q *Queries) getChatroomByUsername(ctx context.Context, username string) (getChatroomByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getChatroomByUsername, username)
	var i getChatroomByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ChatroomID,
		&i.CreatedAt,
		&i.Role,
		&i.ID_2,
		&i.Username,
		&i.Nickname,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt_2,
	)
	return i, err
}

const getUserByChatroomName = `-- name: getUserByChatroomName :one
SELECT uc.id, user_id, chatroom_id, uc.created_at, role, c.id, chatroom_name, description, c.created_at FROM "users_chatrooms" UC
JOIN chatrooms C ON UC.chatroom_id = C.chatroom_id
WHERE
    C.chatroom_name = $1
`

type getUserByChatroomNameRow struct {
	ID           int32              `json:"id"`
	UserID       int32              `json:"user_id"`
	ChatroomID   int32              `json:"chatroom_id"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	Role         string             `json:"role"`
	ID_2         int32              `json:"id_2"`
	ChatroomName string             `json:"chatroom_name"`
	Description  pgtype.Text        `json:"description"`
	CreatedAt_2  pgtype.Timestamptz `json:"created_at_2"`
}

func (q *Queries) getUserByChatroomName(ctx context.Context, chatroomName string) (getUserByChatroomNameRow, error) {
	row := q.db.QueryRow(ctx, getUserByChatroomName, chatroomName)
	var i getUserByChatroomNameRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ChatroomID,
		&i.CreatedAt,
		&i.Role,
		&i.ID_2,
		&i.ChatroomName,
		&i.Description,
		&i.CreatedAt_2,
	)
	return i, err
}

const listChatroomsByUsername = `-- name: listChatroomsByUsername :many
SELECT uc.id, user_id, chatroom_id, uc.created_at, role, u.id, username, nickname, hashed_password, email, u.created_at FROM "users_chatrooms" UC
JOIN users U ON UC.user_id = U.id
WHERE
    U.username = $1
ORDER BY UC.chatroom_id
LIMIT $2
OFFSET $3
`

type listChatroomsByUsernameParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

type listChatroomsByUsernameRow struct {
	ID             int32              `json:"id"`
	UserID         int32              `json:"user_id"`
	ChatroomID     int32              `json:"chatroom_id"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	Role           string             `json:"role"`
	ID_2           int32              `json:"id_2"`
	Username       string             `json:"username"`
	Nickname       string             `json:"nickname"`
	HashedPassword string             `json:"hashed_password"`
	Email          string             `json:"email"`
	CreatedAt_2    pgtype.Timestamptz `json:"created_at_2"`
}

func (q *Queries) listChatroomsByUsername(ctx context.Context, arg listChatroomsByUsernameParams) ([]listChatroomsByUsernameRow, error) {
	rows, err := q.db.Query(ctx, listChatroomsByUsername, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []listChatroomsByUsernameRow{}
	for rows.Next() {
		var i listChatroomsByUsernameRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ChatroomID,
			&i.CreatedAt,
			&i.Role,
			&i.ID_2,
			&i.Username,
			&i.Nickname,
			&i.HashedPassword,
			&i.Email,
			&i.CreatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersByChatroomName = `-- name: listUsersByChatroomName :many
SELECT uc.id, user_id, chatroom_id, uc.created_at, role, c.id, chatroom_name, description, c.created_at FROM "users_chatrooms" UC
JOIN chatrooms C ON UC.chatroom_id = C.chatroom_id
WHERE
    C.chatroom_name = $1
ORDER BY UC.user_id
LIMIT $2
OFFSET $3
`

type listUsersByChatroomNameParams struct {
	ChatroomName string `json:"chatroom_name"`
	Limit        int32  `json:"limit"`
	Offset       int32  `json:"offset"`
}

type listUsersByChatroomNameRow struct {
	ID           int32              `json:"id"`
	UserID       int32              `json:"user_id"`
	ChatroomID   int32              `json:"chatroom_id"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	Role         string             `json:"role"`
	ID_2         int32              `json:"id_2"`
	ChatroomName string             `json:"chatroom_name"`
	Description  pgtype.Text        `json:"description"`
	CreatedAt_2  pgtype.Timestamptz `json:"created_at_2"`
}

func (q *Queries) listUsersByChatroomName(ctx context.Context, arg listUsersByChatroomNameParams) ([]listUsersByChatroomNameRow, error) {
	rows, err := q.db.Query(ctx, listUsersByChatroomName, arg.ChatroomName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []listUsersByChatroomNameRow{}
	for rows.Next() {
		var i listUsersByChatroomNameRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ChatroomID,
			&i.CreatedAt,
			&i.Role,
			&i.ID_2,
			&i.ChatroomName,
			&i.Description,
			&i.CreatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUsersFromChatrooms = `-- name: updateUsersFromChatrooms :one
UPDATE "users_chatrooms" UC
SET 
    role = COALESCE($1,role)
FROM "users" U
WHERE 
    UC.user_id = U.id
    AND U.username = $2
RETURNING u.id, username, nickname, hashed_password, email, u.created_at, uc.id, user_id, chatroom_id, uc.created_at, role
`

type updateUsersFromChatroomsParams struct {
	Role     pgtype.Text `json:"role"`
	Username string      `json:"username"`
}

type updateUsersFromChatroomsRow struct {
	ID             int32              `json:"id"`
	Username       string             `json:"username"`
	Nickname       string             `json:"nickname"`
	HashedPassword string             `json:"hashed_password"`
	Email          string             `json:"email"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	ID_2           int32              `json:"id_2"`
	UserID         int32              `json:"user_id"`
	ChatroomID     int32              `json:"chatroom_id"`
	CreatedAt_2    pgtype.Timestamptz `json:"created_at_2"`
	Role           string             `json:"role"`
}

func (q *Queries) updateUsersFromChatrooms(ctx context.Context, arg updateUsersFromChatroomsParams) (updateUsersFromChatroomsRow, error) {
	row := q.db.QueryRow(ctx, updateUsersFromChatrooms, arg.Role, arg.Username)
	var i updateUsersFromChatroomsRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Nickname,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.ID_2,
		&i.UserID,
		&i.ChatroomID,
		&i.CreatedAt_2,
		&i.Role,
	)
	return i, err
}
