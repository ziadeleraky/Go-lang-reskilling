package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, expectedReturnRate, years float64
	inflationRate := 2.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureInvestmentValue := investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	futureRealInvestmentValue := futureInvestmentValue / math.Pow((1 + inflationRate / 100), years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureInvestmentValue)
	formattedRFV := fmt.Sprintf("Future real value: %.1f\n", futureRealInvestmentValue)

	// fmt.Println("Future investment value: ", futureInvestmentValue)
	// fmt.Println("Future real investment value: ", futureRealInvestmentValue)
	// fmt.Printf("Future investment value: %.1f\nFuture real investment value: %.1f", futureInvestmentValue, futureRealInvestmentValue)
	fmt.Print(formattedFV, formattedRFV)
}