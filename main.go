package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2

	var userKg, userHeigth float64

	//fmt.Print("___Калькулятор индекса массы тела__\n")
	//fmt.Println("___Калькулятор индекса массы тела__")
	//fmt.Print("Введите свой рост в сантиметрах: ")
	//-мультистроки - приется убрать все табы, нарушить кодстайл
	fmt.Print(`___Калькулятор индекса массы тела__
Введите свой рост в сантиметрах: `)
	fmt.Scan(&userHeigth)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeigth/100, IMTPower)
	//вызов созданной функции в рамках одного пакета
	outputResult(IMT)
}

// создание функции
func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
}
