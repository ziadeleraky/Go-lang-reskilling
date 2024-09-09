package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10
	inflationRate := 2.5

	futureInvestmentValue := investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	futureRealInvestmentValue := futureInvestmentValue / math.Pow((1 + inflationRate / 100), years)

	fmt.Println("Future investment value: ", futureInvestmentValue)
	fmt.Println("Future real investment value: ", futureRealInvestmentValue)
}