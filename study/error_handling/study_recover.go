package errorhandling

import "fmt"

func F() {
	fmt.Println("F() start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 - ", r)
		}
	}()
	G()
	fmt.Println("F() end")
}

func G() {
	fmt.Printf("9/3 = %d\n", H(9, 3))
	fmt.Printf("9/0 = %d\n", H(9, 0))
}

func H(a, b int) int {
	if b == 0 {
		panic("b는 0 이면 안됨!")
	}
	return a / b
}
