package errorhandling

import "fmt"

func Divide(a, b int) {
	if b == 0 {
		panic("b는 0일수 없음!")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}
