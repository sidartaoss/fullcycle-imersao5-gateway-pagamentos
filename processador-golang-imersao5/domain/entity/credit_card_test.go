package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidCreditCardNumber(t *testing.T) {
	cc, err := NewCreditCard("4538520392582409", "Jose da Silva", 12, 2024, 123)
	assert.Nil(t, err)
	assert.NotNil(t, cc)
}

func TestInvalidCreditCardNumber(t *testing.T) {
	cc, err := NewCreditCard("400000000000", "Jose da Silva", 12, 2024, 123)
	assert.NotNil(t, err)
	assert.Nil(t, cc)
	assert.Equal(t, "invalid credit card number", err.Error())
}

func TestInvalidCreditCardExpirationMonth(t *testing.T) {
	// arrange
	expirationMonth := 13

	// act
	cc, err := NewCreditCard("4538520392582409", "Jose da Silva", expirationMonth, 2024, 123)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, cc)

	assert.Equal(t, "invalid expiration month", err.Error())
}

func TestValidCreditCardExpirationMonth(t *testing.T) {
	// arrange
	expirationMonth := 11

	// act
	cc, err := NewCreditCard("4538520392582409", "Jose da Silva", expirationMonth, 2024, 123)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, cc)
}

func TestCreditCardExpirationYear(t *testing.T) {
	// arrange
	lastYear := time.Now().AddDate(-1, 0, 0).Year()

	// act
	cc, err := NewCreditCard("4538520392582409", "Jose da Silva", 12, lastYear, 123)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, cc)

	assert.Equal(t, "invalid expiration year", err.Error())

}
