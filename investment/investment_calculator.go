package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 1

	var futureInvestmentValue = investmentAmount * math.Pow((1 + expectedReturnRate / 100), years * 12)

	fmt.Println("Future investment value after 1 year: ", futureInvestmentValue)
}