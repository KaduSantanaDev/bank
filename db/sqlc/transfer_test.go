package db

import (
	"bank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accountFrom, accountTo Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Amount:        util.RandomMoney(),
	}

	newTransfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, newTransfer)

	require.Equal(t, arg.FromAccountID, newTransfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, newTransfer.ToAccountID)
	require.Equal(t, arg.Amount, newTransfer.Amount)

	require.NotZero(t, newTransfer.ID)
	require.NotZero(t, newTransfer.CreatedAt)

	return newTransfer
}

func TestCreateTransfer(t *testing.T) {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)
	createRandomTransfer(t, accountFrom, accountTo)
}

func TestGetTransfer(t *testing.T) {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, accountFrom, accountTo)

	retrievedTransfer, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedTransfer)

	require.Equal(t, transfer1.ID, retrievedTransfer.ID)
	require.Equal(t, transfer1.Amount, retrievedTransfer.Amount)
	require.Equal(t, transfer1.FromAccountID, retrievedTransfer.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, retrievedTransfer.ToAccountID)
	require.WithinDuration(t, transfer1.CreatedAt, retrievedTransfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, accountFrom, accountTo)
	}

	args := ListTransfersParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Limit:         5,
		Offset:        5,
	}

	retrievedTransfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, retrievedTransfers, 5)

	for _, retrievedTransfer := range retrievedTransfers {
		require.NotEmpty(t, retrievedTransfer)
		require.True(t, retrievedTransfer.FromAccountID == accountFrom.ID || retrievedTransfer.ToAccountID == accountTo.ID)
	}
}
