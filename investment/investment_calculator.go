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

	fmt.Println("Future investment value: ", futureInvestmentValue)
	fmt.Println("Future real investment value: ", futureRealInvestmentValue)
}