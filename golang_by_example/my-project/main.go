package main

import "fmt"

// Numeric constant
const MaxBalance = 100_000

// Exported struct (capitalized)
type User struct {
	Name     string
	LastName string
	Age      int
	Balance  float64
	Active   bool
}

// SwapNames returns first and last name swapped (multiple return)
func SwapNames(first, last string) (string, string) {
	return last, first
}

// ComputeStats returns sum and average of balances (named return values)
func ComputeStats(users []User) (sum float64, avg float64) {
	for _, u := range users {
		sum += u.Balance
	}
	if len(users) > 0 {
		avg = sum / float64(len(users)) // type conversion here
	}
	return // naked return
}

func main() {
	// Using var + zero values
	var users []User
	var activeCount int

	// Type inference with :=
	users = []User{
		{Name: "Lucky", LastName: "Baraka", Age: 25, Balance: 5000, Active: true},
		{Name: "Alice", LastName: "Wanjiku", Age: 30, Balance: 7500.50, Active: false},
		{Name: "Bob", LastName: "Karanja", Age: 22, Balance: 3200.75, Active: true},
	}

	// Loop through users
	for _, u := range users {
		if u.Active {
			activeCount++
		}
		// Swap names using multiple return
		u.Name, u.LastName = SwapNames(u.Name, u.LastName)

		// Type conversion: float64 -> int
		balanceInt := int(u.Balance)
		fmt.Printf("User: %s %s | Age: %d | Balance: %d | Active: %t\n",
			u.Name, u.LastName, u.Age, balanceInt, u.Active)
	}

	fmt.Println("Active Users:", activeCount)

	// Compute sum and average balances
	sum, avg := ComputeStats(users)
	fmt.Printf("Total Balance: %.2f | Average Balance: %.2f\n", sum, avg)

	// Demonstrating numeric constants
	if sum > MaxBalance {
		fmt.Println("Warning: Total balance exceeds the maximum allowed!")
	} else {
		fmt.Println("Total balance is within limits.")
	}
}
