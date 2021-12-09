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

func NewCreditCard(
	number string,
	name string,
	expirationMonth int,
	expirationYear int,
	cvv int) (*CreditCard, error) {

	cc := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             cvv,
	}
	err := cc.isValid()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *CreditCard) isValid() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !re.MatchString(c.number) {
		return errors.New("invalid credit card number")
	}

	if !(c.expirationMonth > 0 && c.expirationMonth < 13) {
		return errors.New("invalid expiration month")
	}

	if c.expirationYear <= time.Now().Year() {
		return errors.New("invalid expiration year")
	}
	return nil
}
