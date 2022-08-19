package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		got := wallet.Deposit(Ethereum(10))
		AssertNoError(t, got)
		assertBalance(t, wallet, Ethereum(10.0))
	})

	t.Run("deposit negative value", func(t *testing.T) {
		startBalance := Ethereum(20)
		wallet := Wallet{startBalance}
		err := wallet.Deposit(Ethereum(-10.0))
		assertError(t, err, ErrNegativeValue)
		assertBalance(t, wallet, startBalance)
	})

	t.Run("withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Ethereum(30)}
		got := wallet.Withdraw(Ethereum(10))
		AssertNoError(t, got)
		assertBalance(t, wallet, Ethereum(20.0))
	})

	t.Run("withdraw insufficient balance", func(t *testing.T) {
		startBalance := Ethereum(20)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Ethereum(30))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startBalance)

	})
}

func assertBalance(t testing.TB, w Wallet, want Ethereum) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {

	t.Helper()

	if got == nil {
		t.Fatal("didn't get error but wanted one")
	}

	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}

}
func AssertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("unexpected error")
	}
}
