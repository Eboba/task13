package main

import (
	"fmt"
)

func shift(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	fmt.Println("Задание 2. Функция, принимающая значение по ссылке")

	a := 4
	b := 3

	fmt.Println(a, b)

	shift(&a, &b)

	fmt.Println(a, b)
}
