package main

import (
	"bufio"
	"os"
	"strconv"
)

func addFromFile() map[string]float64 {
	ListCurrencies := make(map[string]float64)
	file, err := os.Open("current.txt")
	if err != nil {
		panic(err)
	}

	var temp []byte
	var a string

	fileText := bufio.NewScanner(file)

	for fileText.Scan() {
		temp = []byte(fileText.Text())
		a = string(temp[1:4])
		b, err := strconv.ParseFloat(string(temp[7:]), 64)
		if err != nil {
			panic(err)
		}
		ListCurrencies[a] = b
	}
	return ListCurrencies
	
}
