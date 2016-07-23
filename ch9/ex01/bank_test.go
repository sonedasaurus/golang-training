package bank_test

import (
	"testing"

	"."
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Withdraw Test
	result := bank.Withdraw(100)

	if got, want := bank.Balance(), 200; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	if got, want := result, true; got != want {
		t.Errorf("Withdraw Result = %t, want %t", got, want)
	}

	result = bank.Withdraw(200)

	if got, want := bank.Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	if got, want := result, true; got != want {
		t.Errorf("Withdraw Result = %t, want %t", got, want)
	}

	result = bank.Withdraw(100)

	if got, want := bank.Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	if got, want := result, false; got != want {
		t.Errorf("Withdraw Result = %t, want %t", got, want)
	}
}
