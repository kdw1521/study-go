package wandoconst

import "fmt"

func Study_Const() {
	const pi1 float64 = 3.141592 // 상수
	var pi2 float64 = 3.141592   // 변수

	// pi1 = 3 // 변경 불가! 컴파일 에서 이미 잡혀버린다.
	pi2 = 4

	fmt.Printf("pi1 %f", pi1)
	fmt.Printf("pi2 %f", pi2)
}

const Wando int = 0
const Minju int = 1
const Marco int = 3

const (
	Wando1 int = iota // 0
	Minju1 int = iota // 1
	Marco1 int = iota // 2
)

const (
	BitFlag1 uint = 1 << iota // 1 = 1 << 0
	BitFlag2                  // 2 = 1 << 1
	BitFlag3                  // 4 = 1 << 2
	BitFlag4                  // 8 = 1 << 3
)

func PrintPeople(people int) {
	if people == Wando {
		fmt.Println("난 완도")
	} else if people == Minju {
		fmt.Println("난 민주야")
	} else if people == Marco {
		fmt.Println("난 마르코야")
	} else {
		fmt.Println("...")
	}
}
