package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	acc1 := createAcc(t) // Sender Account
	acc2 := createAcc(t) // Recieving Account
	amount := int64(20)
	n := 3

	// println("Before Transaction ==> From account balance", acc1.Balance, " || To account balance: ", acc2.Balance)

	errs := make(chan error)
	TxRes := make(chan TransferTxRes)

	for i := 0; i < n; i++ {
		// var tr_names = make(map[int]string)
		var tx_name = fmt.Sprintf("TX: %d", i+1)

		go func(errCh chan error, resCh chan TransferTxRes) {
			ctx := context.WithValue(context.Background(), txkey, tx_name)
			Trtx_res, err := store.TranferTx(ctx, TransferTxPrams{
				From_acc_id: acc1.ID,
				To_acc_id:   acc2.ID,
				Amount:      amount,
			})

			errCh <- err
			resCh <- *Trtx_res

		}(errs, TxRes)
	}

	var existed map[int]bool = make(map[int]bool)

	for i := 0; i < n; i++ {
		// Read results and errors from the channel
		err := <-errs
		res := <-TxRes

		// Ensure no errors
		require.NoError(t, err)
		require.NotEmpty(t, res)

		// Extreact Transfer Object from the result returned by the DB.
		resTr := res.Transfer
		require.Equal(t, resTr.Amount, amount)
		require.NotEmpty(t, resTr)
		require.Equal(t, resTr.FromAccount, acc1.ID)
		require.Equal(t, resTr.ToAccount, acc2.ID)

		require.NotZero(t, resTr.ID)
		require.NotZero(t, resTr.TransactionTime)

		tr1, err := store.GetTransfer(context.Background(), resTr.ID)
		require.NoError(t, err)
		require.NotEmpty(t, tr1)

		toent := res.ToEntry
		require.NotEmpty(t, toent)
		require.Equal(t, toent.AccountID, acc2.ID)
		require.Equal(t, toent.Amount, amount)
		require.NotZero(t, toent.CreatedAt)
		require.NotZero(t, toent.ID)

		from_ent := res.FromEntry
		require.NotEmpty(t, from_ent)
		require.Equal(t, from_ent.AccountID, acc1.ID)
		require.Equal(t, from_ent.Amount, -amount)
		require.NotZero(t, from_ent.ID)

		fromAcc := res.FromAccount
		require.NotEmpty(t, fromAcc)
		require.Equal(t, fromAcc.ID, acc1.ID)

		toAcc := res.ToAccount
		require.NotEmpty(t, toAcc)
		require.Equal(t, toAcc.ID, acc2.ID)

		println("TX: ", i, " || From account balance", fromAcc.Balance, " || To account balance: ", toAcc.Balance)

		diff1 := acc1.Balance - fromAcc.Balance
		diff2 := toAcc.Balance - acc2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 >= 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		// print("The value of k is: ", k, "\n")
		require.True(t, k == i+1)
		require.True(t, k >= 0 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// Check FromAccount balance after the money transfer
	upFromAcc, err := store.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, upFromAcc)
	require.Equal(t, upFromAcc.Balance, acc1.Balance-(int64(n)*amount))

	// Check ToAccount balance after the money transfer
	upToAcc, err := store.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)
	require.NotEmpty(t, upToAcc)
	require.Equal(t, upToAcc.Balance, acc2.Balance+(int64(n)*amount))

	println("After Transacrion ==> From account balance", upFromAcc.Balance, " || To account balance: ", upToAcc.Balance)

}

func TestTransferTxOrder(t *testing.T) {
	store := NewStore(testDB)
	acc1 := createAcc(t)
	acc2 := createAcc(t)
	amount := int64(20)
	n := 6

	println("Before Transaction ==> From account balance", acc1.Balance, " || To account balance: ", acc2.Balance)

	errs := make(chan error)
	// TxRes := make(chan TransferTxRes)

	for i := 0; i < n; i++ {

		go func(errCh chan error) {
			ctx := context.Background()

			fromAccId := acc1.ID
			toAccId := acc2.ID

			if i%2 == 0 {
				fromAccId = acc2.ID
				toAccId = acc1.ID
			}

			_, err := store.TranferTx(ctx, TransferTxPrams{
				From_acc_id: fromAccId,
				To_acc_id:   toAccId,
				Amount:      amount,
			})

			errCh <- err
			// resCh <- *Trtx_res

		}(errs)
	}

	// Read results and errors from the channel || Make sure there are none...
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// Check FromAccount balance after the money transfer
	upFromAcc, err := store.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, upFromAcc)
	require.Equal(t, upFromAcc.Balance, acc1.Balance)

	// Check ToAccount balance after the money transfer
	upToAcc, err := store.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)
	require.NotEmpty(t, upToAcc)
	require.Equal(t, upToAcc.Balance, acc2.Balance)

	println("After Transacrion ==> From account balance", upFromAcc.Balance, " || To account balance: ", upToAcc.Balance)

}
