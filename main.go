package main

import (
	// "fmt"
	"study-go/study/structure1"
)

func main() {
	// q1 := structure1.NewQueue()

	// for i := 1; i < 5; i++ {
	// 	q1.Push(i)
	// }

	// v := q1.Pop()
	// for v != nil {
	// 	fmt.Printf("%v ->", v)
	// 	v = q1.Pop()
	// }

	// stack := structure1.NewStack()
	// books := [5]string{"어린왕자", "완도일대기", "완도일기", "완도와 고랭", "완도네 패밀리"}
	// for i := 0; i < len(books); i++ {
	// 	stack.Push(books[i])
	// }

	// s := stack.Pop()
	// for s != nil {
	// 	fmt.Printf("%v ->", s)
	// 	s = stack.Pop()
	// }
	structure1.Study_Ring()
}
