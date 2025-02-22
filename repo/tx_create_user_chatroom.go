package repo

import (
	"context"
	"time"
)

type CreateUserToChatroomTxParams struct {
	Username     string
	ChatroomName string
}

type CreateUserToChatroomTxResult struct {
	Role      string
	CreatedAT time.Time
}

func (store *Store) createUserToChatroomTx(ctx context.Context, arg CreateUserToChatroomTxParams) (CreateUserToChatroomTxResult, error) {
	var result CreateUserToChatroomTxResult

	err := store.execTx(ctx, store.connPool, func(q *Queries) error {
		user, err := store.getUser(ctx, arg.Username)
		if err != nil {
			return err
		}

		chatroom, err := store.getChatrooms(ctx, arg.ChatroomName)
		if err != nil {
			return err
		}

		userChatroom, err := store.createUsersToChatrooms(ctx, createUsersToChatroomsParams{
			UserID:     user.ID,
			ChatroomID: chatroom.ID,
		})
		if err != nil {
			return err
		}
		result.Role = userChatroom.Role
		result.CreatedAT = user.CreatedAt.Time

		return nil
	})

	if err != nil {
		return result, err
	}

	return result, nil
}
