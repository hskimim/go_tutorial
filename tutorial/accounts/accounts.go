package accounts

import (
	"errors"
)

type bankAccount struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("you can't withdraw money more than you have")

func NewAccount(name string) *bankAccount {
	accounts := bankAccount{owner: name, balance: 0}
	return &accounts
}

func (b *bankAccount) Deposit(money int) {
	b.balance += money
}

func (b bankAccount) ShowBalance() int {
	return b.balance
}

func (b *bankAccount) Withdraw(money int) error {
	if b.balance-money < 0 {
		return errNoMoney
	}
	b.balance -= money
	return nil
}

func (b bankAccount) String() string {
	return "this is about bank account"
}
