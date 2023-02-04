package function

import "fmt"

func Hello(name string, age int8) (string, int8) {
	return "Hello " + name, age
}

func FunctionWando(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a + b
	success = true
	return
}

func Wando_Print_No(n int) {
	if n == 0 { // 재귀 탈출 조건
		return
	}
	fmt.Println(n)
	Wando_Print_No(n - 1)
	fmt.Println("After", n)
}
