package wallet

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Deposit(Bitcoin(10))
		// fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		want := Bitcoin(20)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(0)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("string", func(t *testing.T) {
		btc := Bitcoin(10)

		got := btc.String()
		want := "10 BTC"

		if got != want {
			t.Errorf("%#v: got %q want %q", btc, got, want)
		}
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

	t.Run("withdraw unwrap new error", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := ProcessWithdrawl(&wallet, "acc-012", Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("%#v: got %g want %g", wallet, got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("wanted an error but didn't get one")
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if errors.Is(err, want) {
		t.Errorf("got %q, want %q", err, want)
	}
}
