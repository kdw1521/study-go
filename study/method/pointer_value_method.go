package method

import "fmt"

type diff_account struct {
	balance   int
	firstname string
	lastname  string
}

func (a1 *diff_account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 diff_account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a2 diff_account) withdrawValue2(amount int) diff_account {
	a2.balance -= amount
	return a2
}

func Diff_Runner() {
	var mainA *diff_account = &diff_account{100, "wando", "kim"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	*mainA = mainA.withdrawValue2(20)
	fmt.Println(mainA.balance)
}
