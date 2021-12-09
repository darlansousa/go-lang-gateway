package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_IsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900
	assert.Nil(t, transaction.IsValid())
}

func TestTransaction_IsNotValidAmountGreaterThen1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1200
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "invalid transaction amount", err.Error())
}

func TestTransaction_IsNotValidAmountLessThen1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "invalid transaction amount", err.Error())
}
