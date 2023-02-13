package modulepackage

import (
	"fmt"
	"math/rand"
)

func Runner() {
	fmt.Println(rand.Int())
}

func PrintCustom() {
	fmt.Println("This is Custom pkg")
}

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("exinit.init func", d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrintD() {
	fmt.Println("d:", d)
}
