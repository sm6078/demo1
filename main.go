package main

import (
	"fmt"
	"math"
)

func main() {
	//короткая запись
	//var userHeigth, userKg float64 = 1.85, 100

	//еще более короткая запись, но не можем указать тип переменной придется делать явное преобразование при расчете
	//userHeigth, userKg := 1.85, 100

	//без var
	//userHeigth := 1.85

	//можно задать переменную, а потом писвоить ей значение
	//var userKg float64
	//userKg = 114

	//Константы
	//при объявлении константы можем явно указывать тип
	//const IMTPower int = 2

	//при объявлении константы можем явно не указывать тип - он буде untyped
	const IMTPower = 2

	//можно сделать вот так
	userHeigth := 1.85
	userKg := 100.0

	IMT := userKg / math.Pow(userHeigth, IMTPower)
	fmt.Print(IMT)
}
