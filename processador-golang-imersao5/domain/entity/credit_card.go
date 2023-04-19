package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	cvv             int
}

func NewCreditCard(number, name string, expirationMonth, expirationYear, expirationCVV int) (*CreditCard, error) {
	cc := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             expirationCVV,
	}
	err := cc.IsValid()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (cc *CreditCard) IsValid() error {
	if err := cc.validateNumber(); err != nil {
		return err
	}
	if err := cc.validateMonth(); err != nil {
		return err
	}
	if err := cc.validateYear(); err != nil {
		return err
	}
	return nil
}

func (cc *CreditCard) validateNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !re.MatchString(cc.number) {
		return errors.New("invalid credit card number")
	}
	return nil
}

func (cc *CreditCard) validateMonth() error {
	if cc.expirationMonth < 1 || cc.expirationMonth > 12 {
		return errors.New("invalid expiration month")
	}
	return nil
}

func (cc *CreditCard) validateYear() error {
	if cc.expirationYear < time.Now().Year() {
		return errors.New("invalid expiration year")
	}
	return nil
}
