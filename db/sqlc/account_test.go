package db

import (
	"GolangBackend/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createAcc(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandName(),
		Balance:  util.RandonMoney(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.Owner, args.Owner)
	require.Equal(t, acc.Balance, args.Balance)
	require.Equal(t, acc.Currency, args.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestCreateAccount(t *testing.T) {
	createAcc(t)
}

func TestGetAccount(t *testing.T) {
	acc := createAcc(t)

	acc2, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc.ID, acc2.ID)
	require.Equal(t, acc.Owner, acc2.Owner)
	require.Equal(t, acc.Balance, acc2.Balance)
	require.Equal(t, acc.CreatedAt, acc2.CreatedAt)
	require.Equal(t, acc.Currency, acc2.Currency)
}

func TestAllAccounts(t *testing.T) {
	
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createAcc(t)

	args := UpdateAccountParams{
		ID:       acc1.ID,
		Owner:    acc1.Owner,
		Balance:  util.RandonMoney(),
		Currency: acc1.Currency,
	}

	err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), args.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, acc1.Owner, acc2.Owner)
}
