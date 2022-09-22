package db

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createSingleEntry(t *testing.T) Entry {
	args := CreateEntryParams{
		AccountID: rand.Int63n(10),
		Amount:    100,
		CreatedAt: time.Now(),
	}

	ent, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, ent)

	require.NotZero(t, ent.ID)
	require.Equal(t, ent.AccountID, args.AccountID)
	require.Equal(t, ent.Amount, args.Amount)
	require.NotZero(t, ent.CreatedAt)

	return ent
}

func TestCreateEntry(t *testing.T) {
	createSingleEntry(t)

}
