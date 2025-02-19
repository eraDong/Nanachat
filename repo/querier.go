// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repo

import (
	"context"
)

type Querier interface {
	createChatrooms(ctx context.Context, arg createChatroomsParams) (Chatrooms, error)
	createMessage(ctx context.Context, arg createMessageParams) (Messages, error)
	createUser(ctx context.Context, arg createUserParams) (Users, error)
	createUsersToChatrooms(ctx context.Context, arg createUsersToChatroomsParams) (UsersChatrooms, error)
	deleteChatrooms(ctx context.Context, id int32) error
	deleteMessage(ctx context.Context, id int32) error
	deleteUser(ctx context.Context, id int32) error
	deleteUsersFromChatrooms(ctx context.Context, id int32) error
	getChatroomByUsername(ctx context.Context, username string) (getChatroomByUsernameRow, error)
	getChatrooms(ctx context.Context, chatroomName string) (Chatrooms, error)
	getUser(ctx context.Context, username string) (Users, error)
	getUserByChatroomName(ctx context.Context, chatroomName string) (getUserByChatroomNameRow, error)
	listChatroomsByUsername(ctx context.Context, arg listChatroomsByUsernameParams) ([]listChatroomsByUsernameRow, error)
	listMessagesByChatroomName(ctx context.Context, arg listMessagesByChatroomNameParams) ([]listMessagesByChatroomNameRow, error)
	listMessagesByUsername(ctx context.Context, arg listMessagesByUsernameParams) ([]listMessagesByUsernameRow, error)
	listUsersByChatroomName(ctx context.Context, arg listUsersByChatroomNameParams) ([]listUsersByChatroomNameRow, error)
	updateChatrooms(ctx context.Context, arg updateChatroomsParams) (Chatrooms, error)
	updateUser(ctx context.Context, arg updateUserParams) (Users, error)
	updateUsersFromChatrooms(ctx context.Context, arg updateUsersFromChatroomsParams) (updateUsersFromChatroomsRow, error)
}

var _ Querier = (*Queries)(nil)
