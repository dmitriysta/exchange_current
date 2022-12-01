package main

import (
	"fmt"
)

func main() {
	list := addFromFile()
	var a string
	var b int

	fmt.Println("Введите валюту для конвертации в рубли:")
	fmt.Scan(&a)
	c, err := insertCurrent(a)
	if err != nil {
		panic(err)
	}

	fmt.Println("Введите сумму для конвертации в рубли:")
	fmt.Scan(&b)
	d, er := insertAmount(b)
	if er != nil {
		panic(er)
	}

	for key := range list {
		if key == c {
			fmt.Println("Вы получаете", float64(d)*list[c], "рублей")
		}
	}

}
