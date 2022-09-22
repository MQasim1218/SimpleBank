package db

import (
	"GolangBackend/util"
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createtestTransfer(t *testing.T) Transfer {
	args := CreateTransferParams{
		FromAccount:     rand.Int63n(10),
		ToAccount:       rand.Int63n(10),
		TransactionTime: time.Now(),
		Amount:          util.RandonMoney(),
	}

	tr, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, tr)

	require.NotZero(t, tr.ID)

	require.Equal(t, tr.FromAccount, tr.FromAccount)
	require.Equal(t, tr.ToAccount, tr.ToAccount)
	require.Equal(t, tr.Amount, args.Amount)

	require.NotZero(t, tr.TransactionTime)

	return tr
}

func TestCreateTransfer(t *testing.T) {
	createtestTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	tr := createtestTransfer(t)

	test_tr, err := testQueries.GetTransfer(context.Background(), tr.ID)
	require.NoError(t, err)

	require.Equal(t, test_tr.Amount, tr.Amount)
	require.Equal(t, test_tr.FromAccount, tr.FromAccount)
	require.Equal(t, test_tr.ToAccount, tr.ToAccount)
	require.Equal(t, test_tr.TransactionTime, tr.TransactionTime)
}
