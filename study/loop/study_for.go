package loop

import "fmt"

func Non_If_For() {
	i := 0
	for ; i < 10; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println()
	fmt.Println(i)
}

func Non_After_For() {
	// 이 경우는 i 가 계속 0 니까 무한 루프 가 된다.
	i := 0
	for i < 10 { // 또한 조건문만 있을땐, for i < 10 이렇게만 써줘도 된다.
		fmt.Print(i, ", ")
		// i++ // 이렇게 해주면 되긴한다.
	}
	fmt.Println()
	fmt.Println(i)
}
