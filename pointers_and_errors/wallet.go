package pointersanderrors

import (
	"errors"
	"fmt"
)

type Dogecoin int

type Wallet struct {
	balance Dogecoin
}

type Stringer interface {
	String() string
}

func (d Dogecoin) String() string {
	return fmt.Sprintf("%d DOGE", d)
}

// func takes a pointer to Wallet so that the value is changed instead of a copy
func (w *Wallet) Deposit(amount Dogecoin) {
	w.balance += amount
}

// no need for an explicit dereference of pointer -- beause struct pointer are
// automatically dereferenced
func (w *Wallet) Balance() Dogecoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Dogecoin) error {

	if amount > w.balance {
		return errors.New("you're outa money, bro")
	}

	w.balance -= amount
	return nil
}
