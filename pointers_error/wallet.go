package pointers_error

import (
	"errors"
	"fmt"
)

var insufficientfunds = errors.New("cannot withdraw, insufficient balance")

type Stringer interface {
	String() string
}
type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) WithDraw(amount Bitcoin) error {

	if amount > w.balance {
		return insufficientfunds
	}
	w.balance -= amount
	return nil

}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
