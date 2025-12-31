package db

import (
	"context"
	"testing"
	"time"

	"github.com/mybank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:     utils.RandomOwner(),
		HashPassword: hashedPassword,
		Email:        utils.RandomEmail(),
		FullName:     utils.RandomOwner(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashPassword, user.HashPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashPassword, user2.HashPassword)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.WithinDuration(t, user1.PasswordChangedAt, user1.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

// func TestUpdateUser(t *testing.T) {
// 	user1 := CreateRandomUser(t)

// 	arg := UpdateUserParams{
// 		ID:       user1.ID,
// 		Username: utils.RandomOwner(),
// 		HashPassword:  ("new_secret"),
// 		Email: utils.RandomEmail(),
// 		FullName: utils.RandomOwner(),
// 	}

// 	user2, err := testQueries.UpdateUser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.ID, user2.ID)
// 	require.Equal(t, arg.Username, user2.Username)
// 	require.Equal(t, arg.HashPassword, user2.HashPassword)
// 	require.Equal(t, arg.Email, user2.Email)
// 	require.Equal(t, arg.FullName, user2.FullName)
// 	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
// }
// 		Currency: utils.RandomCurrency(),
// 		Owner: account1.Owner,
// 	}

// 	account2, err := testQueries.UpdateAccount(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, account2)

// 	require.Equal(t, account1.ID, account2.ID)
// 	require.Equal(t, account1.Owner, account2.Owner)
// 	require.Equal(t, arg.Balance, account2.Balance)
// 	require.Equal(t, arg.Currency, account2.Currency)
// 	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
// }

// func TestDeleteAccount(t *testing.T) {
// 	account1 := CreateRandomAccount(t)

// 	err := testQueries.DeleteAccount(context.Background(), account1.ID)
// 	require.NoError(t, err)

// 	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, account2)
// }

// func TestListAccounts(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		CreateRandomAccount(t)
// 	}

// 	arg := ListAccountParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}

// 	accounts, err := testQueries.ListAccount(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, accounts, 5)

// 	for _, account := range accounts {
// 		require.NotEmpty(t, account)
// 	}
// }
