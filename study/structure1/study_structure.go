package structure1

import (
	"container/list"
	"container/ring"
	"fmt"
)

func Study_Structure() {
	v := list.New()       // list 생성
	e4 := v.PushBack(4)   // 4 라는 요소를 맨뒤에 삽입!
	e1 := v.PushFront(1)  // 1 이라는 요소를 맨 앞에 삽입!
	v.InsertBefore(3, e4) // 3 이라는 값을 e4전에 넣어라!
	v.InsertAfter(2, e1)  // 2 라는 값을 e1 후에 넣어라!

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}

// Queue
type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

// Stack
type Stack struct {
	v *list.List
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func Study_Ring() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}
	fmt.Println()
	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
}
