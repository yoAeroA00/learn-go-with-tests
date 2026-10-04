package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amountToAdd Bitcoin) {
	// fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	(*w).balance += amountToAdd
}

func (w *Wallet) Withdraw(amountToRemove Bitcoin) error {

	if amountToRemove > (*w).balance {
		return ErrInsufficientFunds
	}

	(*w).balance -= amountToRemove
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

func ProcessWithdrawl(wallet *Wallet, accountID string, amountToRemove Bitcoin) error {
	if err := (*wallet).Withdraw(amountToRemove); err != nil {
		return fmt.Errorf("processing withdrawl for account %s: %w", accountID, err)
	}
	return nil
}
