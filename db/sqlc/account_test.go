package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	agr := CreateAccountParams{
		Owner:    "Van Long",
		Currency: "USD",
		Balance:  100,
	}

	account, err := testQueries.CreateAccount(context.Background(), agr)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, agr.Owner, account.Owner)
	require.Equal(t, agr.Currency, account.Currency)
	require.Equal(t, agr.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
