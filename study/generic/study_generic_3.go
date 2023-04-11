package generic

import "fmt"

// no generic
// type Node1 struct {
// 	val  interface{}
// 	next *Node1
// }

type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](v T) *Node[T] {
	return &Node[T]{val: v}
}

func (n *Node[T]) Push(v T) *Node[T] {
	node := NewNode(v)
	n.next = node
	return node
}

func RunnerWando5() {
	node1 := NewNode(1)
	node1.Push(2).Push(3).Push(4)

	for node1 != nil {
		fmt.Println(node1.val, " - ")
		node1 = node1.next
	}

	node2 := NewNode("Hi")
	node2.Push("Hello").Push("Wando")

	for node2 != nil {
		fmt.Println(node2.val, " - ")
		node2 = node2.next
	}
}

// 메서드는 지원하지 않음!
// type WandoNode struct {
// 	val int
// 	next *WandoNode
// }

// func (w *WandoNode) Push[T any](a T) {}
