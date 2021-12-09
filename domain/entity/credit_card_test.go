package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("41916194649191949919849111", "José da Silva", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4576010007066439", "José da Silva", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4576010007066439", "José da Silva", 20, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4576010007066439", "José da Silva", 5, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	_, err := NewCreditCard("4576010007066439", "José da Silva", 12, 2012, 123)
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("4576010007066439", "José da Silva", 12, 2024, 123)
	assert.Nil(t, err)
}
