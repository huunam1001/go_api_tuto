package db

import (
	"context"
	"go_api_tuto/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	agr := CreateAccountParams{
		Owner:    util.RandomOwnerName(),
		Currency: util.RandomCurrency(),
		Balance:  util.RandomBalance(),
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
