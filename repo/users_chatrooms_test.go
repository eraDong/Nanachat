package repo

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUserToChatroom(t *testing.T) (UsersChatrooms, Users, Chatrooms) {
	user := createRandomUser(t)
	chatroom := createRandomChatroom(t)
	arg := createUsersToChatroomsParams{
		UserID:     user.ID,
		ChatroomID: chatroom.ID,
	}
	userChatroom, err := testStore.
		createUsersToChatrooms(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, userChatroom)

	require.Equal(t, userChatroom.ChatroomID, arg.ChatroomID)
	require.Equal(t, userChatroom.UserID, arg.UserID)

	require.NotZero(t, userChatroom.Role)
	require.NotZero(t, userChatroom.CreatedAt)
	return userChatroom, user, chatroom
}

func createSpecifyUserToSpecifyChatroom(t *testing.T, user Users, chatroom Chatrooms) UsersChatrooms {
	arg := createUsersToChatroomsParams{
		UserID:     user.ID,
		ChatroomID: chatroom.ID,
	}
	userChatroom, err := testStore.
		createUsersToChatrooms(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, userChatroom)

	require.Equal(t, userChatroom.ChatroomID, arg.ChatroomID)
	require.Equal(t, userChatroom.UserID, arg.UserID)

	require.NotZero(t, userChatroom.Role)
	require.NotZero(t, userChatroom.CreatedAt)
	return userChatroom
}

func TestCreateUserToChatroom(t *testing.T) {
	createRandomUserToChatroom(t)
}

type ChatroomDetails struct {
	ChatroomID   int32
	Description  string
	ChatroomName string
	CreatedAt    time.Time
}

type UserDetails struct {
	UserID    int32
	Username  string
	Nickname  string
	Email     string
	CreatedAt time.Time
}

func TestGetUserChatroom(t *testing.T) {
	userChatroom1, user1, chatroom1 := createRandomUserToChatroom(t)

	t.Run("get users from chatroom", func(t *testing.T) {
		gotUserChatroom, err := testStore.
			getUserByChatroomName(context.Background(), chatroom1.ChatroomName)

		require.NoError(t, err)
		require.NotEmpty(t, gotUserChatroom)

		require.Equal(t, gotUserChatroom.ChatroomUserID, userChatroom1.ID)
		require.Equal(t, gotUserChatroom.ChatroomName, chatroom1.ChatroomName)
		require.Equal(t, gotUserChatroom.Description, chatroom1.Description)
		require.WithinDuration(t, gotUserChatroom.CreatedAt.Time, userChatroom1.CreatedAt.Time, time.Second)

	})

	t.Run("get chatrooms from user", func(t *testing.T) {
		gotChatroomUser, err := testStore.
			getChatroomByUsername(context.Background(), user1.Username)

		require.NoError(t, err)
		require.NotEmpty(t, gotChatroomUser)

		require.Equal(t, gotChatroomUser.ChatroomUserID, userChatroom1.ID)
		require.Equal(t, gotChatroomUser.UserID, user1.ID)
		require.Equal(t, gotChatroomUser.Username, user1.Username)
		require.Equal(t, gotChatroomUser.Nickname, user1.Nickname)
		require.Equal(t, gotChatroomUser.Email, user1.Email)
		require.WithinDuration(t, gotChatroomUser.CreatedAt.Time, user1.CreatedAt.Time, time.Second)
	})

	users := []Users{
		createRandomUser(t),
		createRandomUser(t),
		createRandomUser(t),
	}
	specifyChatroom1 := createRandomChatroom(t)
	specifyChatroom2 := createRandomChatroom(t)
	for _, user := range users {
		createSpecifyUserToSpecifyChatroom(t, user, specifyChatroom1)
	}

	t.Run("list chatrooms by username", func(t *testing.T) {
		createSpecifyUserToSpecifyChatroom(t, users[0], specifyChatroom2)
		arg := listChatroomsByUsernameParams{
			Username: users[0].Username,
			Limit:    10,
			Offset:   0,
		}
		chatrooms, err := testStore.
			listChatroomsByUsername(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, chatrooms)

		expectedChatrooms := map[int32]ChatroomDetails{
			specifyChatroom1.ID: {
				ChatroomID:   specifyChatroom1.ID,
				Description:  specifyChatroom1.Description.String,
				ChatroomName: specifyChatroom1.ChatroomName,
				CreatedAt:    specifyChatroom1.CreatedAt.Time,
			},
			specifyChatroom2.ID: {
				ChatroomID:   specifyChatroom2.ID,
				Description:  specifyChatroom2.Description.String,
				ChatroomName: specifyChatroom2.ChatroomName,
				CreatedAt:    specifyChatroom2.CreatedAt.Time,
			},
		}

		for _, chatroom := range chatrooms {
			require.Contains(t, expectedChatrooms, chatroom.ChatroomID)

			expected := expectedChatrooms[chatroom.ChatroomID]
			require.Equal(t, chatroom.ChatroomID, expected.ChatroomID)
			require.Equal(t, chatroom.ChatroomName, expected.ChatroomName)
			require.Equal(t, chatroom.Description.String, expected.Description)
			require.WithinDuration(t, chatroom.CreatedAt.Time, expected.CreatedAt, time.Second)
		}
	})

	t.Run("list users by chatroom name", func(t *testing.T) {
		arg := listUsersByChatroomNameParams{
			ChatroomName: specifyChatroom1.ChatroomName,
			Limit:        10,
			Offset:       0,
		}
		results, err := testStore.
			listUsersByChatroomName(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, results)

		expectedUsers := map[int32]UserDetails{}

		for _, user := range users {
			expectedUsers[user.ID] = UserDetails{
				UserID:    user.ID,
				Username:  user.Username,
				Nickname:  user.Nickname,
				Email:     user.Email,
				CreatedAt: user.CreatedAt.Time,
			}
		}

		for _, user := range results {
			require.Contains(t, expectedUsers, user.UserID)

			expected := expectedUsers[user.UserID]
			require.Equal(t, user.UserID, expected.UserID)
			require.Equal(t, user.Username, expected.Username)
			require.Equal(t, user.Nickname, expected.Nickname)
			require.Equal(t, user.Email, expected.Email)
			require.WithinDuration(t, user.CreatedAt.Time, expected.CreatedAt, time.Second)
		}
	})

	t.Run("list empty", func(t *testing.T) {
		arg1 := listChatroomsByUsernameParams{
			Username: "",
			Limit:    10,
			Offset:   0,
		}
		chatrooms, err := testStore.
			listChatroomsByUsername(context.Background(), arg1)
		require.NoError(t, err)
		require.Empty(t, chatrooms)

		arg2 := listUsersByChatroomNameParams{
			ChatroomName: "",
			Limit:        10,
			Offset:       0,
		}
		users, err := testStore.
			listUsersByChatroomName(context.Background(), arg2)
		require.NoError(t, err)
		require.Empty(t, users)
	})

}

func TestDeleteUserChatroom(t *testing.T) {
	t.Run("delete user_chatroom", func(t *testing.T) {
		userChatroom, user, chatroom := createRandomUserToChatroom(t)
		err := testStore.
			deleteUsersFromChatrooms(context.Background(), userChatroom.ID)
		require.NoError(t, err)

		result1, err := testStore.
			getChatroomByUsername(context.Background(), user.Username)
		require.Error(t, err)
		require.Empty(t, result1)

		result2, err := testStore.
			getUserByChatroomName(context.Background(), chatroom.ChatroomName)
		require.Error(t, err)
		require.Empty(t, result2)
	})
}

func TestUpdateUserChatroom(t *testing.T) {
	t.Run("update role of users", func(t *testing.T) {
		_, user, _ := createRandomUserToChatroom(t)
		arg := updateUsersFromChatroomsParams{
			Username: user.Username,
			Role: pgtype.Text{
				String: "admin",
				Valid:  true,
			},
		}
		result, err := testStore.
			updateUsersFromChatrooms(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, result)

		require.Equal(t, result.Role, arg.Role.String)
	})
}
