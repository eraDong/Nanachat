package repo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/eraDong/NanaChat/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) Users {
	arg := createUserParams{
		Username:       util.RandomString(6),
		Nickname:       util.RandomString(6),
		HashedPassword: "123",
		Email:          randomEmail(),
	}
	user, err := testStore.createUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEmpty(t, user.ID)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)
	require.Equal(t, user.Nickname, arg.Nickname)
	require.Equal(t, user.Email, arg.Email)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)
	newNickname := util.RandomString(6)
	newEmail := randomEmail()
	arg := updateUserParams{
		Username: user.Username,
		Nickname: pgtype.Text{
			String: newNickname,
			Valid:  true,
		},
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	}
	user, err := testStore.updateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Email, newEmail)
	require.Equal(t, user.Nickname, newNickname)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	gotUser, err := testStore.getUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, gotUser)

	require.Equal(t, user.ID, gotUser.ID)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.Nickname, gotUser.Nickname)
	require.Equal(t, user.Email, gotUser.Email)
	require.Equal(t, user.HashedPassword, gotUser.HashedPassword)
	require.WithinDuration(t, user.CreatedAt.Time, gotUser.CreatedAt.Time, time.Second)

}

func randomEmail() string {
	email := fmt.Sprintf("%s@%s.com", util.RandomString(7), util.RandomString(3))
	return email
}
