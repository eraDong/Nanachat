package repo

import (
	"context"
	"testing"
	"time"

	"github.com/eraDong/NanaChat/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomChatroom(t *testing.T) Chatrooms {
	arg := createChatroomsParams{
		ChatroomName: util.RandomString(6),
		Description: pgtype.Text{
			String: util.RandomString(32),
			Valid:  true,
		},
	}
	chatroom, err := testStore.createChatrooms(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, chatroom)

	require.NotEmpty(t, chatroom.ID)
	require.Equal(t, chatroom.ChatroomName, arg.ChatroomName)
	require.Equal(t, chatroom.Description, arg.Description)

	require.NotZero(t, chatroom.ID)
	require.NotZero(t, chatroom.CreatedAt)
	return chatroom
}

func TestCreateChatroom(t *testing.T) {
	createRandomChatroom(t)
}

func TestGetChatroom(t *testing.T) {
	chatroom := createRandomChatroom(t)

	t.Run("get chatroom", func(t *testing.T) {
		gotChatroom, err := testStore.getChatrooms(context.Background(), chatroom.ChatroomName)
		require.NoError(t, err)
		require.NotEmpty(t, gotChatroom)

		require.Equal(t, chatroom.ID, gotChatroom.ID)
		require.Equal(t, chatroom.ChatroomName, gotChatroom.ChatroomName)
		require.Equal(t, chatroom.Description, gotChatroom.Description)
		require.WithinDuration(t, chatroom.CreatedAt.Time, gotChatroom.CreatedAt.Time, time.Second)
	})

	t.Run("get a empty chatroom", func(t *testing.T) {
		notExistName := util.RandomString(6)
		gotChatroom, err := testStore.getChatrooms(context.Background(), notExistName)
		require.EqualError(t, err, pgx.ErrNoRows.Error())
		require.Empty(t, gotChatroom)
	})
}

func TestDeleteChatroom(t *testing.T) {
	t.Run("delete a chatroom", func(t *testing.T) {
		chatroom := createRandomChatroom(t)
		err := testStore.deleteChatrooms(context.Background(), chatroom.ID)
		require.NoError(t, err)

		chatroom, err = testStore.getChatrooms(context.Background(), chatroom.ChatroomName)
		require.Error(t, err)
		require.Empty(t, chatroom)
	})
}

func TestUpdateChatroom(t *testing.T) {
	t.Run("update description", func(t *testing.T) {
		chatroom := createRandomChatroom(t)
		newDescription := util.RandomString(32)
		arg := updateChatroomsParams{
			ChatroomName: chatroom.ChatroomName,
			Description: pgtype.Text{
				String: newDescription,
				Valid:  true,
			},
		}
		chatroom, err := testStore.updateChatrooms(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, chatroom)

		require.Equal(t, chatroom.Description.String, newDescription)
	})
}
