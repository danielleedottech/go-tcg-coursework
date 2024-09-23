package bank

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"log"
	"os"
	"strconv"
	"time"
)

const filePath = "./bank/balance.txt"

func getBalanceFromFile() float64 {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(filePath, []byte(balanceText), 0644)
}

func Start() {
	var accountBalance float64 = getBalanceFromFile()
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 at", randomdata.PhoneNumber())

	for {
		printOptions()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var amount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&amount)
			exitOnInvalidAmount(amount)
			accountBalance += amount
			writeBalanceToFile(accountBalance)
			fmt.Println("Your balance is", accountBalance)
		case 3:
			var amount float64
			fmt.Print("Your withdrawal: ")
			fmt.Scan(&amount)
			exitOnInvalidAmount(amount)

			if amount > accountBalance {
				fmt.Println("Not enough money")
				os.Exit(0)
			}
			accountBalance -= amount
			writeBalanceToFile(accountBalance)
			fmt.Println("Your balance is", accountBalance)
		default:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		time.Sleep(1 * time.Second)
		fmt.Println("Loading....")
		time.Sleep(1 * time.Second)
	}
}

func exitOnInvalidAmount(amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid amount, Must be greater than 0.")
		os.Exit(0)
	}
}

func printOptions() {
	fmt.Println("What do you want do to?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}
