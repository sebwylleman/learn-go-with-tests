package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(0)

		if got != want {
			t.Errorf("want %s, got %s", got, want)
		}
	})
}
