// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repo

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Chatrooms struct {
	ID           int32              `json:"id"`
	ChatroomName string             `json:"chatroom_name"`
	Description  pgtype.Text        `json:"description"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
}

type Messages struct {
	ID             int32              `json:"id"`
	UserChatroomID int32              `json:"user_chatroom_id"`
	Text           string             `json:"text"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
}

type Users struct {
	ID             int32              `json:"id"`
	Username       string             `json:"username"`
	Nickname       string             `json:"nickname"`
	HashedPassword string             `json:"hashed_password"`
	Email          string             `json:"email"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
}

type UsersChatrooms struct {
	ID         int32              `json:"id"`
	UserID     int32              `json:"user_id"`
	ChatroomID int32              `json:"chatroom_id"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	Role       string             `json:"role"`
}
