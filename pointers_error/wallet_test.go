package pointers_error

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but did not got")
		}
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	t.Run("Bitcoin deposit", func(t *testing.T) {
		w := Wallet{balance: 20}
		w.Deposit(10)
		want := Bitcoin(30)

		assertBalance(t, w, want)

	})

	t.Run("Bitcoin withdraw", func(t *testing.T) {
		w := Wallet{Bitcoin(20)}
		err := w.WithDraw(10)
		want := Bitcoin(10)
		fmt.Println(err)

		assertBalance(t, w, want)


	})

	t.Run("withdraw without funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		w := Wallet{startingBalance}
		err := w.WithDraw(Bitcoin(100))

		assertBalance(t, w, startingBalance)
		assertError(t, err, insufficientfunds)
	})

	t.Run("Bitcoin string", func(t *testing.T) {
		btc := Bitcoin(10)
		got := btc.String()
		want := "10 BTC"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
