package main

import "fmt"

func main() {

	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFanancials(revenue, expenses, taxRate)

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Printf("ebt: %.1f\n", ebt)
	fmt.Printf("profile: %.1f\n", profit)
	fmt.Printf("ratio: %.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFanancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
