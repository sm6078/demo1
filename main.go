package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("___Калькулятор индекса массы тела__")
	userHeigth, userKg := getUserInput()
	//вызов функции с возвращением 2 парапмтеров
	IMT := calculateIMT(userKg, userHeigth)
	//isLean := IMT < 16
	if IMT < 16 {
		fmt.Println("У вас недостаток веса")
	}
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
	IMT := userKg / math.Pow(userHeigth/100, IMTPower)
	return IMT
}

// альтернативный возврат параметра
// func calculateIMT(userKg float64, userHeigth float64) (IMT float64) {
// 	IMT = userKg / math.Pow(userHeigth/100, IMTPower)
// 	return
// }

//функция, которая возвращает несколько параметров

func getUserInput() (float64, float64) {
	var userHeigth float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeigth)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userKg)
	return userHeigth, userKg
}
