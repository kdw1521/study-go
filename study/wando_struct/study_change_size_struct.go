package wandostruct

import (
	"fmt"
	"unsafe"
)

type User2 struct {
	Age   int32
	Score float64
}

func Size_Struct() {
	user := User2{31, 77.2}
	fmt.Println(unsafe.Sizeof(user)) // 메모리 사이즈를 반환해주는 함수 Sizeof
}

type Wando struct {
	A int8 // 1
	B int  // 8
	C int8 // 1
	D int  // 8
	E int8 // 1
}

type Wando2 struct {
	A int8 // 1
	C int8 // 1
	E int8 // 1
	B int  // 8
	D int  // 8
}

func Wando_Struct() {
	wando := Wando2{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(wando))
}
