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
	//fmt.Print("Ваш индекс массы тела: ")
	//fmt.Print(IMT)
	//--вывод строки с переменной
	//fmt.Printf("Ваш индекс массы тела: %v", IMT)

	//форматирование строки 0 символов после точки с округлением
	//fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
	//fmt.Printf("Ваш индекс массы тела: %.2f", IMT)

	//форматирование строки с сохранением в переменную без вывода на консоль
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
}
