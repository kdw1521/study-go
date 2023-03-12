package main

import (
	"fmt"
	"study-go/study/structure2"
)

func main() {
	m := [structure2.M]string{}

	m[structure2.Hash(23)] = "완도"
	m[structure2.Hash(259)] = "완도2"

	fmt.Printf("%d = %s\n", 23, m[structure2.Hash(23)])

	fmt.Printf("%d = %s\n", 259, m[structure2.Hash(259)])
}
