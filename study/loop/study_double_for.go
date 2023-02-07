package loop

import "fmt"

func Write_Star_for() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Write_Reverse_Star_for() {
	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Write_Triangle_for(lineNum int) {
	cnt := 1
	for i := lineNum; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= cnt; k++ {
			fmt.Print("*")
		}
		fmt.Println()
		cnt += 2
	}
}
