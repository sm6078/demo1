package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2

	var userKg, userHeigth float64

	fmt.Print("___Калькулятор индекса массы тела__\n")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeigth)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeigth, IMTPower)
	fmt.Print("Ваш индекс массы тела: ")
	fmt.Print(IMT)
}
