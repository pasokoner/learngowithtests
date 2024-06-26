package pointeirerror

import "testing"

func TestWallet(t *testing.T) {

	asssertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()

		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		asssertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		asssertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(30))

		assertError(t, err, ErrInsufficientFunds)
		asssertBalance(t, wallet, startingBalance)
	})
}
