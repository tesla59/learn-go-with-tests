package pointers_test

import (
	"go-with-tests/pointers"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))
		want := pointers.Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(20))

		err := wallet.Withdraw(pointers.Bitcoin(10))
		want := pointers.Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := pointers.Wallet{}
		err := wallet.Withdraw(pointers.Bitcoin(50))

		assertBalance(t, wallet, 0)
		assertError(t, err, pointers.ErrInsufficientFund)
	})
}

func assertBalance(t *testing.T, w pointers.Wallet, want pointers.Bitcoin) {
	t.Helper()
	if w.Balance() != want {
		t.Errorf("got %s want %s", w.Balance(), want)
	}
}

func assertError(t *testing.T, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("error expected but got nil")
	}
	if want != err {
		t.Errorf("wanted err %q got err %q", want, err)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("didn't expect any error got %q", got)
	}
}
