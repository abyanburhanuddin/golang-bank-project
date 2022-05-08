package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func TestCreateEntry_success(t *testing.T) {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

func createRandomEntry(t *testing.T, acc_id int64) Entry {
	arg := CreateEntryParams{
		AccountID: acc_id,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestGetEntry_success(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account.ID)

	get_entry, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	require.Equal(t, entry.ID, get_entry.ID)
	require.Equal(t, entry.AccountID, get_entry.AccountID)
	require.Equal(t, entry.Amount, get_entry.Amount)
	require.WithinDuration(t, entry.CreatedAt, get_entry.CreatedAt, time.Second)
}

func TestGetListEntry_success(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 30; i++ {
		createRandomEntry(t, account.ID)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     10,
		Offset:    2,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, entries, 10)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
