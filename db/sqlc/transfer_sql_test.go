package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func TestCreateTransfer_success(t *testing.T) {
	acc_1 := createRandomAccount(t)
	acc_2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: acc_1.ID,
		ToAccountID:   acc_2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.CreatedAt)
}

func createRandomTransfer(t *testing.T, from_acc int64, to_acc int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from_acc,
		ToAccountID:   to_acc,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.CreatedAt)

	return transfer
}

func TestGetTransfer_success(t *testing.T) {
	acc_1 := createRandomAccount(t)
	acc_2 := createRandomAccount(t)

	transfer := createRandomTransfer(t, acc_1.ID, acc_2.ID)

	get_transfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)

	require.Equal(t, transfer.ID, get_transfer.ID)
	require.Equal(t, transfer.FromAccountID, get_transfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, get_transfer.ToAccountID)
	require.Equal(t, transfer.Amount, get_transfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, get_transfer.CreatedAt, time.Second)
}

func TestGetListTransfer_success(t *testing.T) {
	acc_1 := createRandomAccount(t)
	acc_2 := createRandomAccount(t)

	for i := 0; i < 30; i++ {
		createRandomTransfer(t, acc_1.ID, acc_2.ID)
	}

	arg := ListTransfersParams{
		FromAccountID: acc_1.ID,
		ToAccountID:   acc_2.ID,
		Limit:         10,
		Offset:        1,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, transfers, 10)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
