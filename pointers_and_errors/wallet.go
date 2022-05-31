package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrNegativeValue = errors.New("amount of ETH must be positive")

type Ethereum float64

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Ethereum
}

func (w *Wallet) Deposit(sum Ethereum) error {
	if sum < 0 {
		return ErrNegativeValue
	}
	w.balance += sum
	return nil
}

func (w *Wallet) Withdraw(sum Ethereum) error {

	if sum < 0 {
		return ErrNegativeValue
	}
	if w.balance < sum {
		return ErrInsufficientFunds
	}
	w.balance -= sum
	return nil
}

func (w *Wallet) Balance() Ethereum {
	return w.balance
}

func (e Ethereum) String() string {
	return fmt.Sprintf("%f ETH", e)
}
