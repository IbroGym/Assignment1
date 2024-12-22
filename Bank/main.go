package main

import (
	"fmt"
)

// Bank account structure
type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

// Deposit adds money to the account
func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("Deposited %.2f. New balance: %.2f\n", amount, b.Balance)
	} else {
		fmt.Println("Invalid deposit amount.")
	}
}

// Withdraw deducts money from the account if sufficient balance
func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= b.Balance {
		b.Balance -= amount
		fmt.Printf("Withdrawn %.2f. New balance: %.2f\n", amount, b.Balance)
	} else if amount <= 0 {
		fmt.Println("Invalid withdrawal amount.")
	} else {
		fmt.Println("Insufficient funds.")
	}
}

// GetBalance returns the current balance
func (b *BankAccount) GetBalance() {
	fmt.Printf("Current balance: %.2f\n", b.Balance)
}

// Transaction processes a slice of transaction amounts
func Transaction(account *BankAccount, transactions []float64) {
	for _, amount := range transactions {
		if amount > 0 {
			account.Deposit(amount)
		} else {
			account.Withdraw(-amount) // Convert negative amount to positive for withdrawal
		}
	}
}

func main() {
	// Create a sample account
	account := BankAccount{
		AccountNumber: "1234567890",
		HolderName:    "John Doe",
		Balance:       1000.00,
	}

	var choice int

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			var amount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.GetBalance()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
