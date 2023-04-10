package generic

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

type Stringer2 interface {
	String() string
}

type Integer2 interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

func Print2(a Stringer2) {
	fmt.Println(a.String())
}

func Print3[T Stringer2](a T) {
	fmt.Println(a.String())
}

type MyString struct {
	name string
}

func (m MyString) String() string {
	return m.name
}

func RunnerWando() {
	m := MyString{"wando"}
	Print2(m)
	Print3(m)
}

type Stringer3 interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
	String() string
}

func PrintMin[T Stringer3](a, b T) {
	if a < b {
		fmt.Println(a.String())
	} else {
		fmt.Println(b.String())
	}
}

type MyInt2 int

func (m MyInt2) String() string {
	return strconv.Itoa(int(m))
}

func RunnerWando2() {
	var m1 MyInt2 = 10
	var m2 MyInt2 = 20
	PrintMin(m1, m2)
}

type ComparableHasher interface {
	comparable // ==, !=
	Hash() uint32
}

type MyString2 string

func (s MyString2) Hash() uint32 {
	f := fnv.New32a()
	f.Write([]byte(s))
	return f.Sum32()
}

func Equal[T ComparableHasher](a, b T) bool {
	if a == b {
		return true
	}
	return a.Hash() == b.Hash()
}

func RunnerWando3() {
	var s1 MyString2 = "Hello"
	var s2 MyString2 = "Wando"
	fmt.Println(Equal(s1, s2))
}

func Map[F, T any](s []F, f func(F) T) []T {
	rst := make([]T, len(s))
	for i, v := range s {
		rst[i] = f(v)
	}
	return rst
}

func RunnerWando4() {
	doubled := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})
	fmt.Println(doubled)

	uppered := Map([]string{"hello", "world", "wando"}, func(v string) string {
		return strings.ToUpper(v)
	})
	fmt.Println(uppered)

	toString := Map([]int{1, 2, 3}, func(v int) string {
		return "str" + strconv.Itoa(v)
	})
	fmt.Println(toString)
}
