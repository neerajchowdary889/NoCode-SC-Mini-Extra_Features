package main

import (
	"testing"
)

func TestCreateAccount(t *testing.T) {
	// Create a new account
	account := NewAccount("Alice")

	// Check that the account was created successfully
	if account == nil {
		t.Errorf("Failed to create account")
	}

	// Check that the account's balance is 0
	if account.Balance != 0 {
		t.Errorf("Account balance should be 0")
	}
}

func TestDeposit(t *testing.T) {
	// Create a new account
	account := NewAccount("Alice")

	// Deposit 100 into the account
	account.Deposit(100)

	// Check that the account's balance is 100
	if account.Balance != 100 {
		t.Errorf("Account balance should be 100")
	}
}

func TestWithdraw(t *testing.T) {
	// Create a new account
	account := NewAccount("Alice")

	// Deposit 100 into the account
	account.Deposit(100)

	// Withdraw 50 from the account
	account.Withdraw(50)

	// Check that the account's balance is 50
	if account.Balance != 50 {
		t.Errorf("Account balance should be 50")
	}
}