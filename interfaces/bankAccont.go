package main

import "fmt"

// Interface
type Account interface {
	Deposit(amount float64)
	Balance() float64
}

// CheckingAccount struct
type CheckingAccount struct {
	balance float64
}

func (c *CheckingAccount) Deposit(amount float64) {
	c.balance += amount
}

func (c *CheckingAccount) Balance() float64 {
	return c.balance
}

// SavingsAccount struct
type SavingsAccount struct {
	balance float64
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.balance += amount
}

func (s *SavingsAccount) Balance() float64 {
	return s.balance
}

func PrintBalance(a Account) {
	fmt.Println("Balance:", a.Balance())
}

func main() {
	ca := &CheckingAccount{balance: 100}
	sa := &SavingsAccount{balance: 500}

	PrintBalance(ca) // Balance: 100
	PrintBalance(sa) // Balance: 500

	ca.Deposit(50)
	sa.Deposit(100)

	PrintBalance(ca) // Balance: 150
	PrintBalance(sa) // Balance: 600
}
