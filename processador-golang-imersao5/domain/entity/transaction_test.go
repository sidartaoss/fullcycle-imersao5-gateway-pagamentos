package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionIsValid(t *testing.T) {
	// arrange
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900

	// act
	err := transaction.IsValid()

	// assert
	assert.Nil(t, err)
}

func TestTransactionIsNotValidWithAmountGreaterThan1000(t *testing.T) {
	// arrange
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1900

	// act
	err := transaction.IsValid()

	// assert
	assert.NotNil(t, err)
	assert.Equal(t, "no limit for this transaction", err.Error())
}

func TestTransactionIsNotValidWithAmountLessThan1(t *testing.T) {
	// arrange
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0

	// act
	err := transaction.IsValid()

	// assert
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than 0", err.Error())
}
