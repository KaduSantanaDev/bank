package db

import (
	"bank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	sut := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), sut.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, sut.ID, account2.ID)
	require.Equal(t, sut.Owner, account2.Owner)
	require.Equal(t, sut.Balance, account2.Balance)
	require.Equal(t, sut.Currency, account2.Currency)
	require.WithinDuration(t, sut.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	sut := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      sut.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, sut.ID, account2.ID)
	require.Equal(t, sut.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, sut.Currency, account2.Currency)
}
