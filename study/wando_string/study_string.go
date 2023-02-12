package wandostring

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func Poet_String() {
	poet1 := "시시시시시시시시시시시시 \n이다."
	poet2 := `시시시시시시시시시시시시 
이다.`

	fmt.Println(poet1)
	fmt.Println(poet2)
}

func Loop_String() {
	str := "Hello 월드"
	for i := 0; i < len(str); i++ {
		fmt.Printf("타입: %T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}
}

func Loop_String2() {
	str := "Hello 월드"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}
}

func Loop_String3() {
	str := "Hello 월드"

	for _, v := range str {
		fmt.Printf("타입: %T 값:%d 문자값:%c\n", v, v, v)
	}
}

func Struct_String() {
	str1 := "Hello 월드"
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}

func Immutable_String() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}

func Immutable_String2() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringHeader.Data)
	fmt.Printf("slice:\t%x\n", sliceHeader.Data)
}

func ToUpper_String(str string) string {
	var rst string

	for _, c := range str {
		if c >= 'a' && c <= 'z' { // 소문자인가?
			rst += string('A' + (c - 'a')) // 65 + (100 - 97) , A에서 몇번째 더가야하는지 구함
		} else {
			rst += string(c)
		}
	}

	return rst
}

func ToUpper2_String(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' { // 소문자인가?
			builder.WriteRune('A' + (c - 'a')) // 65 + (100 - 97) , A에서 몇번째 더가야하는지 구함
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
