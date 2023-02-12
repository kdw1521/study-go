package main

import (
	"fmt"
	wandostring "study-go/study/wando_string"
)

func main() {
	var str string = "Hello World"
	fmt.Println(wandostring.ToUpper_String(str))
	fmt.Println(wandostring.ToUpper2_String(str))
}
