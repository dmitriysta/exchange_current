package main

import (
	"errors"
	"strconv"
)

var (
	incorrectNumber = errors.New("The input amount can only contain digits from 0 to 9")
	bigAmount       = errors.New("The amount cannot be more than 1.000.000")
)

func enterAmount(value int) (int, error) {

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
