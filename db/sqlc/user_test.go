package db

import (
	"bank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	sut := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), sut.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, sut.Username, user2.Username)
	require.Equal(t, sut.HashedPassword, user2.HashedPassword)
	require.Equal(t, sut.FullName, user2.FullName)
	require.Equal(t, sut.Email, user2.Email)
	require.WithinDuration(t, sut.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, sut.CreatedAt, user2.CreatedAt, time.Second)
}
