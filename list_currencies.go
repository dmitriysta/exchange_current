package main

import (
	"bufio"
	"os"
	"strconv"
)

func getCurrenciesFromFile(nameListCurrent string) (listCurrencies map[string]float64) {

	const startNameCurrency = 1
	const endNameCurrency = 4
	const startRateCurrency = 7

	file, err := os.Open(nameListCurrent)
	if err != nil {
		panic(err)
	}

	listCurrencies = make(map[string]float64)

	var temp []byte
	var nameCurrency string

	fileText := bufio.NewScanner(file)

	for fileText.Scan() {
		temp = []byte(fileText.Text())
		nameCurrency = string(temp[startNameCurrency:endNameCurrency])
		rateCurrency, err := strconv.ParseFloat(string(temp[startRateCurrency:]), 64)
		if err != nil {
			panic(err)
		}
		listCurrencies[nameCurrency] = rateCurrency
	}
	return listCurrencies
}
