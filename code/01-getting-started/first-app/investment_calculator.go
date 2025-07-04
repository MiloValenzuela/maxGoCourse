package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment amount: ")
	outputText("Investment amout: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Espected Return Rate: ")
	outputText("Espected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue := calculatedFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue, futureRealValue, years := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	// outputs information
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf(`Future Value: %.0f\nFuture Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculatedFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+expectedReturnRate/100, years)
	return fv, rfv
}
