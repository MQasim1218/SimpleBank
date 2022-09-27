package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (st *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	Tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queryObj := New(Tx)
	if err = fn(queryObj); err != nil {
		if errTx := Tx.Rollback(); errTx != nil {
			return fmt.Errorf("tx err: %v queryerror: %v", errTx, err)
		}
		return err
	}

	return Tx.Commit()
}

// Contains all necessary parameters to transfer Money between 2 accounts;
type TransferTxPrams struct {
	From_acc_id int64 `json:"from_acc_id"`
	To_acc_id   int64 `json:"to_acc_id"`
	Amount      int64 `json:"amount"`
}

//
type TransferTxRes struct {
	FromAccount Account  `json:"from_acc"`
	ToAccount   Account  `json:"to_acc"`
	Transfer    Transfer `json:"transfer"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (st *Store) TranferTx(ctx context.Context, args TransferTxPrams) (*TransferTxRes, error) {
	var res TransferTxRes = TransferTxRes{}

	st.execTx(ctx, func(q *Queries) (err error) {

		// Create a transfer
		tr_args := CreateTransferParams{
			FromAccount:     args.From_acc_id,
			ToAccount:       args.To_acc_id,
			TransactionTime: time.Now(),
			Amount:          args.Amount,
		}
		if res.Transfer, err = q.CreateTransfer(ctx, tr_args); err != nil {
			return err
		}

		// Create an Entry for account 1.

		entry_args := CreateEntryParams{
			AccountID: args.To_acc_id,
			Amount:    args.Amount,
			CreatedAt: time.Now(),
		}
		if res.ToEntry, err = q.CreateEntry(ctx, entry_args); err != nil {
			return err
		}

		entry_args = CreateEntryParams{
			AccountID: args.From_acc_id,
			Amount:    -args.Amount,
			CreatedAt: time.Now(),
		}
		if res.FromEntry, err = q.CreateEntry(ctx, entry_args); err != nil {
			return err
		}

		{
			// Todo: Get The Accounts and Change their Balance
			// up_args := UpdateAccountParams{
			// 	ID:       0,
			// 	Owner:    "",
			// 	Balance:  0,
			// 	Currency: "",
			// }
			// if res.FromAccount, err = q.UpdateAccount(context.Background(), args.From_acc_id); err != nil {
			// 	return err
			// }
			// if res.ToAccount, err = q.GetAccount(context.Background(), args.To_acc_id); err != nil {
			// 	return err
			// }
		}

		return nil
	})

	return &res, nil
}
