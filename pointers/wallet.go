package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += 10
}

func (w *Wallet) Balance() int {
	return w.balance
}
