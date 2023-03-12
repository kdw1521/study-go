package structure2

import "fmt"

func Study_Map() {
	m := make(map[string]string)

	m["완도"] = "서울 은평구"
	m["완도자식"] = "서울 강남구"

	m["완도"] = "서울 강남구" // 값 변경 시
	fmt.Printf("완도의 주소는 %s이다.\n", m["완도"])
}

type Product struct {
	Name  string
	Price int
}

func Study_Map_Loop() {
	m := make(map[int]Product)

	m[16] = Product{"연필", 100}
	m[46] = Product{"지우개", 50}
	m[78] = Product{"자", 1000}
	m[345] = Product{"샤프", 3000}
	m[897] = Product{"샤프심", 20}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func Study_Map_Delete_Check() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	delete(m, 4)

	v, ok := m[3]
	fmt.Println(v, ok)

	v, ok = m[1]
	fmt.Println(v, ok)
}

const M = 10

func Hash(d int) int {
	return d % M
}
