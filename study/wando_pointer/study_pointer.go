package wandopointer

import "fmt"

func Study_Pointer() {
	var a int = 500
	var p *int

	p = &a
	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가르키는 메모리의 값: %d\n", *p)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)
}

func Equal_Pointer() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p2 == p3: %v\n", p2 == p3)
}

type Data1 struct {
	Name string
}

func Init_Struct_Pointer() {
	var data1 Data1
	var p *Data1 = &data1

	var p1 *Data1 = &Data1{}
	fmt.Println(p, p1)
}

func New_Keyword_Pointer() {
	p := &Data1{}
	var p2 = new(Data1)
	// 2개는 동일하다.
	// 초기화 여부만 다르다

	fmt.Println(p, p2)
}
