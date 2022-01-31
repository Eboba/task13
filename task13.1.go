package main

import (
	"fmt"
)

func shift(x, y *int) {
	*x, *y = *y, *x
}

func sumOfEven(e, g int) {
	sum := 0

	if e > g {
		shift(&e, &g)
	}

	for {
		if e%2 == 0 {
			sum += e
		}

		if e == g {
			break
		}
		e++
	}
	fmt.Println("Сумму чётных чисел заданного диапазона:", sum)
}

func main() {
	fmt.Println("Задание 1. Функция, принимающая аргументы")

	var a, b int

	fmt.Println("Введите 1 число:")
	fmt.Scan(&a)

	fmt.Println("Введите 2 число:")
	fmt.Scan(&b)

	sumOfEven(a, b)
}
