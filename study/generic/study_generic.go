package generic

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Runner() {
	var (
		a int = 10
		b int = 20
	)
	fmt.Println(min(a, b))

	var (
		c int16 = 10
		d int16 = 20
	)
	fmt.Println(min(int(c), int(d)))
}

func print[T any](a T) {
	fmt.Println(a)
}

func Runner2() {
	var a int = 10
	print(a)

	var b float32 = 3.14
	print(b)

	var c string = "Wando"
	print(c)
}

type Integer interface {
	~int | int8 | int16 | int32 | int64
}
type Float interface {
	float32 | float64
}

func min2[T Integer | Float](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type MyInt int

func Runner3() {
	var a MyInt = 10
	var b MyInt = 20
	fmt.Println(min2(a, b))
}

func min3[T comparable](a, b T) T {
	if a == b {
		return a
	} else if a != b {
		return b
	}
	return b
}
