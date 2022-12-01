package main

import (
	"fmt"
)

func main() {
	listCurrency := getCurrenciesFromFile()
	var scanCurrent string
	var scanAmount int
	fmt.Println("enter the currency for exchange in rubles:")
	fmt.Scan(&scanCurrent)
	current, err := enterCurrent(scanCurrent)
	if err != nil {
		panic(err)
	}

	fmt.Println("enter the amount for exchange in rubles:")
	fmt.Scan(&scanAmount)
	amount, er := enterAmount(scanAmount)
	if er != nil {
		panic(er)
	}

	for key := range listCurrency {
		if key == current {
			fmt.Println("You get", float64(amount)*listCurrency[current], "rubles")
		}
	}

}
