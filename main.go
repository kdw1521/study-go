package main

import "fmt"

func main() {
	var number1 int = 10
	var number2 float64 = 33.5

	var number3 int = int(number2) // number2는 float64이므로 int type으로 형변환해서 할당, int로 변환하기떄문에 소수는 사라짐

	number4 := float64(number1) * number2

	var number5 int64 = 20
	number6 := number1 * int(number5)

	fmt.Println(number1, number2, number3, number4, number5, number6)
}
