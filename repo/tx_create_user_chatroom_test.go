package repo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUserToChatroomTx(t *testing.T) {
	t.Run("create user to chatroom tx", func(t *testing.T) {
		user := createRandomUser(t)
		chatroom := createRandomChatroom(t)

		arg := CreateUserToChatroomTxParams{
			Username:     user.Username,
			ChatroomName: chatroom.ChatroomName,
		}
		result, err := testStore.createUserToChatroomTx(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, result)

		require.Equal(t, result.Role, "typer")
		require.NotZero(t, result.CreatedAT)
	})

	t.Run("multiple users to chatroom tx", func(t *testing.T) {
		t.Parallel()
		chatroom := createRandomChatroom(t)
		n := 5
		results := make([]CreateUserToChatroomTxResult, n)
		for i := 0; i < n; i++ {
			go func() {
				user := createRandomUser(t)
				arg := CreateUserToChatroomTxParams{
					Username:     user.Username,
					ChatroomName: chatroom.ChatroomName,
				}

				result, err := testStore.createUserToChatroomTx(context.Background(), arg)
				require.NoError(t, err)

				results = append(results, result)
			}()
		}
		require.Len(t, results, n)

	})
}
