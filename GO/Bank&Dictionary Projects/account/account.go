package account

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	// Lowercase private Uppercase public
	owner   string
	balance int
}

var errnoMoney = errors.New("can't withdraw you are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// (a Account) => receiver in go
// (a *Account) => tell go don't copy the whole struct (pointer receiver)

// Deposit  x amount on your Account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your Account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your Account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errnoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the Account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the Account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance(), " baht")
}
