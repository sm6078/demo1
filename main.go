package main

import (
	"fmt"
	"math"
)

func main() {

	//fmt.Print("___Калькулятор индекса массы тела__\n")
	//fmt.Println("___Калькулятор индекса массы тела__")
	//fmt.Print("Введите свой рост в сантиметрах: ")
	//-мультистроки - приется убрать все табы, нарушить кодстайл
	fmt.Println("___Калькулятор индекса массы тела__")
	userHeigth, userKg := getUserInput()

	IMT := calculateIMT(userKg, userHeigth)
	//вызов созданной функции в рамках одного пакета
	outputResult(IMT)
}

// создание функции
func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
}

// создание функции с вовзращаемым параметром
func calculateIMT(userKg float64, userHeigth float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeigth/100, IMTPower)
	return IMT
}

//функция, которая возвращает несколько параметров

func getUserInput() (float64, float64) {
	var userKg float64
	var userHeigth float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeigth)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userKg)
	return userHeigth, userKg
}
