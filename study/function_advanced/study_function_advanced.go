package functionadvanced

import (
	"fmt"
	"os"
)

// ... 키워드
func Sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

// defer 키워드
func Study_Defer() {
	f, err := os.Create("test.txt") // file이 있으면 열어주고 없으면 만든다.
	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}

	defer fmt.Println("파일을 닫으려 한다!")
	defer f.Close()
	defer fmt.Println("파일을 닫았다!")

	fmt.Println("파일에 Hello World를 쓴다!")
	fmt.Fprintln(f, "Hello World")
}

// 함수 타입 변수
func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

type OpFn func(int, int) int

func GetOperator(op string) OpFn {
	if op == "+" {
		return Add
	} else if op == "*" {
		return Mul
	} else {
		return nil
	}
}

// 함수 리터럴
func GetOperator2(op string) OpFn {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

// 함수 리터럴 내부상태
func Inner_State() {
	i := 0
	f := func() {
		i += 10
	}

	i++

	f()

	fmt.Println(i)
}

// 함수 리터럴 내부상태 주의사항
func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

// 의존성 주입
type Writer func(string)

func writerHello(writer Writer) {
	writer("Hello World")
}

func MakeFile() {
	f, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer f.Close()
	writerHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
