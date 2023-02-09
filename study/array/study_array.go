package array

import "fmt"

func Basic_Array() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}
}

func Many_Init_Array() {
	// int의 기본값이 0 이므로 5칸에 0이 할당된다.
	var nums [5]int

	days := [3]string{"a", "b", "c"}

	// 순서대로 24.3, 26.7 이 할당되고 나머지는 기본값인 0.0으로 할당된다.
	var temps [5]float64 = [5]float64{24.3, 26.7}

	// s의 타입은 뒤와 동일.
	// 1:10 -> 인덱스 1에 해당부분을 10으로 할당, 3:30 -> 인덱스 3에 30할당
	var s = [5]int{1: 10, 3: 30}

	// ... -> 길이가 뒤에 나오는 값들의 수와 같다.
	// 길이가 뒤의 초기값으로 고정됨.
	x := [...]int{10, 20, 30}

	// 1. [...]int{10, 20, 30}
	// 2. []int{10, 20, 30}
	// 1,2 는 아예다름
	// 1 은 길이 고정 , 2는 slice라고 동적배열이다.
	fmt.Println(nums, days, temps, s, x)
}

func Const_Array() {
	const Y int = 3

	// x := 5
	// a := [x]int{1, 2, 3, 4, 5} -> 변수는 할당이 안됨.

	b := [Y]int{1, 2, 3}

	var c [10]int

	fmt.Print(b, c)
}

func Literal_Array() {
	nums := [...]int{10, 20, 30, 40, 50}

	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	for i, v := range t {
		fmt.Println(i, v)
	}
}

func Set_Value_Array() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 400, 500}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
	for i, v := range b {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()

	b = a
	for i, v := range b {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
}

func Multi_Array() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	for _, arr := range a {
		for _, v := range arr {
			fmt.Println(v)
		}
	}
}
