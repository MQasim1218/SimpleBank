package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	acc1 := createAcc(t)
	acc2 := createAcc(t)
	amount := int64(20)

	errs := make(chan error)
	TxRes := make(chan TransferTxRes)

	for i := 0; i < 5; i++ {
		go func(errCh chan error, resCh chan TransferTxRes) {
			Trtx_res, err := store.TranferTx(context.Background(), TransferTxPrams{
				From_acc_id: acc1.ID,
				To_acc_id:   acc2.ID,
				Amount:      amount,
			})

			errCh <- err
			resCh <- *Trtx_res

		}(errs, TxRes)
	}

	for i := 0; i < 5; i++ {
		err := <-errs
		res := <-TxRes

		require.NoError(t, err)
		require.NotEmpty(t, res)

		resTr := res.Transfer
		require.Equal(t, resTr.Amount, amount)
		require.NotEmpty(t, resTr)
		require.Equal(t, resTr.FromAccount, acc1.ID)
		require.Equal(t, resTr.ToAccount, acc2.ID)

		require.NotZero(t, resTr.ID)
		require.NotZero(t, resTr.TransactionTime)

		tr1, err :=  store.GetTransfer(context.Background(), resTr.ID)
		require.NoError(t, err)
		require.NotEmpty(t, tr1)
		

	}

}
