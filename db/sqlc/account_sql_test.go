package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func TestCreateAccount_success(t *testing.T) {
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
}

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

func TestGetAccount_success(t *testing.T) {
	account := createRandomAccount(t)

	get_acc, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, get_acc)

	require.Equal(t, account.ID, get_acc.ID)
	require.Equal(t, account.Owner, get_acc.Owner)
	require.Equal(t, account.Balance, get_acc.Balance)
	require.Equal(t, account.Currency, get_acc.Currency)
	require.WithinDuration(t, account.CreatedAt, get_acc.CreatedAt, time.Second)
}

func TestGetListAccounts_success(t *testing.T) {
	for i := 0; i < 30; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 2,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccount_success(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	update_acc, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, update_acc)

	require.Equal(t, update_acc.Balance, arg.Balance)

	require.Equal(t, account.ID, update_acc.ID)
	require.Equal(t, account.Owner, update_acc.Owner)
	require.Equal(t, account.Currency, update_acc.Currency)
	require.WithinDuration(t, account.CreatedAt, update_acc.CreatedAt, time.Second)
}

func TestDeleteAccount_success(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	get_acc, err := testQueries.GetAccount(context.Background(), account.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, get_acc)
}
