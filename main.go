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

	//можно сделать вот так
	userHeigth := 1.85
	userKg := 100.0

	IMT := userKg / math.Pow(userHeigth, 2)
	fmt.Print(IMT)
}
