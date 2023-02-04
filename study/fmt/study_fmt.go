package study_fmt

import "fmt"

func Study_fmt() {
	var num1 int = 10
	var num2 int = 20
	var num3 float64 = 32727327327.8297

	fmt.Print("num1 :", num1, "num2 :", num2)
	fmt.Println("num1 :", num1, "num2 :", num2, "num3 :", num3)
	fmt.Printf("num1 : %d num2 : %d num3 : %f \n", num1, num2, num3)

	var num4 int
	var num5 int
	num, err := fmt.Scanln(&num4, &num5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num, num4, num5)
	}
}
