package repo

import (
	"context"
	"testing"
	"time"

	"github.com/eraDong/NanaChat/util"
	"github.com/stretchr/testify/require"
)

func createRandomMessage(t *testing.T) (Messages, Users, Chatrooms) {
	userChatroom, user, chatroom := createRandomUserToChatroom(t)
	arg := createMessageParams{
		UserChatroomID: userChatroom.ID,
		Text:           util.RandomString(10),
	}
	message, err := testStore.createMessage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, message)
	require.NotZero(t, message.ID)
	require.NotZero(t, message.CreatedAt)

	require.Equal(t, message.UserChatroomID, arg.UserChatroomID)
	require.Equal(t, message.Text, arg.Text)
	return message, user, chatroom
}

func createRandomMessageToSpecifyChatroom(t *testing.T, chatroom Chatrooms) (Messages, Users) {
	user := createRandomUser(t)
	userChatroom := createSpecifyUserToSpecifyChatroom(t, user, chatroom)
	arg := createMessageParams{
		UserChatroomID: userChatroom.ID,
		Text:           util.RandomString(10),
	}
	message, err := testStore.createMessage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, message)
	require.NotZero(t, message.ID)
	require.NotZero(t, message.CreatedAt)

	require.Equal(t, message.UserChatroomID, arg.UserChatroomID)
	require.Equal(t, message.Text, arg.Text)
	return message, user
}

func TestCreateMessage(t *testing.T) {
	createRandomMessage(t)
}

func TestListMessages(t *testing.T) {
	chatroom := createRandomChatroom(t)
	message1, user1 := createRandomMessageToSpecifyChatroom(t, chatroom)
	message2, user2 := createRandomMessageToSpecifyChatroom(t, chatroom)

	t.Run("list message by chatroom_name", func(t *testing.T) {
		arg := listMessagesByChatroomNameParams{
			ChatroomName: chatroom.ChatroomName,
			Limit:        10,
			Offset:       0,
		}
		messages, err := testStore.
			listMessagesByChatroomName(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, messages)
		require.True(t, messages[0].CreatedAt.Time.After(messages[1].CreatedAt.Time))

		require.Equal(t, chatroom.ChatroomName, messages[1].ChatroomName)
		require.Equal(t, user1.Username, messages[1].SenderName)
		require.Equal(t, message1.ID, messages[1].MessageID)
		require.Equal(t, message1.Text, messages[1].Text)
		require.WithinDuration(t, message1.CreatedAt.Time, messages[1].CreatedAt.Time, time.Second)

		require.Equal(t, chatroom.ChatroomName, messages[0].ChatroomName)
		require.Equal(t, user2.Username, messages[0].SenderName)
		require.Equal(t, message2.ID, messages[0].MessageID)
		require.Equal(t, message2.Text, messages[0].Text)
		require.WithinDuration(t, message2.CreatedAt.Time, messages[0].CreatedAt.Time, time.Second)
	})

	t.Run("list message by username", func(t *testing.T) {
		arg := listMessagesByUsernameParams{
			Username: user1.Username,
			Limit:    10,
			Offset:   0,
		}
		messages, err := testStore.
			listMessagesByUsername(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, messages)

		require.Equal(t, chatroom.ChatroomName, messages[0].ChatroomName)
		require.Equal(t, user1.Username, messages[0].SenderName)
		require.Equal(t, message1.ID, messages[0].MessageID)
		require.Equal(t, message1.Text, messages[0].Text)
		require.WithinDuration(t, message1.CreatedAt.Time, messages[0].CreatedAt.Time, time.Second)
	})
}

func TestDeleteMessage(t *testing.T) {
	t.Run("delete exist message", func(t *testing.T) {
		message, user, _ := createRandomMessage(t)
		err := testStore.deleteMessage(context.Background(), message.ID)
		require.NoError(t, err)

		arg := listMessagesByUsernameParams{
			Username: user.Username,
			Limit:    10,
			Offset:   0,
		}
		result, err := testStore.
			listMessagesByUsername(context.Background(), arg)
		require.NoError(t, err)
		require.Empty(t, result)
	})
}
