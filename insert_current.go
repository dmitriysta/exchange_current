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

func validateCurrency(value string, listCurrencies map[string]float64) (string, error) {

	if value == "" {
		return "0", emptyField
	}

	if len(value) != 3 {
		return "0", incorrectLength
	}

	//здесь проверка идет именно на то, что соблюдается именно формат валюты. То есть валюта
	//может быть в Мапе, но к примеру её неправильно вводит пользователь (вместо RUR вводит RUB)
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
