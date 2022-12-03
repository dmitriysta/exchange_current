package main

import (
	"fmt"
)

func main() {
	nameListCurrent := "current.txt"
	listCurrency := getCurrenciesFromFile(nameListCurrent)
	var scanCurrent string
	var scanAmount int
	fmt.Println("enter the currency for exchange in rubles:")
	fmt.Scan(&scanCurrent)
	current, err := validateCurrency(scanCurrent, listCurrency)
	if err != nil {
		panic(err)
	}

	fmt.Println("enter the amount for exchange in rubles:")
	fmt.Scan(&scanAmount)
	amount, er := parcelAmount(scanAmount)
	if er != nil {
		panic(er)
	}

	if _, ok := listCurrency[current]; ok {
		fmt.Println("You get", float64(amount)*listCurrency[current], "rubles")
	}

}
