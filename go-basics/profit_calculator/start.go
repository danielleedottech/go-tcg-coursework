package profit_calculator

import (
	"errors"
	"fmt"
	"os"
)

func Start() {
	calculateAndOutputValues(getFloatInputs())
}

func getFloatInputs() (float64, float64, float64) {
	var revenue = getUserFloatInput("Enter revenue: ")
	var expenses = getUserFloatInput("Enter expenses: ")
	var taxRate = getUserFloatInput("Enter tax rate: ")

	return revenue, expenses, taxRate
}

func getUserFloatInput(message string) float64 {
	var result float64
	fmt.Println(message)
	_, e := fmt.Scan(&result)
	if e != nil {
		panic(e)
	}
	if result <= 0 {
		panic(errors.New("number has to be above zero"))
	}
	return result
}

func calculateAndOutputValues(revenue, expenses, taxRate float64) {
	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)
	resultString := fmt.Sprint(ebt, profit, ratio)
	err := os.WriteFile("./profit_calculator/result.txt", []byte(resultString), 0644)
	if err != nil {
		fmt.Println("Failed to save because of:", err)
	}
	fmt.Println(fmt.Sprintf("ebt: %f, profit: %f, ratio: %f", ebt, profit, ratio))
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	var ebt = revenue - expenses
	var profit = ebt - (ebt * (taxRate / 100.0))
	var ratio = ebt / profit

	return ebt, profit, ratio
}
