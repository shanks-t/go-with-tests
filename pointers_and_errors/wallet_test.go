package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Dogecoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Dogecoin(10))
		assertBalance(t, wallet, Dogecoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Dogecoin(20)}
		wallet.Withdraw(Dogecoin(10))
		assertBalance(t, wallet, Dogecoin(10))

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Dogecoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Dogecoin(100))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)

	})
}
