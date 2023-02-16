package wandoslice

import "fmt"

func Delete_Slice() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	idx := 2 // 지우고 싶은 idx

	for i := idx + 1; i < len(slice1); i++ {
		slice1[i-1] = slice1[i]
	}

	slice1 = slice1[:len(slice1)-1]
	fmt.Println("slice1", slice1)
}

func Delete_Slice2() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	idx := 2 // 지우고 싶은 idx

	slice1 = append(slice1[:idx], slice1[idx+1:]...)
	fmt.Println("slice1", slice1)
}

func Delete_Slice3() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	idx := 2 // 지우고 싶은 idx

	copy(slice1[idx:], slice1[idx+1:])
	slice1 = slice1[:len(slice1)-1]
	fmt.Println("slice1", slice1)
}
