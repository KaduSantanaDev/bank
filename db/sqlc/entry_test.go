package db

import (
	"bank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	newEntry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, newEntry)
	require.Equal(t, account.ID, newEntry.AccountID)
	require.Equal(t, arg.Amount, newEntry.Amount)
	require.NotZero(t, newEntry.ID)
	require.NotZero(t, newEntry.CreatedAt)

	return newEntry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)
	retrievedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedEntry)

	require.Equal(t, entry.AccountID, retrievedEntry.AccountID)
	require.Equal(t, entry.Amount, retrievedEntry.Amount)
	require.Equal(t, entry.ID, retrievedEntry.ID)
	require.WithinDuration(t, entry.CreatedAt, retrievedEntry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	retrievedEntries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, retrievedEntries, 5)

	for _, retrievedEntry := range retrievedEntries {
		require.NotEmpty(t, retrievedEntry)
		require.Equal(t, arg.AccountID, retrievedEntry.AccountID)
	}

}
