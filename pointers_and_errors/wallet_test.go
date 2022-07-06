package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Dogecoin(10))

		assertBalance(t, wallet, Dogecoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Dogecoin(20)}
		err := wallet.Withdraw(Dogecoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Dogecoin(10))

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Dogecoin(20)}
		err := wallet.Withdraw(Dogecoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Dogecoin(20))

	})
}

func assertBalance(t testing.TB, wallet Wallet, want Dogecoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
