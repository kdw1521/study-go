package wandoslice

import (
	"fmt"
	"sort"
)

func Sort_Slice() {
	slice := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(slice)
	fmt.Println(slice)
}

type Student struct {
	Name string
	Age  int
}
type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func Sort_Struct() {
	s := []Student{
		{"wando", 31},
		{"minju", 41},
		{"marco", 21},
		{"lupy", 51},
	}
	sort.Sort(Students(s)) // 요걸 쓰려면 interface여야 하고 그 interface는 Len, Less, Swap 함수를 가져야한다.
	fmt.Println(s)
}
