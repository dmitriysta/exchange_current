package main

import (
	"errors"
	"strconv"
)

func insertAmount(value int) (int, error) {
	var (
		incorrectNumber = errors.New("Сумма ввода может содержать только цифры от 0 до 9")
		bigAmount       = errors.New("Сумма не может быть больше 1.000.000")
	)

	temp := []byte(strconv.Itoa(value))
	for i := 0; i < len(temp); i++ {
		if temp[i] < 48 || temp[i] > 57 {
			return 0, incorrectNumber
		}
	}

	if value > 1000000 {
		return 0, bigAmount
	}

	return value, nil
}
