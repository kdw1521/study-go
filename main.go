package main

import (
	// "fmt"
	"fmt"
	errorhandling "study-go/study/error_handling"
)

func main() {
	// line, err := errorhandling.ReadFile(errorhandling.FILENAME)
	// if err != nil {
	// 	err = errorhandling.WriteFile(errorhandling.FILENAME, "This is WriteFile")
	// 	if err != nil {
	// 		fmt.Println("파일 생성에 실패했습니다.", err)
	// 		return
	// 	}
	// 	line, err = errorhandling.ReadFile(errorhandling.FILENAME)
	// 	if err != nil {
	// 		fmt.Println("파일 읽기에 실패했습니다.", err)
	// 		return
	// 	}
	// }

	// fmt.Println("파일 내용:", line)

	// sqrt, err := errorhandling.Sqrt(-2)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// }
	// fmt.Printf("Sqrt(-2)= %v\n", sqrt)

	// err := errorhandling.RegisterAccount("wando", "1234")

	// if err != nil {
	// 	if errInfo, ok := err.(errorhandling.PasswordError); ok {
	// 		fmt.Printf("%v Len:%d RequireLen: %d\n", errInfo, errInfo.Len, errInfo.RequireLen)
	// 	}
	// } else {
	// 	fmt.Println("회원 가입 완료")
	// }

	// errorhandling.ReadEq("123 3")
	// errorhandling.ReadEq("123 abc")

	// errorhandling.Divide(9, 3)
	// errorhandling.Divide(9, 0)

	errorhandling.F()
	fmt.Println("프로그램 계속 실행")
}
