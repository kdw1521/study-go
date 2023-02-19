package method

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func Runner() {
	a := &account{100}

	withdrawFunc(a, 30)
	fmt.Println(a.balance)
	a.withdrawMethod(30)
	fmt.Println(a.balance)
}

type myInt int

func (m myInt) Add(a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func addFunc(m myInt, a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func Runner_MyInt() {
	var a myInt
	a = a.Add(10)
	fmt.Println(a)
}
