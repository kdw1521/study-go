package wandoslice

import "fmt"

func Basic_Slice() {
	slice := []int{1, 2, 3}

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}
	slice[1] = 10
	fmt.Println(slice)
}

func Append_Slice() {
	slice := []int{1, 2, 3}
	slice2 := append(slice, 4)
	fmt.Println(slice)
	fmt.Println(slice2)
}

func Change_Array(array [5]int) {
	array[2] = 200
}

func Change_Slice(slice []int) {
	slice[2] = 200
}

// Runner에 addNum을 돌리면 Runner의 slice는 변경 안됨. 당연하지.
func addNum(slice []int) {
	slice = append(slice, 4)
}

func addNum2(slice *[]int) {
	*slice = append(*slice, 4)
}

func addNum3(slice []int) []int {
	slice = append(slice, 4)
	return slice
}

func Runner() {
	slice := []int{1, 2, 3}
	slice = addNum3(slice)

	fmt.Println(slice)
}

func Study_Slicing() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2]
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}

func Study_Slicing2() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2]
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100
	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}

func Study_Slicing3() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2]
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100
	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500)
	fmt.Println("After append 500")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}

func Warn_Slice() {
	array := [100]int{1: 1, 2: 2, 99: 100}

	slice1 := array[1:10]
	slice2 := slice1[2:99]

	fmt.Println(slice1)
	fmt.Println(slice2)
}

func Copy_Slice() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[:]

	slice2[1] = 100
	fmt.Println(slice1)
}

func Copy_Slice1() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	slice2[1] = 100
	fmt.Println("slice1", slice1)
	fmt.Println("slice2 ", slice2)
}

func Copy_Slice3() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := append([]int{}, slice1...)

	slice2[1] = 100
	fmt.Println("slice1", slice1)
	fmt.Println("slice2 ", slice2)
}

func Copy_Slice4() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	slice2[1] = 100
	fmt.Println("slice1", slice1)
	fmt.Println("slice2 ", slice2)
}
