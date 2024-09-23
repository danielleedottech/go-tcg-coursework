package investment_calculator

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func Start() {
	investmentAmount := 1000.0
	expectedReturnRate := 5.5
	years := 10.0

	getInitialInvestmentAmount(&investmentAmount)

	fmt.Print("Enter years of investment: ")
	_, err := fmt.Scan(&years)

	if err != nil {
		fmt.Println(fmt.Errorf("error reading years of investment: %v", err))
		return
	}

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.0f\n", futureValue)
	fmt.Printf(formattedFV)
	fmt.Printf("Future real value: %.0f\n", futureRealValue)
}

func getInitialInvestmentAmount(amount *float64) {
	fmt.Print("Enter initial investment amount: ")
	_, err := fmt.Scan(amount)

	if err != nil {
		fmt.Println(fmt.Errorf("error reading initial investment amount: %v", err))
		panic(err)
	}
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1.0+expectedReturnRate/100.0, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100.0, years)
	return futureValue, futureRealValue
}
