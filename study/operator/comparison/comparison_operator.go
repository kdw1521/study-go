package comparison

import "fmt"

func Compatison_Operator() {
	var x int8 = 127

	fmt.Printf("%d < %d+1: %v\n", x, x, x < x+1) // 정수 오버플로우
}
