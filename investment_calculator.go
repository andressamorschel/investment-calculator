package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5

	var investmentAmount float64 = 1000
	var years float64 = 10

	futureValue := calculateValue(investmentAmount, inflationRate, years)

	fmt.Println(futureValue)
}

func calculateValue(investmentAmount float64, inflationRate float64, years float64) float64 {
	return investmentAmount * math.Pow(1+inflationRate/100, years)
}
