package loop

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Infinity_For_1() {
	for true { // 이경우는 true 생략 가능!
		fmt.Print("무한~~")
	}
}

func Infinity_For_2() {
	i := 1
	for { // 1초 마다 로그 찍고 i 에 1씩 추가!
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}

func Continue_Break_For() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(" 6 *", i, "=", 6*i)
	}
}

func Get_Number_For() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력해봐")
		var number int
		// Scanln은 리턴이 입력한갯수,에러 이렇게 반환하는데 갯수는 안쓸거여서 _로 할당해준다. (왜냐하면 go 에서는 안쓰는 변수를 허용을 안하거든)
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력해")

			// 버퍼를 비워주고
			stdin.ReadString('\n')
			// 숫자 입력할때까지 for문을 돌게하자
			continue
		}
		fmt.Printf("%d 입력했구나\n", number)
		if number%2 == 0 {
			break // 짝수이면 for 문을 빠져나가자
		}
		// 짝수가 아니면 다시 for문 시작!
	}
	fmt.Println("for 문 종료!")
}
