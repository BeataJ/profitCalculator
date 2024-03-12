package main

import (
	"errors"
	"fmt"
	"os"
)

const FinancialsFile = "financials.txt"

func writeFanancialsToFile(ebt, profit, ratio float64) {
	financialText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRadio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(FinancialsFile, []byte(financialText), 0644)
}

func main() {

	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// Goals
	// 1) Validate user input
	//  => Show error message & exit if invalid input is provided
	// - no negative numbers
	//  - not  0
	// 2) Store calculated results into file

	revenue, err1 := getUserInput("Revenue: ")
	expenses, err2 := getUserInput("Expenses: ")
	taxRate, err3 := getUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	// if err2 != nil {
	// 	fmt.Println("expenses error")
	// 	return
	// }

	// if err3 != nil {
	// 	fmt.Println("tax rate error")
	// 	return
	// }

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Printf("ebt: %.1f\n", ebt)
	fmt.Printf("profile: %.1f\n", profit)
	fmt.Printf("ratio: %.3f\n", ratio)
	writeFanancialsToFile(ebt, profit, ratio)

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("input must be not 0 or more than 0")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
