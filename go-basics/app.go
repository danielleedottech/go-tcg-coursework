package main

import (
	"fmt"
	"go-basics/bank"
	"go-basics/investment_calculator"
	"go-basics/profit_calculator"
)

func main() {
	var programChoice int

	fmt.Println("Which program do you want?: ")
	fmt.Println("1 => profit calculator")
	fmt.Println("2 => investment calculator")
	fmt.Println("3 => bank")

	fmt.Scan(&programChoice)

	switch programChoice {
	case 0:
		profit_calculator.Start()
	case 1:
		investment_calculator.Start()
	case 2:
		bank.Start()
	default:
		fmt.Println("Program choice not recognized")
	}
}
