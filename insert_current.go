package main

import "errors"

func insertCurrent(value string) (string, error) {
	var (
		emptyField        = errors.New("Поле ввода пустое, необходимо ввести валюту")
		incorrectLength   = errors.New("Валюты ввода должна соответствовать международной спецификации и содержать 3 символа")
		incorrectLanguage = errors.New("Валюта ввода может содержать только буквы английского алфавита")
		noCurrency        = errors.New("Данной валюты нет в списке для конвертации")
	)

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

	listing := addFromFile()

	if _, ok := listing[value]; !ok {
		return "0", noCurrency
	}

	return value, nil
}
