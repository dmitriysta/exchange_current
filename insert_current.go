package main

import (
	"errors"
)

var (
	emptyField        = errors.New("The input field is empty, you need to enter the currency")
	incorrectLength   = errors.New("The input currency must comply with the international specification and contain 3 characters")
	incorrectLanguage = errors.New("The input currency can contain only letters of the English alphabet")
	noCurrency        = errors.New("This currency is not in the list for conversion")
)

func enterCurrent(value string) (string, error) {

	if value == "" {
		return "0", emptyField
	}

	if len(value) != 3 {
		return "0", incorrectLength
	}

	temp := []byte(value)
	for i := 0; i < 3; i++ {
		if temp[i] < 97 || temp[i] > 122 {
			return "0", incorrectLanguage
		}
	}
	if _, ok := listCurrencies[value]; !ok {
		return "0", noCurrency
	}

	return value, nil
}
