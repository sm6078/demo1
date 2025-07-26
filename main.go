package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeigth = 1.85
	var userKg float64 = 114
	// var IMT = float64(userKg) / math.Pow(userHeigth, 2)
	var IMT = userKg / math.Pow(userHeigth, 2)
	fmt.Print(IMT)
}
