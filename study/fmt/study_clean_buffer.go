package study_fmt

import (
	"bufio"
	"fmt"
	"os"
)

func Clean_Buffer() {
	stdin := bufio.NewReader(os.Stdin) // Stdin -> 표준입역을 나타냄, standard input의 약자임

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 표준입력버퍼에서 \n 문자가 나올때 까지 읽어와라 라는 의미!
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 표준입력버퍼에서 \n 문자가 나올때 까지 읽어와라 라는 의미!
	} else {
		fmt.Println(n, a, b)
	}
}
