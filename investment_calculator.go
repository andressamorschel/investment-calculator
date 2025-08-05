package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5

	var investmentAmount float64
	var years float64

	fmt.Print("Enter your investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := calculateValue(investmentAmount, inflationRate, years)

	fmt.Println("Result: ", futureValue)
}

func calculateValue(investmentAmount float64, inflationRate float64, years float64) float64 {
	return investmentAmount * math.Pow(1+inflationRate/100, years)
}
